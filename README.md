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

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the development server:
   ```bash
   go run main.go
   ```

The backend API will be available at `http://localhost:8080`

### Frontend Setup

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

## Development

To run both services simultaneously, open two terminal windows and follow the setup steps above for each.

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

(Add contribution guidelines here)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
