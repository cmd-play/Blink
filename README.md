# Blink

A modern full-stack application built with Go, Gin, ReactJS, and Vite.

## Project Overview

Blink is a full-stack web application that combines a powerful backend API with a fast and responsive frontend.

## Tech Stack

### Backend
- **Language**: Go
- **Framework**: Gin (HTTP web framework)
- **Port**: 8080 (default)

### Frontend
- **Framework**: ReactJS
- **Build Tool**: Vite
- **Port**: 5173 (default Vite dev server)

## Project Structure

```
Blink/
├── backend/          # Go + Gin backend application
├── frontend/         # React + Vite frontend application
├── README.md         # Project documentation
└── LICENSE           # Project license
```

## Getting Started

### Prerequisites

- **Go** 1.21+ (for backend)
- **Node.js** 18+ and npm/yarn (for frontend)

## Backend

### Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Tidy go modules:
   ```bash
   go mod tidy
   ```

### Running the Development Server

Start the Gin server on port 8080:
```bash
go run main.go
```

The backend API will be available at `http://localhost:8080`

### Available Endpoints

- `GET /` - Welcome endpoint with API version info
- `GET /health` - Health check endpoint

### Building for Production

```bash
go build -o blink
./blink
```

## Frontend

### Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Run the development server:
   ```bash
   npm run dev
   ```

The frontend will be available at `http://localhost:5173`

### Building for Production

```bash
npm run build
```

The production build will be generated in the `dist/` directory.

## Development

To run both services simultaneously, open two terminal windows:

**Terminal 1 - Backend:**
```bash
cd backend
go run main.go
```

**Terminal 2 - Frontend:**
```bash
cd frontend
npm run dev
```

Both services will now be running and can communicate with each other.

## Build for Production

### Backend
```bash
cd backend
go build -o blink
```

### Frontend
```bash
cd frontend
npm run build
```

## API Documentation

(Add API documentation details here as your project develops)

## Contributing

### Commit Message Guidelines

When making commits, please use the following prefix to clearly indicate which part of the project your changes affect:

- **BACKEND:** - Use this prefix for all backend (Go/Gin) related changes
  ```
  BACKEND: Add user authentication endpoint
  BACKEND: Fix database connection issue
  ```

- **FRONTEND:** - Use this prefix for all frontend (React/Vite) related changes
  ```
  FRONTEND: Add login form component
  FRONTEND: Update navbar styling
  ```

- **DOCS:** - Use this prefix for documentation updates
  ```
  DOCS: Update README with setup instructions
  ```

### Example Commits

```bash
git commit -m "BACKEND: Create health check endpoint"
git commit -m "FRONTEND: Implement user dashboard page"
```

This helps maintain a clear project history and makes it easy to track changes across the full stack.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
