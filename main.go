package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dee-d-dev/go-gin/controllers"
	"github.com/dee-d-dev/go-gin/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongo          *mongo.Client
	err            error
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	MONGO_URI := os.GetEnv("MONGO_URL")
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(MONGO_URI)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}

	if err = mongoclient.Ping(ctx, redpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongodb has been established")

	usercollection = mongoclient.Database("go-gin").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()

}
func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")

	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
