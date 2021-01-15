package repositories

import "GoMongo/models"

type PostRepo interface{
	GetPosts()([]*models.GetPosts,error)
}