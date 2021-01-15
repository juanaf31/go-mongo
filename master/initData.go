package master

import (
	"GoMongo/master/controllers"
	"GoMongo/master/repositories"
	"GoMongo/master/usecases"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitData(r *mux.Router, db *mongo.Client){
	postsRepo := repositories.InitAccountRepoImpl(db)
	postUsecase := usecases.InitPostUsecase(postsRepo)
	controllers.PostController(r,postUsecase)
}