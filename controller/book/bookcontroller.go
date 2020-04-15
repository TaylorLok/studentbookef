package book

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"studentbookef/config"
	"studentbookef/domain"
	"studentbookef/io/book_io"
	location2 "studentbookef/io/location"
	"studentbookef/io/picture_io"
	"time"
)

func Book(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", NewPostHandler(app))
	r.Post("/postBook", PostBookHandler(app)) //this method receives signUp form
	r.Post("/postPost", PostPostHandler(app))
	return r
}

func PostPostHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userEmail")
		bookId := app.Session.GetString(r.Context(), "book_id")
		if email == "" && bookId == "" {
			app.Session.Put(r.Context(), "userMessage", "post_error_need_to_signup")
			http.Redirect(w, r, "/book/post", 301)
			return
		}
		//Grabbing data from the form
		r.ParseForm()
		description := r.PostFormValue("description")
		place := r.PostFormValue("place")
		placeDescription := r.PostFormValue("place_description")

		//This is the final, if all went well. we now need to create a post
		//todo we will need to implement map that can give us longitude an latitude
		if place != "" && placeDescription != "" {
			location := domain.Location{"", placeDescription, "", "", place}
			newLocation, err := location2.CreateLocation(location)
			if err != nil { // if an error occur
				app.InfoLog.Println(err, "there is an error when creating a post location")
				app.InfoLog.Println(err, "there is an error when creating a post location")
				app.Session.Put(r.Context(), "userMessage", "post_error")
				http.Redirect(w, r, "/book/post", 301)
				return
			}
			post := domain.BookPost{"", email, bookId, time.Now(), newLocation.Id, "", description}
			_, error := book_io.CreatBookPost(post)
			if error != nil { // if an error occur
				_, err := location2.DeleteLocation(newLocation)
				if err != nil {
					app.InfoLog.Println(error, "there is an error when rolling back on location after failing to create location")
				}
				app.InfoLog.Println(error, "there is an error when creating a post location")
				app.Session.Put(r.Context(), "userMessage", "post_error")
				http.Redirect(w, r, "/book/post", 301)
				return
			} else {
				app.Session.Put(r.Context(), "userMessage", "post_successful")
				http.Redirect(w, r, "/", 301)
				return
			}

		} else {
			//todo need to implement this cession error
			app.Session.Put(r.Context(), "userMessage", "post_error")
			http.Redirect(w, r, "/book/post", 301)
			return
		}

	}
}

func PostBookHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//we need to check the user if has created an account first
		email := app.Session.GetString(r.Context(), "userEmail")
		if email == "" {
			app.Session.Put(r.Context(), "userMessage", "post_error_need_to_signup")
			http.Redirect(w, r, "/login", 301)
			return
		}
		fmt.Println("voila we are in book PostBookHandler")
		picture := domain.Picture{}
		r.ParseForm()
		fmt.Println(" reading the file")
		file, handler, err := r.FormFile("file")
		file1, handler, err := r.FormFile("file1")
		fmt.Println(" read successful")
		if err != nil {
			fmt.Println(err, "<<<<<<>>>>>>>")
			app.Session.Put(r.Context(), "userMessage", "post_error")
			http.Redirect(w, r, "/book/postBook", 301)
			return
		}

		title := r.PostFormValue("title")
		language := r.PostFormValue("language")
		edition := r.PostFormValue("edition")
		price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64) // converting string into double

		//Creating a book first
		book := domain.Book{"", title, language, edition, price}
		newBook, err := book_io.CreateBook(book)
		if err != nil { //if an error occurs we interrupt everything here and return an error message to the user
			app.Session.Put(r.Context(), "userMessage", "post_error")
			http.Redirect(w, r, "/book/postBook", 301)
			return
		}
		//converting to []byte and into slice array
		fmt.Println(" converting to []byte and into slice array", handler)
		reader := bufio.NewReader(file)
		reader1 := bufio.NewReader(file1)
		content, _ := ioutil.ReadAll(reader)
		content1, _ := ioutil.ReadAll(reader1)

		//converting the file into an slice of byte
		sliceOfImage := [][]byte{content, content1}
		//we now looping the slice of byte to create picture
		for index, pic := range sliceOfImage {
			if index == 0 {
				picture, err = sendingPicture(pic, "fist picture")
			} else {
				picture, err = sendingPicture(pic, "second picture")
			}
			if err != nil { //if an error occurs when creating a picture here we return an error message and we delete the book and it picture
				app.InfoLog.Println(err)
				_, err := book_io.DeleteBook(newBook)
				if err != nil {
					app.InfoLog.Println(err, "  when trying to delete failed book: ", newBook)
				}
				app.Session.Put(r.Context(), "userMessage", "post_error")
				http.Redirect(w, r, "/book/postBook", 301)
				return
			} else { //in case when there is no error when creating a picture, now we will need to create BookImage
				bookImage := domain.BookImage{newBook.Id, picture.Id, ""}
				newBookImage, err := book_io.CreatBookImage(bookImage)
				if err != nil {
					app.InfoLog.Println(err)
					_, err := book_io.DeleteBook(newBook)
					if err != nil {
						app.InfoLog.Println(err, "  when trying to delete failed book: ", newBook)
					}
					_, errr := picture_io.DeletePicture(picture)
					if errr != nil {
						app.InfoLog.Println(err, "  when trying to delete failed Pictue: ", picture)
					}
					app.Session.Put(r.Context(), "userMessage", "post_error")
					http.Redirect(w, r, "/book/postBook", 301)
					return
				} else {
					fmt.Println("book Successful: ", newBookImage)
					app.Session.Put(r.Context(), "userMessage", "post_successful")
					http.Redirect(w, r, "/post/postBook", 301)
					return
				}
			}
		}
	}
}

func NewPostHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("voila we are in")
		files := []string{
			app.Path + "book/book_post.html",
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
func sendingPicture(picture []byte, desc string) (domain.Picture, error) {
	newPicture := domain.Picture{"", picture, desc}
	result, err := picture_io.CreatePicture(newPicture)
	return result, err
}
