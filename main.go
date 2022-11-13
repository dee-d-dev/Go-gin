package main

import (
	"context"

	"github.com/dee-d-dev/go-gin/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo" 
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

func init(){
	ctx = context.TODO()

	mongoconn := optiosns.Client().ApplyURI()
}
func main() {

}
