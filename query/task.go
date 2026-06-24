package query

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
func QueryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method Rejected")
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	atext, _ := strconv.Atoi(a)
	btext, _ := strconv.Atoi(b)

	if op == "add" {
		fmt.Fprintln(w, atext+btext)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

}

func CountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Send a POST request with text to count word")
		return
	}
	if r.Method == http.MethodPost {
		read, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusContinue)
			fmt.Fprintln(w, "continue")
			return
		} else {
			fmt.Fprintln(w, len(read))
		}

	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Method Not Found")
	//return

}

func AgentHandler(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")
	if agent == "" {
		agent = "Unknown"
	}
	fmt.Fprintf(w, "You are visiting us using: %s", agent)

}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "This Method Is not Allowed")
		return
	}
	defer r.Body.Close()

	read, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error reading file")
		return
	}
	
	if len(read) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "body cannot be empty")
		return
	}

	
	w.Write(read)
}


