package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g-vega-cl/rezumebackend/models"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) Home(w http.ResponseWriter, r *http.Request) {
	u := models.TestModel{Message: "data"}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error marshall", err)
	}
	// w.Header().Set("content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "%s\n", uj)
	w.Write(uj)
}
