package controllers

import (
	"github.com/savsgio/atreugo/v10"
	"go-api/class/functions"
	_"encoding/json"
	// "fmt"
)
func User(ctx *atreugo.RequestCtx) error  {
    /*name  := ctx.Params().Get("name")
    name2 := ctx.URLParam("name")*/
	return ctx.TextResponse(string("you become in users of godgiven") )
}

func Login(ctx *atreugo.RequestCtx) error {
	username := string( functions.Validate_data(string(ctx.FormValue("username")) ))
	password := string( functions.Validate_data(string(ctx.FormValue("password")) ))

	// if functions.IsEmpty(username) || functions.IsEmpty(password){
	// 	return ctx.TextResponse(string("error") )
	// }
	return ctx.TextResponse( string(username+password) )
}