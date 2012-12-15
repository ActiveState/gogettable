package main

import (
	"code.google.com/p/gorilla/mux"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var Options struct {
	Import        string
	RepoURLFormat string
}

func main() {
	flag.StringVar(&Options.Import, "import", "", "Import prefix to handle (eg: go.company.com)")
	flag.StringVar(&Options.RepoURLFormat, "repo", "", "Repo URL template (eg: git://git.company.com:/%s.git)")
	flag.Parse()
	if Options.Import == "" || Options.RepoURLFormat == "" {
		flag.Usage()
		os.Exit(1)
	}

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
	fmt.Fprintf(w, "<html><head>%s</head></html>\n", getMetaTag(vars["name"]))
}

func getMetaTag(name string) string {
	imp := Options.Import + "/" + name
	repo := fmt.Sprintf(Options.RepoURLFormat, name)
	return fmt.Sprintf("<meta name=\"go-import\" content=\"%s git %s\">", imp, repo)
}
