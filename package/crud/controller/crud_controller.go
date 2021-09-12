package controller

import (
	"time"

	"github.com/novalwardhana/golang-boiler-plate/package/crud/model"
	"github.com/novalwardhana/golang-boiler-plate/package/crud/service"
	"golang.org/x/crypto/bcrypt"
)

type controller struct {
	service service.Service
}

type Controller interface {
	Add(user model.User) <-chan model.Result
	Update(user model.User, id int) <-chan model.Result
	Info(id int) <-chan model.Result
	Delete(id int) <-chan model.Result
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) Add(user model.User) <-chan model.Result {
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

		addResult := <-c.service.Add(user)
		if addResult.Error != nil {
			output <- model.Result{Error: addResult.Error}
			return
		}

		output <- model.Result{Data: addResult.Data}
	}()
	return output
}

func (c *controller) Update(user model.User, id int) <-chan model.Result {
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

		updateResult := <-c.service.Update(user, id)
		if updateResult.Error != nil {
			output <- model.Result{Error: updateResult.Error}
			return
		}

		output <- model.Result{Data: updateResult.Data}
	}()
	return output
}

func (c *controller) Info(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
	}()
	return output
}

func (c *controller) Delete(id int) <-chan model.Result {
	output := make(chan model.Result)
	go func() {
		defer close(output)
	}()
	return output
}
