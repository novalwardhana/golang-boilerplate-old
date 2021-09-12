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
	Update() <-chan model.Result
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
		output <- model.Result{Data: user}

	}()
	return output

}

func (s *service) Update() <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
	}()
	return output
}

func (s *service) Delete(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
	}()
	return output
}
