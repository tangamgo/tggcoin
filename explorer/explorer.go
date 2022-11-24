package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"tggcoin/blockchain"
)

const (
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(writer http.ResponseWriter, request *http.Request) {
	data := homeData{"Home", nil}
	templates.ExecuteTemplate(writer, "home", data)
}

func add(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		templates.ExecuteTemplate(writer, "add", nil)
	case "POST":
		blockchain.Blockchain().AddBlock()
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}

}

func Start(port int) {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), handler))
}
