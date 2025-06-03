# 💸 OptimalFlow - Simple Transfer API with JWT Auth and PostgreSQL

A lightweight backend service built in **Go (Fiber)** to manage users and securely transfer balance between them. Designed with clean architecture and Docker support.

---

## 🚀 Features

- ✅ JWT-based Authentication
- ✅ Register & Login endpoints
- ✅ Transfer balance between users
- ✅ Validate sufficient balance
- ✅ Atomic database operations (via GORM Transactions)
- ✅ Environment-based configuration (`.env`, `.env.local`)
- ✅ Docker & Docker Compose setup
- ✅ RESTful APIs with clear structure
- ✅ Unit tested core logic

---

## 🧪 API Endpoints

| Method | Endpoint         | Description                  
|--------|------------------|------------------------------
| POST   | `/users`         | Register new user            
| POST   | `/login`         | Login with email & password  
| GET    | `/users`         | List all users               
| GET    | `/users/:id`     | Get user by ID               
| POST   | `/transfer`      | Transfer balance             

> 💡 Use JWT token from `/login` response in the `Authorization: Bearer <token>` header

---

## 🧱 Project Structure

.
├── cmd/main.go # Entry point
├── internal/
│ ├── handler/ # HTTP handlers
│ ├── service/ # Business logic
│ ├── repository/ # DB interaction
│ ├── model/ # GORM models
├── pkg/middleware/ # JWT Middleware
├── .env # Docker environment
├── .env.local # Local development environment
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
└── README.md


---

## 🐳 Running with Docker

```bash
# Build & start PostgreSQL and app
docker compose up --build
App runs at: http://localhost:3001
PostgreSQL runs at: localhost:5433
🛠️ Running Locally (without Docker)

Make sure PostgreSQL is running on port 5432 locally.

# Run app locally
go run cmd/main.go
⚙️ Make sure .env.local is set to use localhost:5432
🧪 Running Tests

go test ./internal/...
🧰 Makefile Commands

make run        # Run the app locally
make test       # Run unit tests
make docker-up  # Start app + db with Docker Compose
make docker-down  # Stop all containers
📈 Scaling to 10x Traffic

To support a 10x increase in load:

🧩 Run multiple app instances with Docker Compose or Kubernetes
🔀 Add a load balancer (e.g., Nginx or cloud ALB)
🧠 Optimize PostgreSQL with connection pooling and indexing
🚀 Cache frequently accessed data using Redis
📬 Use message queues for async processing
📊 Add monitoring with Prometheus or Grafana
🔒 Enable rate-limiting to prevent abuse
📬 Example Login Response

{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
Use this in Postman or curl headers:

Authorization: Bearer <token>