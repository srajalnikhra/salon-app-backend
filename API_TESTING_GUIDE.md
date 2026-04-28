# 🎯 API Testing Guide - Quick Reference

**Base URL:** `http://localhost:3000/api/v1`  
**PgAdmin:** `http://localhost:5050` | DB: `salon_app`

---

## 🔐 Auth (No Token Needed)

| Endpoint | Method | Request | Check in DB |
|----------|--------|---------|------------|
| `/admin/signup` | POST | `{name, email, password, business_id}` | **admins** table |
| `/admin/login` | POST | `{email, password}` | Returns token ✅ |
| `/staff/login` | POST | `{phone, pin}` | Returns token ✅ |

---

## 📅 Bookings (Token Required ✅)

| Endpoint | Method | Request | Check in DB |
|----------|--------|---------|------------|
| `/admin/bookings` | POST | `{customer:{name, phone}, service_id, staff_id?, start_time, notes?}` | **bookings** & **customers** table |
| `/admin/bookings/{id}/approve` | PUT | `{}` | **bookings** - status → `confirmed` |
| `/admin/bookings/{id}/cancel` | PUT | `{}` | **bookings** - status → `cancelled` |
| `/bookings` | GET | - | Lists all bookings |

### ⚠️ Booking Rules:
- Staff must be assigned to service (use endpoint below)
- Staff must have availability set (use endpoint below)
- End time = start_time + service.duration (in minutes)

---

## 👨‍💼 Staff Services (Token Required ✅)

| Endpoint | Method | Request | Check in DB |
|----------|--------|---------|------------|
| `/admin/staff/{staffId}/services/{serviceId}` | POST | `{}` | **staff_services** table |
| `/admin/staff/{staffId}/services/{serviceId}` | DELETE | `{}` | Row deleted ✅ |

---

## ⏰ Staff Availability (Token Required ✅)

| Endpoint | Method | Request | Check in DB |
|----------|--------|---------|------------|
| `/admin/staff/{staffId}/availability` | POST | `{day_of_week: 0-6, start_time: "HH:mm", end_time: "HH:mm"}` | **staff_availabilities** table |

**Day codes:** 0=Sun, 1=Mon, 2=Tue, 3=Wed, 4=Thu, 5=Fri, 6=Sat

---

## 📊 Database Tables Cheat Sheet

| Table | Key Columns | Used For |
|-------|------------|----------|
| **admins** | id, email, password, is_active | Admin accounts |
| **staffs** | id, business_id, phone, pin, is_active | Staff members |
| **services** | id, business_id, name, duration, price | Services offered |
| **customers** | id, business_id, phone, name | Customer records |
| **bookings** | id, customer_id, service_id, staff_id, status, start_time, end_time | Booking records |
| **staff_services** | staff_id, service_id | Staff-service links |
| **staff_availabilities** | staff_id, day_of_week, start_time, end_time | Staff working hours |

---

## 🚀 Quick Test Flow

1. **Login:** `POST /admin/login` → Save token
2. **Assign Service:** `POST /admin/staff/2/services/1` → Check staff_services table
3. **Set Availability:** `POST /admin/staff/2/availability` → `{day_of_week: 3, start_time: "09:00", end_time: "18:00"}` → Check staff_availabilities table
4. **Create Booking:** `POST /admin/bookings` → `{customer:{name, phone}, service_id:1, staff_id:2, start_time:"2026-04-30T14:00:00Z"}` → Check bookings & customers tables
5. **Approve:** `PUT /admin/bookings/5/approve` → Status changes to "confirmed"

---

## ❌ Quick Troubleshooting

| Error | Solution |
|-------|----------|
| Staff not found | Verify in **staffs** table |
| Service not found | Verify in **services** table |
| Staff not assigned to service | Run staff services POST endpoint first |
| Staff not available at time | Run staff availability POST endpoint first |
| Unauthorized | Missing/invalid token - login first |

---

**v1.0 - April 27, 2026**

