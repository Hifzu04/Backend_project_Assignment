# 🚀 Scalable Ecommerce REST API (Golang)

This project is a high-performance, secure backend system built for the **Backend Developer Intern Assignment**. It features a modular architecture, JWT-based authentication with Role-Based Access Control (RBAC), and a scalable MongoDB integration.

---

## 🛠️ Tech Stack
* **Language:** Go (Golang 1.22+)
* **Framework:** Gin Gonic (HTTP Web Framework)
* **Database:** MongoDB (NoSQL)
* **Security:** JWT (JSON Web Tokens), Bcrypt (Password Hashing)
* **Validation:** Go-Playground Validator
* **Tools:** Godotenv, Postman

---

## ✨ Core Features
* **Authentication:** User registration and login with secure password hashing.
* **Authorization (RBAC):** Middleware-level restriction for `Admin` vs. `User` roles.
* **CRUD Operations:** Full product management (Create, Read, Update, Delete).
* **Security:** * CORS Middleware for frontend integration.
    * Input sanitization and structural validation.
    * Protected routes requiring valid Bearer Tokens.
* **API Versioning:** Structured under `/api/v1` for future-proofing and scalability.

---

## 📂 Project Structure
```text
/backend
├── Controller/      # Request handling logic (Auth & Products)
├── Middleware/      # JWT Authentication & RBAC (Admin-only checks)
├── Models/          # MongoDB Schemas & Data Structures
├── Database/        # DB Connection & Collection Initialization
├── Utils/           # Helper functions (JWT generation, Hashing)
├── docs/            # Postman Collection & API documentation
├── main.go          # Entry point & Route definitions
└── .env             # Environment configuration




