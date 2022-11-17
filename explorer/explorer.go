package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"nomadcoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(writer http.ResponseWriter, request *http.Request) {

	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(writer, "home", data)
}

func add(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		templates.ExecuteTemplate(writer, "add", nil)
	case "POST":
		request.ParseForm()
		data := request.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
	}

}

func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
