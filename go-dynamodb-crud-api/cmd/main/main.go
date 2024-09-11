package main

import "github.com/datarohit/go-dynamodb-crud-api/utils/logger"

func main(){
	logger.InitializeLogger()
	log := logger.GetLogger()
}
