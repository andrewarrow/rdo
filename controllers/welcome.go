package controllers

import (
	"net/http"

	"rdo/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
var flash = ""
var user *models.User

func ValidAdminUser(c *gin.Context) bool {
	json, _ := c.Cookie("user")
	user = models.DecodeUser(json)
	if user == nil || user.Flavor != "admin" {
		SetFlash("you need to login", c)
		c.Redirect(http.StatusFound, "/sessions/new")
		c.Abort()
		return false
	}
	return true
}
func BeforeAll(flavor string, c *gin.Context) bool {
	flash, _ = c.Cookie("flash")
	SetFlash("", c)

	json, _ := c.Cookie("user")
	user = models.DecodeUser(json)

	if flavor == "" {
		return true
	}

	if flavor == "user" {
		if user == nil {
			SetFlash("you need to login", c)
			c.Redirect(http.StatusFound, "/sessions/new")
			c.Abort()
			return false
		}
		return true
	}
	if user == nil || user.Flavor != "admin" {
		SetFlash("you need to login", c)
		c.Redirect(http.StatusFound, "/sessions/new")
		c.Abort()
		return false
	}
	return true
}

func SetFlash(s string, c *gin.Context) {
	c.SetCookie("flash", s, 3600, "/", "", false, true)
}

func WelcomeIndex(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	if user == nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"user":  nil,
			"flash": flash,
		})
		return
	}

	c.HTML(http.StatusOK, "homepage.tmpl", gin.H{
		"user":  user,
		"flash": flash,
	})
}
