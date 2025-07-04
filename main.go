package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/addition-form/", http.StripPrefix("/addition-form", fs))

	http.HandleFunc("/add-form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Can't parse the body", http.StatusBadRequest)
		}

		num1Str := r.FormValue("a")
		num2Str := r.FormValue("b")

		num1Int, err1 := strconv.Atoi(num1Str)
		num2Int, err2 := strconv.Atoi(num2Str)

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		result := num1Int + num2Int
		response := map[string]interface{}{
			"result": result,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	})

	http.HandleFunc("/add-form-with-parse", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var input struct {
			A float64 `json:"a"`
			B float64 `json:"b"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		result := input.A + input.B
		response := map[string]interface{}{
			"result": result,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

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
