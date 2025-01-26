# Go CRUD with Gin, Gorm -example
This project demonstrates a simple CRUD API built with **Go**, **Gin** web framework, and **Gorm** ORM for managing a `User` table.

## Technologies Used

- **Go**: The programming language used for backend development.
- **Gin**: A high-performance web framework for Go, providing fast routing and middleware support.
- **Gorm**: An Object-Relational Mapping (ORM) library for Go, simplifying database operations.
- **Fx**: A Dependency Injection (DI) framework from **Uber**, designed to help with building and managing application components in Go. It provides a way to organize your code by separating concerns and making components easily testable and replaceable.

## Features

### 1. Gin Web Framework
**Gin** is a lightweight, fast, and flexible web framework for Go. It provides:
- High performance for building scalable web applications.
- Efficient routing and support for middleware.
- A clean and simple API for handling HTTP requests and responses.

### 2. Gorm ORM
**Gorm** simplifies interactions with databases by providing:
- An object-relational mapping system that abstracts SQL queries into Go structs.
- Automatic migrations that create or update database tables based on model definitions, removing the need for manual SQL schema updates.
- Built-in features like relationships, query building, and model associations.

### 3. JSON Binding with Gin
- **Gin's ShouldBindJSON** is used to automatically bind JSON data from the request body into Go structs.
- This functionality performs automatic validation based on struct tags, such as ensuring required fields are provided or ensuring the format of email addresses.
- If the data is invalid or cannot be parsed, **Gin** will return a helpful error message to the client.

### 4. Automatic Migrations with Gorm
- **Gorm's AutoMigrate** function allows automatic creation or updating of database tables based on Go struct models.
- When the application starts, **AutoMigrate** ensures that the necessary tables, columns, and relationships are correctly created or updated.
- This simplifies the setup process, as you don't need to manually write SQL migrations for simple schema changes.

### 5. Error Handling with Gin
- **Gin** provides an easy way to handle errors using custom functions.
- When handling requests, you can ensure consistent error responses by defining a centralized error handler.
- This improves maintainability and ensures that error responses are structured uniformly throughout the API.

### 6. Uber's Fx Framework
- **Fx** is a dependency injection framework for Go, designed to help you manage and organize complex applications. It enables you to structure your code by defining explicit dependencies and decoupling components, which makes your application easier to scale, maintain, and test.
- With **Fx**, you can manage lifecycle events, wire dependencies automatically, and cleanly separate concerns across your application, making it easier to build and manage large, production-ready systems.


## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/MaluMSiza/go-crud-gin.git
   cd go-crud-gin
   ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Run the application:
    ```bash
    go run main.go
    ```
The server will start on port 8080 by default.