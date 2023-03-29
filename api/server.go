package api

import (
	"github.com/deep0ne/firstWBProject/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Database
	router *gin.Engine
}

func NewServer(store db.Database) (*Server, error) {
	server := &Server{
		store: store,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/orders/:id", server.getOrder)
	router.Static("/static", "./static/")

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
