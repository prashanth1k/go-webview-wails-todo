package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"time"

	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	db  gorm.DB
}

type todo struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Due         time.Time `json:"due" gorm:"type:datetime"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created" gorm:"autoCreateTime;type:datetime"`
	Updated     time.Time `json:"updated" gorm:"autoUpdateTime;type:datetime"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database" + err.Error())
	}
	a.db = *db
	log.Println("Connected to database. db exists: " + fmt.Sprintf("%t", db != nil))

	// AutoMigrate will create the todos table if it does not exist
	if !db.Table("todos").Migrator().HasTable("todos") {
		err = db.AutoMigrate(&todo{})
		if err != nil {
			log.Fatal("Failed to migrate database" + err.Error())
		}
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ListTodos() []todo {
	// return []todo{{Id: 1, Description: "Code Go", Due: "2023-10-15", Status: "In Progress", Created: "2023-10-01 18:30"}}
	var todos []todo
	result := a.db.Find(&todos)
	if result.Error != nil {
		panic("Failed to retrieve todos: " + result.Error.Error())
	}
	return todos
}

func (a *App) CreateTodo(todoItem todo) todo {
	log.Println("Create: " + prettyPrint(todoItem))
	// todoItem.Due, _ = time.Parse(time.RFC3339, todoItem.Due.String())

	result := a.db.Create(&todoItem)

	if result.Error != nil {
		log.Println("Failed to create todo item: " + result.Error.Error())
		return todo{}
	}

	log.Println("db result: " + prettyPrint(result))
	return todoItem
}
func (a *App) UpdateTodo(todoItem todo) todo {
	log.Println("Update: " + prettyPrint(todoItem))

	result := a.db.Save(&todoItem)
	if result.Error != nil {
		log.Println("Failed to update todo item: " + result.Error.Error())
		return todo{}
	}

	log.Println("update result: " + prettyPrint(todoItem))
	return todoItem
}
func (a *App) DeleteTodo(todoItem todo) todo {
	log.Println("Delete: " + prettyPrint(todoItem))

	result := a.db.Where("id = ?", todoItem.Id).Delete(&todoItem)
	if result.Error != nil {
		log.Println("Failed to delete todo item: " + result.Error.Error())
		return todo{}
	}
	log.Println("delete result: " + prettyPrint(todoItem))
	return todoItem
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
