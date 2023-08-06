# Golang API Project - Book Store

This project implements a Golang API for a Book Store, providing various routes for managing books.

## Routes

### Create Book

Endpoint: `POST /book/`

Creates a new book in the store.

### Get Book

Endpoint: `GET /book/`

Retrieves a list of all books in the store.

### Get Book by ID

Endpoint: `GET /book/{bookId}`

Retrieves details of a specific book based on the provided `bookId`.

### Update Book

Endpoint: `PUT /book/{bookId}`

Updates the details of a specific book based on the provided `bookId`.

### Delete Book

Endpoint: `DELETE /book/{bookId}`

Deletes a specific book from the store based on the provided `bookId`.

## Usage

To use this API, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/Cyberguru1/Go-BookStore
   ```


2. Install required dependencies:

   ```bash
   cd Go-BookStore
   go mod download
   ```
3. Configure the database:
   Update the database connection details in `config/config.go`.
4. Build and run the application:

   ```bash
   go build
   ./cmd/main/main
   ```
5. Access the API:
   Use tools like `curl`, Postman, or your web browser to interact with the API using the provided routes.

## Dependencies

This project uses the following libraries/frameworks:

- [gorilla/mux](https://github.com/gorilla/mux): A powerful HTTP router and URL matcher for building web applications.
- [gorm](https://github.com/jinzhu/gorm)

## Contributing

Contributions are welcome! If you find any issues or have suggestions, feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
