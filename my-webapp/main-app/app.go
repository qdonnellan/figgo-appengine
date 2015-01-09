package app

import (
    "net/http"
    "my-webapp/views"
)

func init() {
    http.HandleFunc("/", views.FrontPageView)
    http.HandleFunc("/about", views.AboutPageView)
    http.HandleFunc("/static/", views.StaticFileViewHandler)
}
