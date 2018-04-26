package main

const (
	username = "root"
	password = "root"
	database = "rest_api_example"
)

func main() {
	a := App{}

	a.Initialize(username, password, database)

	a.Run(":8080")
}
