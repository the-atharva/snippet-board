package main

import (
		"bytes"
		"fmt"
		"net/http"
		"runtime/debug"
		"time"

		"snippetbox.basic/pkg/models"

		"github.com/justinas/nosurf"
)

func (app *application) authenticatedUser(r *http.Request) *models.User {
		user, ok := 	r.Context().Value(contextKeyUser).(*models.User)
		if !ok {
				return nil
		}
		return user
}

func (app *application) serverError(w http.ResponseWriter, err error) {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		app.errorLog.Output(2, trace)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int){
		http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter){
		app.clientError(w, http.StatusNotFound)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
		if td == nil {
				td = &templateData{}
		}
		td.AuthenticatedUser = app.authenticatedUser(r)
		td.CSRFToken = nosurf.Token(r) 
		td.CurrentYear = time.Now().Year()
		td.Flash = app.session.PopString(r, "flash")
		return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
		ts, ok := app.templateCache[name]
		if !ok {
				app.serverError(w, fmt.Errorf("The template %s does not exisit", name))
				return
		}
		buf := new(bytes.Buffer)
		err := ts.Execute(buf, app.addDefaultData(td, r))
		if err != nil {
				app.serverError(w, err)
				return
		}
		buf.WriteTo(w)


}
























 
