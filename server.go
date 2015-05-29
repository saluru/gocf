package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "github.com/gocf/julienschmidt/httprouter"
    "github.com/gocf/saluru/gorest-cf"
  )

func main(){

const (
	DEFAULT_PORT = "4001"
)
 

var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		fmt.Printf("Warning, VCAP_APP_PORT not set. Defaulting to %+v\n", DEFAULT_PORT)
		port = DEFAULT_PORT
	}
 


  r := httprouter.New()
 // port := os.Getenv("PORT")

  r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request,  p httprouter.Params){
    u := gorestcf.User {
            Name:   "Bob Smith",
            Gender: "male",
            Age:    50,
            Id:     p.ByName("id"),
    }

    uj, _ := json.Marshal(u)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", uj)
    })

    http.ListenAndServe(":" + port, r)
}
