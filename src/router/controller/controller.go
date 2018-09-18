package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"

	"gitlab.com/dpcat237/flisy/src/module/user"
	"gitlab.com/dpcat237/flisy/src/service"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Collector struct {
	StCnt *seatController
	UsCnt *userController
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func Init(sCll service.Collector) Collector {
	return Collector{
		StCnt: newSeatController(sCll.FlHnd, sCll.StHnd),
		UsCnt: newUserController(),
	}
}

func createError(s string) (string, error) {
	bytes, err := json.Marshal(ErrorMsg{Message: s})
	return string(bytes[:]), err
}

func getUser(r *http.Request) (user.User, error) {
	var u user.User
	decoded := context.Get(r, "decoded")
	return u, mapstructure.Decode(decoded.(jwt.MapClaims), &u)
}

func getVariable(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}

func returnErrorWithStatus(w http.ResponseWriter, msg string, s int) {
	msg, err := createError(msg)
	if err != nil {
		returnServerFailed(w, "")
	}
	http.Error(w, msg, s)
}

func returnJson(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
}

func returnNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func returnUnauthorized(w http.ResponseWriter, s string) {
	msg, err := createError(s)
	if err != nil {
		returnServerFailed(w, "")
	}
	http.Error(w, msg, http.StatusUnauthorized)
}

func returnPreconditionFailed(w http.ResponseWriter, s string) {
	msg, err := createError(s)
	if err != nil {
		returnServerFailed(w, "")
	}
	http.Error(w, msg, http.StatusPreconditionFailed)
}

func returnServerFailed(w http.ResponseWriter, s string) {
	defaultMsg := "Internal server error"
	if len(s) < 1 {
		msg, err := createError(defaultMsg)
		if err != nil {
			http.Error(w, defaultMsg, http.StatusInternalServerError)
		}
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	msg, err := createError(s)
	if err != nil {
		http.Error(w, defaultMsg, http.StatusInternalServerError)
	}
	http.Error(w, msg, http.StatusInternalServerError)
}
