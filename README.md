# Frontend Masters - Basics of Go

This repository contains hands-on projects and exercises from the [Frontend Masters "Basics of Go" course](https://frontendmasters.com/courses/go-basics/) taught by Maximiliano Firtman.

## 📚 Course Overview

The course provides a comprehensive introduction to Go programming language, covering fundamental concepts through practical projects. Topics include:

-   Variables, constants, and data types
-   Functions and control structures
-   Error handling and pointers
-   Structs, interfaces, and methods
-   Goroutines and channels for concurrency
-   HTTP servers and API development
-   File I/O operations
-   JSON parsing and web templates

## 🗂️ Project Structure

### 1. Calculator

**Location:** `Calculator/`

A simple command-line calculator that demonstrates:

-   Basic input/output with `fmt.Scanf` and `fmt.Scanln`
-   Control structures (`switch` statements)
-   Variable declarations and arithmetic operations

**Features:**

-   Supports add, subtract, multiply, and divide operations
-   Interactive command-line interface
-   Error handling for invalid operations

**Run:**

```bash
cd Calculator
go run main.go
```

### 2. CryptoMasters

**Location:** `cryptoMasters/`

A cryptocurrency price tracker that showcases:

-   HTTP API requests
-   JSON parsing with structs
-   Goroutines for concurrent API calls
-   WaitGroups for synchronization
-   Error handling

**Features:**

-   Fetches real-time cryptocurrency prices
-   Concurrent API calls for multiple currencies
-   Custom data structures for API responses

**Run:**

```bash
cd cryptoMasters
go run main.go
```

### 3. Files

**Location:** `files/`

File I/O operations demonstrating:

-   Reading and writing files
-   Package organization
-   Error handling with file operations
-   Custom utility functions

**Features:**

-   Text file reading and writing
-   Structured package organization with `utils/` package
-   Type definitions and data structures

**Run:**

```bash
cd files
go run main.go
```

### 4. InputOutput

**Location:** `InputOutput/`

Basic I/O operations and functions covering:

-   Function definitions and parameters
-   Data collections (arrays, slices, maps)
-   Variable scope and visibility

**Run:**

```bash
cd InputOutput
go run main.go
```

### 5. Museum

**Location:** `museum/`

A complete web server application featuring:

-   HTTP server with multiple endpoints
-   HTML templates with Go templating
-   Static file serving
-   API handlers for GET and POST requests
-   Structured project organization

**Features:**

-   Web interface with CSS styling
-   RESTful API endpoints
-   Template rendering with dynamic data
-   Static asset serving (images, CSS, JS)
-   Museum exhibition management

**Run:**

```bash
cd museum
go run main.go
```

Then visit `http://localhost:3001` in your browser.

### 6. Routines

**Location:** `routines/`

Concurrency examples demonstrating:

-   Goroutines for parallel execution
-   Channels for communication between goroutines
-   Synchronization patterns

**Run:**

```bash
cd routines
go run main.go
```

### 7. Structs

**Location:** `stucts/`

Advanced type system features including:

-   Struct definitions and methods
-   Interfaces and polymorphism
-   Type embedding and composition
-   Factory patterns

**Run:**

```bash
cd stucts
go run main.go
```

## 🚀 Getting Started

### Prerequisites

-   Go 1.19 or later installed on your system
-   Basic understanding of programming concepts

### Installation

1. Clone this repository:

```bash
git clone https://github.com/KiranNamawar/FEMBasicsOfGo.git
cd FEMBasicsOfGo
```

2. Navigate to any project directory and run:

```bash
go mod tidy  # Download dependencies
go run main.go  # Run the project
```

## 🛠️ Technologies Used

-   **Go 1.x** - Main programming language
-   **Standard Library Packages:**
    -   `fmt` - Formatted I/O operations
    -   `net/http` - HTTP client and server
    -   `html/template` - HTML templating
    -   `encoding/json` - JSON encoding/decoding
    -   `os` - Operating system interface
    -   `sync` - Synchronization primitives

## 📖 Learning Outcomes

After completing these projects, you will understand:

1. **Go Fundamentals**

    - Syntax and language features
    - Package system and module management
    - Variable declarations and type system

2. **Concurrency**

    - Goroutines for parallel execution
    - Channels for safe communication
    - Synchronization with WaitGroups

3. **Web Development**

    - Building HTTP servers
    - Creating RESTful APIs
    - Template rendering and static file serving

4. **Best Practices**
    - Error handling patterns
    - Code organization and package structure
    - Testing and debugging techniques

## 🔗 Resources

-   [Course Link](https://frontendmasters.com/courses/go-basics/)
-   [Go Official Documentation](https://golang.org/doc/)
-   [Go Tour](https://tour.golang.org/)
-   [Effective Go](https://golang.org/doc/effective_go.html)

## 👨‍💻 Author

**Kiran Namawar**

-   GitHub: [@KiranNamawar](https://github.com/KiranNamawar)

## 📄 License

This project is for educational purposes as part of the Frontend Masters course completion.

---

_Happy Learning! 🎉_
