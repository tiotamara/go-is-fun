package controllers

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "step4/TaskListGo/models"
    util "step4/TaskListGo/utils"
    "strconv"
    "encoding/json"
)

var Hello = func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello World from controller")
    response := util.MetaMsg(true, "Hello World")
    util.Respond(w, response)
}

var ListToDos = func(w http.ResponseWriter, r *http.Request) {
    // params := mux.Vars(r)

    data := models.GetToDo()

    response := util.MetaMsg(true, "Success")
    response["data"] = data
    util.Respond(w, response)
}

var ListToDo = func(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    toDoId, err := strconv.Atoi(params["toDoId"])

    if err != nil {
        util.Respond(w, util.MetaMsg(false, "Params is invalid"))
        return
    }

    data := models.SingleToDo(uint(toDoId))

    response := util.MetaMsg(true, "Success")
    response["data"] = data
    util.Respond(w, response)
}

var CreateToDo = func(w http.ResponseWriter, r *http.Request) {
    todo := &models.Todo{}

    err := json.NewDecoder(r.Body).Decode(todo) //decode the request body into struct and failed if any error occur
    if err != nil {
        fmt.Println(err)
        util.Respond(w, util.MetaMsg(false, "Invalid request"))
        return
    }

    response := todo.Create()
    util.Respond(w, response)
}

var ActionToDo = func (w http.ResponseWriter, r *http.Request) {
    todo := &models.Todo{}

    params := mux.Vars(r)
    toDoId, err := strconv.Atoi(params["toDoId"])

    if err != nil {
        util.MetaMsg(false, "Param is invalid")
        return
    }

    todo.ID = uint(toDoId)

    err = json.NewDecoder(r.Body).Decode(todo) //decode the request body into struct and failed if any error occur
    if err != nil {
        fmt.Println(err)
        util.Respond(w, util.MetaMsg(false, "Invalid request"))
        return
    }

    response := todo.ActionToDo()
    util.Respond(w, response)
}

var EditToDo = func (w http.ResponseWriter, r *http.Request) {
    todo := &models.Todo{}

    params := mux.Vars(r)

    toDoId, err := strconv.Atoi(params["toDoId"])

    if err != nil {
        util.MetaMsg(false, "Param is invalid")
        return
    }

    todo.ID = uint(toDoId)

    err = json.NewDecoder(r.Body).Decode(todo) //decode the request body into struct and failed if any error occur
    if err != nil {
        fmt.Println(err)
        util.Respond(w, util.MetaMsg(false, "Invalid request"))
        return
    }

    response := todo.EditToDo()
    util.Respond(w, response)
}

var DeleteToDo = func (w http.ResponseWriter, r *http.Request) {
    todo := &models.Todo{}
    params := mux.Vars(r)

    toDoId, err := strconv.Atoi(params["toDoId"])

    if err != nil {
        util.MetaMsg(false, "Param is invalid")
        return
    }

    todo.ID = uint(toDoId)

    response := todo.DeleteToDo()
    util.Respond(w, response)
}