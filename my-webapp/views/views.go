package views

import (
    "net/http"
    "my-webapp/controllers"
)

// FrontPageView handles request for the Home page.
func FrontPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "frontPage") 
}

// AboutPageView handles request for the About page. 
func AboutPageView(w http.ResponseWriter, r *http.Request) {
    controllers.RenderBasicPage(w, "aboutPage")
}

// StaticFileViewHandler handles all /static/* requests. 
func StaticFileViewHandler(w http.ResponseWriter, r *http.Request) {
    staticFileName := r.URL.Path[1:]
    http.ServeFile(w, r, staticFileName)
}