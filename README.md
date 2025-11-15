# Fyne窗口应用程序

一个使用Go语言和Fyne框架开发的跨平台GUI应用程序，展示了各种Fyne UI组件的使用方法。

## 项目概述

本项目是一个基于Fyne框架的GUI应用程序演示，旨在展示Fyne框架的各种UI组件和功能特性。应用程序包含了丰富的交互元素，如按钮、标签、输入框、表格、下拉菜单、单选框、复选框、进度条、滑块、列表、多行文本框和选项卡等。

## 功能特性

### UI组件展示
- **标签组件**：支持多种文本样式和对齐方式
- **输入框组件**：单行和多行文本输入
- **按钮组件**：支持图标和自定义点击事件
- **表格组件**：自定义样式和数据展示
- **下拉选择框**：选项列表和选择事件处理
- **单选框和复选框**：选项选择和状态管理
- **进度条和滑块**：数值显示和交互控制
- **列表组件**：动态数据展示
- **选项卡组件**：多页面内容组织
- **自定义主题**：蓝色调的深色主题

### 交互功能
- 按钮点击事件处理
- 文本输入实时反馈
- 选择组件状态变化响应
- 滑块数值动态更新

## 技术栈

- **编程语言**：Go 1.25.3
- **GUI框架**：Fyne v2.7.1
- **构建工具**：Go原生构建工具

## 安装说明

### 环境要求

1. **Go语言环境**：版本1.25.3或更高
   - 下载地址：[https://golang.org/dl/](https://golang.org/dl/)
   
2. **C编译器**：Fyne需要CGO支持
   - Windows推荐：TDM-GCC或MinGW-w64
   - 下载地址：[https://jmeubank.github.io/tdm-gcc/](https://jmeubank.github.io/tdm-gcc/)

### 克隆项目

```bash
git clone https://github.com/yuan71058/go_fyne.git
cd go_fyne
```

### 安装依赖

```bash
go mod tidy
```

## 构建说明

本项目提供了两种构建方式：

### 1. 使用批处理文件（Windows）

#### 快速构建GUI版本
```bash
.\build-gui.bat
```
- 输出文件：`fyne-app-gui.exe`
- 特点：无控制台窗口，适合作为发布版本

#### 使用交互式构建菜单
```bash
.\build.bat
```
提供以下选项：
1. 构建GUI版本（无控制台窗口）
2. 构建控制台版本
3. 运行GUI版本
4. 运行控制台版本
5. 清理构建文件
6. 退出

### 2. 使用命令行

#### 构建GUI版本（无控制台窗口）
```bash
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-H=windowsgui" -o fyne-app-gui.exe
```

#### 构建控制台版本
```bash
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -o fyne-app.exe
```

## 运行说明

### 直接运行
```bash
# 运行GUI版本
.\fyne-app-gui.exe

# 运行控制台版本
.\fyne-app.exe
```

### 开发模式运行
```bash
go run main.go
```

## 项目结构

```
go_fyne/
├── main.go                 # 主程序入口
├── go.mod                  # Go模块定义
├── go.sum                  # 依赖校验和
├── build.bat               # 交互式构建脚本
├── build-gui.bat           # GUI版本快速构建脚本
├── BUILD_README.md         # 构建详细说明
├── assets/                 # 资源文件目录
└── docs/                   # 项目文档
    └── fyne_window_app/    # 项目开发文档
```

## 开发指南

### 添加新组件

1. 在`main.go`中的`main`函数内创建新组件
2. 设置组件的属性和事件处理
3. 将组件添加到适当的容器中
4. 更新UI布局

### 自定义主题

项目使用自定义主题`myTheme`，可以通过修改以下方法来自定义外观：

```go
func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource
func (m myTheme) Size(name fyne.ThemeSizeName) float32
```

## 贡献指南

欢迎提交问题报告和功能请求！如果您想贡献代码，请遵循以下步骤：

1. Fork本项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建Pull Request

## 许可证

本项目采用MIT许可证 - 查看[LICENSE](LICENSE)文件了解详情。

## 相关资源

- [Fyne官方文档](https://fyne.io/develop/)
- [Fyne API参考](https://pkg.go.dev/fyne.io/fyne/v2)
- [Go语言官方网站](https://golang.org/)

## 联系方式

如有问题或建议，请通过以下方式联系：

- 提交Issue：[https://github.com/yuan71058/go_fyne/issues](https://github.com/yuan71058/go_fyne/issues)
- 项目主页：[https://github.com/yuan71058/go_fyne](https://github.com/yuan71058/go_fyne)