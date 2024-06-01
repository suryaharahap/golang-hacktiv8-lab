package main

import "golang-hactiv8-lab/chapter6/practice/routers"

func main() {
	var PORT = ":8089"

	routers.StartServer().Run(PORT)

}
