package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func timeHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		date := time.Now()
		fmt.Fprintf(w, date.Format("3h4"))
	}
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}

		author := req.FormValue("author")
		entry := req.FormValue("entry")

		file, err := os.OpenFile("bdd.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		defer file.Close() // on ferme automatiquement à la fin de notre programme

		if err != nil {
			panic(err)
		}

		_, err = file.WriteString(author + ": " + entry + "\n") // écrire dans le fichier
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%s: %s", author, entry)
	}
}

func entriesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {

		data, err := ioutil.ReadFile("bdd.txt") // lire le fichier text.txt
		if err != nil {
			fmt.Println(err)
		}

		entries := strings.Split(string(data), "\n")

		for _, dataEntry := range entries {
			if dataEntry != "" {
				output := strings.Split(dataEntry, ": ")
				fmt.Fprintf(w, output[1]+"\n")
			}
		}
	}
}

func main() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/entries", entriesHandler)
	http.ListenAndServe(":4567", nil)
}
