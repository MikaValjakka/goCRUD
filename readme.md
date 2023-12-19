# Go Backend with Gin and MySQL

This is a basic example of a Go backend using the Gin framework and MySQL database with GORM.

- Go, often referred to as Golang, is a programming language developed by Google. It is known for its simplicity, efficiency, and strong support for concurrent programming.
- Gin is a web framework written in Go. It's lightweight and fast, making it a popular choice for building web applications and APIs in the Go programming language.
- Gorm is an Object Relational Mapping (ORM) library for Go. It provides a convenient way to interact with databases by mapping Go structs to database tables, making database operations more straightforward and idiomatic.
- MySQL is a popular open-source relational database management system (RDBMS). It is widely used for managing structured data and is compatible with various programming languages, including Go.

## Prerequisites

Before you start, ensure you have the following installed:

- Go: [https://golang.org/dl/](https://golang.org/dl/)
- MySQL: [https://dev.mysql.com/downloads/](https://dev.mysql.com/downloads/)

## Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd your-project-directory
   ```

2. Set up your MySQL database:

Create a new database (e.g., your_database_name).

3. Configure your environment:

Create a .env file in the project root:

    PORT=8080
    DB_STRING=username:password@tcp(127.0.0.1:3306)/your_database_name?  charset=utf8mb4&parseTime=True&loc=Local

Replace username, password, and your_database_name with your MySQL credentials.

4. Run the server:
   ```
   go run main.go
   ```

The server will start on http://localhost:8080.

## API Endpoints

### Create a Book

- Endpoint: POST /books

Example Payload:

```
{
  "title": "The Lord of the Rings",
  "author": "J. R. R. Tolkien",
  "pages": 1137
}
```

Response (Notice GORM created_at, updated_at and deleted_at):

```
{
"ID": 1,
"CreatedAt": "2023-12-18T13:36:20.786+02:00",
"UpdatedAt": "2023-12-18T13:36:20.786+02:00",
"DeletedAt": null,
"Title": "Lord of the Rings",
"Author": "J. R. R. Tolkien",
"Pages": 1137
}
```

### Read Books

- Endpoint: GET /books

Response

```
{
"ID": 1,
"CreatedAt": "2023-12-18T13:36:20.786+02:00",
"UpdatedAt": "2023-12-18T13:36:20.786+02:00",
"DeletedAt": null,
"Title": "Lord of the Rings",
"Author": "J. R. R. Tolkien",
"Pages": 1137
},
{
"ID": 2,
"CreatedAt": "2023-12-19T06:16:10.988+02:00",
"UpdatedAt": "2023-12-19T06:47:47.172+02:00",
"DeletedAt": null,
"Title": "Hobbit",
"Author": "J. R. R. Tolkien",
"Pages": 300
},
{
"ID": 4,
"CreatedAt": "2023-12-19T06:17:45.33+02:00",
"UpdatedAt": "2023-12-19T06:17:45.33+02:00",
"DeletedAt": null,
"Title": "Elric of Melniboné",
"Author": "Michael Moorcock",
"Pages": 752
} ...
```
### Edit Book

- Endpoint: PUT /books/{id}

- Example Payload:

```
{
  "title": "Silmarillion",
}
```

Response (Notice the creation and update timestamps)
```
{
  "ID": 2,
  "CreatedAt": "2023-12-19T06:16:10.988+02:00",
  "UpdatedAt": "2023-12-19T12:35:38.33+02:00",
  "DeletedAt": null,
  "Title": "Silmarillion",
  "Author": "J. R. R. Tolkien",
  "Pages": 300
}
```
