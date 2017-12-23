package main

import(
  "net/http"
  "path/filepath"
  "sync"
  "text/template"
  "github.com/gin-gonic/gin"
)

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func() {
    t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
  })
  t.templ.Execute(w, nil)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "1111",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}