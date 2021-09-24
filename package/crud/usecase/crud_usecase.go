package usecase

import (
	"math"
	"time"

	"github.com/novalwardhana/golang-boilerplate/package/crud/model"
	"github.com/novalwardhana/golang-boilerplate/package/crud/repository"
	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	repository repository.Repository
}

type Usecase interface {
	List(params model.Params) <-chan model.Result
	Add(user model.User) <-chan model.Result
	Update(user model.User, id int) <-chan model.Result
	Info(id int) <-chan model.Result
	Delete(id int) <-chan model.Result
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) List(params model.Params) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		countDataResult := <-u.repository.CountData()
		if countDataResult.Error != nil {
			output <- model.Result{Error: countDataResult.Error}
			return
		}
		countData := countDataResult.Data.(int64)

		getDataResult := <-u.repository.GetData(params)
		if getDataResult.Error != nil {
			output <- model.Result{Error: getDataResult.Error}
			return
		}
		data := getDataResult.Data.([]model.User)

		var userList model.UserList
		userList.Total = int(countData)
		userList.PerPage = params.Limit
		userList.Page = params.Page
		userList.Data = data
		lastPage := math.Ceil(float64(countData) / float64(params.Limit))
		userList.LastPage = int(lastPage)

		output <- model.Result{Data: userList}
	}()
	return output
}

func (u *usecase) Add(user model.User) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		user.IsActive = true
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		user.CreatedAt = timeNow
		user.UpdatedAt = timeNow

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), model.PasswordHashCost)
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		user.Password = string(encryptPassword)

		addResult := <-u.repository.Add(user)
		if addResult.Error != nil {
			output <- model.Result{Error: addResult.Error}
			return
		}

		output <- model.Result{Data: addResult.Data}
	}()
	return output
}

func (u *usecase) Update(user model.User, id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		timeNow := time.Now().Format("2006-01-02 15:04:05")
		user.UpdatedAt = timeNow

		encryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), model.PasswordHashCost)
		if err != nil {
			output <- model.Result{Error: err}
			return
		}
		user.Password = string(encryptPassword)

		updateResult := <-u.repository.Update(user, id)
		if updateResult.Error != nil {
			output <- model.Result{Error: updateResult.Error}
			return
		}

		output <- model.Result{Data: updateResult.Data}
	}()
	return output
}

func (u *usecase) Info(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		infoResult := <-u.repository.Info(id)
		if infoResult.Error != nil {
			output <- model.Result{Error: infoResult.Error}
			return
		}

		output <- model.Result{Data: infoResult.Data}
	}()
	return output
}

func (u *usecase) Delete(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)

		deleteResult := <-u.repository.Delete(id)
		if deleteResult.Error != nil {
			output <- model.Result{Error: deleteResult.Error}
			return
		}

		output <- model.Result{}
	}()
	return output
}
