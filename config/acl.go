package config
type AclVal struct {
	Value   uint64 `json:"value"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
}
var AclList = map[string]AclVal{
	"GET /":				    {Value:0x1 , Title:"Root"},
	"GET /login":				{Value:0x1 , Title:"login"},
	"GET /user": 				{Value:0x2 , Title:"show user"},
	"POST /user/%s/show_user": 	{Value:0x4 , Title:"insert user"},
	"POST /user/%s/show_detail":{Value:0x8 , Title:"insert user"},
	"POST /user/%s/kill": 		{Value:0x10 , Title:"insert user"},
	"POST /user/%s/edit": 		{Value:0x20 , Title:"insert user"},
}