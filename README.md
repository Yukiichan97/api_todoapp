# API Todo App

Một RESTful API đơn giản để quản lý công việc (Todo) được viết bằng Go, sử dụng Echo framework.

## 📁 Cấu trúc thư mục

```
api_todoapp/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── server/
│   │   └── server.go            # Server configuration & middleware
│   └── tasks/
│       ├── model.go             # Task struct definition
│       ├── handler.go           # HTTP handlers
│       ├── service.go           # Business logic
│       └── repository.go        # Data access layer
└── README.md
```


## 📋 Data Model

```go
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Complete    bool   `json:"complete"`
}
```
## 🛠️ API Endpoints

### Tạo công việc mới
```http
POST /tasks
Content-Type: application/json

{
    "title": "Tasks",
    "description": "Tasks description",
}
```

**Response:**
```json
{
    "id": 1,
    "title": "Tasks",
    "description": "Tasks description",
    "complete": false
}
```

### Lấy tất cả công việc
```http
GET /tasks
```

**Response:**
```json
[
    {
        "id": 1,
        "title": "Tasks",
        "description": "Tasks description",
        "complete": false
    }
]
```

### Cập nhật công việc
```http
PUT /tasks/{id}
Content-Type: application/json

{
    "complete": true
}
```

### Xóa công việc
```http
DELETE /tasks/{id}
```

## 🚀 Cài đặt và chạy

### Yêu cầu
- Go 1.19 hoặc cao hơn
- Git

### Bước 1: Clone repository
```bash
git clone <repository-url>
cd api_todoapp
```

### Bước 2: Cài đặt dependencies
```bash
go mod tidy
```

### Bước 3: Chạy server
```bash
go build -o server.out ./cmd/server
./server.out
```

## 🧪 Test API

### Sử dụng curl

**Tạo task:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Task","description":"This is a test"}'
```

**Lấy tất cả tasks:**
```bash
curl http://localhost:8080/tasks
```

**Cập nhật task:**
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"complete":true}'
```

**Xóa task:**
```bash
curl -X DELETE http://localhost:8080/tasks/1
```
