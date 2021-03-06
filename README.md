# 说明

## 实现功能

参考 [DangoOCR](https://github.com/PantsuDango/DangoOCR)

## 开发参考

### 安装 vcpkg

- 参考 [vcpkg](https://github.com/microsoft/vcpkg)

- 环境变量 `Path` 增加路径 `[vcpkg root]/installed/x64-windows/bin`

### 安装 pkg-config

- Linux 直接安装 pkg-config，Windows 下用使用 `[vcpkg root]/vcpkg install pkgconf:x64-windows`

- 配置 go 环境（Linux不用配置） `go env -w PKG_CONFIG="[vcpkg root]/installed/x64-windows/tools/pkgconf/pkgconf"`

- 添加环境变量 `PKG_CONFIG_PATH` = `[vcpkg root]/installed/x64-windows/lib/pkgconfig`

### 安装依赖

- [tesseract](https://github.com/tesseract-ocr/tesseract) 和 [leptonica](https://github.com/DanBloomberg/leptonica) 构建依赖于 [sw](https://github.com/SoftwareNetwork/sw)，请提前配置完成

- `[vcpkg root]/vcpkg install tesseract:x64-windows`

### 构建相关

依赖的 [tesseract](https://github.com/tesseract-ocr/tesseract) 在运行时会读取对应的训练模型，从 [tessdata_fast](https://github.com/tesseract-ocr/tessdata_fast) 自行下载放入执行目录下的 `tessdata` 目录
