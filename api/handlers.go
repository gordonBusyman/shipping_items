package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"module-path/internal"
	"module-path/util"
)

// PackItemsHandler handles the GET /pack_items endpoint
func (api API) PackItemsHandler(w http.ResponseWriter, r *http.Request) {
	itemsOrdered, err := strconv.Atoi(chi.URLParam(r, "items"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid input")

		return
	}

	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())

		return
	}

	s.PacksAvailable = p
	if len(s.PacksAvailable) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "no packs available")

		return
	}

	result := s.CalculatePacks(itemsOrdered)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, util.PrepareMapResponse(result))
}

// AvailablePacksHandler handles the GET /available_packs endpoint
func (api API) AvailablePacksHandler(w http.ResponseWriter, _ *http.Request) {
	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())

		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, util.PrepareSliceResponse(p))
}

// DeletePackItemsHandler handles the DELETE /pack/{items} endpoint
func (api API) DeletePackItemsHandler(w http.ResponseWriter, r *http.Request) {
	size, err := strconv.Atoi(chi.URLParam(r, "items"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid input")

		return
	}

	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())

		return
	}

	// Check if pack size exists
	for k, v := range p {
		if size == v {
			// Write to storage without the pack size
			if err := s.WritePacks(append(p[:k], p[k+1:]...)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, err.Error())

				return
			}

			w.WriteHeader(http.StatusOK)

			return
		}
	}

	// Pack size not found, means it doesn't exist
	w.WriteHeader(http.StatusBadRequest)
}

// PostPackItemHandler handles the POST /pack/{items} endpoint
func (api API) PostPackItemHandler(w http.ResponseWriter, r *http.Request) {
	size, err := strconv.Atoi(chi.URLParam(r, "items"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid input")

		return
	}

	s := internal.NewStore(api.DBName)

	p, err := s.AvailablePacks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())

		return
	}

	// Check if pack size already exists
	for _, v := range p {
		if size == v {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
	}

	// Add new pack size and write to storage
	if err := s.WritePacks(append(p, size)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())

		return
	}

	w.WriteHeader(http.StatusCreated)
}
