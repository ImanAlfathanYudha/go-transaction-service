# Bank Statement Viewer - Go + Next.js

A small full-stack app to upload bank statement CSVs, view insights, and inspect transaction issues.  
Backend: Go (Gin), Frontend: React/Next.js  

---

## ⚙️ Backend Setup (Go)

### 1. Clone the repository
```bash
git clone https://github.com/ImanAlfathanYudha/go-transaction-service.git
cd go-transaction-service
```
### 2. Import the library
```bash
go mod tidy
```
If you're prefer using live reload, make sure live reload library installed. User use this lib:

*go install github.com/air-verse/air@latest
```bash
go install github.com/air-verse/air@latest
```
### 3. Run the program
You can run the program in normal way or using live-reload.
#### Running normally
```bash
go run main.go
```
#### Run using live-reload
```bash
air
```
### 3. Run the unit test
To run the unit test, you can use this command. 
*Notes: user only wrote the unit test on service part
```bash
go test ./services/transaction/service_test -v
```
## 🏗️ Architecture Decisions

This project follows a **modular and layered architecture** for clarity, scalability, and testability.

### 1. Layered Structure

| Layer | Folder | Responsibility |
|-------|---------|----------------|
| **Controller** | `/controllers` | Handles HTTP requests and responses using Gin. Validates input and calls service functions. |
| **Service** | `/services` | Contains business logic, data validation, and processing (e.g., CSV parsing, balance calculation). |
| **Repository** | `/repositories` | Handles data storage and retrieval. Abstracted through an interface for easier testing and mock usage. |
| **Model** | `/models` | Defines data structures used across layers (e.g., `Transaction`). |
| **Response** | `/response` | Standardizes HTTP responses for consistent API output. |

This separation allows clear responsibility between layers and makes it easier to maintain or test each component independently.

---

### 2. Design Principles

- **Dependency Injection**  
  Each service receives its repository through an interface (e.g., `IRepositoryRegistry`), allowing flexible swapping of implementations and easy mocking during unit tests.

- **Error Handling**  
  The service layer validates CSV format and data consistency before saving to the repository.  
  Errors (like missing fields or invalid formats) are wrapped with clear messages for debugging.

- **Testability**  
  Unit tests mock the repository layer using `testify/mock` to verify service logic without needing a real database.

- **Extensibility**  
  New features (e.g., analytics or additional endpoints like `/issues`) can be added by extending the service and controller layers without breaking existing functionality.

- **Framework Choice: Gin**  
  Gin is lightweight and high-performance, ideal for simple REST APIs.  
  It provides easy routing, middleware handling, and JSON response support.

---

### 3. Data Flow Example

1. **Controller** receives a request to `/upload` with a CSV file.  
2. **Service** parses the CSV, validates data, and converts it into transaction objects.  
3. **Repository** stores the data (mocked or real DB).  
4. **Response** returns success or error JSON to the client.
