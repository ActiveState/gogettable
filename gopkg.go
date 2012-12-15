package main

import (
	"code.google.com/p/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	// HACK: we really want to match /name/*
	router.HandleFunc("/{name}", Handler)
	router.HandleFunc("/{name}/{y}", Handler)
	router.HandleFunc("/{name}/{y}/{z}", Handler)
	http.Handle("/", router)

	addr := ":" + os.Getenv("PORT")
	log.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "<html><head>%s</head></html>", getMetaTag(vars["name"]))
}

func getMetaTag(name string) string {
	return fmt.Sprintf("<meta name=\"go-import\" content=\"go.stacka.to/%s git ssh://gitolite@gitolite.activestate.com/%s\">",
		name, name)
}
