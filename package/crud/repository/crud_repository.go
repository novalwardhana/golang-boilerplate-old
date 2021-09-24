package repository

import (
	"github.com/novalwardhana/golang-boilerplate/package/crud/model"
	"gorm.io/gorm"
)

type repository struct {
	dbMasterRead  *gorm.DB
	dbMasterWrite *gorm.DB
}

type Repository interface {
	CountData() <-chan model.Result
	GetData(params model.Params) <-chan model.Result
	Add(user model.User) <-chan model.Result
	Update(user model.User, id int) <-chan model.Result
	Info(id int) <-chan model.Result
	Delete(id int) <-chan model.Result
}

func NewRepository(dbMasterRead *gorm.DB, dbMasterWrite *gorm.DB) Repository {
	return &repository{
		dbMasterRead:  dbMasterRead,
		dbMasterWrite: dbMasterWrite,
	}
}

func (r *repository) CountData() <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
		var count int64

		sql := `select count(*) from users`
		if err := r.dbMasterRead.Raw(sql).Count(&count).Error; err != nil {
			output <- model.Result{Error: err}
			return
		}

		output <- model.Result{Data: count}
	}()
	return output
}

func (r *repository) GetData(params model.Params) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		offset := (params.Page - 1) * params.Limit
		orderBy := ` id DESC `

		db := r.dbMasterRead.Order(orderBy)
		db = db.Limit(params.Limit)
		db = db.Offset(offset)

		var datas []model.User
		if err := db.Find(&datas).Error; err != nil {
			output <- model.Result{Error: err}
			return
		}

		for i := range datas {
			datas[i].Password = ""
		}

		output <- model.Result{Data: datas}
	}()
	return output
}

func (r *repository) Add(user model.User) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		tx := r.dbMasterWrite.Begin()
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

func (r *repository) Update(user model.User, id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		sql := ` update users set name = ?, username = ?, password = ?, is_active = ?, updated_at = ? where id = ? `
		tx := r.dbMasterWrite.Begin()
		if err := tx.Exec(sql, user.Name, user.Username, user.Password, user.IsActive, user.UpdatedAt, id).Error; err != nil {
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

func (r *repository) Info(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		var user model.User
		sql := ` select name, username, email, password, is_active, created_at, updated_at from users where id = ? `
		if err := r.dbMasterRead.Raw(sql, id).First(&user).Error; err != nil {
			output <- model.Result{Error: err}
			return
		}

		user.Password = ""
		output <- model.Result{Data: user}
	}()
	return output
}

func (r *repository) Delete(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		sql := ` delete from users where id = ? `
		tx := r.dbMasterWrite.Begin()
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
