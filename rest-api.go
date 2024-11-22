package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getIpAddress(r *http.Request) string {
	ipaddr := r.Header.Get("X-Real-IP")
	return ipaddr
}

func startRest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/app", getApp)
	router.HandleFunc("/api/user", getUser)
	router.HandleFunc("/api/companionapp", getCompanionApp)
	router.HandleFunc("/api/loginpass", loginPass)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", router)
}

func getApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["app_id"]
	if len(appID) == 0 {
		appID = r.FormValue("app_id")
	}
	appName := vars["app_name"]
	if len(appName) == 0 {
		appName = r.FormValue("app_name")
	}
	a := GlassApp{
		AppID:   appID,
		AppName: appName,
	}
	a.Retrieve()
	resp := APIResponse{
		Code:     200,
		Valid:  true,
		Response: a,
	}
	ret, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error")
	}
	fmt.Fprintf(w, "%+v", string(ret))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	if len(username) == 0 {
		username = r.FormValue("username")
	}
	identifier := vars["identifier"]
	if len(identifier) == 0 {
		identifier = r.FormValue("identifier")
	}
	email := vars["email"]
	if len(email) == 0 {
		email = r.FormValue("email")
	}
	twitter := vars["twitter"]
	if len(twitter) == 0 {
		twitter = r.FormValue("twitter")
	}
	github := vars["github"]
	if len(github) == 0 {
		github = r.FormValue("github")
	}
	keybase := vars["keybase"]
	if len(keybase) == 0 {
		keybase = r.FormValue("keybase")
	}
	discord := vars["discord"]
	if len(discord) == 0 {
		discord = r.FormValue("discord")
	}
	linkedin := vars["linkedin"]
	if len(linkedin) == 0 {
		linkedin = r.FormValue("linkedin")
	}
	website := vars["website"]
	if len(website) == 0 {
		website = r.FormValue("website")
	}
	reddit := vars["reddit"]
	if len(reddit) == 0 {
		reddit = r.FormValue("reddit")
	}

	a := User{
		Username:   username,
		Identifier: identifier,
		Email:      email,
		Twitter:    twitter,
		Github:     github,
		Keybase:    keybase,
		Discord:    discord,
		LinkedIn:   linkedin,
		Reddit:     reddit,
		Website:    website,
	}
	a.Retrieve()
	resp := APIResponse{
		Code:     200,
		Valid:  true,
		Response: a,
	}
	ret, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error")
	}
	fmt.Fprintf(w, "%+v", string(ret))
}

func getCompanionApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["app_id"]
	if len(appID) == 0 {
		appID = r.FormValue("app_id")
	}
	a := CompanionApp{
		GlassAppID: appID,
	}
	a.Retrieve()
	resp := APIResponse{
		Code:     200,
		Valid:  true,
		Response: a,
	}
	ret, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error")
	}
	fmt.Fprintf(w, "%+v", string(ret))
}

func loginPass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipAddr := getIpAddress(r)

	username := vars["username"]
	if len(username) == 0 {
		username = r.FormValue("username")
	}
	password := vars["password"]
	if len(password) == 0 {
		password = r.FormValue("password")
	}
	email := ""
	if strings.Contains(username, "@") {
		email = username
		username = ""
	}
	a := AuthToken{
		Current:  "true",
		Email:    email,
		Username: username,
	}
	success, _ := TryPassword(&a, password, ipAddr)
	if !success {
		resp := APIResponse{
			Code: 401,
			Valid: true,
			Response: "Invalid Login Attempt",
		}
		ret, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("Error")
		}
		fmt.Fprintf(w, "%+v", string(ret))
		return
	}
	// Success! Returning token
	resp := APIResponse{
		Code:     200,
		Valid:  true,
		Response: a,
	}
	ret, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error")
	}
	fmt.Fprintf(w, "%+v", string(ret))
}
