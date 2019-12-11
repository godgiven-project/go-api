package main

import (
	/**************************************************/
	/**************************************************/
	/**************************************************/
	"github.com/savsgio/atreugo/v10"
    "go-api/app/controllers"
    "go-api/class"
    /**************************************************/
	/**************************************************/
	/**************************************************/
)
func main() {
	/**************************************************/
	/****************** config server *****************/
	/**************************************************/
	// config port and addres
    config := &atreugo.Config{ Addr: "0.0.0.0:9002", }
	server := atreugo.New(config)
	// run before rout
	server.UseBefore(Acl.Auth)
	
    /**************************************************/
	/********************* router *********************/
	/**************************************************/
	/*server.PathWithFilters("GET","/",,filters)
	server.PathWithFilters("GET","/user",controllers.User,filters)*/
	server.Path("GET", "/", controllers.Root)
	server.Path("POST", "/Login", controllers.Login)
	server.Path("GET", "/user", controllers.User)
    /**************************************************/
	/**************** run server ggn ******************/
	/**************************************************/
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
