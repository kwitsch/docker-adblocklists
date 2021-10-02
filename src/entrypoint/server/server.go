package server

import (
	"adblocklists/config"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config       *config.Config
	block        string
	blockEntries int
	allow        string
	allowEntries int
}

func New(conf *config.Config) *Server {
	res := &Server{
		conf,
		"",
		0,
		"",
		0,
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

func (s *Server) UpdateBlocklist(list string, entries int) {
	s.block = list
	s.blockEntries = entries
}

func (s *Server) UpdateAllowlist(list string, entries int) {
	s.allow = list
	s.allowEntries = entries
}

func (s *Server) getBlocklist(c *gin.Context) {
	s.returnString(c, s.block, s.blockEntries)
}
func (s *Server) getAllowlist(c *gin.Context) {
	s.returnString(c, s.allow, s.allowEntries)
}
func (s *Server) getHealthcheck(c *gin.Context) {
	s.returnString(c, "ok", 1)
}

func (s *Server) returnString(c *gin.Context, value string, entries int) {
	if entries > 0 {
		c.Data(http.StatusOK, "text/plain", []byte(value))
		s.Config.Resolver.VPrint("Returned " + strconv.Itoa(entries) + " entries")
	} else {
		c.AbortWithStatus(http.StatusTooEarly)
		s.Config.Resolver.VPrint("List isen't ready yet")
	}
}
