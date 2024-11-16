package main

import (
	"chitchatv3/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /login
// Show the login page
func login(c *gin.Context) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	t.Execute(c.Writer, nil)
}

// GET /signup
// Show the signup page
func signup(c *gin.Context) {
	generateHTML(c.Writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	c.Redirect(http.StatusFound, "/login")
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(c *gin.Context) {
	err := c.Request.ParseForm()
	user, err := data.UserByEmail(c.PostForm("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(c.PostForm("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, &cookie)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

// GET /logout
// Logs the user out
func logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	c.Redirect(http.StatusFound, "/")
}
