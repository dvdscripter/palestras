package main

import (
	"html/template"
	"log"
	"net/http"
)

const header = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<style>
body {
	border-style: solid;
}
</style>
<body>`

const footer = `
</p>
</body>
</html>`

func main() {
	http.HandleFunc("/boleto", boleto)

	log.Println("localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func boleto(w http.ResponseWriter, r *http.Request) {
	doc := r.URL.Query().Get("numerodoc")
	content := `<p>Número do documento: `
	content += doc
	html := header + content + footer
	w.Write([]byte(html))
}

// just to load html/template
func placeholder(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("content").Parse(`{{.}}`)
	t.Execute(w, "")
}
