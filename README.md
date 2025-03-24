# TestDBMock

一个使用Go语言开发的数据库操作示例项目，展示了如何使用依赖注入和数据库模拟进行单元测试。

## 功能特点

- 使用SQLite作为数据库存储
- 采用依赖注入模式（使用Wire工具）
- 包含用户服务的CRUD操作
- 完整的单元测试示例
- 展示了如何使用mock进行数据库操作的测试

## 安装

1. 克隆项目
```bash
git clone <your-repository-url>
cd testDBMock
```

2. 安装依赖
```bash
go mod download
```

## 使用方法

1. 运行项目
```bash
go run main.go
```

2. 项目结构
```
.
├── internal/
│   ├── repository/    # 数据库操作层
│   ├── service/       # 业务逻辑层
│   ├── wire.go        # 依赖注入配置
│   └── wire_gen.go    # Wire生成的依赖注入代码
└── main.go           # 程序入口
```

## 测试

运行测试：
```bash
go test ./...
```

项目包含了完整的单元测试示例，展示了如何：
- 使用mock对象进行测试
- 测试数据库操作
- 测试业务逻辑

## 依赖

- Go 1.21+
- github.com/google/wire - 依赖注入工具
- github.com/mattn/go-sqlite3 - SQLite驱动
- github.com/stretchr/testify - 测试框架

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件