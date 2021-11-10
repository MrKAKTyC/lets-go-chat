package serv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MrKAKTyC/lets-go-chat/pkg/services"

	"github.com/MrKAKTyC/lets-go-chat/client/auth"
	"github.com/gorilla/mux"
)

func Serve(port string) {
	fmt.Println("Running on port:", port)
	router := mux.NewRouter()
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/login", getUser).Methods("POST")

	http.ListenAndServe(":"+port, router)
}

func createUser(w http.ResponseWriter, req *http.Request) {
	login, password := req.FormValue("userName"), req.FormValue("password")
	resp, err := services.RegisterUser(*auth.NewUserRequest(login, password))
	if err != nil {
		sendError(w, err)
		return
	}
	jsonResponse, _ := json.Marshal(resp)
	sendJSONResponse(w, jsonResponse)
}

func getUser(w http.ResponseWriter, req *http.Request) {
	login, password := req.FormValue("userName"), req.FormValue("password")
	resp, err := services.AuthorizeUser(auth.New(login, password))
	if err != nil {
		sendError(w, err)
		return
	}
	jsonResponse, _ := json.Marshal(resp)
	w.Header().Set("X-Rate-Limit", "120")
	w.Header().Set("X-Expires-After", time.Now().AddDate(0, 0, 1).UTC().String())
	sendJSONResponse(w, jsonResponse)
}

func sendError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}

func sendJSONResponse(w http.ResponseWriter, jsonResponse []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
