package main

import (
	"gochat/router"
)

func main() {
	r := router.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
