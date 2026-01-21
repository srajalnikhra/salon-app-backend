# Salon & Beauty Parlour Management App
## Backend Task Plan (Go + Fiber + PostgreSQL)
_Single Developer Execution Plan_

---

## Phase 0 — Backend Foundation & Setup

### 0.1 Repository & Project Setup
- Initialize GitHub repository
- Add `.gitignore`, `README.md`
- Initialize Go module (`go mod init`)
- Decide base folder structure
- Add environment config loader (`.env`)
- Setup basic logging strategy

### 0.2 Tech Stack Confirmation
- Backend: Go + Fiber
- Database: PostgreSQL (Neon)
- Auth: JWT
- API Style: REST
- Architecture: Controller → Service → Repository

---

## Phase 1 — Core Backend Architecture

### 1.1 Project Structure Setup
- Create folders:
  - `config/`
  - `db/`
  - `models/`
  - `controllers/`
  - `services/`
  - `routes/`
  - `middleware/`
  - `utils/`
  - `docs/`
- Setup `main.go`
- Setup Fiber app initialization
- Setup graceful shutdown

### 1.2 Database Connection
- PostgreSQL connection setup
- Connection pooling
- Health check endpoint
- DB connection error handling

---

## Phase 2 — Multi-Tenant Database Design (VERY IMPORTANT)

### 2.1 Core Tables Design
- Admins
- Businesses (tenant)
- Staff
- Services
- Customers
- Bookings
- Payments

### 2.2 Database Migrations
- Choose migration tool (Goose / sql-migrate / manual)
- Write migration files for:
  - Create tables
  - Foreign keys
  - Indexes
- Add:
  - `business_id` in all tenant-scoped tables
  - Unique constraints (phone per business)
  - Booking time indexes

### 2.3 Seed Data (Optional)
- Seed admin
- Seed sample business
- Seed services & staff

---

## Phase 3 — Authentication & Authorization

### 3.1 Admin Authentication
- Admin signup
- Admin login
- Password hashing
- JWT token generation
- Token refresh strategy (simple)

### 3.2 Staff Authentication
- Staff login using phone + PIN
- PIN hashing
- JWT generation for staff
- Role identification in token

### 3.3 Role-Based Access Control (RBAC)
- Admin-only routes
- Staff-limited routes
- Business isolation enforcement

---

## Phase 4 — Business & Staff Management APIs

### 4.1 Business APIs
- Create business
- List admin businesses
- Update business details
- Trial period logic

### 4.2 Staff APIs
- Add staff
- Update staff
- Activate / deactivate staff
- Assign staff to business
- Staff listing per business

---

## Phase 5 — Service Management

### 5.1 Service APIs
- Create service
- Update service
- Enable / disable service
- List services per business
- Service duration & price handling

---

## Phase 6 — Customer Management

### 6.1 Customer Logic
- Auto-create customer on booking
- Customer lookup by phone
- Customer visit history
- Last visit tracking

### 6.2 Customer APIs
- List customers per business
- View customer profile
- View booking history

---

## Phase 7 — Booking Engine (CORE LOGIC)

### 7.1 Booking Rules
- Pending → Admin approval → Confirmed
- Customer selects staff (optional)
- Admin assigns staff if missing
- Same working hours every day

### 7.2 Slot & Conflict Handling
- Service duration based slot calculation
- Prevent overlapping bookings
- Staff availability check
- Business time zone handling

### 7.3 Booking APIs
- Create booking
- Approve booking
- Reschedule booking
- Cancel booking (admin only)
- Complete booking

---

## Phase 8 — Billing & Payments (MVP)

### 8.1 Payment Logic
- Create bill from booking
- Payment modes: Cash / UPI / Gateway
- Store payment records
- Daily revenue calculation

### 8.2 Billing APIs
- Create payment
- Get booking bill
- Daily revenue endpoint

---

## Phase 9 — WhatsApp Integration (Later Phase)

### 9.1 WhatsApp Setup
- Meta Cloud API setup
- Template approval
- Environment variable storage

### 9.2 WhatsApp APIs
- Booking confirmation trigger
- Reminder trigger (scheduler)
- Manual message link generation

---

## Phase 10 — AI Integration (Optional / Controlled)

### 10.1 AI Scope
- Admin-only usage
- Offer suggestions
- Simple business insights

### 10.2 Safety
- Rate limiting
- Max calls per business per day
- Fail-safe fallback

---

## Phase 11 — API Documentation

### 11.1 Swagger Setup
- Swagger config
- API annotations
- Auth headers documentation

### 11.2 API Versioning
- `/api/v1/`
- Deprecation-ready structure

---

## Phase 12 — Security & Stability

### 12.1 Security
- JWT validation middleware
- Business data isolation
- Input validation
- SQL injection safety

### 12.2 Error Handling
- Centralized error responses
- Meaningful HTTP status codes
- Logging critical failures

---

## Phase 13 — Testing & Readiness

### 13.1 Manual Testing
- Auth flows
- Booking conflicts
- Role enforcement

### 13.2 Production Readiness
- Env separation
- Logging enabled
- DB backup strategy (basic)

---

## Final Goal

- Production-ready backend
- Clean Go architecture
- Multi-tenant SaaS experience
- Strong Golang backend portfolio project
