# Các bước cài đặt server:

###Bước 1. Cài đặt Database
Cài đặt postgresSQL

Sau khi cài xong thì tạo user để truy cập vào database(có thể sử dụng default của postgres or tự tạo)

###Bước 2. 
Cài đặt Golang, link: https://golang.org/dl/

###Bước 3. Chạy source server

Chỉnh thông tin kết nối database trong dbconfig.yml và env.dev.yml

gõ lệnh sql-migrate up -config=dbconfig.yml -env="development" để tạo các table

###Bước 4 : Chạy server
Đi tới thư mục cmd/app và cmd/gen rồi gõ lệnh: go run main.go
(dùng lệnh cd cmd -> cd app rồi gõ lệnh)