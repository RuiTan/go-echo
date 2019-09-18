package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"top.guitoubing/gotest/controller/user"
)

type Server struct {
	Addr	string
	e		*echo.Echo
}

func NewServer(Addr string) *Server {
	return &Server{
		Addr: Addr,
		e:    echo.New(),
	}
}

func (s *Server) Init() (err error) {
	s.e.Debug = true
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	//g := s.e.Group("")
	err = user.NewController().Initialize(s.e)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Start()  {
	s.e.Logger.Fatal(s.e.Start(s.Addr))
}