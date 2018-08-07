rest
-

[![CircleCI](https://circleci.com/gh/thepkg/rest.svg?style=svg)](https://circleci.com/gh/thepkg/rest)
[![Go Report Card](https://goreportcard.com/badge/github.com/thepkg/rest)](https://goreportcard.com/report/github.com/thepkg/rest)
[![godoc](https://godoc.org/github.com/thepkg/rest?status.svg)](https://godoc.org/github.com/thepkg/rest)

Package `rest` it's tiny lightweight package which helps to work with RESTful API and JSON API.

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