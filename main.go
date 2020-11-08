// PURE GO
//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//func main()  {
//	http.HandleFunc("/hello",hello)
//	http.ListenAndServe(":8080",nil)
//}
//
//func hello(writer http.ResponseWriter, request *http.Request) {
//fmt.Fprint(writer,"hello\n")}


// FIBER
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// basic server set up
	//app.Get("/", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello, World!")
	//})


	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("Hello,  " + c.Params("value"))
	})

	app.Listen(":8080")
}