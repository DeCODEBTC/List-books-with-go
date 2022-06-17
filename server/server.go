package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine // É o que o gin usa para iniciar o server
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() { //Metodo que roda nosso server
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
