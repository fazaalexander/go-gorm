package main

import (
	"github.com/fazaalexander/go-gorm/config"
	m "github.com/fazaalexander/go-gorm/middleware"
	"github.com/fazaalexander/go-gorm/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
