# Hexagonal Architecture – Product Service (Go)

A small, production-leaning example of **Hexagonal Architecture (a.k.a. Ports and Adapters)** implemented in Go.  
It models a simple Product domain with a CLI adapter and a SQLite persistence adapter, emphasizing **testability, decoupling, and clear separation of concerns**.

---

## 🚀 What This Project Demonstrates
- A pure domain core with business rules and validation.
- **Ports (interfaces)** that define what the core expects from the outside world.
- **Adapters (CLI, DB)** that plug into those ports without leaking infrastructure concerns into the domain.
- A simple service layer orchestrating use cases like **Create, Get, Enable, and Disable**.

---

## 🎯 Why Hexagonal Here
- **Domain-centric**: Business logic lives in the core and does not depend on frameworks or databases.
- **Replaceable Adapters**: CLI or DB can be swapped (e.g., HTTP REST, gRPC, Postgres) without changing the domain code.
- **Testable**: The core can be unit-tested in isolation using interfaces and mocks.
- **Maintainable**: Clear boundaries reduce coupling and support incremental evolution.

---

## 🛠 Tech Stack
- **Language**: Go 1.24
- **CLI**: Cobra
- **DB**: SQLite (`github.com/mattn/go-sqlite3`)
- **Testing**: Testify
- **Dependency management**: Go modules

---

## 🧩 Core Concepts
- **Domain Model**: Product with attributes like ID, Name, Price, Status.
- **Business Rules**:
  - Enable requires `Price > 0`.
  - Disable requires `Price == 0`.
  - Status must be one of the allowed values (`enabled` / `disabled`).

- **Ports**:
  - Application service interface for product operations.
  - Repository interface for persistence.

- **Adapters**:
  - CLI adapter to invoke use cases via command-line.
  - DB adapter implementing the repository against SQLite.

---

## ⚡ Getting Started

### Prerequisites
- Go 1.24 installed
- SQLite runtime available on your system

### Install Dependencies
```bash
go mod download
```

### Database Setup
Create a SQLite database and initialize the products table:
```sql
CREATE TABLE IF NOT EXISTS products (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  price REAL NOT NULL,
  status TEXT NOT NULL
);
```

Apply it:
```bash
sqlite3 ./products.db < products.sql
```

### Build
```bash
go build -o bin/app
```

### Run (CLI)
The CLI exposes basic product operations:

```bash
# Create a product
./bin/app cli -a create -n "Sample Product" -p 29.90

# Get product by ID
./bin/app cli -i "<PRODUCT_ID>"

# Enable a product
./bin/app cli -a enable -i "<PRODUCT_ID>"

# Disable a product
./bin/app cli -a disable -i "<PRODUCT_ID>"
```

---

## ✅ Testing
Run all tests:
```bash
go test ./...
```

Validations covered:
- Enable requires `price > 0`
- Disable requires `price == 0`
- Status must be either `enabled` or `disabled`
- ID must be non-empty

---

## 🏗 Design Notes
- **Separation of Concerns**: Domain model and services are pure Go, adapters implement interfaces defined by the core.
- **Dependency Inversion**: Core depends on abstractions (ports), not concrete implementations.
- **Error Handling**: Business rule violations return descriptive errors.
- **Persistence**: Repository uses parameterized queries; implements upsert-like flow.
- **CLI**: Thin adapter mapping flags to use cases.

---

## 🔮 Extensibility
- Add a REST or gRPC adapter without touching domain rules.
- Replace SQLite with another database by implementing the repository interface.
- Introduce CQRS or event-driven extensions via new outbound ports.

---

## 📚 Example: Programmatic Usage
```go
package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, _ := sql.Open("sqlite3", "./products.db")
    repo := /* new repository adapter using db */
    service := /* new product service using repo */

    product, err := service.Create("Programmatic Product", 42.0)
    if err != nil {
        // handle error
    }

    _ = product
}
```

---

## 📂 Project Structure
- `application`: domain entities, interfaces (ports), services (use cases)
- `adapters/cli`: CLI adapter mapping flags to use cases
- `adapters/db`: repository adapter for SQLite
- `cmd`: Cobra commands and entrypoint wiring
- `main`: application bootstrap

---

## 🔐 Security & Configuration
- Use environment variables for config (e.g., DB path).
- Example:
  ```bash
  export APP_DB_PATH=./products.db
  ```
- Validate inputs at the edges (CLI/HTTP) in addition to domain rules.

---

## 📄 License
This project is licensed under the MIT License – see the [LICENSE](LICENSE) file for details.

---

## 📬 Contact
For questions or suggestions, feel free to open an issue or reach me at:  
📧 matheuslehnen@gmail.com
