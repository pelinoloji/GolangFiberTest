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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}