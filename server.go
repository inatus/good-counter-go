package main

import (
	"encoding/json"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"os"
)

const (
	MONGO_DB_NAME = "app21819484"
)

var (
	mgoSession *mgo.Session
)

type count struct {
	Count int `json:"count" bson:"count"`
}

func main() {
	sess, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	mgoSession = sess

	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/count", countHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var cnt count
		c := mgoSession.DB(MONGO_DB_NAME).C("count")
		err := c.Find(bson.M{}).One(&cnt)
		if err != nil {
			log.Println(err)
		}
		response, err := json.Marshal(cnt)
		if err != nil {
			log.Println(err)
		}
		w.Write(response)
	}
	if r.Method == "POST" {
		var cnt count
		c := mgoSession.DB(MONGO_DB_NAME).C("count")
		change := mgo.Change{
			Update:    bson.M{"$inc": bson.M{"count": 1}},
			ReturnNew: true,
		}
		info, err := c.Find(bson.M{}).Apply(change, &cnt)
		if err != nil {
			log.Println(err)
			log.Println(info)
		}
		response, err := json.Marshal(cnt)
		if err != nil {
			log.Println(err)
		}
		w.Write(response)
	}
}
