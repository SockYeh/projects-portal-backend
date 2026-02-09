# Projects Portal Backend

Backend service for the Projects Portal, developed by IEEE Computer Society, VITC. This application is built using Go, the Echo framework, and PostgreSQL.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

-   **Go** (version 1.25 or higher) - [Download Go](https://go.dev/doc/install)
-   **PostgreSQL** - [Download PostgreSQL](https://www.postgresql.org/download/)
-   **Goose** (Database Migration Tool) - [Installation Guide](https://github.com/pressly/goose#installation)

    ```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
    ```

## Getting Started

Follow these steps to set up the project locally.

### 1. Clone the Repository

```bash
git clone https://github.com/ComputerSocietyVITC/projects-portal-backend.git
cd projects-portal-backend
```

### 2. Environment Configuration

Copy the example environment file to create your local configuration:

```bash
cp .env.example .env
```

Open the `.env` file and update the database credentials and other settings as needed:

```dotenv
APP_ENV=development
PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_user
DB_PASSWORD=your_postgres_password
DB_NAME=projects_portal
DB_SSLMODE=disable
```

### 3. Database Setup

1.  Start your PostgreSQL service.
2.  Create a new database named `projects_portal` (or whatever you set in `DB_NAME`).
3.  Apply the database migrations using `goose`.

    Run the migrations from the project root:

    ```bash
    # Syntax: goose -dir migrations postgres "CONNECTION_STRING" up
    goose -dir migrations postgres "user=compsoc password=secretpassword dbname=projects_portal sslmode=disable" up
    ```

    *Note: Make sure to replace the connection string values with your actual database credentials.*

### 4. Install Dependencies

Download the required Go modules:

```bash
go mod tidy
```

### 5. Run the Application

Start the development server:

```bash
go run main.go
```

The server should now be running at `http://localhost:8080` or any other port you set in your .env file.

## Project Structure

```
├── main.go           # Application entry point
├── cmd/
├── internal/
│   ├── config/       # Configuration logic (DB, env)
│   ├── handlers/     # HTTP request handlers
│   ├── logger/       # Logging setup
│   ├── middleware/   # Echo middleware
│   ├── models/       # Database models/structs
│   ├── repository/   # Data access layer
│   └── service/      # Business logic
├── migrations/       # SQL migration files
└── go.mod           # Go module definition
```

## Contributing Guidelines

Please follow these steps to contribute:

1.  **Clone the repo** locally.
2.  **Create a new branch** for your feature or bugfix (`git checkout -b feature/amazing-feature`).
3.  **Make your changes**.
4.  **Format your code**: Before committing, strictly ensure your code is formatted according to Go standards.
    ```bash
    go fmt ./...
    ```
5.  **Commit your changes** (`git commit -m 'feat: added amazing feature'`). Follow [**Conventional Commits**](https://www.conventionalcommits.org/)
6.  **Push to the branch** (`git push origin feature/amazing-feature`).
7.  **Open a Pull Request**.
