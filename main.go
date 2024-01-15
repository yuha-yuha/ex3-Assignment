package main

import "ej-ex3-backend/controller"

func main() {
	r := controller.Router()

	r.Run()
}
