package controllers

import (
    "fmt"
    util "project.golang/studycrud/TaskListGo/utils"
    "project.golang/studycrud/TaskListGo/models"
    "net/http"
    "encoding/json"
)

var Register = func(w http.ResponseWriter, r *http.Request) {
    account := &models.Account{}

    err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
    if err != nil {
        fmt.Println(err)
        util.Respond(w, util.MetaMsg(false, "Invalid request"))
        return
    }

    response := account.CreateAccount()
    util.Respond(w, response)
}

var Login = func(w http.ResponseWriter, r *http.Request) {
    account := &models.Account{}

    err := json.NewDecoder(r.Body).Decode(account)
    if err != nil {
        fmt.Println(err)
        util.Respond(w, util.MetaMsg(false, "Invalid request"))
        return
    }

    response := account.Login()
    util.Respond(w, response)
}