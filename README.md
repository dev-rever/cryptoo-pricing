# Welcome to the Cryptoo Pricing Home Page

[中文版](http://cryptoo-pricing.com/readme.zh)

## 🔗 Quick Links

- **GitHub Repository**  
  Browse the source code on GitHub.  
  ➡️ [@dev-rever/cryptoo-pricing](https://github.com/dev-rever/cryptoo-pricing)

- **API Documentation (Swagger)**  
  View available endpoints and test them via an interactive interface.  
  ➡️ [cryptoo-pricing.com/docs](http://cryptoo-pricing.com/docs)

---

## 📌 Introduction

This is my first backend project built with Go, created to practice backend architecture design, development workflows, and deployment operations.  
I used Docker Compose to ensure consistency and stability of the runtime environment, and integrated Nginx as a reverse proxy to simulate load balancing scenarios.  
The system supports user registration and login, and implements JWT-based authentication. Cryptocurrency data is retrieved via the free public API provided by [CoinGecko](https://www.coingecko.com/).

As this is a beginner-level project, there may be design flaws or areas for improvement — feedback and suggestions are always welcome. Thank you!

---

## ⚙️ System Architecture

The following diagram illustrates the overall system architecture, including client interaction, DNS resolution, EC2 infrastructure, containerized services, and backend components.

```
┌───────────────────┐
│   Client / User   │
└──────┬────────────┘
       │  DNS Lookup
       ▼
 ┌───────────────┐
 │   Route 53    │
 └──────┬────────┘
        │
        ▼
 ┌──────────────────────┐
 │      AWS EC2         │
 │   (Linux Host VM)    │
 └──────┬───────────────┘
        │
        ▼
 ┌───────────────────┐
 │   Nginx (Docker)  │  ← Reverse Proxy
 └──────┬────────────┘
        │
        ▼
 ┌────────────────────────────────┐
 │     Go Backend API (Docker)    │
 │   ─ Gin / JWT / Wire / Resty   │
 └──────┬────────────────┬────────┘
        │                │
   ┌────▼──────┐    ┌────▼──────┐
   │PostgreSQL │    │   Redis   │
   │ (Docker)  │    │ (Docker)  │
   └───────────┘    └───────────┘
```

## 🛠️ Tech Stack

### Server Environment
- AWS EC2
- AWS Route 53

### Containerization
- Docker
- Docker Compose
- Nginx Reverse Proxy (to simulate load balancing)

### Databases
- **Relational Database**: PostgreSQL  
- **Cache Storage**: Redis (primarily for caching)

### Core Libraries & Tools
- Web Framework: [Gin](https://github.com/gin-gonic/gin)  
- Dependency Injection: [Google Wire](https://github.com/google/wire)  
- HTTP Client: [Resty](https://github.com/go-resty/resty)  
- Authentication: JWT (JSON Web Token)

---

## 🙋 About Me

- Email: rever.developer@gmail.com  
- Cake: [Rever Chen (陳彥均)](https://www.cake.me/rever-dev_rever)  
- GitHub: [@dev-rever](https://github.com/dev-rever)
