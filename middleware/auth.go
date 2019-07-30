package middleware

import (
    "net/http"
    jwt "github.com/dgrijalva/jwt-go"
    "os"
    "context"
    "fmt"
    util "project.golang/studycrud/TaskListGo/utils"
    "strings"
    "project.golang/studycrud/TaskListGo/models"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        noAuthPath := []string{"/api/user/register", "/api/user/login"}
        requestPath := r.URL.Path

        for _, path := range noAuthPath {
            if path == requestPath {
                next.ServeHTTP(w, r)
                return
            }
        }

        response := make(map[string] interface{})
        tokenHeader := r.Header.Get("Authorization")

        if tokenHeader == "" {
            response = util.MetaMsg(false, "Token is not present")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            util.Respond(w, response)
            return
        }

        headerAuthorizationString := strings.Split(tokenHeader, " ")
        if len(headerAuthorizationString) != 2 {
            response = util.MetaMsg(false, "Invalid/Malformed auth token")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            util.Respond(w, response)
            return
        }

        tokenValue := headerAuthorizationString[1]
        tk := &models.Token{}

        token, err := jwt.ParseWithClaims(tokenValue, tk, func(token *jwt.Token) (interface{}, error){
                        return []byte(os.Getenv("jwt_secret")), nil
                      })

        if err != nil {
            response = util.MetaMsg(false, "Malformed authentication token")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            util.Respond(w, response)
            return
        }

        if !token.Valid {
            response = util.MetaMsg(false, "Token is not valid.")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            util.Respond(w, response)
            return
        }

        fmt.Sprintf("User Id is %s", tk.UserId)
        ctx := context.WithValue(r.Context(), "user", tk.UserId)
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
    });
}