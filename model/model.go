package model

import (
	"github.com/labstack/echo"
	"top.guitoubing/gotest/db"
)

const (
	ContextUser = "CtxUser"
)

type RunningMode string

const (
	ModeNoauth RunningMode = "noauth"
	ModeDev RunningMode = "dev"
	ModeProd RunningMode = "prod"
)

const (
	FAPIResponsePre = "APIResponsePre"
)

type MonthEnum int

type EchoServer interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

type Entity interface {
	GenerateID()
	PersisPre()
	APIResponsePre()
	CollectionName() string
}

type Helper struct {
	t Entity
}

func NewHelper(entity Entity) *Helper  {
	return &Helper{
		t:entity,
	}
}

func (h *Helper) GetCollection() db.GetCollection {
	return db.Global.Collection(h.t.CollectionName())
}

func (h *Helper) All(result interface{}) (err error) {
	collection, closeConn := h.GetCollection()()
	defer closeConn()
	return collection.Find(nil).All(result)
}