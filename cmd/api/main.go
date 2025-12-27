package main

import "github.com/ariel-rubilar/photography-api/cmd/api/bootstrap"

func main() {

	if err := bootstrap.Run(); err != nil {
		panic(err)
	}
}
