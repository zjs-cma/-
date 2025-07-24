
1. 环境配置

1. **安装必要工具**
   - 确保已安装 Go 1.16+ 和 SQL Server
   - 安装 SQL Server 驱动：`go get github.com/denisenkom/go-mssqldb`



2. **连接字符串配置**
   - 准备类似这样的连接字符串：
   `server=localhost;user id=sa;password=yourpassword;database=FileService`

## 核心功能开发

1. **初始化GORM连接**
   - 使用 `gorm.io/driver/sqlserver` 驱动
   - 注意SQL Server特有的配置项，如连接超时

2. **模型定义**
   ```go
   type File struct {
     ID        string    `gorm:"primaryKey"`
     Name      string
     Path      string
     Size      int64
     CreatedAt time.Time
   }

