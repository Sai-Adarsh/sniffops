package bucket

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
func bucket() {
	LoadEnv()
	awsAccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	fmt.Println("My access key ID is ", awsAccessKeyID)
}
