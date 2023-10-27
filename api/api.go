package api

// API is the API struct
type API struct {
	DBName string
}

// NewAPI returns a new API struct
func NewAPI(db string) API {
	return API{
		DBName: db,
	}
}
