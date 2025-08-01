package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
func saveUser(c echo.Context) error {
	// Logic to save user
	return c.String(http.StatusOK, "User saved")
}

func getUser(c echo.Context) error {
	// Logic to get user by ID
	return c.String(http.StatusOK, "User retrieved")
}

func updateUser(c echo.Context) error {
	// Logic to update user by ID
	return c.String(http.StatusOK, "User updated")
}

func deleteUser(c echo.Context) error {
	// Logic to delete user by ID
	return c.String(http.StatusOK, "User deleted")
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
