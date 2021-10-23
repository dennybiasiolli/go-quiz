# Contributing

## Checking for outdated dependencies

Launch `go list -m -u all` for all dependencies, or for project's dependencies:

```sh
go list -m -u github.com/go-playground/validator/v10
go list -m -u github.com/gofiber/fiber/v2
go list -m -u github.com/gofiber/jwt/v3
go list -m -u github.com/golang-jwt/jwt/v4
go list -m -u github.com/joho/godotenv
go list -m -u golang.org/x/oauth2
go list -m -u gorm.io/driver/postgres
go list -m -u gorm.io/gorm
```

## Updating dependency

`go get -u <dependency_name>`, for example

```sh
go get -u gorm.io/gorm
```


## Testing

`go test -v ./tests`
