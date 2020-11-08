package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var users []*User

func main() {
	app := fiber.New()

	//app.Get("/:value", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello,  " + c.Params("value"))
	//})

	app.Post("/users", createUser)
	app.Get("/users", readUsers)
	app.Get("/users/:id", readUser)
	app.Put("/users/", updateUser)

	app.Listen(":8080")
}

func updateUser(ctx *fiber.Ctx) error {

	//id := ctx.Params("id")

	user := new(User)
	err := ctx.BodyParser(user)

	if err != nil {
		return err
	}

	for i, userUpdate := range users {
		if userUpdate.ID == user.ID {
			users[i].UserName = user.UserName
			users[i].Pass = user.Pass
			return ctx.JSON(userUpdate)
		}
	}
	return fiber.NewError(fiber.StatusBadRequest, "UPSSS!")

}

func readUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	for _, user := range users {
		if user.ID == id {
			return ctx.JSON(user)
		}
	}
	return fiber.NewError(fiber.StatusBadRequest, "NOOOOOOOOOOOOOOO!")
}

func readUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(&users)
}

func createUser(ctx *fiber.Ctx) error {
	//userName := ctx.FormValue("username")
	//pass := ctx.FormValue("password")
	user := new(User)
	err := ctx.BodyParser(user)

	if err != nil {
		return err
	}
	users = append(users, user)
	fmt.Printf("username: %s password: %s \n", user.UserName, user.Pass)
	return ctx.SendString("DONE!")
}

type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Pass     string `json:"password"`
}
