package controller

import (
	"github.com/novalwardhana/golang-boiler-plate/package/crud/model"
	"github.com/novalwardhana/golang-boiler-plate/package/crud/service"
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

		addResult := <-c.Add(user)
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
