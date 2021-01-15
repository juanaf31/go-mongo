package usecases

import (
	"GoMongo/master/repositories"
	"GoMongo/models"
)

type PostUsecaseImpl struct{
	postRepo repositories.PostRepo
}

func InitPostUsecase(postRepo repositories.PostRepo) PostUsecase {
	return &PostUsecaseImpl{
		postRepo: postRepo,
	}
}

func (u *PostUsecaseImpl)GetPosts()([]*models.GetPosts,error){
	posts,err := u.postRepo.GetPosts()
	if err != nil {
		return nil,err
	}
	return posts, nil
}