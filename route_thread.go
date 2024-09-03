package main

import (
	"chitchatv2/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /threads/new
// Show the new thread form page
func newThread(c *gin.Context) {
	_, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		generateHTML(c.Writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /signup
// Create the user account
func createThread(c *gin.Context) {
	sess, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		err = c.Request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := c.PostForm("topic")
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		c.Redirect(http.StatusFound, "/")
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(c *gin.Context) {
	uuid := c.Query("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		error_message(c, "Cannot read thread")
	} else {
		_, err := session(c)
		if err != nil {
			generateHTML(c.Writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			generateHTML(c.Writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func postThread(c *gin.Context) {
	sess, err := session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		err = c.Request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		body := c.PostForm("body")
		uuid := c.PostForm("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			error_message(c, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		c.Redirect(http.StatusFound, url)
	}
}
