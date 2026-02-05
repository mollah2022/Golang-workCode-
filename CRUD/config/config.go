package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host string
	Port int
	Name string
	User string
	Password string
	EnableSSLMODE bool
}


type Config struct{
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecretKey string
	DB *DBConfig
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

	jwtSecretKey:= os.Getenv("JWTSECRETKEY")
	if jwtSecretKey == "" {
		fmt.Println("Service name is Required")
		os.Exit(1)
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		fmt.Println("DbHost Port is Required")
		os.Exit(1)
	}

	dbPOrt := os.Getenv("DB_PORT")
	if dbPOrt == "" {
		fmt.Println("DB_Port Port is Required")
		os.Exit(1)
	}

	dbPort,dberror := strconv.ParseInt(dbPOrt,10,64)
	if dberror != nil{
		fmt.Println("DBPort must be Number")
		os.Exit(1)
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		fmt.Println("DB_Name is Required")
		os.Exit(1)
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		fmt.Println("Db_User  is Required")
		os.Exit(1)
	}

	dbpassword := os.Getenv("DB_PASSWORD")
	if dbpassword == "" {
		fmt.Println("DB_Password  is Required")
		os.Exit(1)
	}

	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")

	enblSSLMode,err := strconv.ParseBool(enableSslMode)
	if err!=nil {
		fmt.Println("Invalid enable ssl mode value",err)
	}

	dbCongif := &DBConfig{
		Host: dbhost,
		Port: int(dbPort),
		Name: dbname,
		User: dbuser,
		Password: dbpassword,
		EnableSSLMODE: enblSSLMode,
	}

	configuration = &Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: int(port),
		JwtSecretKey: jwtSecretKey,
		DB:dbCongif,
	}
}

func GetConfig() *Config{
	// loadConfig()
	if configuration == nil {
		loadConfig()
	}
	return configuration
}