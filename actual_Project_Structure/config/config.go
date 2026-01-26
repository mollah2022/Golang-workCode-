package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

var configuration Config


type Config struct{
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig(){

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load the env variables",err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("Version is Required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("Service name is Required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		fmt.Println("Http Port is Required")
		os.Exit(1)
	}

	port,error := strconv.ParseInt(httpPort,10,64)

	if error != nil{
		fmt.Println("Port must be Number")
		os.Exit(1)
	}

	configuration = Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: int(port),
	}
}

func GetConfig() Config{
	loadConfig()
	return configuration
}