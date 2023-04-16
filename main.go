package main

import (
	"fmt"
	"log"
	"net/http"
	"petProject/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	err := r.Run(":9010")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting Server")
	// Engine.SetTrustedProxies(nil)
}
