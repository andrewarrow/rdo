package controllers

import (
	"net/http"
	"rdo/models"

	"github.com/gin-gonic/gin"
)

func UsersIndex(c *gin.Context) {
	if !BeforeAll("user", c) {
		return
	}
	users, err := models.SelectUsers(Db)

	c.HTML(http.StatusOK, "users__index.tmpl", gin.H{
		"flash": err,
		"users": users,
		"name":  "name",
	})

}
func UsersCreate(c *gin.Context) {
	if false {
		SetFlash("try again", c)
		c.Redirect(http.StatusFound, "/")
	} else {
		//companyName := c.PostForm("company_name")
		//name := c.PostForm("name")
		email := c.PostForm("email")
		flavor := c.PostForm("flavor")

		//db.Insert("users", flavor, companyName, name, email)
		user := models.User{}
		user.Email = email
		user.Flavor = flavor
		c.SetCookie("user", user.Encode(), 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusFound, "/")
	}
}
func UsersNew(c *gin.Context) {
	c.HTML(http.StatusOK, "users__new.tmpl", gin.H{
		"flash":  "",
		"user":   map[string]string{},
		"flavor": c.Query("flavor"),
	})

}

func UsersShow(c *gin.Context) {
	c.HTML(http.StatusOK, "users__show.tmpl", gin.H{
		"flash": "",
		"user":  map[string]string{},
		"name":  "name",
	})

}
