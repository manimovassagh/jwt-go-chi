
# JWT Authentication with Go and Chi

A simple Go API that demonstrates JWT authentication using the [Chi](https://github.com/go-chi/chi) router and the [jwt-go](https://github.com/dgrijalva/jwt-go) library for secure token handling.

## Features
- **JWT Authentication**: Protects API routes with JWT-based authentication.
- **Middleware**: Utilizes Chi middleware for easy route handling and token verification.
- **Modular Code**: Organized for readability and scalability.

## Getting Started

### Prerequisites
- **Go** (version 1.16 or higher is recommended)
- **Git**

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/manimovassagh/jwt-go-chi.git
   cd jwt-go-chi
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the server**
   ```bash
   go run main.go
   ```

### Environment Variables

Define the following environment variables in a `.env` file or export them in your shell:
- `JWT_SECRET`: Secret key for signing JWT tokens.

Example `.env`:
```
JWT_SECRET=your_secret_key
```

## Usage

- **Generate a Token**: Access the `/login` endpoint with user credentials to receive a JWT token.
- **Access Protected Routes**: Use the JWT token in the `Authorization` header as `Bearer <token>` to access protected routes like `/protected`.

### Example Requests

1. **Login**:
   ```http
   POST /login
   Content-Type: application/json
   {
       "username": "your_username",
       "password": "your_password"
   }
   ```

2. **Access Protected Route**:
   ```http
   GET /protected
   Authorization: Bearer <your_jwt_token>
   ```

## Download

To download this repository directly, you can use this command:
```bash
curl -L https://github.com/manimovassagh/jwt-go-chi/archive/refs/heads/main.zip -o jwt-go-chi.zip
```

Then unzip and navigate into the project directory:
```bash
unzip jwt-go-chi.zip
cd jwt-go-chi-main
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments
- [Chi Router](https://github.com/go-chi/chi)
- [JWT-Go](https://github.com/dgrijalva/jwt-go)

---

Happy coding!
