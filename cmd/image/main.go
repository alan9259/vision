package main

import (
	"fmt"
	"miu-vision-api-v1/internal/router"
)

func main() {
	fmt.Println("You think I'm a child of Ultron.")
	r := router.New()
	//v1 := r.Group("/api")

	//d := platform.New()    //db connection
	//platform.AutoMigrate(d) //Create DB's if they don't exist

	/*
		add in stores
		image only one for this api
	*/
	/*
		create handler initialised with stores
	*/

	//h.Register(v1)  //register the group with the handler
	r.Logger.Fatal(r.Start("127.0.0.1:8586")) // diff port to auth api
}
