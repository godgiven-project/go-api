# GodGiven Project GoLang REST API
This is a fast golang api that response more than 26K request/s according to [TechEmpower](https://www.techempower.com/benchmarks/#section=data-r18&hw=ph&test=query) that provide DgraphClient package to connect Dgraph database.<br/>
Dgraph is a fast noSQL-GraphQL database [DgraphOfficalWebSite](https://dgraph.io/)<br/>

To start first you should install golang and dgraph on your server following by their documentaion also we provide our documentation under Document directory of our project.<br/>

## project structure
<pre>
app/
    controllers/

class/
    auth.go
    functions/
            functions.go
    dgraph/
        DgraphClient.go

config/
    config.go
    acl.go
    
documents/
    GOLANG_README.md
    README.md
    DGRAPH_README.md

main.go
go.sum
go.mod
</pre>

## routing
you can define any routes as you like in main.go as follows:
<pre>
import "go-api/app/controllers"

/**************************************************/
/********************* router *********************/
/**************************************************/
server.Path("GET", "/", controllers.Root)
server.Path("GET", "/user", controllers.User)
...
</pre>
you must update acl.go file under config directory to define rules of permissions as follows:<br/>
"`METHOD`     `ROUTE`":			        {Value:`Basis2_NUM` , Title:"`Root`"}<br/>
LIKE:
<pre>
var AclList = map[string]AclVal{
    "GET        /":                     {Value:0x1  , Title:"Root"},
    "GET        /about":                {Value:0x2  , Title:"about us"},
    "POST       /user/%s/update_user": 	{Value:0x4  , Title:"update user"},
    "POST       /user/%s/show_user":    {Value:0x8  , Title:"show user"},
    "POST       /user/%s/drop_user": 	{Value:0x10 , Title:"drop user"},
    "POST       /user/%s/all_users": 	{Value:0x20 , Title:"show all users"},
}
</pre>
>  if parameter exist in any routes you should replace it with `%s`

by running `go run main.go` project starts working. 

## Dgraph database

package `go-api/class/dgraphClient/` is a package that contain dgo version2 database stuffs.
`NewClient` is a public function to create connection to database also by passing the q (stands for query variable) to `RunQuery` function, it return data you wanted.
<pre>
q := `{
	user(func: has(persons.userID)) {
		persons.userID,
		persons.Fname,
		persons.Lname,
		persons.email,
		persons.address
	}
}`
result := dgraphClient.RunQuery(q)
fmt.Println(result)
</pre>




