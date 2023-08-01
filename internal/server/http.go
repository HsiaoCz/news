package server

import (
	"net/http"
	"time"

	"github.com/HsiaoCz/news/internal/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	us *service.UserService
}

func NewServer(us *service.UserService) *Server {
	return &Server{
		us: us,
	}
}

func (s *Server) RegisterHTTPServer(r *gin.Engine) error {
	r.GET("/api/v1/user/register", s.us.UserRegister)
	srv := http.Server{
		Handler:      r,
		Addr:         ":9091",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return srv.ListenAndServe()
}
