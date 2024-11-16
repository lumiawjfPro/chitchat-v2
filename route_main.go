package main

import (
	"chitchatv3/data"

	"github.com/gin-gonic/gin"
)

// GET /err?msg=
// shows the error message page
func err(c *gin.Context) {
	msg := c.Query("msg")
	_, err := session(c)
	if err != nil {
		generateHTML(c.Writer, msg, "layout", "public.navbar", "error")
	} else {
		generateHTML(c.Writer, msg, "layout", "private.navbar", "error")
	}
}

func index(c *gin.Context) {
	threads, err := data.Threads()
	if err != nil {
		error_message(c, "Cannot get threads")
	} else {
		_, err := session(c)
		if err != nil {
			generateHTML(c.Writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(c.Writer, threads, "layout", "private.navbar", "index")
		}
	}
}
