package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dennybiasiolli/go-quiz/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

func getSignedClaimsFromUser(user *User) (string, string, error) {
	// Set claims
	claims := JwtCustomClaims{
		TokenType: "access",
		UserId:    user.ID,
		UserInfo: JwtUserInfo{
			Email:      user.Email,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			FullName:   strings.TrimSpace(user.FirstName + " " + user.LastName),
			PictureUrl: user.PictureUrl,
			Locale:     user.Locale,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(common.JWT_ACCESS_TOKEN_LIFETIME_MINUTES))),
			ID:        fmt.Sprintf("%v", time.Now().UnixMilli()),
		},
	}

	// Create token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// change to refresh token
	claims.TokenType = "refresh"
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(common.JWT_REFRESH_TOKEN_LIFETIME_MINUTES)))

	// Create token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded tokens and send it as response.
	a, err := accessToken.SignedString([]byte(common.JWT_HMAC_SAMPLE_SECRET))
	if err != nil {
		return "", "", err
	}
	r, err := refreshToken.SignedString([]byte(common.JWT_HMAC_SAMPLE_SECRET))
	if err != nil {
		return "", "", err
	}
	return a, r, nil
}

func TokenRefresh(c *fiber.Ctx) error {
	input := new(TokenRefreshInput)
	c.BodyParser(input)
	if err := validator.New().Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := jwt.Parse(input.Refresh, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common.JWT_HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claimss, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid && claimss["token_type"] == "refresh") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Refresh token is not valid",
		})
	}

	db := common.GetDB()
	var user User = User{
		IsActive: true,
	}
	if err := db.Where(&user).First(&user, claimss["user_id"]).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	access, refresh, err := getSignedClaimsFromUser(&user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"access":  access,
		"refresh": refresh,
	})
}

func GoogleOauth2Login(c *fiber.Ctx) error {
	redirect_uri := c.Query("redirect_uri")
	conf := getGoogleOauth2Config()
	if redirect_uri != "" {
		conf.RedirectURL = redirect_uri
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func GoogleOauth2Callback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Redirect("/oauth2/google?code=" + url.QueryEscape(code))
}

func GoogleOauth2(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	redirect_uri := c.Query("redirect_uri")
	conf := getGoogleOauth2Config()
	if redirect_uri != "" {
		conf.RedirectURL = redirect_uri
	}
	tok, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	response, err := http.Get(
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var userInfo GoogleUserInfo
	json.Unmarshal(contents, &userInfo)

	// check existing user
	db := common.GetDB()
	var user User = User{
		Email: userInfo.Email,
	}

	// `Assign` attributes to the record regardless it is found or not and save them back to the database.
	err = db.Attrs(User{
		FirstName:  userInfo.GivenName,
		LastName:   userInfo.FamilyName,
		PictureUrl: userInfo.Picture,
		Locale:     userInfo.Locale,
	}).FirstOrCreate(&user).Error
	if err != nil || user.IsActive == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	access, refresh, err := getSignedClaimsFromUser(&user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"access":  access,
		"refresh": refresh,
	})
}
