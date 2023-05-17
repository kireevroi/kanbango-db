package onstart

import (
	"github.com/joho/godotenv"
)

func LoadEnv(path string) error {
	return godotenv.Load(path)
}