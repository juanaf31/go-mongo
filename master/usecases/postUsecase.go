package usecases

import "GoMongo/models"

type PostUsecase interface{
	GetPosts()([]*models.GetPosts,error)
}