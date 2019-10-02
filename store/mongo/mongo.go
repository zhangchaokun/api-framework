package mongo


//var MongoDB *mgo.Database

/*
 * mgo文档 http://labix.org/mgo
 * https://godoc.org/gopkg.in/mgo.v2
 * https://godoc.org/gopkg.in/mgo.v2/bson
 * https://godoc.org/gopkg.in/mgo.v2/txn
 */
/*
func InitMongo() {
	if config.MongoConfig.URL == "" {
		return
	}
	session, err := mgo.Dial(config.MongoConfig.URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	MongoDB = session.DB(config.MongoConfig.Database)
}
*/