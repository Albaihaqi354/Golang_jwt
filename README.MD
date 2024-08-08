# Golang JWT

This project is a simple implementation of authentication and authorization using JSON Web Tokens (JWT) in Go. It includes JWT token creation and validation, as well as middleware to protect API endpoints.

## Features

- **Token Creation**: Generates JWT tokens with user claims.
- **Token Validation**: Verifies JWT tokens to ensure their validity and integrity.
- **Authentication Middleware**: Provides middleware to protect endpoints with JWT authentication

## Struktur Proyek

- `controllers/`: Contains the logic for handling HTTP requests.
- `helpers/`: Includes helper functions for creating and validating JWT tokens.
- `middleware/`: Contains middleware for authentication and authorization.
- `models/`: Defines the data models used in the application.
- `routes/`: Sets up API routing and applies middleware.

## Instalation

1. **Clone Repositori**:

   ```bash
   git clone https://github.com/Albaihaqi354/Golang_jwt.git
   cd Golang_jwt

2. **Instalasi Dependensi**
    ```bash
    go mod tidy

## Usage
To start using the project, follow these steps to clone the repository and install dependencies. You can then modify the code to fit your application's needs, such as customizing token claims or adding additional middleware.

## Contributing
Contributions are welcome! If you have suggestions for improvements or encounter any issues, please feel free to submit a pull request or open an issue.