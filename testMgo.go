package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Avmovie struct {
	Author string
	Tag string
	Url string
	Title string
}

var (
	session *mgo.Session
	err error
)

func init() {
	session, err = mgo.Dial("localhost")
	if err != nil{
		fmt.Errorf("init error: %v ", err)
	}
}

func main() {
	adddata()
	findall()
	defer session.Clone()
}

func findall()  {
	c := session.DB("av").C("movies")
	var movies []Avmovie
	c.Find(bson.M{}).All(&movies)

	for _, m := range movies{
		fmt.Println(m.Title)
	}
	if len(movies) <=0 {
		fmt.Println("there is no data")
	}
}

func adddata()  {
	c := session.DB("av").C("movies")
	movie := Avmovie{"rion0", "tag0", "http0", "shome0"}
	movie2 := Avmovie{"rion2", "tag2", "http2", "shome2"}
	arr := []Avmovie{movie, movie2}
	err := c.Insert(arr)
	if err!= nil {
		fmt.Println(err)
	}
}


