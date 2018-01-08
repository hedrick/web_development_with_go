package main

import (
	"net/http"

	"./controllers"
	"./views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
var signUpView *views.View

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	// takes path and usersC's ResponseWriter
	r.HandleFunc("/signup", usersC.New)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
