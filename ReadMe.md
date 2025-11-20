# HRMS – Student Management & Attendance Reporting System

HRMS is a backend API designed to help faculties and organizers efficiently manage student data, track attendance, and automate weekly/monthly attendance reports.

The system consists of **three core components**:
- **Students API** — Manage student records  
- **Attendance API** — Track and fetch attendance  
- **Cron Job Service** — Dispatch scheduled attendance reports  

---

## 1. Students API

This API provides endpoints for creating, retrieving, updating, and deleting student records.

### **Endpoints**

| Method | Endpoint           | Description                                   |
|--------|---------------------|-----------------------------------------------|
| POST   | `/students`         | Create a new student                          |
| GET    | `/students`         | List all students                             |
| GET    | `/students/:id`     | Fetch a specific student                      |
| PUT    | `/students/:id`     | Update student details                        |
| DELETE | `/students/:id`     | Delete a student                              |

---

## 2. Attendance API

This API allows marking and retrieving attendance for students.

### **Endpoints**

| Method | Endpoint                       | Description                                      |
|--------|---------------------------------|--------------------------------------------------|
| POST   | `/attendance/mark`              | Mark attendance for the student                  |
| GET    | `/attendance/:student_id`       | Fetch historical attendance of a student         |

## Swagger documentation Link:
{BASE_URL}/swagger/index.html#/

---

## 3. Cron Job Service

A standalone service responsible for:

- Generating **weekly attendance reports**
- Generating **monthly attendance reports**
- Dispatching attendance summaries over email

This service runs on a schedule using a cron expression.

---

# Setup

Follow these steps to run the project locally.

---

## **1. Prerequisites**

- Install the latest version of **Go (Golang)**
- Install **MySQL Server**

---

## **2. Clone the Repository**

```bash
git clone https://github.com/kapil-puranik05/HRMS
```

## 3. Setup Environment Variables

Create a `.env` file inside both modules:

- `/server`
- `/cron`

### Server Module `.env`

| Variable        | Description               |
|-----------------|---------------------------|
| DB_USER         | Database username         |
| DB_PASSWORD     | Database password         |
| DB_HOST         | Database host             |
| DB_PORT         | Database port             |
| DB_NAME         | Database name             |
| TEST_EMAIL      | Login email for testing   |
| TEST_PASSWORD   | Login password for testing|
| JWT_KEY         | JWT signing key           |
| PORT            | Server port               |

---

### Cron Module `.env`

| Variable     | Description                        |
|--------------|------------------------------------|
| MAIL         | Gmail address used for sending reports |
| PASSWORD     | 16-digit Gmail App Password        |
| DB_USER      | Database username                  |
| DB_PASSWORD  | Database password                  |
| DB_HOST      | Database host                      |
| DB_PORT      | Database port                      |
| DB_NAME      | Database name                      |

---

## 4. Run the Server Module

```bash
cd server
go mod tidy
go run ./cmd
```
This will start the backend server.

## 5. Run the Cron Module

```bash
cd cron
go mod tidy
go run ./cmd
```
This will start the cron job.

# Steps to run test cron job
1. Insert a student payload into the database with your actual email, using the Student API.  
2. Mark the attendance of this student entity.  
3. Now remove the comment of the CronTest() Method in cron/cmd/main.go
4. Run the command
```bash
cd cron
go mod tidy
go run ./cmd
```

