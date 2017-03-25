package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"time"

	"github.com/gorilla/mux"
	"github.com/timkellogg/todo/api/models"
)

// TodoIndex : shows all todos
func TodoIndex(w http.ResponseWriter, h *http.Request) {
	SetHeaders(w)

	todos := models.FindAllTodos()

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow : shows one todo
func TodoShow(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	var (
		err error
		id  int
	)
	if id, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}

	todo := models.FindOneTodo(id)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

// TodoCreate : creates one todo from json post
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	err := r.ParseForm()
	HandleError(err)

	name := r.PostFormValue("name")

	completed, err := strconv.ParseBool(r.PostFormValue("completed"))
	HandleError(err)

	due, err := time.Parse("2006-01-02 00:00:00", r.PostFormValue("due"))
	HandleError(err)

	todo := models.CreateTodo(name, completed, due)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

// TodoUpdate : updates a todo's name, due date, or mark as completed
func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)

	err := r.ParseForm()
	HandleError(err)

	vars := mux.Vars(r)
	var id int

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")

	completed, err := strconv.ParseBool(r.PostFormValue("completed"))
	HandleError(err)

	due, err := time.Parse("2006-01-02 00:00:00", r.PostFormValue("due"))
	HandleError(err)

	todo := models.Todo{ID: id, Name: name, Completed: completed, Due: due}
	updatedTodo := models.UpdateTodo(todo)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedTodo); err != nil {
		panic(err)
	}
}

// TodoDestroy : removes a todo by the given id
func TodoDestroy(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	var (
		err error
		id  int
	)

	if id, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}

	models.DestroyOneTodo(id)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{}); err != nil {
		panic(err)
	}
}
