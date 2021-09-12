package service

import (
	"github.com/novalwardhana/golang-boiler-plate/package/crud/model"
	"gorm.io/gorm"
)

type service struct {
	dbMasterRead  *gorm.DB
	dbMasterWrite *gorm.DB
}

type Service interface {
	Add(user model.User) <-chan model.Result
	Update(user model.User, id int) <-chan model.Result
	Info(id int) <-chan model.Result
	Delete(id int) <-chan model.Result
}

func NewService(dbMasterRead *gorm.DB, dbMasterWrite *gorm.DB) Service {
	return &service{
		dbMasterRead:  dbMasterRead,
		dbMasterWrite: dbMasterWrite,
	}
}

func (s *service) Add(user model.User) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		tx := s.dbMasterWrite.Begin()
		if err := tx.Create(user).Error; err != nil {
			tx.Rollback()
			output <- model.Result{Error: err}
			return
		}

		tx.Commit()
		user.Password = ""
		output <- model.Result{Data: user}

	}()
	return output

}

func (s *service) Update(user model.User, id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		sql := ` update users set name = ?, password = ?, is_active = ?, updated_at = ? where id = ? `
		tx := s.dbMasterWrite.Begin()
		if err := tx.Exec(sql, user.Name, user.Password, user.IsActive, user.UpdatedAt, id).Error; err != nil {
			tx.Callback()
			output <- model.Result{Error: err}
			return
		}

		tx.Commit()
		user.Password = ""
		output <- model.Result{Data: user}
	}()
	return output
}

func (s *service) Info(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		var user model.User
		sql := ` select name, username, email, password, is_active, created_at, updated_at from users where id = ? `
		if err := s.dbMasterRead.Raw(sql, id).First(&user).Error; err != nil {
			output <- model.Result{Error: err}
			return
		}

		user.Password = ""
		output <- model.Result{Data: user}
	}()
	return output
}

func (s *service) Delete(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		sql := ` delete from users where id = ? `
		tx := s.dbMasterWrite.Begin()
		if err := tx.Exec(sql, id).Error; err != nil {
			tx.Callback()
			output <- model.Result{Error: err}
			return
		}

		tx.Commit()
		output <- model.Result{}
	}()
	return output
}
