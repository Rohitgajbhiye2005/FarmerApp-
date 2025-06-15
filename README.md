✅ README.md for go-jwt-backend
markdown
Copy
Edit
# 🐐 Go JWT Backend – Farmer App

This is a backend service for a Farmer App built using **Go**, **Gin**, **JWT authentication**, and **Supabase** as the database.

---

## 📦 Features

- 🔐 User Signup & Login with JWT
- 🛒 Add to Cart API
- 🧾 Cart tracking by user
- 🧰 RESTful architecture
- 🍪 JWT token via cookies (for auth-protected routes)

---

## 📁 Folder Structure

go-jwt/
├── controllers/
├── initializers/
├── middleware/
├── models/
├── routes/
├── main.go
└── go.mod

yaml
Copy
Edit

---

## 🧰 Prerequisites

- Go 1.20+ installed
- Supabase project created
- Supabase table for users & cart setup
- Git installed
- Postman or Thunder Client (for testing)

---

## ⚙️ Setup Instructions

### 1. Clone the project

```bash
git clone https://github.com/yourusername/go-jwt-backend.git
cd go-jwt-backend
2. Install dependencies
bash
Copy
Edit
go mod tidy
3. Setup Environment Variables
Create a .env file in the root:

env
Copy
Edit
PORT=3000
DB_URL=your_supabase_postgres_url
JWT_SECRET=your_very_secret_key
✅ Note: Ask the project owner for .env values or Supabase keys.

4. Run the server
bash
Copy
Edit
go run main.go
Server will run at:

arduino
Copy
Edit
http://localhost:3000
🔐 API Authentication Flow
Signup
POST /signup
Body: { "email": "test@example.com", "password": "123456" }

Login
POST /login
Returns: Sets JWT cookie

Add to Cart (Protected)
POST /cart
Requires JWT cookie from login response

🍪 How to Test Protected Routes
Using Postman / Thunder Client
Login first → Copy Set-Cookie from response

On Add-to-Cart request → Paste it in the Cookie tab:

ini
Copy
Edit
Authorization=eyJhbGciOi...<your token here>;
Set request body like:

json
Copy
Edit
{
  "product_id": 1,
  "quantity": 2
}
🧪 Supabase DB Schema
Users Table:
Column	Type
id	UUID (Primary Key)
email	Text
password	Text

CartItems Table:
Column	Type
id	BigInt (Auto)
user_id	UUID (FK to users.id)
product_id	Int
quantity	Int

🤝 Contributors
Rohit Gajbhiye – @roh1tgajbhiye

Your Friend – Add name after contributing!

📄 License
MIT License. Use freely, but give credit where due.

yaml
Copy
Edit

---

## ✅ Next Steps

After you add this file to your root:
```bash
touch README.md
# Paste the content
git add README.md
git commit -m "Added README with setup instructions"
git push
