package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

//var users []*User

func main() {
	db, err := Connect()
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(0)
		return
	}
	userCollection := db.Collection("users")

	app := fiber.New()

	//app.Get("/:value", func(c *fiber.Ctx) error {
	//	return c.SendString("Hello,  " + c.Params("value"))
	//})

	app.Post("/users", createUser(userCollection))
	//app.Get("/users", readUsers)
	//app.Get("/users/:id", readUser)
	//app.Put("/users/", updateUser)
	//app.Delete("/users/:id", deleteUser)
	app.Listen(":8080")
}

// Fiber Delete User

//func deleteUser(ctx *fiber.Ctx) error {
//	id := ctx.Params("id")
//
//	for i, userDelete := range users {
//		if userDelete.ID == id {
//			remove(users, i) //not working properly
//		}
//	}
//	return ctx.JSON(fiber.Map{"status": "DELETED!!"})
//}

// Fiber Update User

//func updateUser(ctx *fiber.Ctx) error {
//
//	user := new(User)
//	err := ctx.BodyParser(user)
//
//	if err != nil {
//		return err
//	}
//
//	for i, userUpdate := range users {
//		if userUpdate.ID == user.ID {
//			users[i].UserName = user.UserName
//			users[i].Pass = user.Pass
//			return ctx.JSON(userUpdate)
//		}
//	}
//	return fiber.NewError(fiber.StatusBadRequest, "UPSSS!")
//
//}

// Fiber Read User

//func readUser(ctx *fiber.Ctx) error {
//	id := ctx.Params("id")
//	for _, user := range users {
//		if user.ID == id {
//			return ctx.JSON(user)
//		}
//	}
//	return fiber.NewError(fiber.StatusBadRequest, "NOOOOOOOOOOOOOOO!")
//}

// Fiber Read Users

//func readUsers(ctx *fiber.Ctx) error {
//	return ctx.JSON(&users)
//}

// Fiber Create User

//func createUser(ctx *fiber.Ctx) error {
//	user := new(User)
//	err := ctx.BodyParser(user)
//
//	if err != nil {
//		return err
//	}
//	users = append(users, user)
//	fmt.Printf("username: %s password: %s \n", user.UserName, user.Pass)
//	return ctx.JSON(user)
//}

func createUser(uc *mongo.Collection) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := new(User)
		err := ctx.BodyParser(user)

		if err != nil {
			logrus.Error(err.Error())
			return err
		}
		user.ID = primitive.NewObjectID() // generates new ID
		_, err = uc.InsertOne(context.Background(), user)

		if err != nil {
			logrus.Error(err.Error())
			return err
		}
		return ctx.JSON(user)
	}
}

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	UserName string             `bson:"username" json:"username"`
	Pass     string             `bson:"password" json:"password"`
}

//Mongo DB Connection Code Block
func Connect() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	db := client.Database("test")
	return db, nil

}
