package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func DiscussionViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetAllDiscussions()); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}
func DiscussionViewByProspectId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	prospectid := vars["id"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetDiscussionByProspectId(prospectid)); err != nil {
		fmt.Println("Err")
		panic(err)
	}
}

func DiscussionQuestionAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var discussion Discussion
	err = json.Unmarshal(body, &discussion)
	if err != nil {
		panic(err)
	}
	err = discussion.Write()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Notify(NPQuestionPosted, discussion)
}

func DiscussionUpdateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var t Discussion
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	err = t.Update()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DiscussionAnswerAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var discussion Discussion
	err = json.Unmarshal(body, &discussion)
	if err != nil {
		panic(err)
	}
	err = discussion.AddAnswer()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Notify(NPQuestionAnswered, discussion)
}
