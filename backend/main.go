package main

import (
	"ej-ex3-backend/controller"
	"ej-ex3-backend/database"
)

func init() {
	database.DBconn()
}
func main() {
	r := controller.Router()

	r.Run(":8080")
}
