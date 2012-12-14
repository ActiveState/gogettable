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
	router.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "<html><head>%s</head></html>", getMetaTag(vars["name"]))
	})
	http.Handle("/", router)

	addr := ":" + os.Getenv("PORT")
	log.Println("Listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getMetaTag(name string) string {
	return fmt.Sprintf("<meta name=\"go-import\" content=\"gopkg.iprod.activestate.com/%s git ssh://gitolite@gitolite.activestate.com/%s\">",
		name, name)
}
