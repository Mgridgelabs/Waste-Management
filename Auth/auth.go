package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/nedpals/supabase-go"
	// "log"/
	"context"
)

type AuthService struct {
	client *supabase.Client
}

// NewAuthService creates and returns an AuthService
func NewAuthService(client *supabase.Client) *AuthService {
	return &AuthService{client: client}
}

// Signup handles user registration via email/password
func (a *AuthService) Signup(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Pass context.Background() to SignUp
	user, err := a.client.Auth.SignUp(context.Background(), supabase.UserCredentials{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}


// Login handles user login via email/password
func (a *AuthService) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Pass context.Background() and UserCredentials to SignIn
	session, err := a.client.Auth.SignIn(context.Background(), supabase.UserCredentials{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User logged in successfully",
		"session": session,
	})
}

