package main

import (
	"GoMongo/configs"
	"GoMongo/master"
	"log"
)

func main()  {
	db,err,host,port := configs.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	router:=configs.CreatRouter()

	master.InitData(router,db)
	configs.RunServer(router,host,port)
}