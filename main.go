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

type Todo struct {
	ID        primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection
func main() {
	fmt.Println("Todo List Application")

	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file",err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	// Create a new MongoDB client
	clienOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clienOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()
	
	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodos)
	app.Patch("/api/todos/:id", patchTodos)
	app.Delete("/api/todos/:id", deleteTodos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000" // Default port if not set in .env
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo
	cursor, err :=collection.Find(context.Background(),bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch todos",
		})
	}

	defer cursor.Close(context.Background())
	
	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to decode todo",
			})
		}
		todos = append(todos, todo)
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}
func createTodos(c *fiber.Ctx) error {
	todo := new(Todo)  // this is a pointer to Todo struct
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	if todo.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Todo body cannot be empty",
		})
	}
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(todo)
}
func patchTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": bson.M{"completed":true}}
	result, err := collection.UpdateOne(context.Background(), filter,update)
	if err != nil || result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo updated successfully",	"todo_id": id})
}
func deleteTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	filter := bson.M{"_id": ObjectID}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted successfully", "todo_id": id})

}