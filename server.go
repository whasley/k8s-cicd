package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}


func main() {
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "Projeto DevOps - Alexandre",
            Todos: []Todo{
                {Title: "Git Push para o GitHub", Done: true},
                {Title: "Login no DockerHub", Done: true},
                {Title: "Build da imagem Docker", Done: true},
				{Title: "Push da imagem para o Docker Hub", Done: true},
				{Title: "Deploy da Imagem no Kubernetes", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":3000", nil)
}



/*func main() {

http.HandleFunc("/", Home)
	http.ListenAndServe(":3000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Olá mundo - Kubernetes</h1>"))
}
*/
