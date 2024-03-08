package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
    // load .env file
    env_error := godotenv.Load(".env")
    if env_error != nil {
        fmt.Print("Error loading .env file")
    }
        // Return the value of the variable
    return os.Getenv(key)
}
