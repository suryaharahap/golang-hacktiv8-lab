package main

import (
	"golang-hactiv8-lab/chapter7/practice/go-jwt/database"
	"golang-hactiv8-lab/chapter7/practice/go-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8083")
}
