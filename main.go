// skins project main.go
package main

import (
	"github.com/gorilla/mux"
	"github.com/minotar/minecraft"
	"image"
	"image/png"
	"io"
	"net/http"
)

func SkinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	username := vars["username"]

	skin, _ := minecraft.GetSkin(minecraft.User{Name: username})

	w.Header().Add("Content-Type", "image/png")

	WritePNG(w, skin.Image)
}

func WritePNG(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{username:"+minecraft.ValidUsernameRegex+"}.png", SkinHandler)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
