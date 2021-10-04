package repository

import (
	"github.com/novalwardhana/golang-boilerplate/package/upload-file/model"
	"gorm.io/gorm"
)

type repository struct {
	dbMasterRead  *gorm.DB
	dbMasterWrite *gorm.DB
}

type Repository interface {
	SaveFileInfo(file model.File) <-chan model.Result
	CSVToDatabase(users []model.User) <-chan model.Result
}

func NewRepository(dbMasterRead, dbMasterWrite *gorm.DB) Repository {
	return &repository{
		dbMasterRead:  dbMasterRead,
		dbMasterWrite: dbMasterWrite,
	}
}

func (r *repository) SaveFileInfo(file model.File) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		tx := r.dbMasterWrite.Begin()
		if err := tx.Create(file).Error; err != nil {
			output <- model.Result{Error: err}
			return
		}
		tx.Commit()
		output <- model.Result{}

	}()
	return output
}

func (r *repository) CSVToDatabase(users []model.User) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		tx := r.dbMasterWrite.Begin()
		if err := tx.Create(users).Error; err != nil {
			output <- model.Result{Error: err}
			return
		}
		tx.Commit()
		output <- model.Result{}

	}()
	return output
}
