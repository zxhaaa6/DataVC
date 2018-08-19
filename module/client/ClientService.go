package client

import (
	"github.com/zxhaaa6/DataVC/module/common_type"
	"github.com/zxhaaa6/DataVC/module/mongo"
)

func HandleConnection(options common_type.ConnectOptionType) error {
	if options.DbType == "mongo" {
		_, err:=mongo.ConnectMongo(options)
		if err!=nil {
			return err
		}
	}
	return nil
}
