package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// UsersShow default implementation.
func UsersShow(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/show.html"))
}

// UsersIndex default implementation.
func UsersIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/index.html"))
}

// UsersNew default implementation.
func UsersNew(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/new.html"))
}

// UsersEdit default implementation.
func UsersEdit(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/edit.html"))
}

// UsersCreate default implementation.
func UsersCreate(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/create.html"))
}

// UsersUpdate default implementation.
func UsersUpdate(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/update.html"))
}

// UsersDelete default implementation.
func UsersDelete(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("users/delete.html"))
}

