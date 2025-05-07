# 歡迎來到 Cryptoo Pricing 首頁

[English version](http://cryptoo-pricing.com/readme.en)

## 🔗 快速連結

- **GitHub 原始碼庫**  
  瀏覽專案原始碼。  
  ➡️ [@dev-rever/cryptoo-pricing](https://github.com/dev-rever/cryptoo-pricing)

- **API 文件（Swagger）**  
  查看可用的 API 端點，並透過互動式介面進行測試。  
  ➡️ [cryptoo-pricing.com/docs](http://cryptoo-pricing.com/docs)

## 📌 專案簡介

這是我第一個使用 Go 開發的後端專案，目的是練習後端架構設計、開發流程與部署操作。

我使用 Docker Compose 部署整個專案，確保運行環境的穩定性與一致性，並透過 Nginx 反向代理模擬附載均衡的情境。

系統具備會員註冊與登入功能，並實作 JWT 驗證機制。加密貨幣的資料則是透過 [CoinGecko](https://www.coingecko.com/) 提供的免費 API 進行查詢。

由於這是一個初學練習性質的專案，若有任何設計不佳或可改進之處，還請不吝指教，感謝！

## ⚙️ 系統架構圖

下圖展示了整體系統架構，包括使用者互動流程、DNS 解析、EC2 部署環境、容器化服務與後端元件的分工。

```
┌───────────────────┐
│   Client / User   │
└──────┬────────────┘
       │  DNS Lookup
       ▼
 ┌───────────────┐
 │   Route 53    │
 └──────┬────────┘
        │
        ▼
 ┌──────────────────────┐
 │      AWS EC2         │
 │   (Linux Host VM)    │
 └──────┬───────────────┘
        │
        ▼
 ┌───────────────────┐
 │   Nginx (Docker)  │  ← Reverse Proxy
 └──────┬────────────┘
        │
        ▼
 ┌────────────────────────────────┐
 │     Go Backend API (Docker)    │
 │   ─ Gin / JWT / Wire / Resty   │
 └──────┬────────────────┬────────┘
        │                │
   ┌────▼──────┐    ┌────▼──────┐
   │PostgreSQL │    │   Redis   │
   │ (Docker)  │    │ (Docker)  │
   └───────────┘    └───────────┘
```

## 🛠️ 技術棧

### 伺服器環境
- AWS EC2
- AWS Route 53

### 容器化技術
- Docker
- Docker Compose
- Nginx 反向代理（模擬附載均衡）

### 資料庫
- **關聯式資料庫**：PostgreSQL  
- **快取儲存**：Redis（主要用於資料快取）

### 使用的主要函式庫與工具
- Web 框架：[Gin](https://github.com/gin-gonic/gin)  
- 依賴注入：[Google Wire](https://github.com/google/wire)  
- HTTP 客戶端：[Resty](https://github.com/go-resty/resty)  
- 驗證機制：JWT（JSON Web Token）

## 🙋 關於我

- Email：rever.developer@gmail.com  
- Cake：[陳彥均 (Rever Chen)](https://www.cake.me/rever-dev_rever)  
- GitHub：[@dev-rever](https://github.com/dev-rever)
