package main

import (
	"fmt"
	"net/http"
)

// TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
func respondToRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World!")
}
func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(respondToRequest))
	if err != nil {
		panic(err)
	}
}
