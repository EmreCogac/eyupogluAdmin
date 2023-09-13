package main

import (
	"admin-panel/admin-panel/router"
)

func main() {

	r := router.SetupRouter()

	r.Run(":8080")
}
