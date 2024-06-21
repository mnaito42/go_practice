package main

import (
	"fmt"
	"net/http"

)

func add(writer http.ResponseWriter,  req *http.Request){

	fmt.Fprintln(writer, "***初めての関数***")
}