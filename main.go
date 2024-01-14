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
	// Create packs.json if it doesn't exist
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
	// Serve static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Serve the index page
	r.Get("/", a.Index)

	// API endpoints
	r.Get("/pack_items/{items}", a.PackItems)
	r.Get("/available_packs", a.AvailablePkgSizes)
	r.Delete("/pack/{items}", a.DeletePkgSize)
	r.Post("/pack/{items}", a.CreatePkgSize)

	fmt.Println("server listening on 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
