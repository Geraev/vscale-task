package handlers

import (
	"net/http"
	"strconv"
	"vscale-task/cmd/manager"

	"vscale-task/cmd/providers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port    string
	manager manager.APIManager
}

func NewServer(port string, manager *manager.APIManager) *Server {
	return &Server{
		port:    port,
		manager: *manager,
	}
}

func (s *Server) Run() error {
	r := gin.Default()

	r.POST("/:number", s.CreateServer)
	r.DELETE("/:group")

	return r.Run(":" + s.port)
}

func (s *Server) CreateServer(c *gin.Context) {
	var param = c.Param("number")
	number, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
	}

	var servReq providers.CreateServerRequest
	if err := c.ShouldBindJSON(&servReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var (
		groupID    int64
		err        error
		isAccepted = make(chan struct{}, 1)
	)

	go func() {
		if groupID, err = s.manager.CreateServerGroup(&servReq, number); err != nil {
			switch err {
			case manager.ErrNeedRollback:
			default:

			}
		}

	}()

}
