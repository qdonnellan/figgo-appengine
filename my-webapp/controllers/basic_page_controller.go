package controllers

import (
    "net/http"
    "my-webapp/models"
)

var templateTitleMap = map[string]string{
    "frontPage" : "[Go] Google App Engine on Docker; Demo",
    "aboutPage" : "Sample About Page",
}

// RenderBasicPage renders a simple template.
func RenderBasicPage(w http.ResponseWriter, templateName string) {
    basicPageModel := models.WebPage{
        Title : templateTitleMap[templateName],
        TemplateName : templateName,
    }
    RenderTemplateFromPageModel(w, basicPageModel)
}