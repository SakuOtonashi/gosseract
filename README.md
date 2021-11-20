# 说明

## 实现功能

参考 [DangoOCR](https://github.com/PantsuDango/DangoOCR)

## 开发参考

### 安装 vcpkg

参考 [vcpkg](https://github.com/microsoft/vcpkg)

### 安装 pkg-config

- `[vcpkg root]/vcpkg install pkgconf:x64-windows`

- 添加环境变量 `PKG_CONFIG_PATH` = `[vcpkg root]/installed/x64-windows/lib/pkgconfig`

- 环境变量 `Path` 增加路径 `[vcpkg root]/installed/x64-windows/bin`

- 配置 go 环境 `go env -w PKG_CONFIG="[vcpkg root]/installed/x64-windows/tools/pkgconf/pkgconf"`

### 安装依赖

`[vcpkg root]/vcpkg install tesseract:x64-windows`

### 构建相关

依赖的 [tesseract](https://github.com/tesseract-ocr/tesseract) 在运行时会读取对应的训练模型，从 [tessdata_fast](https://github.com/tesseract-ocr/tessdata_fast) 自行下载放入执行目录下的 `tessdata` 目录
