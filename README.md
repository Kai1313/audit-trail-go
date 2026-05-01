# Audit Trail Service

A high-performance, asynchronous logging microservice built with **Go** and **Echo v4**. This service acts as a centralized collector for audit logs from various microservices, utilizing a worker pool pattern to ensure high availability and non-blocking database writes.

## 🚀 Features

*   **Asynchronous Processing**: Uses Go channels to handle incoming logs. The API returns a `202 Accepted` immediately, while a background worker manages the MySQL insertion.
*   **Flexible Data Payload**: Supports `json.RawMessage` for `old_values` and `new_values`, allowing microservices to send any JSON structure without breaking the schema.
*   **Auto IP Tracking**: Automatically captures the sender's IP address if not explicitly provided in the payload.
*   **Swagger Integration**: Fully documented API with an interactive UI for testing.
*   **Layered Architecture**: Clean separation of concerns between Handlers, Services, and Repositories.

---

## 🛠️ Tech Stack

*   **Framework:** [Echo v4](https://github.com/labstack/echo)
*   **Database:** MySQL 8.0 (Supports `JSON` column types)
*   **Documentation:** [swaggo/swag](https://github.com/swaggo/swag)
*   **Config:** [joho/godotenv](https://github.com/joho/godotenv)

---

## 📋 Database Schema

The service is mapped to the following MySQL structure (referenced from `image_4d155c.png`):

| Column | Type | Nullable | Default |
| :--- | :--- | :--- | :--- |
| `id` | BIGINT (PK) | No | Auto Increment |
| `service_source` | VARCHAR(100) | No | - |
| `actor_id` | VARCHAR(100) | Yes | NULL |
| `action` | VARCHAR(50) | No | - |
| `entity_type` | VARCHAR(50) | Yes | NULL |
| `entity_id` | VARCHAR(100) | Yes | NULL |
| `old_values` | JSON | Yes | NULL |
| `new_values` | JSON | Yes | NULL |
| `ip_address` | VARCHAR(45) | Yes | NULL |
| `status` | ENUM('SUCCESS','FAILED')| Yes | 'SUCCESS' |
| `error_message` | TEXT | Yes | NULL |
| `request_id` | VARCHAR(100) | Yes | NULL |
| `created_at` | TIMESTAMP | Yes | CURRENT_TIMESTAMP |

---

## 📥 Getting Started

### 1. Clone & Install
```bash
git clone https://github.com/Kai1313/audit-trail-go.git
cd auditservice
go mod tidy
go mod vendor