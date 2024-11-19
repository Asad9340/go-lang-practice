// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// )

// type ToDo struct {
// 	ID          int    `json:"id"`
// 	Completed   bool   `json:"completed"`
// 	Description string `json:"description"`
// }

// var todos []ToDo // Persistent in-memory storage

// func main() {
// 	fmt.Println("Hello, World!")
// 	app := fiber.New()
// 	err:=godotenv.Load(".env")
// 	if err!= nil {
//     log.Fatal("Error loading.env file")
//   }
// 	PORT:=os.Getenv("PORT")
// 	// Define routes
// 	app.Get("/", handlerFunc)
// 	app.Get("/api/todos", handlerGetTodosFunc)
// 	app.Post("/api/todos", handlerTodosFunc)
// 	app.Patch("/api/todos/:id",handlerFuncPatchTodos)
// 	app.Delete("api/todos/:id", handlerDeleteTodos)
// 	// Start server
// 	log.Fatal(app.Listen(":"+PORT))
// }

// // Handler for the root endpoint
// func handlerFunc(c *fiber.Ctx) error {
// 	return c.Status(200).JSON(fiber.Map{"message": "API is working successfully"})
// }

// // Handler to get all todos
// func handlerGetTodosFunc(c *fiber.Ctx) error {
// 	return c.Status(200).JSON(fiber.Map{"data": todos})
// }

// // Handler to create a new todo
// func handlerTodosFunc(c *fiber.Ctx) error {
// 	todo := &ToDo{}
// 	if err := c.BodyParser(todo); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
// 	}

// 	if todo.Description == "" {
// 		return c.Status(400).JSON(fiber.Map{"error": "ToDo description is required"})
// 	}

// 	// Assign a unique ID and add to the list
// 	todo.ID = len(todos) + 1
// 	todos = append(todos, *todo)

// 	return c.Status(201).JSON(todo)
// }
// func handlerFuncPatchTodos(c *fiber.Ctx) error  {
// 	id:=c.Params("id")
// 	for i,todo := range todos {
// 		if(fmt.Sprint(todo.ID)==id){
// 			todos[i].Completed = true
// 			return c.Status(200).JSON(todos[i])
// 		}
// 	}
// 	return c.Status(404).JSON(fiber.Map{"error": "ToDo not found"})
// }
// func handlerDeleteTodos(c *fiber.Ctx) error  {
// 	id:=c.Params("id");
// 	for i, todo := range todos {
// 		if(fmt.Sprint(todo.ID)==id){
// 			todos = append(todos[:i], todos[i+1:]...)
//       return c.Status(204).JSON(fiber.Map{"success":"true"})
// 		}
// 	}
// 	return c.Status(404).JSON(fiber.Map{"error": "ToDo not found"})
// }

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDo struct {
	ID          primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

var collection *mongo.Collection

func main()  {
	fmt.Println("Hello world!")
 err:=godotenv.Load(".env")
 if err!= nil {
    log.Fatal("Error loading.env file")
  }
	// PORT := os.Getenv("PORT")
	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client , err :=mongo.Connect(context.Background(),clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(),nil)
		if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MONGO server")
	collection = client.Database("golang_db").Collection("todos");

	app:=fiber.New()
	app.Get("/api/v1/todos",getAllTodos)
	app.Post("/api/v1/todos",createTodo)
	app.Patch("/api/v1/todos/:id",updateTodo)
	// app.Delete("/api/v1/todos/:id",deleteTodo)
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}
	log.Fatal(app.Listen(":"+port))
}

func getAllTodos(c *fiber.Ctx) error  {
		var todos []ToDo
		result , err := collection.Find(context.Background(),bson.M{})
		if err!= nil {
      return c.Status(500).SendString(err.Error())
    }
		defer result.Close(context.Background())
		for result.Next(context.Background()){
			var todo ToDo
      err:= result.Decode(&todo)
      if err!= nil {
        return c.Status(500).SendString(err.Error())
      }
      todos = append(todos, todo)
		}
		return c.JSON(todos)
}
func createTodo(c *fiber.Ctx) error  {
		todo:=new(ToDo)
		if err:= c.BodyParser(todo);err!= nil{
			return c.Status(400).SendString("Invalid request body")
		}
		if todo.Description=="" {
		return c.Status(400).SendString("Description is required")
		}
		result, err:=collection.InsertOne(context.Background(),todo)
		if err!=nil {
			return err
		}
		todo.ID=result.InsertedID.(primitive.ObjectID)
		return c.Status(201).JSON(todo)
}
func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Object ID"})
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No ToDo found"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "ToDo updated successfully"})
}

// func deleteTodo(c *fiber.Ctx) error  {

// }