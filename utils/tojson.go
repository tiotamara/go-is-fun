package utils

import (
    "encoding/json"
    "net/http"
)

func MetaMsg(stat bool, msg string) (map[string]interface{}) {
    return map[string]interface{} {"status": stat, "message": msg}
}

func Respond(w http.ResponseWriter, rsp map[string]interface{}){
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(rsp)
}