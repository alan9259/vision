package main

import (
	"fmt"
	"vision/internal/handler"
	"vision/internal/platform"
	"vision/internal/router"
	"vision/internal/store"
)

func main() {
	fmt.Println("You think I'm a child of Ultron.")
	r := router.New()
	v1 := r.Group("/api")

	d := platform.New() //db connection
	//platform.AutoMigrate(d) //Create DB's if they don't exist

	cs := store.NewAzureStore(d)
	/*
		add in stores
		image only one for this api
	*/
	/*
		create handler initialised with stores
	*/

	h := handler.NewHandler(cs)

	h.Register(v1)                            //register the group with the handler
	r.Logger.Fatal(r.Start("127.0.0.1:8586")) // diff port to auth api
}
