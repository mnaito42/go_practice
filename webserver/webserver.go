package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter,
	req *http.Request){

	fmt.Fprintln(writer, "レッツゴー\n")
}