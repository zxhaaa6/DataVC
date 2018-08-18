package client

import "github.com/zxhaaa6/DataVC/module/mongo"

func HandleConnection(options ConnectOptionType) error {
	if options.dbType == "mongo" {
		_, err:=mongo.ConnectMongo(options)
		if err!=nil {
			return err
		}
	}
	return nil
}
