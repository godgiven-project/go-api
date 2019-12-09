package controllers

import (
	"github.com/savsgio/atreugo/v10"
)
func User(ctx *atreugo.RequestCtx) error  {
    /*name  := ctx.Params().Get("name")
    name2 := ctx.URLParam("name")*/
	return ctx.TextResponse(string("you become in users of godgiven") )
}

func Login(ctx *atreugo.RequestCtx) error {
	username :=string( functions.Strip_tags(ctx.URLParam("username")))
	password :=string( functions.Strip_tags(ctx.URLParam("password")))

	if functions.IsEmpty(username) || functions.IsEmpty(password){
		return ctx.TextResponse(string("error") )
	}
}