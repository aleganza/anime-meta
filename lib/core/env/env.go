package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
}

func GetVar(name string) (string, error) {

	value, exists := os.LookupEnv(name)

	if !exists {
		return "", fmt.Errorf(`Environment variable "%s" not found`, name)
	}

	return value, nil
}
