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
	if c.Redis.Enabled() {
		c.Resolver.VPrint("Redis client enabled")
		c.Redis.Init()
		bval, bent, berr := c.Redis.GetBlock()
		if berr == nil {
			s.UpdateBlocklist(bval, bent)
		}
		aval, aent, aerr := c.Redis.GetBlock()
		if aerr == nil {
			s.UpdateAllowlist(aval, aent)
		}
	}
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
			bv, bc := blockL.ToString()
			s.UpdateBlocklist(bv, bc)
			if c.Redis.Enabled() {
				c.Redis.SetBlock(bv, bc)
			}
			c.Resolver.VPrint("---------------------")
			allowL.AddOnlineMap(httpClient, c.Allow.Lists, c.Resolver.Verbose)
			av, ac := allowL.ToString()
			s.UpdateAllowlist(av, ac)
			if c.Redis.Enabled() {
				c.Redis.SetAllow(av, ac)
			}
			c.Resolver.VPrint("---------------------")
			time.Sleep(c.Refresh)
		}
	} else {
		fmt.Println(rinitErr.Error())
		os.Exit(1)
	}
}
