package main

import (
	"context"
	"github.com/gofiber/fiber/v2" // HTTP methods
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson" // mongo-driver using for saving the sate from Mongo DB
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
	app.Get("/users", readUsers(userCollection))
	app.Get("/users/:id", readUser(userCollection))
	app.Get("/users/un/:username", readUserByUsername(userCollection))
	app.Put("/users/", updateUser(userCollection))
	app.Delete("/users/:id", deleteUser(userCollection))

	app.Listen(":8080")
}

//Mongo DB Remove user by username
func readUserByUsername(uc *mongo.Collection) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		username := ctx.Params("username")
		var user User
		err := uc.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
		if err != nil {
			return err
		}
		return ctx.JSON(user)
	}
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

//Mongo DB Delete User
func deleteUser(uc *mongo.Collection) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")                   // Fiber holds this id params, with this code we're getting id params from fiber
		ID, err := primitive.ObjectIDFromHex(id) // return string to object
		if err != nil {
			return err
		}

		one, err := uc.DeleteOne(context.Background(), bson.M{"_id": ID})

		if err != nil {
			return err
		}
		return ctx.JSON(fiber.Map{"status": one.DeletedCount})
	}
}

// Fiber Update User
// func updateUser(ctx *fiber.Ctx) error {
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
// }

// Mongo DB Update User
func updateUser(uc *mongo.Collection) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user User
		err := ctx.BodyParser(&user) //I'm gonna receive user
		if err != nil {
			return err
		}
		_, err = uc.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})

		if err != nil {
			return err
		}
		return ctx.JSON(user)
	}
}

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

// Mongo DB Read User
func readUser(uc *mongo.Collection) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		ID, err := primitive.ObjectIDFromHex(id) // We get id as a string but we need to set id to object. Why? We set the object as a type in struct area.
		if err != nil {
			return err
		}
		var user User
		err = uc.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&user) // searching inside the user Collection
		if err != nil {
			return err
		}
		return ctx.JSON(user)
	}
}

// Fiber Read Users
// func readUsers(ctx *fiber.Ctx) error {
//	return ctx.JSON(&users)
// }

// Mongo DB Read Users
func readUsers(uc *mongo.Collection) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var users []User
		cursor, err := uc.Find(context.Background(), bson.D{}) //connect mongo DB and bring all data from DB

		if err != nil {
			return err
		}
		for cursor.Next(context.TODO()) { //return all data with cursor method, and it's gonna add user List all returning datas
			var user User
			_ = cursor.Decode(&user)
			users = append(users, user)
		}
		return ctx.JSON(users) //Show user list as a json format
	}
}

// Fiber Create User
// func createUser(ctx *fiber.Ctx) error {
//	user := new(User)
//	err := ctx.BodyParser(user)
//
//	if err != nil {
//		return err
//	}
//	users = append(users, user)
//	fmt.Printf("username: %s password: %s \n", user.UserName, user.Pass)
//	return ctx.JSON(user)
// }

// Mongo DB Create User
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
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserName    string             `bson:"username" json:"username"`
	Pass        string             `bson:"password" json:"password"`
	CreatedDate time.Time          `bson:"createdDate" json:"createdDate"`
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
