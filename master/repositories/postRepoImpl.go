package repositories

import (
	"GoMongo/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepoImpl struct{
	db *mongo.Client
}

func InitAccountRepoImpl(db *mongo.Client) PostRepo {
	return &PostRepoImpl{db: db}
}

func (rp *PostRepoImpl)GetPosts()([]*models.GetPosts,error){
	var posts []*models.GetPosts

	collection := rp.db.Database("juan").Collection("posts")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := collection.Find(ctx,bson.M{})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx){
		var post models.GetPosts
		cursor.Decode(&post)
		posts = append(posts,&post)
	}
	if err := cursor.Err();err!=nil{
		log.Println(err)
		return nil,err
	}
	
	return posts,nil
}