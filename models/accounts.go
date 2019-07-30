package models

import (
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
    util "project.golang/studycrud/TaskListGo/utils"
    jwt "github.com/dgrijalva/jwt-go"
    "strings"
    "os"
)

type Account struct {
    gorm.Model
    Email string `json:"email"`
    Password string `json:"password"`
    Token string `json:"token"`  //skipping this field does not work in mysql
}

func (account *Account) Validate() (map[string]interface{}, bool) {
    if !strings.Contains(account.Email, "@") {
        return util.MetaMsg(false, "Email address format is incorrect"), false
    }

    if len(account.Password) < 6 {
        return util.MetaMsg(false, "Password is minimum 6 character"), false
    }

    temp := &Account{}

    err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return util.MetaMsg(false, "Connection error. Please retry"), false
    }
    if temp.Email != "" {
        return util.MetaMsg(false, "Email address already in use by another user."), false
    }

    return util.MetaMsg(true, "Requirement passed"), true
}

func (account *Account) CreateAccount() (map[string]interface{}) {
    if rsp, status := account.Validate(); !status {
        return rsp
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
    account.Password = string(hashedPassword)

    GetDB().Create(account)

    if account.ID <= 0 {
        return util.MetaMsg(false, "Failed to create account")
    }

    tk := &Token{UserId: account.ID}
    token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
    tokenString, _ := token.SignedString([]byte(os.Getenv("jwt_secret")))

    account.Token = tokenString
    account.Password = ""

    response := util.MetaMsg(true, "Account is successfully created")
    response["account"] = account
    return response
}

func (account *Account) Login() (map[string]interface{}) {
    registeredAccount := &Account{}
    err := GetDB().Table("accounts").Where("email = ?", account.Email).First(registeredAccount).Error

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return util.MetaMsg(false, "Account is not recognized")
        }
        return util.MetaMsg(false, "There is something error")
    }

    err = bcrypt.CompareHashAndPassword([]byte(registeredAccount.Password), []byte(account.Password))

    if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
        return util.MetaMsg(false, "Password is Invalid")
    }

    tk := &Token{UserId: registeredAccount.ID}
    token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
    tokenString, _ := token.SignedString([]byte(os.Getenv("jwt_secret")))

    registeredAccount.Token = tokenString
    registeredAccount.Password = ""

    response := util.MetaMsg(true, "Successfully Login")
    response["account"] = registeredAccount
    return response
}
