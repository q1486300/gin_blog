## Gin 實現 - 簡易版 Blog

使用 Golang 的 Gin 框架實現簡易版的部落格後台;\
資料庫使用 MySQL，透過 gorm 進行資料存取與映射。

## 目錄結構

```shell
gin_blog/
    ├   .gitignore
    ├   docker-compose.yml
    ├   Dockerfile
    ├   ginblog_backup.sql  資料庫備份檔案
    ├   go.mod              專案依賴
    ├   go.sum
    ├   main.go             主函數
    ├   README.md
    ├── api
    ├── config              專案配置
    ├── log                 日誌相關的程式碼
    ├── logs                日誌文件
    ├── middleware          中間件
    ├── model               資料模型層
    ├── routes              路由
        └── routes.go
    ├── uploads
    └── utils               專案公共工具庫
        ├   setting.go
        ├── err_msg
        └── validator
```

## 執行 && 部屬

1. 複製專案

```shell
git clone https://github.com/q1486300/gin_blog.git
```

2. 切換至專案根目錄

```shell
cd ./gin_blog
```

3. 建立 docker image

```shell
docker build -t gin_blog:latest .
```

4. 透過 docker compose 啟動容器

```shell
docker compose up -d
```

5. 啟動後即可訪問應用

http://{伺服器 IP}:3000/ \
預設綁定在 3000 port

6. 如需進入容器內操作資料庫

```shell
docker exec -it ginblog_mysql bash
mysql -u ginblog -p
```
預設資料庫用戶: ginblog\
預設資料庫密碼: admin123456


>容器相關配置可參考`./docker-compose.yml`文件\
專案相關配置可參考`./config/config.ini`文件

## 實現功能

1. 簡單的用戶權限管理設置
2. 用戶密碼加密儲存
3. 自定義文章分類
4. List 分頁
5. 上傳檔案
6. JWT 驗證
7. 自定義日誌功能
8. 跨域 CORS 設置

## 使用的工具、框架或套件

- Golang
    - Gin web framework
    - gorm
    - jwt-go
    - scrypt
    - logrus
    - gin-contrib/cors
    - go-playground/validator/v10
    - go-ini
- MySQL version:8.0.29

## API 文件

[https://documenter.getpostman.com/view/22826972/2s935hQ6Zf](https://documenter.getpostman.com/view/22826972/2s935hQ6Zf)