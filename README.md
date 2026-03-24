# Ghost Language

一个用 Go 编写的简洁、易学的解释型编程语言实现，专注于教学和学习编程语言设计。

## 更新速递

### v0.3.0 - 2026/3/24

新特性
- 范围运算符 ..，支持创建整数范围（如 1..5）
- 包含检查运算符 contains，用于检查元素是否存在
- 命名空间 namespace，支持模块化代码组织
- 多变量赋值与解构赋值
- foreach 循环语法，用于遍历可索引对象
- 科学计数法语法糖
- 二进制、八进制、十六进制数支持（0b、0o、0x 前缀）

改进
- 错误提示更清晰直观

## 项目图标

从 ghost v0.2.0 开始，Ghost Language 拥有了自己的图标：

![Ghost Language Icon](assets/image.png)

此图标位于 `assets/image.png`，setup 脚本会使用此图标为应用程序创建桌面快捷方式和文件关联。

## 快速开始

### 安装

有三种方式可以使用 Ghost：

1. 使用 PowerShell 脚本构建（推荐）：
```bash
# 克隆项目
git clone https://github.com/Ghost-Xiao/ghost-lang.git
cd ghost-lang

# 使用 PowerShell 脚本构建（Windows）
./build.ps1

# 构建后的可执行文件位于 bin/ 目录下
```

2. 使用 Go 工具直接构建：
```bash
go build -o ghost cmd/ghost/main.go
```

3. 从 Release 界面下载压缩包：
```bash
# 1. 访问项目的 GitHub Release 页面
# 2. 下载对应操作系统的压缩包
# 3. 解压压缩包
# 4. 运行解压目录中的 setup 脚本进行安装
```

### 运行 REPL

```bash
./ghost repl
# 或
./ghost -r
```

### 执行脚本文件

```bash
./ghost run script.gh
```

### setup 脚本使用指南

#### Windows 平台

**权限配置**
- 以管理员身份运行：所有脚本需要以管理员身份运行，否则可能会遇到权限不足的问题
- PowerShell 执行策略：确保 PowerShell 执行策略允许运行脚本
  - 可以通过以下命令查看当前执行策略：`Get-ExecutionPolicy`
  - 如果执行策略为 `Restricted`，可以临时设置为 `RemoteSigned`：`Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope Process`

**执行步骤**
1. 构建项目：运行 `./build.ps1` 构建项目
2. 安装设置：运行 `./setup.ps1` 进行安装
3. 脚本会自动处理图标安装和环境变量配置

**验证安装成功的方法**
- 在命令行中运行 `ghost -r` 启动 REPL（交互式解释器）
- 如果 REPL 成功启动并显示提示符，则表示安装成功

#### macOS 平台

**权限配置**
- 添加执行权限：需要给 `setup.command` 脚本添加执行权限
  - 运行 `chmod +x setup.command` 来添加执行权限

**执行步骤**
1. 构建项目：运行 `go build -o ghost cmd/ghost/main.go` 构建项目
2. 安装设置：双击 `setup.command` 文件或在终端中运行 `./setup.command`
3. 脚本会创建一个带图标的 `ghost.app` 并放入应用程序文件夹

**验证安装成功的方法**
- 在应用程序文件夹中找到 `ghost.app` 并双击启动
- 或在终端中运行 `ghost -r` 启动 REPL
- 如果 REPL 成功启动并显示提示符，则表示安装成功

#### Linux 平台

**权限配置**
- 添加执行权限：需要给 `setup.sh` 脚本添加执行权限
  - 运行 `chmod +x setup.sh` 来添加执行权限
- 管理员权限：某些操作（如安装图标到系统目录）可能需要 `sudo` 权限

**执行步骤**
1. 构建项目：运行 `go build -o ghost cmd/ghost/main.go` 构建项目
2. 安装设置：运行 `./setup.sh`（可能需要 `sudo ./setup.sh`）
3. 脚本会安装图标、创建 `.desktop` 文件并配置文件关联
4. 脚本会询问是否将 `ghost` 添加到环境变量中

**验证安装成功的方法**
- 在应用菜单中搜索 "Ghost" 并启动
- 或在终端中运行 `ghost -r` 启动 REPL
- 如果 REPL 成功启动并显示提示符，则表示安装成功

## 文档导航

为了提供更清晰、更详细的文档，我们将语法说明和示例移至 `docs/` 目录：

### 语法文档
- [表达式语法](docs/syntax/expressions.md) - 详细的表达式语法定义和示例
- [语句语法](docs/syntax/statements.md) - 详细的语句语法定义和示例
- [函数语法](docs/syntax/functions.md) - 详细的函数语法定义和示例
- [控制流语法](docs/syntax/control-flow.md) - 详细的控制流语法定义和示例

### 示例和架构
- [代码示例](docs/examples/fibonacci.md) - 斐波那契数列等示例代码
- [项目架构](docs/architecture.md) - 项目的整体架构和模块说明

### 其他文档
- [内置函数说明](docs/builtins.md) - 详细的内置函数文档

## 如何贡献

我们欢迎任何形式的贡献！无论是报告 bug、提出新功能建议，还是提交代码改进。

1. Fork 本项目
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个 Pull Request

## 测试

项目包含单元测试，可以使用以下命令运行：

```bash
go test ./...
```

这将运行所有包中的测试，包括词法分析器、语法分析器和解释执行器的测试。