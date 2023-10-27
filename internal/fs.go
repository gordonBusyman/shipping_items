package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrInternalServer = fmt.Errorf("internal server error")
)

// AvailablePacks reads the packs available from the file
func (s *Store) AvailablePacks() ([]int, error) {
	file, err := os.Open(s.DBName)
	if err != nil {
		log.Printf("error opening file: %v", err)

		return nil, ErrInternalServer
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Printf("error reading file: %v", err)

		return nil, ErrInternalServer
	}

	var packs []int

	if err := json.Unmarshal(byteValue, &packs); err != nil {
		log.Printf("error unmarshalling json: %v", err)

		return nil, ErrInternalServer
	}

	return packs, nil
}

func (s *Store) WritePacks(packs []int) error {
	// Open the file. This will create the file if it doesn't exist, or truncate it to zero length if it does.
	file, err := os.Create(s.DBName)
	if err != nil {
		log.Printf("error opening file: %v", err)

		return ErrInternalServer
	}
	defer file.Close()

	jsonData, _ := json.Marshal(packs)

	writer := io.Writer(file)
	if _, err = writer.Write(jsonData); err != nil {
		log.Printf("error writing to file: %v", err)

		return ErrInternalServer
	}

	return nil
}
