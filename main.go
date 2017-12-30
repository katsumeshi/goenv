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

func convertBase64ToImage(b64 string) error {

	dataIndex := strings.Index(b64, ",") + 1
	unbased, err := base64.StdEncoding.DecodeString(b64[dataIndex:])
	if err != nil {
		return err
	}

	index := strings.Index(b64, ",")
	switch strings.TrimSuffix(b64[5:index], ";base64") {
	case "image/png":
		return convertBase64ToPng(unbased)
	case "image/jpeg":
		return convertBase64ToJpeg(unbased)
	case "image/gif":
		return convertBase64ToGif(unbased)
	}

	return nil
}

func convertBase64ToPng(unbased []byte) error {
	f, err := os.Create("../storage/test.png")
	if err != nil {
		return err
	}

	img, err := png.Decode(bytes.NewReader(unbased))
	png.Encode(f, img)
	if err != nil {
		return err
	}

	f.Close()
	return err
}

func convertBase64ToJpeg(unbased []byte) error {
	f, err := os.Create("../storage/test.jpg")
	if err != nil {
		return err
	}

	img, err := jpeg.Decode(bytes.NewReader(unbased))
	var opt jpeg.Options
	opt.Quality = 80

	jpeg.Encode(f, img, &opt)
	if err != nil {
		return err
	}

	f.Close()
	return err
}

func convertBase64ToGif(unbased []byte) error {
	f, err := os.Create("../storage/test.gif")
	if err != nil {
		return err
	}

	img, err := gif.DecodeAll(bytes.NewReader(unbased))
	var opt gif.Options
	opt.NumColors = 256

	gif.EncodeAll(f, img)
	if err != nil {
		return err
	}

	f.Close()
	return err
}