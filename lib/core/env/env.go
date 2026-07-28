package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func GetVar(name string) (string, error) {
	value, exists := os.LookupEnv(name)

	if !exists {
		return "", fmt.Errorf(`Environment variable "%s" not found`, name)
	}

	return value, nil
}
