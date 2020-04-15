package httpserver

import (
	"net/http"
	"strconv"
	"vscale-task/cmd/manager"
	"vscale-task/cmd/storage"

	"vscale-task/cmd/providers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port    string
	manager manager.APIManager
}

func NewServer(port string, manager manager.APIManager) *Server {
	return &Server{
		port:    port,
		manager: manager,
	}
}

func (s *Server) Run() error {
	r := gin.Default()

	r.POST("/:number", s.CreateServerGroup)
	r.DELETE("/:groupid", s.DeleteServerGroup)
	r.GET("/:groupid", s.StatusServerGroup)

	return r.Run(":" + s.port)
}

func (s *Server) CreateServerGroup(c *gin.Context) {
	var param = c.Param("number")
	number, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	var servReq providers.CreateServerRequest
	if err := c.ShouldBindJSON(&servReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var (
		groupID    int64
		chAccepted = make(chan int64, 1)
		chFailed   = make(chan struct{}, 1)
	)

	go func() {
		if err = s.manager.CreateServerGroup(chAccepted, &servReq, int64(number)); err != nil {
			chFailed <- struct{}{}
		}
	}()

	select {
	case groupID = <-chAccepted:
		c.JSON(http.StatusAccepted, APICreateRespone{
			Status:  "accepted",
			GroupID: groupID,
		})
	case <-chFailed:
		c.JSON(http.StatusInternalServerError, APICreateRespone{
			Status:  "failed",
			GroupID: groupID,
		})
	}
}

func (s *Server) DeleteServerGroup(c *gin.Context) {
	var param = c.Param("groupid")
	groupID, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	status, ok := s.manager.Storage.GetGroupStatus(int64(groupID))
	if !ok {
		c.JSON(
			http.StatusNotFound,
			gin.H{"error": "server group not found"},
		)
		return
	}
	if status != storage.StatusComplete {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "server group not complete"},
		)
		return
	}

	if err = s.manager.DeleteServerGroup(int64(groupID)); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) StatusServerGroup(c *gin.Context) {
	var param = c.Param("groupid")
	groupID, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	status, ok := s.manager.Storage.GetGroupStatus(int64(groupID))
	if !ok {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "server group not found"},
		)
		return
	}

	c.JSON(http.StatusOK,
		APICreateRespone{
			Status:  status,
			GroupID: int64(groupID),
		},
	)
}
