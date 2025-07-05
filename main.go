package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	Addition()

	http.HandleFunc(("/add"), func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		a := r.FormValue("a")
		b := r.FormValue("b")

		aInt, err1 := strconv.Atoi(a)
		bInt, err2 := strconv.Atoi(b)

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Addition of %d and %d is %d", aInt, bInt, aInt+bInt)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
