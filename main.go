package main

import (
	/**************************************************/
	/**************************************************/
	/**************************************************/
	"github.com/savsgio/atreugo/v10"
    "go-api/app/controllers"
    ggn_config "go-api/config"
    "go-api/class"
    //"fmt"
    /**************************************************/
	/**************************************************/
	/**************************************************/
)
func main() {
	/**************************************************/
	/****************** config server *****************/
	/**************************************************/
	// config port and addres
    config := &atreugo.Config{ Addr: ggn_config.Server_name+":"+ggn_config.PORT, }
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
