# Salon & Beauty Parlour Management App

## Backend Task Plan (Go + Fiber + PostgreSQL)

_Single Developer • Backend-First • Multi-Tenant SaaS_

This document is the **single source of truth** for backend development.
It explains **what to build, in what order, and why**, so that:

- Future changes are easy
- A new contributor (or future me) can understand everything
- Backend architecture stays clean and scalable

---

## Core Product Decisions (Locked for MVP)

- Multi-tenant SaaS (every business is isolated by `business_id`)
- Backend-first development
- Single backend service
- Default language: **English**
- Optional language: **Hindi**
- Language selection is **client-driven** (header-based)
- No feature is 100% final — changes are allowed if documented here

---

## Tech Stack

- Language: **Go**
- Framework: **Fiber**
- Database: **PostgreSQL (Neon)**
- Auth: **JWT**
- API Style: **REST**
- Architecture: **Controller → Service → Repository**
- Localization: **Backend-driven messages (i18n-ready)**

---

## Folder Structure (Target)

```
salon-app-backend/
├── main.go
├── internal/
│ ├── config/
│ ├── db/
│ ├── models/
│ ├── controllers/
│ ├── services/
│ ├── routes/
│ ├── middleware/
│ ├── i18n/
│ └── utils/
├── docs/
├── README.md
├── Backend_Tasks.md
├── go.mod
└── .gitignore
```

---

## Phase 0 — Backend Foundation & Setup

### 0.1 Repository & Project Setup

- Initialize GitHub repository

- Add `.gitignore`, `README.md`

- Initialize Go module

- Keep `main.go` at project root

- Setup basic logging

- Prepare `.env` based config loading

### 0.2 Application Bootstrapping

- Initialize Fiber app

- Add `/health` endpoint

- Setup graceful shutdown

- Verify server starts correctly

---

## Phase 1 — Configuration & Database Layer

### 1.1 Configuration Management

- Environment variable loader

- App config struct

- DB config struct

- JWT config struct

### 1.2 Database Connection

- PostgreSQL connection setup

- Connection pooling

- Health check

- Central DB access layer

---

## Phase 2 — Multi-Tenant Database Design (CRITICAL)

### 2.1 Core Tables (All tenant data scoped by `business_id`)

- Admins

- Businesses

- Staff

- Services

- Customers

- Bookings

- Payments

### 2.2 Database Migrations

- Choose migration approach

- Create tables with:
  - Proper foreign keys

  - Indexes for performance

  - Unique constraints (phone per business)

- Booking time conflict-safe indexes

### 2.3 Seed Data (Optional)

- Seed admin

- Seed sample business

- Seed staff and services

---

## Phase 3 — Authentication & Authorization

### 3.1 Admin Authentication

- Signup

- Login

- Password hashing

- JWT generation

### 3.2 Staff Authentication

- Phone + PIN login

- PIN hashing

- JWT with role and business context

### 3.3 Role-Based Access Control

- Middleware for:
  - Admin-only routes

  - Staff-limited routes

- Enforce business isolation at API level

---

## Phase 4 — Localization (English + Hindi)

### 4.1 Localization Strategy (IMPORTANT)

- Default language: **English**

- Optional language: **Hindi**

- Language selected via request header:

`Accept-Language: en | hi`

- Backend controls:

- Error messages

- Success messages

- System responses

### 4.2 i18n Structure

- Create `internal/i18n/`

- Language files:

- `en.json`

- `hi.json`

- Key-based messages:

`booking.created`

`auth.invalid_credentials`

`staff.not_found`

### 4.3 Localization Middleware

- Read `Accept-Language` header

- Default to English if missing/invalid

- Inject language context into request

### 4.4 Usage Rules

- No hardcoded user-facing strings in controllers

- All responses use message keys

- Frontend does NOT translate backend errors

---

## Phase 5 — Business & Staff Management APIs

### 5.1 Business APIs

- Create business

- List businesses for admin

- Update business details

- Trial period tracking

### 5.2 Staff APIs

- Add staff

- Update staff

- Activate / deactivate staff

- Assign staff to business

- List staff per business

---

## Phase 6 — Service Management

### 6.1 Service APIs

- Create service

- Update service

- Enable / disable service

- List services per business

- Duration and pricing logic

---

## Phase 7 — Customer Management

### 7.1 Customer Logic

- Auto-create customer on booking

- Lookup by phone number

- Track visit history

- Track last visit

### 7.2 Customer APIs

- List customers

- View customer profile

- View booking history

---

## Phase 8 — Booking Engine (CORE LOGIC)

### 8.1 Booking Rules

- Booking flow:

`Pending → Admin Approved → Confirmed → Completed`

- Customer may select staff (optional)

- Admin assigns staff if missing

- Fixed daily working hours

### 8.2 Slot & Conflict Handling

- Service-duration-based slots

- Prevent overlapping bookings

- Staff availability checks

- Timezone-safe logic

### 8.3 Booking APIs

- Create booking

- Approve booking

- Reschedule booking

- Cancel booking (admin)

- Complete booking

---

## Phase 9 — Billing & Payments (MVP)

### 9.1 Billing Logic

- Generate bill from booking

- Payment modes:

- Cash

- UPI

- Gateway (future)

- Daily revenue calculation

### 9.2 Billing APIs

- Create payment

- Get bill for booking

- Daily revenue summary

---

## Phase 10 — WhatsApp Integration (Later Phase)

### 10.1 WhatsApp Setup

- Meta Cloud API

- Template approval

- Secure credential storage

### 10.2 WhatsApp APIs

- Booking confirmation

- Appointment reminder

- Manual message links

---

## Phase 11 — AI Integration (Optional & Controlled)

### 11.1 AI Scope

- Admin-only usage

- Offer suggestions

- Simple insights

### 11.2 Safety

- Rate limiting

- Max calls per business per day

- Safe fallback when AI fails

---

## Phase 12 — API Documentation

### 12.1 Swagger

- Swagger setup

- API annotations

- Auth header documentation

### 12.2 Versioning

- `/api/v1`

- Backward-compatible structure

---

## Phase 13 — Security, Stability & Testing

### 13.1 Security

- JWT validation

- Business data isolation

- Input validation

- SQL injection protection

### 13.2 Error Handling

- Centralized error responses

- Localized error messages

- Proper HTTP status codes

### 13.3 Testing & Readiness

- Manual testing

- Auth & role checks

- Booking conflict testing

- Production env separation

---

## Final Goal

- Production-ready Go backend

- Multi-tenant SaaS architecture

- Localization-ready system (English + Hindi)

- Strong backend portfolio project

- Easy future scaling (features, languages, platforms)
