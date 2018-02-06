package main


import (
	"net/http"
	"./models"
)

func loadDbPage(title string) (*models.Page, error) {
	page := models.GetPage(db, title)
	return &models.Page{Title: title, Body: page.Body}, nil
}

func saveDbHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/db/"):]
	body := r.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	err := p.Save(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/db/"+title, http.StatusFound)
}

func editDbHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/db/"):]
	p, err := loadDbPage(title)
	if err != nil {
		p = &models.Page{Title: title}
	}
	renderTemplate(w, "view/db/edit", p)
}

func viewDbHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/db/"):]
	p, _ := loadDbPage(title)
	renderTemplate(w, "view/db/view", p)
}
