package api

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/go-chi/chi"

	"module-path/internal"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// TemplateData a struct for the template data
type TemplateData struct {
	APIEndpoint string
}

func (api API) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := path.Join("templates", "index.html")
	templates := template.Must(template.ParseFiles(tmpl))

	data := TemplateData{
		APIEndpoint: "http://localhost:8080", // Modify as needed
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// PackItems handles the GET /pack_items endpoint.
func (api API) PackItems(w http.ResponseWriter, r *http.Request) {
	itemsOrdered, err := strconv.Atoi(chi.URLParam(r, "items"))
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "invalid input")

		return
	}

	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	s.PacksAvailable = p
	if len(s.PacksAvailable) == 0 {
		sendErrorResponse(w, http.StatusInternalServerError, "no packs available")

		return
	}

	result := s.CalculatePacks(itemsOrdered)

	json.NewEncoder(w).Encode(result)
	w.Header().Set("Content-Type", "application/json")
}

// AvailablePkgSizes handles the GET /available_packs endpoint.
func (api API) AvailablePkgSizes(w http.ResponseWriter, _ *http.Request) {
	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	json.NewEncoder(w).Encode(p)
	w.Header().Set("Content-Type", "application/json")
}

// DeletePkgSize handles the DELETE /pack/{items} endpoint
// This is a DELETE request because we are deleting a pack size.
func (api API) DeletePkgSize(w http.ResponseWriter, r *http.Request) {
	size, err := strconv.Atoi(chi.URLParam(r, "items"))
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "invalid input")

		return
	}

	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	// Check if pack size exists
	for k, v := range p {
		if size == v {
			// Write to storage without the pack size
			if err := s.WritePacks(append(p[:k], p[k+1:]...)); err != nil {
				sendErrorResponse(w, http.StatusInternalServerError, err.Error())

				return
			}

			return
		}
	}

	// Pack size not found, means it doesn't exist
	sendErrorResponse(w, http.StatusBadRequest, "pack size not found")
}

// CreatePkgSize handles the POST /pack/{items} endpoint
// This is a POST request because we are creating a new pack size.
func (api API) CreatePkgSize(w http.ResponseWriter, r *http.Request) {
	size, err := strconv.Atoi(chi.URLParam(r, "items"))
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "invalid input")

		return
	}

	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	// Check if pack size already exists
	for _, v := range p {
		if size == v {
			sendErrorResponse(w, http.StatusBadRequest, "pack size already exists")

			return
		}
	}

	// Add new pack size and write to storage
	if err := s.WritePacks(append(p, size)); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	w.WriteHeader(http.StatusCreated)
}
