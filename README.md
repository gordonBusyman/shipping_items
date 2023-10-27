# Shipping Items API

## Description
This project is a simple example of coding.
It is a simple shipping items project.

## API "Documentation"
The API has 4 endpoints:
- `GET` `{base_url}/pack_items/{items}` Calculates the best way to pack the items
- `GET` `{base_url}/available_packs` Returns the available packs
- `DELETE` `{base_url}/pack/{items}` Deletes the pack
- `POST` `{base_url}/pack/{items}` Creates a new pack

## NOTES
- After restart application state resets to default (`250`, `500`, `1000`, `2000`, `5000` package sizes).
- There is some very basic test coverage.
- There is validation for the input data.
- There is a simple logging.
- There is a simple error handling.
- The calculation won't work efficiently for pack sizes that are not multiples of each other (e.g. `234`, `501`, `502`).
