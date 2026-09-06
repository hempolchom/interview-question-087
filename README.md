# Full-Stack Authentication System (Angular + Go + PostgreSQL)

ระบบพิสูจน์ตัวตน (Authentication System) สำหรับสมัครสมาชิกและเข้าสู่ระบบ พัฒนาด้วย **Angular (Standalone Components)** สำหรับ Frontend และ **Go (Gin / GORM / PostgreSQL)** สำหรับ Backend

---

## ความต้องการของระบบ (Prerequisites)

* **กรณีใช้งานผ่าน Docker:**
  * Docker Desktop (ติดตั้งพร้อม Docker Compose)
* **กรณีไม่ใช้งาน Docker (Manual Setup):**
  * Node.js (v18 ขึ้นไป) และ npm
  * Go (v1.20 ขึ้นไป)
  * PostgreSQL (v14 ขึ้นไป)

---

## วิธีที่ 1: การติดตั้งและใช้งานด้วย Docker (แนะนำ)

เหมาะสำหรับผู้ที่ต้องการรันระบบทั้งหมด (Frontend, Backend, Database) ผ่าน คอนเทนเนอร์โดยไม่ต้องลง Environment ในเครื่อง

1. **เปิดใช้งานระบบทั้งหมด**

   ```bash
   docker-compose up -d --build
   ```

2. **เข้าใช้งานแอปพลิเคชัน**

   * **Frontend (Angular):** http://localhost:4200
   * **Backend API (Go):** http://localhost:8080
   * **PostgreSQL Database:** `localhost:5432`

3. **ปิดการทำงานระบบ**

   ```bash
   docker-compose down
   ```

---

## วิธีที่ 2: การติดตั้งและใช้งานแบบรันแยกด้วยตัวเอง (Manual Setup)

สำหรับผู้ที่ไม่ได้ใช้ Docker หรือต้องการพัฒนาโค้ดแบบ Live-Reload ในเครื่อง

### 1. จัดเตรียม Database (PostgreSQL)

1. สร้าง Database ใหม่ ชื่อ `interview_db`
2. ระบบจะทำการ **Auto-Migrate** สร้างตาราง `users` พร้อม Sequence `id` (Auto Increment) และ Unique Index บน `username` ให้โดยอัตโนมัติเมื่อสั่งรัน Go Backend

---

### 2. ตั้งค่าและรัน Go Backend

1. เข้าไปที่โฟลเดอร์ Backend
   ```bash
   cd go-backend
   ```

2. ตั้งค่าไฟล์ `.env` สำหรับเชื่อมต่อ Database
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=interview_db
   JWT_SECRET=your_secret_key
   PORT=8080
   ```

3. ติดตั้ง Dependencies และรัน Server
   ```bash
   go mod download
   go run main.go
   ```
   * Backend API พร้อมใช้งานที่ http://localhost:8080

---

### 3. ตั้งค่าและรัน Angular Frontend

1. เข้าไปที่โฟลเดอร์ Frontend
   ```bash
   cd angular-frontend
   ```

2. ติดตั้ง Dependencies
   ```bash
   npm install
   ```

3. รัน Angular Development Server
   ```bash
   ng serve
   ```
   * เปิดเบราว์เซอร์ไปที่ http://localhost:4200

---

## ฟีเจอร์หลักของระบบ

* **Password Hashing:** เข้ารหัสรหัสผ่านด้วย `bcrypt` ก่อนบันทึกลงฐานข้อมูล
* **Session Management:** ใช้ **JWT (JSON Web Token)** ในการยืนยันตัวตน
* **Input Sanitization:** มี `SqlCleanDirective` กรองอักขระอันตราย (`'`, `"`, `;`, `<`, `>`) บน Frontend ทันทีขณะพิมพ์
* **Validation & UI UX:** 
  * รองรับเฉพาะตัวอักษรภาษาอังกฤษ ตัวเลข และขีด (`a-z`, `A-Z`, `0-9`, `_`, `-`) ในช่อง User
  * แสดงดอกจันสีแดง `*` สำหรับฟิลด์ที่จำเป็น
  * แสดงกรอบและพื้นหลังสีแดงอ่อนเมื่อไม่ได้กรอกข้อมูล หรือกรอกรหัสผ่านไม่ตรงกัน
  * ปุ่มสลับเปิด-ปิดดูรหัสผ่าน (Toggle Password Visibility)
  * แจ้งเตือนสถานะการทำงานด้วย **SweetAlert2**