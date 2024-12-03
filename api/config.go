package api

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	//	println(godotenv.Unmarshal("env"))
	println(godotenv.Load(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	//	println(godotenv.Unmarshal("env"))
	println(godotenv.Load())
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("AUTH_TOKEN")
}

func envSERVICESID() string {
	err :=  godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("SERVICE_SID")
}