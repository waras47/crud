package main

import "crud_test/routers"

func main() {

	var PORT = ":8080"
	routers.StarttServer().Run(PORT)
}