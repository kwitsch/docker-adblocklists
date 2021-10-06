package server

import (
	"adblocklists/config"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const plain = "text/plain"

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
	if entries > 0 {
		s.block = list
		s.blockEntries = entries
	}
}

func (s *Server) UpdateAllowlist(list string, entries int) {
	if entries > 0 {
		s.allow = list
		s.allowEntries = entries
	}
}

func (s *Server) getBlocklist(c *gin.Context) {
	s.Config.Resolver.VPrint("Request /blocklist")
	for i := 0; i < 120; i++ {
		if s.blockEntries > 0 {
			c.Data(http.StatusOK, plain, []byte(s.block))
			s.Config.Resolver.VPrint("Returned " + strconv.Itoa(s.blockEntries) + " entries")
			c.Done()
			return
		} else {
			sleepContext(c, time.Second)
		}
	}
}
func (s *Server) getAllowlist(c *gin.Context) {
	s.Config.Resolver.VPrint("Request /allowlist")
	for i := 0; i < 120; i++ {
		if s.allowEntries > 0 {
			c.Data(http.StatusOK, plain, []byte(s.allow))
			s.Config.Resolver.VPrint("Returned " + strconv.Itoa(s.allowEntries) + " entries")
			c.Done()
			return
		} else {
			sleepContext(c, time.Second)
		}
	}
}

func (s *Server) getHealthcheck(c *gin.Context) {
	c.Data(http.StatusOK, plain, []byte("ok"))
}

func sleepContext(c *gin.Context, delay time.Duration) {
	select {
	case <-c.Request.Context().Done():
	case <-time.After(delay):
	}
}
