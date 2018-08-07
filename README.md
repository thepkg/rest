rest
-

Package `rest` it's tiny lightweight package which help to work with RESTful API and JSON API.

## Installation

`go get -u github.com/thepkg/rest`

## Usage

````
import "github.com/thepkg/rest"

func main() {
	rest.GET("/users", func(w http.ResponseWriter, r *http.Request) {
		rest.Success(w, http.StatusOK, "find")
	})

	rest.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		rest.Success(w, http.StatusOK, "create")
	})

	rest.PUT("/users", func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusMethodNotAllowed, "update not allowed")
	})

	rest.DELETE("/users", func(w http.ResponseWriter, r *http.Request) {
		rest.Error(w, http.StatusMethodNotAllowed, "delete not allowed")
	})

	http.ListenAndServe(":8080", nil)
}
````
å