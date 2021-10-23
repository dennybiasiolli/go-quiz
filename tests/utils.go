package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/dennybiasiolli/go-quiz/common"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupTests(setupFiberRoutes func(*fiber.App), models ...interface{}) (*fiber.App, *gorm.DB) {
	// reading test envs
	common.GetEnvVariables("../.env.tests", "../.env.tests.default")

	// connecting to testing db
	common.ConnectDb()
	db := common.GetDB()
	for _, model := range models {
		db.AutoMigrate(&model)
	}

	// clear tables
	stmt := &gorm.Statement{DB: db}
	// for _, model := range []interface{}{
	// 	citazioni.Citazione{},
	// } {
	for _, model := range models {
		stmt.Parse(&model)
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY;", stmt.Schema.Table))
	}

	app := fiber.New()
	setupFiberRoutes(app)
	return app, db
}

func GetJsonRequest(method string, url string, jsonBody interface{}) *http.Request {
	body, _ := json.Marshal(jsonBody)
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func GetJsonResponse(app *fiber.App, req *http.Request) (*http.Response, error) {
	resp, err := app.Test(req)
	return resp, err
}

func NewJsonRequest(
	app *fiber.App,
	method string, url string, jsonBody interface{},
) (*http.Response, error) {
	req := GetJsonRequest(method, url, jsonBody)
	return GetJsonResponse(app, req)
}
