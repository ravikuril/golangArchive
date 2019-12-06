
// All the components are explained below
//========================================================Sample code===========================================
package main
import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
)
type Post struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
}
var posts []Post
func getPosts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(posts)
}
func createPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var post Post
  _ = json.NewDecoder(r.Body).Decode(&post)
  post.ID = strconv.Itoa(rand.Intn(1000000))
  posts = append(posts, post)
  json.NewEncoder(w).Encode(&post)
}
func getPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range posts {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Post{})
}
func updatePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range posts {
    if item.ID == params["id"] {
      posts = append(posts[:index], posts[index+1:]...)
      var post Post
      _ = json.NewDecoder(r.Body).Decode(&post)
      post.ID = params["id"]
      posts = append(posts, post)
      json.NewEncoder(w).Encode(&post)
      return
    }
  }
  json.NewEncoder(w).Encode(posts)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range posts {
    if item.ID == params["id"] {
      posts = append(posts[:index], posts[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(posts)
}
func main() {
  router := mux.NewRouter()
  posts = append(posts, Post{ID: "1", Title: "My first post", Body:      "This is the content of my first post"})
  router.HandleFunc("/posts", getPosts).Methods("GET")
  router.HandleFunc("/posts", createPost).Methods("POST")
  router.HandleFunc("/posts/{id}", getPost).Methods("GET")
  router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
  router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
http.ListenAndServe(":8000", router)
}




//===============================================EXPLANATION====================================================
// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
//     "encoding/json"
// 	"github.com/gorilla/mux"
// 	"math/rand"
// 	"strconv"
// )

// // this is the post struct which will send the data as per the format
// type Post struct {
// 	ID string `json:"id"`
// 	Title string `json:"title"`
// 	Body string `json:"body"`
//   }
//   var posts []Post			// array of struct which will contain the post

// // Existing code from above
// func handleRequests() {
//     // creates a new instance of a mux router
//    //** myRouter := mux.NewRouter().StrictSlash(true)							//** are of same code block unblock it
//     // replace http.HandleFunc with myRouter.HandleFunc
// 	//**myRouter.HandleFunc("/", homePage)
// 	//router.HandleFunc("/<your-url>", <function-name>).Methods("<method>")
// 	// myRouter.HandleFunc("/all", returnAllArticles)
// 	//**myRouter.HandleFunc("/article/{id}", returnSingleArticle)
// 	//fmt.Println("Id is:",id)

// 	//ANOTHER BLOCK OF EXAMPLE 
// 		router := mux.NewRouter()
// 		router.HandleFunc("/posts", getPosts).Methods("GET")
// 		router.HandleFunc("/posts", createPost).Methods("POST")
// 		router.HandleFunc("/posts/{id}", getPost).Methods("GET")
// 		router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
// 		router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

//     log.Fatal(http.ListenAndServe(":10000", router))
// }

//  func getPosts(w http.ResponseWriter, r *http.Request) {
//  	w.Header().Set("Content-Type", "application/json")
//  	json.NewEncoder(w).Encode(posts)
//    }

// // this function checks if the passed id in the request is same in the array of post struct 
// func getPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)									// params contains the request parameters
// 	for _, item := range posts { 							// looping from the post and takes the 
// 	  if item.ID == params["id"] {							//item and checks if it is given id and breaks if it founds
// 		json.NewEncoder(w).Encode(item)
// 		break
// 	  }
// 	  return
// 	}
// 	json.NewEncoder(w).Encode(&Post{})
//   }

// //This is the handler that creates a new post, we start by creating a new instance of the struct Post.
//   func createPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var post Post
// 	_ = json.NewDecoder(r.Body).Decode(post)				// the request payload will have same format as the struct 
// 	post.ID = strconv.Itoa(rand.Intn(1000000))				//
// 	posts = append(posts, post)								// appends the post array which the current data 
// 	json.NewEncoder(w).Encode(&post)						// sending the post array with current updated data
//   }

// //For the update handler, we loop through our posts array to find the post to update. 
// //When it matches we remove that post from the array and create a new post with the same ID (using params[“id”]) with the new values from the request body
// func updatePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range posts {
// 	  if item.ID == params["id"] {
// 		posts = append(posts[:index], posts[index+1:]...)
// 		var post Post
// 		_ = json.NewDecoder(r.Body).Decode(post)
// 		post.ID = params["id"]
// 		posts = append(posts, post)
// 		json.NewEncoder(w).Encode(&post)
// 		return
// 	  }
// 	}
// 	json.NewEncoder(w).Encode(posts)
//   }

// // Here we are also looping through our posts array to find the post we want to delete. When we do, we delete the post with this line:

// func deletePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, item := range posts {
// 	  if item.ID == params["id"] {
// 		posts= append(posts[:index], posts[index+1]...)
// 		break
// 	  }
// 	}
// 	json.NewEncoder(w).Encode(books)
//   }

// func main() {
//     fmt.Println("Rest API v2.0 - Mux Routers")
//     Articles = []Article{
//         Article{Title: "Hello", Desc: "Article Description", Content	: "Article Content"},
//         Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
//     }
//     handleRequests()
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request){
//     vars := mux.Vars(r)
//     key := vars["id"]

//     fmt.Fprintf(w, "Key: " + key)
// }
