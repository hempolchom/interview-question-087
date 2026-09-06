# Full-Stack Authentication System (Angular + Go + PostgreSQL)

ระบบพิสูจน์ตัวตน (Authentication System) สำหรับสมัครสมาชิกและเข้าสู่ระบบ พัฒนาด้วย **Angular (Standalone Components)** สำหรับ Frontend และ **Go (Gin / GORM / PostgreSQL)** สำหรับ Backend

---

## 📂 System Documentation Dashboard

โปรเจกต์นี้มาพร้อมกับชุดเอกสารออกแบบระบบและแผนภาพ UML ในรูปแบบ HTML สามารถเปิดดูศูนย์รวมเอกสารได้ที่:

* **Dashboard URL:** `http://localhost:5500/documets/index.html` (หรือเปิดผ่าน Live Server ใน VS Code)
* **รายการเอกสารในระบบ:**
  * **API Specs:** `documets/api_document.html`
  * **Functional Spec (FSD):** `documets/functional_specification_document.html`
  * **Business Flow:** `documets/business_flow_document.html`
  * **Unit Test Strategy:** `documets/unit_test_document.html`
  * **System Diagrams:** Use Case, Activity, Sequence, Class, Component UML, DFD, ERD

---

## 💻 กรณีรันแบบ Manual (ไม่ใช้ Docker)

สำหรับผู้ที่ต้องการรันแยกบริการเพื่อทดสอบหรือพัฒนาแบบ Live-Reload บนเครื่องคอมพิวเตอร์โดยตรง

### ความต้องการของระบบ (Prerequisites)
* **Node.js:** v18.x ขึ้นไป และ `npm`
* **Go:** v1.20 ขึ้นไป
* **PostgreSQL:** v15.x หรือ v16.x (รันเป็น Local Service)

### โครงสร้างพอร์ตกรณีไม่ใช้ Docker (Local Setup)
| Service | Local Host / URL | Internal App Port | Description |
| :--- | :--- | :---: | :--- |
| **Angular Frontend** | http://localhost:4200 | `4200` | Angular Dev Server (`ng serve`) |
| **Go Backend API** | http://localhost:9087 | `9087` | Go REST API Server (`go run main.go`) |
| **PostgreSQL DB** | `localhost:5432` | `5432` | Local Database Service |

---

### ขั้นตอนการรันระบบแบบ Manual

#### 1. จัดเตรียม Database (PostgreSQL)
1. เปิด PostgreSQL ในเครื่อง และสร้าง Database ใหม่ชื่อ `interview_db`
2. ระบบจะทำการ **Auto-Migrate** สร้างตาราง `users` พร้อม Sequence `id` และ Unique Index บน `username` ให้โดยอัตโนมัติเมื่อสั่งรัน Go Backend

#### 2. ตั้งค่าและรัน Go Backend
1. เข้าไปที่โฟลเดอร์ Backend
   ```bash
   cd go-backend
   ```
2. สร้างหรือแก้ไขไฟล์ `.env` สำหรับเชื่อมต่อ Database ในเครื่อง:
   ```env
   PORT=9087
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=12345
   DB_NAME=interview_db
   JWT_SECRET=your_secret_key
   ```
3. ติดตั้ง Dependencies และสั่งรัน Server:
   ```bash
   go mod download
   go run main.go
   ```
   * Backend API พร้อมใช้งานที่ `http://localhost:9087`

#### 3. ตั้งค่าและรัน Angular Frontend
1. เข้าไปที่โฟลเดอร์ Frontend
   ```bash
   cd angular-frontend
   ```
2. ติดตั้ง Dependencies:
   ```bash
   npm install
   ```
3. ตรวจสอบไฟล์ `src/environments/environment.ts` ให้ชี้ API ไปยัง Backend Local:
   ```typescript
   export const environment = {
     production: false,
     apiUrl: 'http://localhost:9087/api/v1'
   };
   ```
4. สั่งรัน Angular Development Server:
   ```bash
   ng serve
   ```
   * เข้าใช้งานแอปพลิเคชันผ่านเบราว์เซอร์ที่ `http://localhost:4200`

---

## 🐳 กรณีรันด้วย Docker (Docker Compose)

เหมาะสำหรับผู้ที่ต้องการรันระบบทั้งหมดผ่าน Container โดยไม่ต้องลง Go, Node.js หรือ PostgreSQL ในเครื่อง

### โครงสร้าง Docker Architecture

| Service | Container Name | Internal Port | Exposed Port (Host) | Description |
| :--- | :--- | :---: | :---: | :--- |
| **Frontend** | `my_angular_frontend` | `4200` | `9070` | http://localhost:9070 |
| **Backend** | `my_go_backend` | `9087` | `9071` | http://localhost:9071 |
| **Database** | `my_postgres_db` | `5432` | `9072` | `localhost:9072` |

---

### การตั้งค่า Docker Configuration Files

#### 1. ไฟล์ `docker-compose.yml`

```yaml
version: '3.8'

services:
  postgres-db:
    image: postgres:16-alpine
    container_name: my_postgres_db
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 12345
      POSTGRES_DB: interview_db
    ports:
      - "9072:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  go-backend:
    build: ./go-backend
    container_name: my_go_backend
    ports:
      - "9071:9087"
    environment:
      - PORT=9087
      - DB_HOST=postgres-db
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=12345
      - DB_NAME=interview_db
    depends_on:
      - postgres-db

  angular-frontend:
    build: ./angular-frontend
    container_name: my_angular_frontend
    ports:
      - "9070:4200"
    depends_on:
      - go-backend

volumes:
  pgdata:
```

#### 2. ไฟล์ `go-backend/Dockerfile`

```dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 9087
CMD ["./main"]
```

#### 3. ไฟล์ `angular-frontend/Dockerfile`

```dockerfile
FROM node:18-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/dist/angular-frontend /usr/share/nginx/html
EXPOSE 4200
CMD ["nginx", "-g", "daemon off;"]
```

---

### ขั้นตอนการรันระบบด้วย Docker

1. **สั่งเปิดใช้งานระบบทั้งหมด (Build & Run)**
   ```bash
   docker-compose up -d --build
   ```

2. **เข้าใช้งานแอปพลิเคชันผ่าน Port ที่กำหนด**
   * **Frontend (Angular):** http://localhost:9070
   * **Backend API (Go):** http://localhost:9071
   * **PostgreSQL Database:** `localhost:9072` (User: `postgres`, Pass: `12345`, DB: `interview_db`)

3. **สั่งปิดการทำงานระบบ**
   ```bash
   docker-compose down
   ```

4. **สั่งลบข้อมูล Database ทั้งหมด (Reset Data)**
   ```bash
   docker-compose down -v
   ```

---

## ✨ ฟีเจอร์หลักของระบบ

* **Password Hashing:** เข้ารหัสรหัสผ่านด้วย `bcrypt` ก่อนบันทึกลงฐานข้อมูล
* **Session Management:** ใช้ **JWT (JSON Web Token)** ในการยืนยันตัวตน
* **Input Sanitization:** มี `SqlCleanDirective` กรองอักขระอันตราย (`'`, `"`, `;`, `<`, `>`) บน Frontend ทันทีขณะพิมพ์
* **Validation & UI UX:** 
  * รองรับเฉพาะตัวอักษรภาษาอังกฤษ ตัวเลข และขีด (`a-z`, `A-Z`, `0-9`, `_`, `-`) ในช่อง User
  * แสดงดอกจันสีแดง `*` สำหรับฟิลด์ที่จำเป็น
  * แสดงกรอบและพื้นหลังสีแดงอ่อนเมื่อไม่ได้กรอกข้อมูล หรือกรอกรหัสผ่านไม่ตรงกัน
  * ปุ่มสลับเปิด-ปิดดูรหัสผ่าน (Toggle Password Visibility)
  * แจ้งเตือนสถานะการทำงานด้วย **SweetAlert2**