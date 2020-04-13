package home

import (
	"errors"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"studentbookef/config"
	"studentbookef/controller/user"
	"studentbookef/domain"
	user2 "studentbookef/io/user"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))

	return r
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myUser := domain.User{}
		cessionData := user.Message{}
		//just_login_message:=user.Message{}
		//just_login:=app.Session.GetString(r.Context(),"just_login") 		// we checking if the user just login?
		//if just_login!=""{
		//	just_login_message=user.GetMessage(just_login)
		//}
		sessionType := app.Session.GetString(r.Context(), "userMessage")
		if sessionType != "" {
			cessionData = user.GetMessage(sessionType)
		}
		email := app.Session.GetString(r.Context(), "userEmail") //We are checking of the user has login?
		if email != "" {
			err := errors.New("")
			myUser, err = user2.ReadUser(email)
			if err != nil {
				app.InfoLog.Println(err)
			}
		}
		type PageData struct {
			PageMessage user.Message
			User        domain.User
		}
		data := PageData{cessionData, myUser}
		files := []string{
			app.Path + "index.html",
			app.Path + "template/navigator.html",
			app.Path + "template/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
