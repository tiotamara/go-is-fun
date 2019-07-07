package main

import (
	"workshop/TaskListGo/controllers"
	"workshop/TaskListGo/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
        log.Fatal("$PORT must be set")
    }

	router := mux.NewRouter()

	router.Use(middleware.JwtAuthentication)

	router.HandleFunc("/", controllers.Hello).Methods("GET")
	router.HandleFunc("/api/todos", controllers.ListToDos).Methods("GET")
	router.HandleFunc("/api/todos/{toDoId}", controllers.ListToDo).Methods("GET")
	router.HandleFunc("/api/todos", controllers.CreateToDo).Methods("POST")
	router.HandleFunc("/api/todos/{toDoId}/action", controllers.ActionToDo).Methods("POST")
	router.HandleFunc("/api/todos/{toDoId}/edit", controllers.EditToDo).Methods("POST")
	router.HandleFunc("/api/todos/{toDoId}/delete", controllers.DeleteToDo).Methods("POST")

	router.HandleFunc("/api/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Login).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
