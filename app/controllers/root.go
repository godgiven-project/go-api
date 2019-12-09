package controllers

import (
	/**************************************************/
	/**************************************************/
	/**************************************************/
	"github.com/savsgio/atreugo/v10"
)
func Root(ctx *atreugo.RequestCtx) error {
	return ctx.TextResponse(string("you become in root of godgiven") )
}