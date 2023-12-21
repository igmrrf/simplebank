package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/igmrrf/simplebank/db/sqlc"
)

const name string = "Francis"

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.GET("/", server.home)

	fmt.Printf("Hello %s, this is a test", name)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) home(ctx *gin.Context) {
	fmt.Printf("Hello %s, this is the home route", name)
	ctx.JSON(http.StatusOK, "You're welcome home")
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()

		latency := time.Since(t)

		fmt.Printf("%s %s %s %s\n", c.Request.Method, c.Request.RequestURI, c.Request.Proto, latency)
	}
}

func ResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Next()

		fmt.Printf("%d %s %s\n", c.Writer.Status(), c.Request.Method, c.Request.RequestURI)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
