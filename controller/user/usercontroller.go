package user

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"studentbookef/config"
	"studentbookef/domain"
	"studentbookef/io/user"
	"time"
)

func User(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Get("/login", logInHandler(app))
	r.Post("/loginpost", LoginPostHandler(app))
	r.Get("/signup", SignUpHandler(app))
	r.Post("/register", RegisterHandler(app)) //this method receives signUp form
	r.Get("/userAccount",userAccountHandler(app)) // done by Taylor
	return r
}

type Message struct {
	Message string
	Class   string
}

func GetMessage(Type string) Message {
	switch Type {
	case "sign_up_error":
		text := "An error has occurred during sign up. You may have already signed up"
		return Message{text, "warning"}
	case "sign_up_success":
		text := "You have successfully sign up, please verify your email we have sent your an email with your temporary password"
		return Message{text, "info"}
	case "just_login":
		text := "Welcome back"
		return Message{text, "info"}
	case "login_error":
		text := "An error has occurred, Please try again with correct details"
		return Message{text, "warning"}
	case "post_error":
		text := "An error has occurred, Please try again later"
		return Message{text, "warning"}
	case "post_error_need_to_signup":
		text := "An error has occurred, Please try to sign in first and again later"
		return Message{text, "warning"}
	case "post_empty_error":
		text := "Please make sure that you have filled in all the required fields"
		return Message{text, "warning"}
	case "post_image_error":
		text := "Please try again uploading light pictures "
		return Message{text, "warning"}

	case "userAccount_error": // done by Taylor
		text := "An error has occurred. please check your input and try again"
		return Message{text,"warning"}

	case "userAccount_successful_added":
		text := "Thanks for your time, your account was successfully created"
		return Message{text,"info"}

	}
	return Message{}
}

/****
When the user press submit button on sign up form this method will execute.
we will collect all the data in the form with r.ParseForm() method now we getting each input by passing the input name(html name).
we then create a user with only email and name other attributes will remain empty until when the user update his profile.
if an error occurs we will redirect the url address to /user/signup.
this Url will return a sign up page on user's interface with a proper error Message
But if there no errors, we will direct the user on home page with a notification Message for him/her to check the email to confirm registration.
*/
func LoginPostHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myuser := domain.UserAccount{}
		r.ParseForm()
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		if password != "" || email != "" {
			myuser = domain.UserAccount{email, password, "", time.Now()}
			fmt.Println(myuser)
			result, err := user.UserLog(myuser)
			fmt.Println(result, " result")
			if result.Email != "" {
				// If there is no error we save the login details in the cession so that we can authenticate the user during her/his cession period
				app.Session.Put(r.Context(), "userEmail", result.Email)
				app.Session.Put(r.Context(), "userMessage", "just_login")
				http.Redirect(w, r, "/", 301)
				return
			} else {
				app.InfoLog.Println(err)
				app.Session.Put(r.Context(), "userMessage", "login_error")
				//app.Session.Put(r.Context(), "userMessage","just_login")
				http.Redirect(w, r, "/user/login", 301)
				return
			}

		} else {
			app.InfoLog.Println("error with password or username")
			app.Session.Put(r.Context(), "userMessage", "login_error")
			//app.Session.Put(r.Context(), "userMessage","just_login")
			http.Redirect(w, r, "/user/login", 301)
			return
		}

	}
}

func RegisterHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myuser := domain.User{} //creating an empty object
		r.ParseForm()           //Now we grabbing the contents of the form by call the name of the input(html)
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		if email != "" {
			myuser = domain.User{email, name, "", ""}
			_, err := user.CreateUser(myuser)
			if err != nil { //when an error occurs when signing up
				app.Session.Put(r.Context(), "userMessage", "sign_up_error")
				http.Redirect(w, r, "/user/signup", 301)
				return
			} else {
				app.Session.Put(r.Context(), "userMessage", "sign_up_success")
				http.Redirect(w, r, "/", 301)
				return
			}
		}
	}
}

func SignUpHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Type := Message{}
		sessionType := app.Session.GetString(r.Context(), "userMessage") // we are checking what could be in the cession.
		app.Session.Remove(r.Context(), "userMessage")
		if sessionType != "" { //if there is something in the cession we need to check what it is
			Type = GetMessage(sessionType)
		}
		files := []string{
			app.Path + "user/sign_up.html",
			app.Path + "template/navigator.html",
			app.Path + "template/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, Type)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func logInHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("voila we are in")
		files := []string{
			app.Path + "user/loginpage.html",
			app.Path + "template/navigator.html",
			app.Path + "template/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "user/loginpage.html",
			app.Path + "template/navigator.html",
			app.Path + "template/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func userAccountHandler(app *config.Env) http.HandlerFunc {  //done by Taylor
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("voila we are in")
		files := []string{
			app.Path + "user/userAccount.html",
			app.Path + "template/navigator.html",
			app.Path + "template/footer.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}