
# 💳 Fintech Event-Driven Microservices Platform

## **....................WILL UPDATE AS THE PROJECT PROGRESSES....................**

## 🧠 Overview
An event-driven digital wallet platform that enables users to send and receive money, powered by a ledger-based transaction system and real-time notifications.

The system is built using a microservices architecture, where services communicate asynchronously via **RabbitMQ (event bus)** for maximum scalability, decoupling, and resilience.

---

## 🏗️ Architecture

### 🔁 Event-Driven Core
- **Central Message Broker:** RabbitMQ

#### **Domain Events**
- `user.created`
- `payment.processed`
- `transaction.recorded`

---

## 🧩 Core Services

| Service               | Responsibility                          | Emits / Consumes                      | Stack        |
|----------------------|----------------------------------------|--------------------------------------|-------------|
| **User Service**     | Identity & Account Management          | Emits: `user.created`                | PostgreSQL  |
| **Payment Service**  | Transaction Validation & Processing    | Emits: `payment.processed`           | PostgreSQL  |
| **Transaction Service** | Immutable Ledger (Source of Truth) | Consumes: `payment.processed`<br>Emits: `transaction.recorded` | PostgreSQL  |
| **Notification Service** | Real-Time Updates & WebSockets     | Consumes: All events<br>Pushes: WebSocket / Email | Redis Pub/Sub |

---

## 🗄️ Data Architecture

We follow a **Database-per-Service pattern** to ensure strict ownership and prevent tight coupling.

- No shared databases  
- **Asynchronous Consistency:** All inter-service communication happens via events  

### Storage by Service

| Service                | Storage                          |
|------------------------|----------------------------------|
| User Service           | PostgreSQL (Users DB)           |
| Payment Service        | PostgreSQL (Payments DB)        |
| Transaction Service    | PostgreSQL (Ledger DB)          |
| Notification Service   | Redis (Pub/Sub)                 |

---

## ⚡ Real-Time Layer

The **Notification Service** acts as the bridge between backend events and the end-user.

- **Consume:** Picks up events from RabbitMQ  
- **Publish:** Relays updates to specific Redis channels  
- **Push:** WebSocket server pushes updates to clients instantly  

### Use Cases
- Live payment status updates  
- Instant balance refreshes  
- In-app user alerts  

---

## 🛡️ Reliability & Standards

### Consistency & Resilience
- **Retry Policies:** Implemented with exponential backoff  
- **Dead-Letter Queues (DLQ):** Handle unprocessable messages in RabbitMQ  
- **Idempotency:** Prevents double-charging in payment operations  
- **Graceful Shutdown:** Ensures no active transactions are dropped during deployments  

### Health & Observability
- **Endpoints:**  
  - `/health` (Liveness)  
  - `/ready` (Readiness)  

- **Structured Logging:** JSON format for easy parsing  
- **Correlation IDs:** Track requests across the system  

- **Telemetry:**  
  - OpenTelemetry + Jaeger (distributed tracing)  
  - Prometheus & Grafana (metrics)  

---

## ☁️ Deployment & CI/CD

- **Containerization:** Docker  
- **Orchestration:** Kubernetes (Deployments, Services, ConfigMaps, Secrets)  

### Pipeline
- **CI:**  
  - Linting  
  - Unit / Integration / Contract Testing  
  - Image Building  

- **CD:**  
  - Automated Kubernetes deployment  
  - Rollback strategies  

---

## 🧭 Getting Started

### Local Development

- Clone the repository
sh ```git clone <repo-url>```
sh ```cd fintechy-microservices```

- Start the entire stack (RabbitMQ, Postgres, Redis, Services)
sh ```docker-compose up --build```

## 🧾 License

This project is licensed under the MIT License.
