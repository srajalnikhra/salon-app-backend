# Salon Management Backend

A backend API for a Salon & Beauty Parlour Management System built with Go and Fiber. The application is designed as a multi-tenant SaaS backend, where salon/business data is isolated using `business_id`.

---

## Tech Stack

- **Language:** Go 1.23.4
- **Web Framework:** Fiber v2
- **Database:** PostgreSQL
- **ORM:** GORM
- **Authentication:** JWT Authentication, bcrypt
- **Documentation:** Swagger
- **Configuration:** godotenv

---

## Features

- Admin signup and login
- Staff login
- JWT-based authentication
- Role-based authorization
- Multi-tenant business isolation
- Salon service management
- Staff management
- Staff-service assignment
- Staff availability management
- Customer creation/find during booking
- Booking creation
- Booking approval and cancellation
- Staff availability validation
- Booking conflict validation
- Swagger API documentation
- Automatic database migration
- Development seed data

---

## Project Architecture

The project follows a layered backend architecture:

```
Client
  ↓
Routes
  ↓
Middleware
  ↓
Controller
  ↓
Service
  ↓
Database
  ↓
PostgreSQL
```

### Main Layers

- **Routes:** Defines API endpoints and groups public, protected, and admin-only routes.
- **Middleware:** Validates JWT, extracts user information from the token, and restricts admin-only endpoints.
- **Controllers:** Handles HTTP requests and responses, validates request data, and calls the required service logic.
- **Services:** Contains business logic, handles booking rules, staff/service relationships, and performs database operations through GORM.
- **Models:** Defines database entities and relationships.

---

## Project Structure

```text
salon-app-backend/
│
├── main.go
├── go.mod
├── go.sum
├── .env
├── .gitignore
├── README.md
├── API_TESTING_GUIDE.md
├── Backend_Tasks.md
│
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
└── internal/
    ├── config/
    │   ├── app.go
    │   ├── config.go
    │   ├── db.go
    │   └── jwt.go
    │
    ├── controllers/
    │   ├── admin_auth_controller.go
    │   ├── booking_controller.go
    │   ├── responses.go
    │   ├── service_controller.go
    │   ├── staff_auth_controller.go
    │   ├── staff_availability_controller.go
    │   └── staff_service_controller.go
    │
    ├── db/
    │   ├── gorm.go
    │   ├── migrate.go
    │   └── seed/
    │       └── seed.go
    │
    ├── middleware/
    │   └── jwt_middleware.go
    │
    ├── models/
    │   ├── admin.go
    │   ├── booking.go
    │   ├── business.go
    │   ├── customer.go
    │   ├── service.go
    │   ├── staff.go
    │   ├── staff_availability.go
    │   └── staff_service.go
    │
    ├── routes/
    │   └── routes.go
    │
    ├── services/
    │   ├── admin_auth_service.go
    │   ├── booking_service.go
    │   ├── create_staff_service.go
    │   ├── list_booking_service.go
    │   ├── service_service.go
    │   ├── staff_auth_service.go
    │   ├── staff_availability_service.go
    │   └── staff_service_assignment.go
    │
    └── utils/
        ├── jwt.go
        └── password.go
```

---

## Database Design

### Main Entities

```
Admin
  │
  └── Business
        │
        ├── Staff
        │     ├── Staff Availability
        │     └── Staff Services
        │
        ├── Services
        │
        ├── Customers
        │
        └── Bookings
              ├── Customer
              ├── Service
              └── Staff
```

### Main Tables

| Table | Purpose |
|---|---|
| `admins` | Stores admin accounts |
| `businesses` | Stores salon/business information |
| `staffs` | Stores salon staff |
| `services` | Stores salon services |
| `customers` | Stores customers |
| `bookings` | Stores appointment bookings |
| `staff_services` | Maps staff members to services |
| `staff_availabilities` | Stores staff working hours |

---

## API Endpoints

**Base URL:** `http://localhost:3000/api/v1`

### Authentication

| Method | Endpoint | Access |
|---|---|---|
| `POST` | `/admin/signup` | Public |
| `POST` | `/admin/login` | Public |
| `POST` | `/staff/login` | Public |

### Bookings

| Method | Endpoint | Access |
|---|---|---|
| `POST` | `/admin/bookings` | Admin |
| `GET` | `/bookings` | Admin / Staff |
| `PUT` | `/admin/bookings/:id/approve` | Admin |
| `PUT` | `/admin/bookings/:id/cancel` | Admin |

### Services

| Method | Endpoint | Access |
|---|---|---|
| `POST` | `/admin/services` | Admin |
| `GET` | `/services` | Admin / Staff |
| `GET` | `/services/:id` | Admin / Staff |
| `PUT` | `/admin/services/:id` | Admin |
| `DELETE` | `/admin/services/:id` | Admin |

### Staff

| Method | Endpoint | Access |
|---|---|---|
| `POST` | `/admin/staff` | Admin |
| `POST` | `/admin/staff/:staffId/services/:serviceId` | Admin |
| `DELETE` | `/admin/staff/:staffId/services/:serviceId` | Admin |
| `POST` | `/admin/staff/:staffId/availability` | Admin |

---

## Authentication

Protected endpoints require a JWT token passed in the header:

```http
Authorization: Bearer <your_jwt_token>
```

### Token Payload Examples

**Admin Token Payload:**
```json
{
  "admin_id": 1,
  "business_id": 1,
  "role": "admin"
}
```

**Staff Token Payload:**
```json
{
  "staff_id": 1,
  "business_id": 1,
  "role": "staff"
}
```

The authentication middleware validates the token and stores user context in Fiber.

---

## Booking Flow & Rules

### Flow Diagram

```
Create Booking
      ↓
Validate Service
      ↓
Calculate End Time
      ↓
Check Staff-Service Assignment
      ↓
Check Staff Availability
      ↓
Check Booking Conflict
      ↓
Find/Create Customer
      ↓
Create Booking
      ↓
Pending
      ↓
Admin Approves
      ↓
Confirmed
```

### Booking Rules

1. The selected service must exist.
2. If a staff member is selected, they must be assigned to the requested service.
3. The staff member must be available during the requested timeframe.
4. Existing bookings are checked for time conflicts.
5. Booking end time is automatically calculated using service duration.
6. Customers are identified or created automatically using their phone number.

---

## Staff Management

### Staff Availability

Configured using `day_of_week`, `start_time`, and `end_time`.

**Day Codes:**
- `0` = Sunday
- `1` = Monday
- `2` = Tuesday
- `3` = Wednesday
- `4` = Thursday
- `5` = Friday
- `6` = Saturday

**Example Request Body:**
```json
{
  "day_of_week": 1,
  "start_time": "09:00",
  "end_time": "18:00"
}
```

### Staff-Service Assignment

A staff member must be assigned to a service before they can be assigned to bookings for that service.

**Example:**
- **Staff:** Rahul
- **Services:** Haircut, Shave, Hair Spa
- Relationship stored in `staff_services` table.

---

## Environment Variables

Create a `.env` file in the project root:

```env
APP_NAME=Salon Management Backend
APP_ENV=development
APP_PORT=3000

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=salon_app

JWT_SECRET=your_secret_key
JWT_EXPIRY_DAYS=7
```

> **Note:** Never commit real credentials or JWT secrets to version control.

---

## Prerequisites & Installation

### Prerequisites

- **Go:** 1.23+ (`go version`)
- **Database:** PostgreSQL
- **Tools:** Git

### Installation Steps

1. **Create the PostgreSQL Database:**
   ```sql
   CREATE DATABASE salon_app;
   ```

2. **Clone Repository & Navigate:**
   ```bash
   git clone https://github.com/srajalnikhra/salon-app-backend.git
   cd salon-app-backend
   ```

3. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

4. **Configure Environment:** Create `.env` based on the configuration above.

---

## Run the Application

Start the server:
```bash
go run main.go
```
The server will run on: `http://localhost:3000`

### Health Check

**Endpoint:** `GET /health`

**Example Response:**
```json
{
  "success": true,
  "message": "Server is healthy",
  "data": {
    "app": "Salon Management Backend",
    "env": "development",
    "db": "connected"
  }
}
```

---

## Swagger Documentation

Swagger UI is available at:
`http://localhost:3000/swagger/index.html`

Use Swagger to:
- Browse API endpoints and models
- Execute request tests directly
- Authorize requests with JWT tokens

---

## Database Migration & Seed Data

### Automatic Migration
The application automatically runs GORM `AutoMigrate()` on server startup to maintain table schema.

### Seed Data
Development seed data is populated on startup:

- **Admin Account:**
  - **Email:** `admin@salon.com`
  - **Password:** `admin123`
- **Demo Business:**
  - **Name:** Demo Salon
  - **Phone:** `9999999999`
  - **Address:** Demo Address
  - **Timezone:** Asia/Kolkata
- **Demo Services:**
  - Haircut — 30 minutes — ₹200
  - Shave — 15 minutes — ₹100

*(For local testing only. Do not use default credentials in production.)*

---

## API Testing Flow

```
1. Start PostgreSQL
        ↓
2. Start Go server
        ↓
3. Login as Admin
        ↓
4. Copy JWT token
        ↓
5. Create Staff
        ↓
6. Assign Service to Staff
        ↓
7. Set Staff Availability
        ↓
8. Create Booking
        ↓
9. Approve Booking
        ↓
10. View Bookings
```

---

## Development Commands

| Task | Command |
|---|---|
| Run Application | `go run main.go` |
| Sync Dependencies | `go mod tidy` |
| Build Executable | `go build` |
| Run Unit Tests | `go test ./...` |
| Regenerate Swagger Docs | `swag init` |

---

*Built with Go + Fiber + PostgreSQL.*