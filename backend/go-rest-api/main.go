package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Post struct {
	ID            int    `json:"id"`
	Body          string `json:"body"`
	Author        string `json:"author"`
	Date          string `json:"date"`
	Avatar        string `json:"avatar"`
	CommentCount  string `json:"commentCount"`
	FavoriteCount string `json:"favoriteCount"`
}
type Posts []Post

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
	Favorites []int  `json:"favorites"`
}
type Users []User

type Comment struct {
	ID           int    `json:"id"`
	Comment      string `json:"comment"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
	Date         string `json:"date"`
	PostID       int    `json:"postid"`
	CommentCount int    `json:"commentcount"`
}

type Favorite struct {
	ID         int  `json:"userid"`
	IsFavorite bool `json:"isFavorite"`
	Count      int  `json:"count"`
	PostID     int  `json:"postid"`
}

type Comments []Comment

func allPosts(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "db/database.db")
	table, _ := db.Query("select * from posts")
	var id int
	var body string
	var author string
	var date string
	var avatar string
	var commentCount string
	var favoriteCount string

	var posts Posts

	for table.Next() {
		error := table.Scan(&id, &body, &author, &date, &avatar, &commentCount, &favoriteCount)
		if error == nil {
			posts = append(posts, Post{ID: id, Body: body, Author: author, Date: date, Avatar: avatar, CommentCount: commentCount, FavoriteCount: favoriteCount})
		} else {
			fmt.Fprintf(w, "Something went wrong.")
		}

	}

	table.Close()
	db.Close()

	fmt.Println("GET request granted on /posts")
	json.NewEncoder(w).Encode(posts)
}

func sendPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request granted on /posts")
	var post Post

	_ = json.NewDecoder(r.Body).Decode(&post)

	db, _ := sql.Open("sqlite3", "db/database.db")

	pros, _ := db.Prepare("insert into posts (body,author,date,avatar) values (?,?,?,?)")

	pros.Exec(post.Body, post.Author, post.Date, post.Avatar)

	db.Close()

	fmt.Println(post)
	json.NewEncoder(w).Encode(&post)

}

func getComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query()["id"][0])
	fmt.Println("GET request granted on /comments")
	db, _ := sql.Open("sqlite3", "db/database.db")
	rows, _ := db.Query(`SELECT * FROM comments WHERE postId=$1`, r.URL.Query()["id"][0])
	var id int
	var comment string
	var username string
	var avatar string
	var date string
	var postid int
	var comments Comments

	for rows.Next() {
		error := rows.Scan(&id, &comment, &username, &avatar, &date, &postid)
		if error == nil {
			comments = append(comments, Comment{ID: id, Comment: comment, Username: username, Avatar: avatar, Date: date, PostID: postid})
		}
	}

	rows.Close()
	db.Close()

	json.NewEncoder(w).Encode(comments)
}

func sendComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request granted on /comments")
	var comment Comment

	_ = json.NewDecoder(r.Body).Decode(&comment)

	// post comment
	db, _ := sql.Open("sqlite3", "db/database.db")

	pros, _ := db.Prepare("insert into comments (comment,username,avatar,date,postId) values (?,?,?,?,?)")

	pros.Exec(comment.Comment, comment.Username, comment.Avatar, comment.Date, comment.PostID)

	// update posts commentCount

	pros, _ = db.Prepare("UPDATE posts SET commentCount=? WHERE id=?")
	pros.Exec((comment.CommentCount + 1), comment.PostID)

	db.Close()

	fmt.Fprintf(w, "Request complated.")

}
func getFavorite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET request granted on /favorite")

	var user User
	var favorites []int
	_ = json.NewDecoder(r.Body).Decode(&user)
	db, _ := sql.Open("sqlite3", "db/database.db")
	rows, _ := db.Query(`SELECT postid FROM favorites WHERE userid=$1`, user.ID)
	for rows.Next() {
		var id int
		var userid int
		var postid int
		rows.Scan(&id, &userid, &postid)
		favorites = append(favorites, postid)
	}
	user.Favorites = favorites
	rows.Close()
	db.Close()
	json.NewEncoder(w).Encode(user)

}

func setFavorite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request granted on /favorite")

	var favorite Favorite
	_ = json.NewDecoder(r.Body).Decode(&favorite)
	fmt.Println(favorite)
	db, _ := sql.Open("sqlite3", "db/database.db")
	pros, _ := db.Prepare("UPDATE posts SET favoriteCount=? WHERE id=?")
	pros.Exec(favorite.Count, favorite.PostID)
	if favorite.IsFavorite {
		pros, _ = db.Prepare("INSERT INTO favorites (userid,postid) VALUES (?,?)")
		pros.Exec(favorite.ID, favorite.PostID)
	} else {
		fmt.Println(favorite.PostID)
		pros, _ = db.Prepare("DELETE FROM favorites WHERE postid=? and userid=?")
		pros.Exec(favorite.PostID, favorite.ID)
	}

	db.Close()

	fmt.Fprintf(w, "Request complated.")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request granted on /login")
	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)

	db, _ := sql.Open("sqlite3", "db/database.db")
	rows, _ := db.Query(`SELECT * FROM users WHERE username=$1`, user.Username)

	var id int
	var username string
	var email string
	var password string
	var avatar string
	var favorites []int

	for rows.Next() {
		rows.Scan(&id, &username, &email, &password, &avatar)

	}
	rows.Close()
	fmt.Println(username, password)
	if password == user.Password {
		fmt.Println("Login işlemi başarılı...")
		rows, _ = db.Query("SELECT * from favorites WHERE userid=$1", id)
		for rows.Next() {
			var favid int
			var favuserid int
			var favpostid int
			rows.Scan(&favid, &favuserid, &favpostid)
			favorites = append(favorites, favpostid)
		}
		rows.Close()
		db.Close()
		user.ID = id
		user.Username = username
		user.Password = ""
		user.Avatar = avatar
		user.Favorites = favorites

		token, _ := GenerateJWT(user)
		user.Token = token
		json.NewEncoder(w).Encode(&user)

	} else {
		fmt.Fprintf(w, "wrong")
		db.Close()
	}

}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request granted on /register")
	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Username == "" && len(user.Username) < 6 && len(user.Username) > 16 &&
		user.Password == "" && len(user.Password) < 6 && len(user.Password) > 16 &&
		user.Email == "" && !validateEmail(user.Email) {
		fmt.Fprintf(w, "wrong")
	} else {
		db, _ := sql.Open("sqlite3", "db/database.db")
		rows, _ := db.Query(`SELECT username FROM users WHERE username=$1 or email=$2`, user.Username, user.Email)

		var username string
		for rows.Next() {
			rows.Scan(&username)
		}
		if username != "" {
			fmt.Fprintf(w, "wrong")
		} else {

			pros, _ := db.Prepare("insert into users (username,email,password,avatar) values (?,?,?,?)")
			pros.Exec(user.Username, user.Email, user.Password, user.Avatar)
			rows, _ := db.Query(`Select id from users where username=$1`, user.Username)
			var id int
			for rows.Next() {
				rows.Scan(&id)
			}
			rows.Close()
			var favorites []int
			rows, _ = db.Query("SELECT * from favorites WHERE userid=$1", id)
			for rows.Next() {
				var favid int
				var favuserid int
				var favpostid int
				rows.Scan(&favid, &favuserid, &favpostid)
				favorites = append(favorites, favpostid)
			}
			user.ID = id
			user.Favorites = favorites
			user.Password = ""
			user.Email = ""

			token, _ := GenerateJWT(user)
			user.Token = token
			json.NewEncoder(w).Encode(&user)
		}
		db.Close()
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")

}

func GenerateJWT(user User) (string, error) { // JWT Generator
	var mySigningKey = []byte("mysuperduperubersecretkey")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = user.Username
	claims["avatar"] = user.Avatar
	claims["id"] = user.ID
	claims["favorites"] = user.Favorites
	claims["exp"] = time.Now().Add(time.Minute*30).Unix() * 1000 // Duration

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Routes...
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/posts", allPosts).Methods("GET")
	myRouter.HandleFunc("/posts", sendPost).Methods("POST")
	myRouter.HandleFunc("/comments", getComments).Methods("GET")
	myRouter.HandleFunc("/comments", sendComment).Methods("POST")
	myRouter.HandleFunc("/favorite", getFavorite).Methods("GET")
	myRouter.HandleFunc("/favorite", setFavorite).Methods("POST")
	myRouter.HandleFunc("/login", login).Methods("POST")
	myRouter.HandleFunc("/register", register).Methods("POST")

	handler := handlers.CORS( // CORS Settings...
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Origin", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(10),
		handlers.AllowCredentials(),
	)(myRouter)
	var httpAddr = flag.String("http", ":8081", "Listen address")
	fmt.Println("Server listen on", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, handler))

}

func main() {
	handleRequests()
}
