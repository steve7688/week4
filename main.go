package main

import "Week4/internal/router"

func main() {

	router := router.InitRouter()
	router.Run(":8867")

}
