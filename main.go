package main

import(
  "net/http"
  "path/filepath"
  "sync"
  "text/template"
  "github.com/gin-gonic/gin"
	"bytes"
	"image/png"
	"os"
	"strings"
	"image/jpeg"

	"image/gif"
	"encoding/base64"
	//"fmt"
)

type templateHandler struct {
  once sync.Once
  filename string
  templ *template.Template
}

type Image struct {
	Base64     string `form:"base64" json:"base64" binding:"required"`
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
			"message": "1121",
		})
	})

	v1 := r.Group("/v1")
	{
		//v1.GET("/login", GetLoginView)
		v1.POST("/postImage", postImage)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

func postImage(c *gin.Context) {
	var json Image
	if err := c.Bind(&json); err == nil {
		convertBase64ToImage(json.Base64)
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func convertBase64ToImage(b64 string) {

	dataIndex := strings.Index(b64, ",")
	dataIndex++

	unbased, err := base64.StdEncoding.DecodeString(b64[dataIndex:])
	if err != nil {
		panic("Bad png")
	}

	coI := strings.Index(b64, ",")

	//fmt.Print(strings.TrimSuffix(b64[5:coI], ";base64"))
	switch strings.TrimSuffix(b64[5:coI], ";base64") {
	case "image/png":
		f, err := os.Create("./test.png")
		if err != nil {
			panic("Cannot create file")
		}

		pngI, err := png.Decode(bytes.NewReader(unbased))
		png.Encode(f, pngI)
		if err != nil {
			panic("Bad png")
		}

		f.Close()
	case "image/jpeg":
		f, err := os.Create("./test.jpg")
		if err != nil {
			panic("Cannot create file")
		}

		jpgI, err := jpeg.Decode(bytes.NewReader(unbased))
		var opt jpeg.Options
		opt.Quality = 80

		jpeg.Encode(f, jpgI, &opt)
		if err != nil {
			panic("Bad png")
		}

		f.Close()

	case "image/gif":
		f, err := os.Create("./test.gif")
		if err != nil {
			panic("Cannot create file")
		}

		img, err := gif.DecodeAll(bytes.NewReader(unbased))
		var opt gif.Options
		opt.NumColors = 256

		gif.EncodeAll(f, img)
		if err != nil {
			panic("Bad png")
		}

		f.Close()
	}


}