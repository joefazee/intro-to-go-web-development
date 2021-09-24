package main

import (
	"html/template"
	"net/http"

	"abahjoseph.com/books/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	templateFiles := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}

}

func (app *application) books(w http.ResponseWriter, r *http.Request) {

	templateFiles := []string{
		"./ui/html/books.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(templateFiles...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := struct {
		Users []*models.User
	}{
		Users: users,
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
