package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}
//NewDB returns pointer to new Database
func NewDB() *DB {
	return &DB{}
}

func (db *DB) Connect(URL string) error {
	gormdb, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	db.DB = gormdb
	if err == nil {
		log.Println("Connected to Postgres Database")
	}
	return err
}

func (db *DB) CreateUser(u User) error {
	x := db.Create(&u)
	return x.Error
}

func (db *DB) GetUser(username string) (User, error) {
	var u User
	x := db.Where("username = ?", username).First(&u)
	return u, x.Error
}

func (db *DB) GetUserIdByUuid(uuid string) (uint, error) {
	var s Session
	x := db.Where("session = ?", uuid).First(&s)
	return s.User_ID, x.Error
}

func (db *DB) SetSession(s Session) error {
	x := db.Create(&s)
	return x.Error
}

func (db *DB) DeleteSession(s Session) error {
	x := db.Where("session = ?", s.Session).Delete(&Session{})
	return x.Error
}

func (db *DB) DeleteUser(User_ID uint) error {
	x := db.Where("id = ?", User_ID).Delete(&User{})
	return x.Error
}

func (db *DB) GetAllSessions(User_ID uint) ([]Session, error) {
	var sessions []Session
	x := db.Where("user_id = ?", User_ID).Find(&sessions)
	return sessions, x.Error
}

func (db *DB) DeleteAllSessions(User_ID uint) error {
	x := db.Where("user_id = ?", User_ID).Delete(&Session{})
	return x.Error
}


