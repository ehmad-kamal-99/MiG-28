package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const userContext = "USER"

// UserInfo - struct holding user information.
type UserInfo struct {
	Type           string `json:"type"`
	Status         string `json:"status"`
	FirebaseUserID string `json:"user_id" validate:"required"`
	GraphUserID    string `json:"graph_id" validate:"required"`
	DocumentUserID string `json:"document_id" validate:"required"`
}

// GetUser - get user info.
func GetUser(c *gin.Context) UserInfo {
	return c.MustGet(userContext).(UserInfo)
}

// FirebaseAuth - auth struct for firebase.
type FirebaseAuth struct {
	auth      *auth.Client
	validator *validator.Validate
}

// FirebaseConfig - struct holding values for firebase app.
type FirebaseConfig struct {
	AuthOverride     *map[string]interface{}
	DatabaseURL      string
	ProjectID        string
	ServiceAccountID string
	StorageBucket    string
}

// New - sets up new firebase app.
func New(cfg *FirebaseConfig) *FirebaseAuth {
	var err error

	var app *firebase.App
	if cfg != nil {
		app, err = firebase.NewApp(context.Background(), &firebase.Config{
			AuthOverride:     cfg.AuthOverride,
			DatabaseURL:      cfg.DatabaseURL,
			ProjectID:        cfg.ProjectID,
			ServiceAccountID: cfg.ServiceAccountID,
			StorageBucket:    cfg.StorageBucket,
		})
		if err != nil {
			log.Print(err)
		}
	} else {
		app, err = firebase.NewApp(context.Background(), nil)
		if err != nil {
			log.Print(err)
		}
	}

	client := &FirebaseAuth{
		validator: validator.New(),
	}

	client.auth, err = app.Auth(context.Background())
	if err != nil {
		log.Print(err)
	}

	return client
}

// Check - authenticate incoming request.
func (a *FirebaseAuth) Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			log.Print("Authorization header is missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Authorization header is missing",
			})

			return
		}

		idToken, err := a.auth.VerifyIDToken(context.Background(), strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})

			return
		}

		var user UserInfo

		claims, _ := json.Marshal(idToken.Claims)
		if err := json.Unmarshal(claims, &user); err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": err.Error(),
			})

			return
		}

		c.Set(userContext, user)
		c.Next()
	}
}
