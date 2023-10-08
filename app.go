package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// App struct
type App struct {
	ctx context.Context
}

type todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Due         string `json:"due"`
	Status      string `json:"status"`
	Created     string `json:"created"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ListTodos() []todo {
	return []todo{{Id: 1, Description: "Code Go", Due: "2023-10-15", Status: "In Progress", Created: "2023-10-01 18:30"}}
}
func (a *App) CreateTodo(todoItem todo) todo {
	log.Println("Create: " + prettyPrint(todoItem))
	return todo{Id: 1, Description: "Code Go", Due: "2023-10-15", Status: "In Progress", Created: "2023-10-01 18:30"}
}
func (a *App) UpdateTodo(todoItem todo) todo {
	log.Println("Update: " + prettyPrint(todoItem))
	return todo{Id: 1, Description: "Code Go", Due: "2023-10-15", Status: "In Progress", Created: "2023-10-01 18:30"}
}
func (a *App) DeleteTodo(todoItem todo) todo {
	log.Println("Delete: " + prettyPrint(todoItem))
	return todo{Id: 1, Description: "Code Go", Due: "2023-10-15", Status: "In Progress", Created: "2023-10-01 18:30"}
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
