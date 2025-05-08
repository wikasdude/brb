# BRB Mid-Svc Platform

A microservice backend written in Go for managing service bookings, vendors, and users. It includes vendor summaries, booking logic, and integrates PostgreSQL without using an ORM like GORM.

## Features

- Vendor CRUD operations
- User management
- Service listing and management
- Booking creation with validation
- Vendor booking summary endpoint
- PostgreSQL integration via `database/sql`
- Clean layered architecture: Handler → Usecase → Repository → DB

## Tech Stack

- **Language:** Go (Golang)
- **Database:** PostgreSQL
- **API Tooling:** net/http
- **Documentation:** Swagger (optional)
- **Others:** Git, Docker (optional)

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/wikasdude/brb.git
cd brb
