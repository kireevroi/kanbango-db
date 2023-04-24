package main

import (
	"os"


	"github.com/kireevroi/kanbango/auth/internal/db"
	"github.com/kireevroi/kanbango/auth/pkg/onstart"
)

func main() {
	onstart.LoadEnv()
	d := db.NewDB()
	d.Connect(os.Getenv("DBURL"))
	d.AutoMigrate(&db.User{})
	d.AutoMigrate(&db.Session{})
}