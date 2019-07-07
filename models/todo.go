package models

import (
    "github.com/jinzhu/gorm"
    "fmt"
    util "step4/TaskListGo/utils"
)

type Todo struct {
    gorm.Model
    Name string `json:"name"`
    Priority int `json:"priority"`
    Completed bool `json:"completed"`
}

func GetToDo() ([]*Todo) {
    todos := make([]*Todo, 0)
    err := GetDB().Table("todos").Find(&todos).Error

    if err != nil {
        fmt.Println(err)
        return nil
    }

    return todos
}

func SingleToDo(id uint) (*Todo) {
    todo := &Todo{}
    err := GetDB().Table("todos").Where("id = ?", id).First(todo).Error

    if err != nil {
        fmt.Println(err)
        return nil
    }

    return todo
}

func (todo *Todo) Create() (map[string] interface{}) {
    GetDB().Create(todo)

    if todo.ID <= 0 {
        return util.MetaMsg(false, "Todo is not created")
    }

    response := util.MetaMsg(true, "Created Successfully")
    response["data"] = todo
    return response
}

func (todo *Todo) ActionToDo() (map[string] interface{}){
    actionStatus := todo.Completed
    err := GetDB().Table("todos").Where("id = ?", todo.ID).First(todo).Error

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return util.MetaMsg(false, "Task is not recognized")
        }
        return util.MetaMsg(false, "Error on DB Query")
    }

    GetDB().Model(todo).Update("Completed", actionStatus)

    response := util.MetaMsg(true, "Task updated successfully")
    response["data"] = todo
    return response
}

func (todo *Todo) EditToDo() (map[string] interface{}) {
    editedTodo := &Todo{}
    err := GetDB().Table("todos").Where("id = ?", todo.ID).First(editedTodo).Error

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return util.MetaMsg(false, "Task is not recognized")
        }
        return util.MetaMsg(false, "Error on DB Query")
    }

    editedTodo.Name = todo.Name
    if todo.Priority != 0 {
        editedTodo.Priority = todo.Priority
    }

    GetDB().Save(editedTodo)

    response := util.MetaMsg(true, "Task is edited successfully")
    response["data"] = editedTodo
    return response
}

func (todo *Todo) DeleteToDo() (map[string]interface{}) {
    err := GetDB().Table("todos").Where("id = ?", todo.ID).First(todo).Error

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return util.MetaMsg(false, "Task is not recognized")
        }
        return util.MetaMsg(false, "Error on DB Query")
    }

    GetDB().Delete(todo)

    response := util.MetaMsg(true, "Task is deleted successfully")
    response["data"] = todo
    return response
}

