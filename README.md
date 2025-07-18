# Schedule Gateway Service

Dịch vụ này là API Gateway cho xác thực, phân quyền và quản lý resource, định tuyến request đến các microservice backend qua gRPC. Xây dựng bằng Go, Gin, tích hợp PostgreSQL, Redis, RabbitMQ.

## Cấu trúc dự án

- `cmd/server/main.go` – Điểm khởi động ứng dụng.
- `internal/` – Logic nghiệp vụ và module lõi:
  - `client/` – Wrapper client gRPC cho các backend service (auth, permission, role, token).
  - `controller/` – HTTP controller cho từng API endpoint.
  - `grpc/` – Code gRPC sinh từ protobuf.
  - `helper/` – Hàm tiện ích (quản lý resource, backup).
  - `initialize/` – Khởi tạo app: config, logger, router...
  - `middlewares/` – Gin middleware (log, error, check quyền).
  - `models/` – Struct dùng chung.
  - `routers/` – Định nghĩa nhóm route cho authentication/authorization.
- `pkg/` – Package dùng chung:
  - `loggers/` – Thiết lập logger zap.
  - `response/` – Chuẩn hóa response HTTP.
  - `settings/` – Cấu trúc config.
  - `utils/` – Tiện ích.
- `config/` – File cấu hình YAML.
- `storage/logs/` – Log file.
- `script/` – Script tiện ích (clear log, generate proto).
- `tests/` – Bộ test API (Postman).

## Tính năng chính

- **API Gateway**: Điểm vào trung tâm cho các API xác thực, phân quyền.
- **gRPC Integration**: Giao tiếp backend qua gRPC client.
- **Quản lý Role & Permission**: Endpoint quản lý vai trò, quyền, gán quyền/vai trò.
- **Middleware**: Log request, xử lý lỗi, kiểm tra quyền.
- **Configurable**: Đa môi trường qua YAML config.
- **Logging**: Log cấu trúc ra file và console.

## Bắt đầu

### Yêu cầu

- Go 1.20+
- PostgreSQL
- Redis
- RabbitMQ

### Thiết lập

1. **Clone repository**

   ```sh
   git clone <your-repo-url>
   cd schedule_gateway
   ```
2. **Cấu hình ứng dụng**
   - Sửa `config/dev.yaml` với thông tin database, Redis, RabbitMQ, server.
3. **Cài đặt dependencies**

   ```sh
   go mod tidy
   ```
4. **Khởi động dịch vụ**

   ```sh
   go run cmd/server/main.go
   ```
   Dịch vụ sẽ chạy trên cổng cấu hình trong file config.

### API Endpoints

- `GET /v1/api/checkStatus` – Health check.
- `POST /v1/api/auth/register` – Đăng ký tài khoản.
- `POST /v1/api/auth/login` – Đăng nhập.
- `POST /v1/api/auth/logout` – Đăng xuất.
- `POST /v1/api/auth/reset-password` – Đổi mật khẩu.
- `POST /v1/api/auth/forgot-password` – Quên mật khẩu.
- `POST /v1/api/auth/confirm-email` – Xác nhận email.
- `POST /v1/api/auth/confirm-forgot-password` – Xác nhận quên mật khẩu.
- `GET /v1/api/roles` – Lấy danh sách vai trò.
- `POST /v1/api/roles` – Tạo vai trò.
- `PUT /v1/api/roles/:id` – Cập nhật vai trò.
- `DELETE /v1/api/roles/:id` – Xóa vai trò.
- `PUT /v1/api/roles/:id/disable-or-enable` – Kích hoạt/vô hiệu hóa vai trò.
- `POST /v1/api/roles/assign-role-to-user` – Gán vai trò cho user.
- `GET /v1/api/permissions` – Lấy danh sách quyền.
- `POST /v1/api/permissions` – Tạo quyền.
- `PUT /v1/api/permissions/:id` – Cập nhật quyền.
- `DELETE /v1/api/permissions/:id` – Xóa quyền.
- `POST /v1/api/permissions/assign-permission-to-role` – Gán quyền cho vai trò.
- `POST /v1/api/token/refresh` – Làm mới token.
- `POST /v1/api/token/revoke` – Thu hồi token.

## Lưu ý phát triển

- **gRPC Clients**: Trong `internal/client/auth/`, mỗi service có client riêng.
- **Controllers**: Xử lý HTTP request và gọi gRPC client.
- **Routers**: Gom nhóm và đăng ký endpoint với middleware.
- **Logging**: Dùng zap cho structured logging.
- **Permission Checks**: Middleware kiểm tra quyền resource/action.
- **Scripts**: Dùng `script/clear_log.sh` để dọn log cũ.

## Lệnh hữu ích

- `go run cmd/server/main.go` – Khởi động gateway server.
- `bash script/clear_log.sh` – Xóa log cũ trong `storage/logs/`.

## Đóng góp

- Tuân thủ cấu trúc hiện tại khi phát triển tính năng mới.
- Sử dụng dependency injection cho controller/service/client.
- Viết test cho logic mới nếu có thể.

---

Nếu có thắc mắc, hãy đọc comment trong code hoặc liên hệ maintainer dự án.
