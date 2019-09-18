package user

import (
	"top.guitoubing/gotest/crud"
	"top.guitoubing/gotest/model"
)

const BaseURL = "/api/users"

type Controller struct {}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Initialize(e model.EchoServer) error  {
	b := crud.BasicCRUD{
		BaseURL: BaseURL,
	}
	b.SetRecordType(&model.User{})
	err := b.Check()
	if err != nil {
		return err
	}
	e.GET(BaseURL, b.All)
	return nil
}

