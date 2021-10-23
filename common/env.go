package common

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var HTTP_LISTEN string

var PG_HOST string
var PG_PORT string
var PG_USER string
var PG_PASSWORD string
var PG_SSLMODE string
var PG_DATABASE string

var JWT_HMAC_SAMPLE_SECRET string
var JWT_ACCESS_TOKEN_LIFETIME_MINUTES int = 5
var JWT_REFRESH_TOKEN_LIFETIME_MINUTES int = 1440

var GOOGLE_OAUTH2_CLIENT_ID string
var GOOGLE_OAUTH2_CLIENT_SECRET string
var GOOGLE_OAUTH2_DEFAULT_REDIRECT_URL string

func GetEnvVariables(mainFile string, fallbackFile string) {
	/*
		loading .env files in this order, if a variable is not set in `mainFile`,
		it's read from `fallbackFile`
	*/
	errEnv := godotenv.Load(mainFile)
	errEnvDefault := godotenv.Load(fallbackFile)
	if errEnvDefault != nil && errEnv != nil {
		log.Fatalf("Error loading %s or %s file", mainFile, fallbackFile)
	}

	HTTP_LISTEN = os.Getenv("HTTP_LISTEN")

	PG_HOST = os.Getenv("PG_HOST")
	PG_PORT = os.Getenv("PG_PORT")
	PG_USER = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	PG_SSLMODE = os.Getenv("PG_SSLMODE")
	PG_DATABASE = os.Getenv("PG_DATABASE")

	JWT_HMAC_SAMPLE_SECRET = os.Getenv("JWT_HMAC_SAMPLE_SECRET")
	if val, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_LIFETIME_MINUTES")); err == nil {
		JWT_ACCESS_TOKEN_LIFETIME_MINUTES = val
	}
	if val, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_LIFETIME_MINUTES")); err == nil {
		JWT_REFRESH_TOKEN_LIFETIME_MINUTES = val
	}

	GOOGLE_OAUTH2_CLIENT_ID = os.Getenv("GOOGLE_OAUTH2_CLIENT_ID")
	GOOGLE_OAUTH2_CLIENT_SECRET = os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET")
	GOOGLE_OAUTH2_DEFAULT_REDIRECT_URL = os.Getenv("GOOGLE_OAUTH2_DEFAULT_REDIRECT_URL")
}
