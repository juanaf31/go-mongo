package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetEnv(key, defaultValue string) string {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}

func InitDB()(*mongo.Client,error,string,string){
	var client *mongo.Client

	MongoURI := GetEnv("MONGOURI","yourmongouri")
	host := GetEnv("HOST","yourhost")
	port := GetEnv("PORT","yourport")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client,_ = mongo.Connect(ctx,options.Client().ApplyURI(MongoURI))

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println(err)
		return nil,err,host,port
	}
	fmt.Println("database connected")
	return client,err,host,port
}

