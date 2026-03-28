# 🛒 Ecommerce Frontend (Supportive UI)

This is the frontend component of the **Backend Developer Intern Assignment**. Built with **React.js** and **Vite**, this UI serves as a functional dashboard to demonstrate and interact with the Golang REST APIs.

## 🖼️ UI Preview & API Verification
> **Note:** The screenshot below demonstrates the functional dashboard alongside the browser **Inspect** console. This proves successful API communication, JWT handling in the Network tab, and role-based rendering.

![UI and Inspect Page](./ui-preview.png)


-----

## ✨ Key Features
* **Secure Authentication:** Dedicated flows for User Registration and Login.
* **JWT Management:** Automatic token storage in `localStorage`. 
* **Axios Interceptors:** Implemented a central API client that automatically attaches the `Bearer Token` to every outgoing request for protected routes.
* **Role-Based Access Control (RBAC):**
    * **Admin Mode:** Full access to view the product catalog and the "Add Product" administrative form.
    * **User Mode:** Restricted "Read-Only" view; the administrative form is hidden from the UI to prevent unauthorized interaction.
* **Dynamic Feedback:** Real-time success/error toast notifications based on API status codes (200, 201, 403, etc.).

## 🛠️ Tech Stack
* **Framework:** React.js (Vite)
* **Styling:** Tailwind CSS (v4)
* **Icons:** Lucide-React
* **State & API:** Axios, React Hooks (useState, useEffect)
* **Routing:** React Router DOM

---

## 🚀 Getting Started

### 1. Prerequisites
Ensure you have **Node.js** installed. The **Go Backend** must be running simultaneously on `http://localhost:8080`.

### 2. Installation
```bash
cd frontend
npm install