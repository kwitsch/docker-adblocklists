package server

import (
	"adblocklists/config"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config *config.Config
	block  string
	allow  string
}

func New(conf *config.Config) *Server {
	res := &Server{
		conf,
		"",
		"",
	}
	return res
}

func (s *Server) Run() {
	go func() {
		router := gin.Default()
		router.GET("/blocklist", s.getBlocklist)
		router.GET("/allowlist", s.getAllowlist)
		router.GET("/healtcheck", s.getHealthcheck)
		port := "80"
		if runtime.GOOS == "windows" {
			port = "8080"
		}
		router.Run(":" + port)
	}()
}

func (s *Server) UpdateBlocklist(list string) {
	s.block = list
}

func (s *Server) UpdateAllowlist(list string) {
	s.allow = list
}

func (s *Server) getBlocklist(c *gin.Context) {
	returnRaw(c, s.block)
}
func (s *Server) getAllowlist(c *gin.Context) {
	returnRaw(c, s.allow)
}
func (s *Server) getHealthcheck(c *gin.Context) {
	returnRaw(c, "ok")
}

func returnRaw(c *gin.Context, value string) {
	c.Data(http.StatusOK, "", []byte(value))
}
