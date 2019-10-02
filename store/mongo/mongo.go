package mongo

import (
	"api-framework/core/config"
	"github.com/globalsign/mgo"
	"fmt"
	"os"
)

var MongoDB *mgo.Database


func InitMongo() {
	if config.MongoConfig.URL == "" {
		return
	}
	session, err := mgo.Dial(config.MongoConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	session.SetMode(mgo.Monotonic, true)
	MongoDB = session.DB(config.MongoConfig.Database)
}
