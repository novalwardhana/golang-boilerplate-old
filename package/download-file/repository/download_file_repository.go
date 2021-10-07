package repository

import (
	"github.com/novalwardhana/golang-boilerplate/package/download-file/model"
	"gorm.io/gorm"
)

type repository struct {
	dbMasterRead  *gorm.DB
	dbMasterWrite *gorm.DB
}

type Repository interface {
	GetFileInfo(filename string) <-chan model.Result
}

func NewRepository(dbMasterRead, dbMasterWrite *gorm.DB) Repository {
	return &repository{
		dbMasterRead:  dbMasterRead,
		dbMasterWrite: dbMasterWrite,
	}
}

func (r *repository) GetFileInfo(filename string) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		var file model.File
		if err := r.dbMasterRead.Where("name", filename).First(&file).Error; err != nil {
			output <- model.Result{Error: err}
			return
		}

		output <- model.Result{Data: file}
	}()
	return output
}
