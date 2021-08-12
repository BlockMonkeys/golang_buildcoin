package explorer

import (
	"fmt"
	"github.com/blockmonkeys/nomadCoin/blockchain"
	"net/http"
	"html/template"
	"log"
)

var templates *template.Template


const (
	port string = ":4000"
	templateDir string ="explorer/templates/"
)

type homeData struct{
	PageTitle		string
	Blocks 			[]*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request){
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData") // Parse Body (HTML)
		blockchain.GetBlockchain().AddBlock(data)
		//redirect
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}


func Start(){
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.html"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.html"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Println("Listening on :", port)
	log.Fatal(http.ListenAndServe(port, nil))
}