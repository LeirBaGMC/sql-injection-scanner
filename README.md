# 🛡️ SQL Injection & Vulnerability Scanner
> **Educational Security Tool** | Automated Detection Engine | Dockerized Architecture

![Go](https://img.shields.io/badge/Backend-Go_1.21-00ADD8?logo=go&logoColor=white)
![React](https://img.shields.io/badge/Frontend-React_Vite-61DAFB?logo=react&logoColor=black)
![Docker](https://img.shields.io/badge/Deploy-Docker_Compose-2496ED?logo=docker&logoColor=white)
![Status](https://img.shields.io/badge/Status-Academic_Research-success)

## 📋 Overview
This project is a modular vulnerability scanner designed to detect security flaws in web applications within controlled environments (DVWA). It employs a **multi-engine architecture** orchestrated by a high-performance Go backend, capable of identifying SQL Injections, exposed sensitive files, and insecure HTTP configurations.

**⚠️ Disclaimer:** *This tool is developed strictly for educational and defensive purposes. Usage against targets without prior mutual consent is illegal.*

---

## 🚀 Key Features
The scanning engine implements **Defense-in-Depth** analysis vectors:

### 💉 1. SQL Injection Engine
* **Error-Based Detection:** Identifies database syntax errors via payload injection (e.g., `' OR 1=1 --`).
* **Blind SQLi (Time-Based):** Detects vulnerabilities by analyzing server response latency using time-delay payloads (e.g., `SLEEP(5)` logic).

### 🔍 2. Infrastructure Reconnaissance
* **Sensitive File Exposure:** Scans for leaked configuration files (`/.git/HEAD`, `config.php.bak`, `/.env`).
* **Port Scanning:** TCP connectivity checks for critical services (SSH/22, MySQL/3306, FTP/21).
* **Header Analysis:** Audits missing security headers (`Content-Security-Policy`, `X-Frame-Options`, `HSTS`).

### 📊 3. Reporting & Visualization
* **PDF Generation:** Automated technical reports using `gofpdf`.
* **Real-Time UI:** React-based dashboard for scan monitoring and result visualization.

---

## 🛠️ Architecture

The system operates on a microservices-based architecture orchestrated via Docker Compose:

| Service | Technology | Role |
| :--- | :--- | :--- |
| **Backend** | Go (Gin) | Core logic, scan orchestration, and REST API. |
| **Frontend** | React + Vite | Interactive UI for scan management. |
| **Database** | MySQL 8.0 | Persistence for scan history and vulnerability logs. |
| **Target** | DVWA | Deliberately Vulnerable Web App for testing. |

---

## ⚡ Quick Start

### Prerequisites
* Docker & Docker Compose

### Deployment
1. **Clone the repository:**
   ```bash
   git clone [https://github.com/LeirBaGMC/sql-injection-scanner.git](https://github.com/LeirBaGMC/sql-injection-scanner.git)
   cd sql-injection-scanner
