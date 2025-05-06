# 歡迎來到 Cryptoo Pricing 首頁

[English version](http://cryptoo-pricing.com/readme.en)

## 前言

這是我第一個使用 Go 開發的專案，目的是練習後端架構設計、開發流程與部署操作。
我使用了 Docker Compose 來部署專案，確保運行環境的穩定及一致性，也使用了 Nginx 來做反向代理，模擬附載均衡場景。 
系統具備會員註冊與登入功能，並採用 JWT 驗證機制。加密貨幣的相關資料則透過 [CoinGecko](https://www.coingecko.com/) 提供的免費 API 查詢取得。  

由於這是一個初學練習性質的專案，若當中有任何設計不當或尚待改進之處，還請多多指教，謝謝！

## 技術棧

### 伺服器環境
- AWS EC2
- AWS Route 53

### 容器化技術
- Docker
- Docker Compose
- Nginx Reverse Proxy（模擬附載均衡）

### 資料庫
- **關聯式資料庫**：PostgreSQL  
- **快取儲存**：Redis（主要用於資料快取）

### 使用的主要函式庫與工具
- Web 框架：[Gin](https://github.com/gin-gonic/gin)  
- 依賴注入：[Google Wire](https://github.com/google/wire)  
- HTTP 客戶端：[Resty](https://github.com/go-resty/resty)  
- 驗證機制：JWT（JSON Web Token）

## 關於我

- 📧 Email：rever.developer@gmail.com  
- 🎂 Cake：[陳彥均 (Rever Chen)](https://www.cake.me/rever-dev_rever)  
- 💻 GitHub：[@dev-rever)](https://github.com/dev-rever/cryptoo-pricing)  
- 📘 API 文件（Swagger）：[cryptoo-pricing.com/docs](http://cryptoo-pricing.com/docs)
