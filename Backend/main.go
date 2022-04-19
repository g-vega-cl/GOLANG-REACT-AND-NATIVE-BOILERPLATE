package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"github.com/g-vega-cl/rezumebackend/controllers"
	"github.com/rs/cors"
	mgo "gopkg.in/mgo.v2"
)

// How to deploy
// https://dev.to/heavykenny/how-to-deploy-a-golang-app-to-heroku-5g1j

func main() {
	fmt.Println("Running main()")
	uc := controllers.NewUserController(getSession())
	mux := http.NewServeMux()

	mux.HandleFunc("/", uc.Home)

	// port := os.Getenv("PORT") //  heroku config:set PORT=8080  <- env
	port := "8000"
	handler := cors.Default().Handler(mux)
	fmt.Println("run ", port)
	http.ListenAndServe(":"+port, handler)
	fmt.Println("Listening and serving in port: ", port)
}

func createConnection() (*mgo.Session, error) {
	dialInfo := mgo.DialInfo{
		Addrs: []string{
			"cluster0-shard-00-00.5r4op.mongodb.net:27017",
			"cluster0-shard-00-01.5r4op.mongodb.net:27017",
			"cluster0-shard-00-02.5r4op.mongodb.net:27017"},
		Username: "default",
		Password: "default",
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	return mgo.DialWithInfo(&dialInfo)
}

func getSession() *mgo.Session {
	s, err := createConnection()

	if err != nil {
		fmt.Println("Error connecting to mongo")
		panic(err)
	}
	return s
}
