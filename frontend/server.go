package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type handler struct {
	Template map[string]*template.Template
	ApiUrl   string
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	apiPath := fmt.Sprintf("%s/users", h.ApiUrl)

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(apiPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := []*User{}
	if err = json.Unmarshal(body, &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Template["user_list.html"].Execute(w, users)
}

func (h *handler) userDetail(w http.ResponseWriter, r *http.Request) {
	apiPath := fmt.Sprintf("%s/users/%s", h.ApiUrl, mux.Vars(r)["id"])

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(apiPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Stop here if the response code is not OK
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Invalid ID or user not found", http.StatusNotFound)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := User{}
	if err = json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Template["user_detail.html"].Execute(w, user)
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	// Check if environment variables are set, otherwise use a default value
	portEnv := getEnv("SERVER_PORT", "8080")
	apiUrl := getEnv("API_URL", "http://127.0.0.1:8000")

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%d", port)

	// Parse templates or panic if they cannot be found
	tmpl := make(map[string]*template.Template)
	tmpl["user_list.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/user_list.html"))
	tmpl["user_detail.html"] = template.Must(template.ParseFiles("templates/base.html", "templates/user_detail.html"))

	// Initialize handler
	h := handler{
		Template: tmpl,
		ApiUrl:   apiUrl,
	}

	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Routes
	r.HandleFunc("/", h.index).Methods("GET")
	r.HandleFunc("/user/{id}", h.userDetail).Methods("GET")

	log.Printf("Listening on %s\n", address)
	log.Fatal(srv.ListenAndServe())
}
