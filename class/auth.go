package Acl

import (	
	"github.com/savsgio/atreugo/v10"
	"go-api/class/ggn_encryption"
	"go-api/config"
	"encoding/base64"
	"encoding/json"
	"strings"
	"regexp"
	// "fmt"
)
var aclList = config.AclList
type Token_data struct {
	User_id uint64
	Premission uint64
	Username string
	Name string
}
func Auth(ctx *atreugo.RequestCtx) error {
	var user_permissions uint64 = 1;
	sKey := config.Security_key
	reqToken := string(ctx.Request.Header.Peek("Authorization"))
	Token := strings.TrimSpace( strings.ReplaceAll(reqToken,"Bearer " , "") )
	//Token = "EEsnEhwZaSoSZnJLeVJicG1HUHMFX3dGTW4JBx0MDT4LBiIjZ31heENMYns7RTU7HhQXME1CZg0qAVlcOUBKcgY7AVRTfhI9XBBULCwCInUW"
	//fmt.Println(string(sKey))
	TokenDbase64 , _ := base64.StdEncoding.DecodeString( Token )
	TokenD := ggn_encryption.Decrypt(string(TokenDbase64), sKey )
	
	//fmt.Println(TokenD)
	var Valid_Token Token_data
	//var Valid_Token map[string]interface{}
	Invalid_Token := json.Unmarshal([]byte(TokenD), &Valid_Token)
	if(Invalid_Token != nil){
		user_permissions = 1
	}else{
		//fmt.Println("ok")
		user_permissions = Valid_Token.Premission
	}
	// Auth function for test IsAllow test
	current_route   := string(ctx.Path())
	current_method := string(ctx.Method())
	allow := IsAllow(user_permissions,current_route,current_method)
	//fmt.Println(allow)
	if(allow){
		return ctx.Next()
	}
	// return ctx.TextResponse(string("you become in "+ current_route+ " by method: "+ current_method) )
	return ctx.TextResponse("result")
	//return ctx.Next()
}
func IsAllow( permissions uint64 , route string , method string ) bool {
	// what is current user ????
	current_route   := string(route)
	current_route 	= strings.ReplaceAll(current_route, "/", "\\/") 
	current_route    = strings.ReplaceAll(current_route, "%s", ".*") 
	current_route 	= string("^"+current_route+"$")
	current_method := string(method)
	//fmt.Println(current_route)
	//multi_part_current_rout := strings.Split(current_route, "/")
	Allow := false
	for loop_acl := range aclList {
		//************************  fix space  *****************************
		fixspace := regexp.MustCompile(`\s+`);
		loop_acl2 := fixspace.ReplaceAllString(loop_acl," ");
		current_acl:= strings.Split(loop_acl2, " ")
		//******************************************************************
		//******************************************************************
		//******************************************************************
		if Allow != false{
			break
		}
		//******************************************************************
		//******************************************************************
		acl_method := string(current_acl[0])
		acl_path   := string(current_acl[1])
		route_match , _ := regexp.Compile(current_route)
		if( route_match.MatchString(string(acl_path)) ) {
			if acl_method == current_method {
				//fmt.Println(aclList[loop_acl].Value)
				//fmt.Println(permissions)
				if(aclList[loop_acl].Value&permissions!=0){
					Allow = true
					return Allow
				}
			}
		}
	}
	return Allow
}
