package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"module-path/api"
)

func init() {
	//
	file, err := os.Create("packs.json")
	if err != nil {
		log.Printf("error opening file: %v", err)

		panic(err)
	}
	defer file.Close()

	// reset to initial state
	jsonData, _ := json.Marshal([]int{250, 500, 1000, 2000, 5000})

	writer := io.Writer(file)
	if _, err = writer.Write(jsonData); err != nil {
		log.Printf("error writing to file: %v", err)

		panic(err)
	}
}

func main() {
	r := chi.NewRouter()

	a := api.NewAPI("packs.json")

	r.Get("/pack_items/{items}", a.PackItemsHandler)
	r.Get("/available_packs", a.AvailablePacksHandler)
	r.Delete("/pack/{items}", a.DeletePackItemsHandler)
	r.Post("/pack/{items}", a.PostPackItemHandler)

	fmt.Println("server listening on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
