package main

import (
	"fmt"
	reaper "github.com/ramr/go-reaper"
)

func init() {
	go reaper.Reap()
	fmt.Println("go-reaper started")
}
