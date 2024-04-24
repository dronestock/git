package main

import (
	"github.com/dronestock/drone"
	"github.com/dronestock/git/internal"
)

func main() {
	drone.New(internal.New).Boot()
}
