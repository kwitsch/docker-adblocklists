package main

import (
	"adblocklists/config"
	"adblocklists/list"
	"adblocklists/server"
	"fmt"
	"os"
	"time"

	_ "github.com/kwitsch/go-dockerutils"
)

func main() {
	c := config.Get()
	s := server.New(c)
	s.Run()
	c.Resolver.VPrint("---------------------")
	rinitErr := c.Resolver.Init()
	httpClient, _ := c.Resolver.GetHttpClient()
	if rinitErr == nil {
		blockL := list.New()
		blockL.AddMap(c.Block.Entries)

		allowL := list.New()
		allowL.AddMap(c.Allow.Entries)

		c.Resolver.VPrint("---------------------")
		for {
			blockL.AddOnlineMap(httpClient, c.Block.Lists, c.Resolver.Verbose)
			s.UpdateBlocklist(blockL.ToString())
			c.Resolver.VPrint("---------------------")
			allowL.AddOnlineMap(httpClient, c.Allow.Lists, c.Resolver.Verbose)
			s.UpdateAllowlist(allowL.ToString())
			c.Resolver.VPrint("---------------------")
			time.Sleep(c.Refresh)
		}
	} else {
		fmt.Println(rinitErr.Error())
		os.Exit(1)
	}
}
