# Go POS API

นี่คือโปรเจกต์ API สำหรับระบบ Point of Sale (POS) ของร้านค้าที่เขียนด้วยภาษา Go ซึ่งสามารถใช้เช็ค stock สินค้า, ตัด stock, ขายสินค้า, ชำระเงิน, ออกใบเสร็จรับเงิน, แยกร้านค้าสาขา, แยกพนักงานแต่ละสาขา และสรุปยอดขายได้

## การรันโปรเจกต์

### ขั้นตอนที่ 1: การติดตั้ง Docker

โปรดติดตั้ง Docker บนเครื่องของคุณถ้ายังไม่ได้ทำ

### ขั้นตอนที่ 2: การเริ่มต้น

1. คล๊อนโปรเจกต์จาก GitHub:

    ```bash
    git clone https://github.com/yourusername/go-pos.git
    ```

2. เข้าไปในไดเร็กทอรีของโปรเจกต์:

    ```bash
    cd go-pos
    ```

3. สร้างไฟล์ .env และกำหนดค่า DSN ของฐานข้อมูล PostgreSQL ของคุณ:

    ```bash
    cp .env.example .env
    ```

4. รันคำสั่ง docker-compose เพื่อสร้างและเริ่มต้นบริการ Docker:

    ```bash
    docker-compose up --build
    ```

### ขั้นตอนที่ 3: การเข้าใช้งาน API

โปรเจกต์จะเริ่มต้น API ที่ http://localhost:8080

คุณสามารถใช้ Postman หรือเบราว์เซอร์เพื่อเข้าถึงเส้นทางต่างๆของ API ได้

## การใช้งาน API

### Products API

- `GET /api/v1/products`: เรียกดูรายการสินค้าทั้งหมด
- `POST /api/v1/products`: เพิ่มสินค้าใหม่
- `PUT /api/v1/products/:id`: แก้ไขสินค้าตาม ID
- `DELETE /api/v1/products/:id`: ลบสินค้าตาม ID

### Orders API

- `POST /api/v1/orders`: สร้างคำสั่งซื้อใหม่
- `GET /api/v1/orders/:id`: เรียกดูรายละเอียดคำสั่งซื้อตาม ID
- `POST /api/v1/orders/:id/pay`: ชำระเงินสำหรับคำสั่งซื้อตาม ID
- `GET /api/v1/sales-summary`: สรุปยอดขายทั้งหมด

### Authentication API

- `POST /api/v1/auth/register`: ลงทะเบียนพนักงานใหม่
- `POST /api/v1/auth/login`: เข้าสู่ระบบเพื่อรับ Token

## หมายเหตุ

โปรดตรวจสอบไฟล์ `.env` เพื่อกำหนดค่าการเชื่อมต่อฐานข้อมูลของคุณอย่างถูกต้อง
