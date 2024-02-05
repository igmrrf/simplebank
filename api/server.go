package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/igmrrf/simplebank/db/sqlc"
	"github.com/igmrrf/simplebank/token"
	"github.com/igmrrf/simplebank/util"
)

const name string = "Francis"

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/", server.home)

	userGroup := router.Group("/users")
	userGroup.POST("", server.createUser)
	userGroup.POST("login", server.loginUser)

	accountGroup := router.Group("/accounts").Use(authMiddleware(server.tokenMaker))
	accountGroup.POST("", server.createAccount)
	accountGroup.GET(":id", server.getAccount)
	accountGroup.GET("", server.listAccounts)

	transferGroup := router.Group("/transfers").Use(authMiddleware(server.tokenMaker))
	transferGroup.POST("", server.createTransfer)

	server.router = router
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
