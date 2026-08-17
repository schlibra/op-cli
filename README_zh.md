# op-cli

[English](README.md) | [中文](README_zh.md)

`op-cli` 是一个运行在终端中的 [OpenList](https://github.com/OpenListTeam/OpenList) 客户端，支持交互式菜单和命令行两种使用方式，可完成服务器地址配置、用户认证、文件浏览、文件信息查询和文件下载。

## 功能

- 配置 OpenList 服务器 Base URL。
- 使用用户名和密码登录，并支持可选的 TOTP 验证码。
- 查询当前登录用户信息。
- 按路径列出文件和目录。
- 查看文件或目录的元数据。
- 使用进度条下载文件。
- 在交互式文件浏览器中浏览目录、查看文件信息并发起下载。
- 显示版本号、Git 提交和构建时间。
- 为 Linux、macOS、Windows、BSD 及其他类 Unix 平台构建静态二进制文件。

## 环境要求

- Go 1.26.6 或兼容版本；具体版本声明在 `go.mod` 中。
- 一个提供本客户端所需 API 的 OpenList 服务器。
- 使用交互式菜单时，需要支持终端输入。
- 仅使用 Makefile 构建目标时需要 GNU Make。

当前客户端调用以下 OpenList 接口：

| 操作 | 接口 |
| --- | --- |
| 登录 | `POST /api/auth/login` |
| 当前用户 | `GET /api/me` |
| 登出 | `POST /api/auth/logout` |
| 文件列表 | `POST /api/fs/list` |
| 文件信息 | `POST /api/fs/get` |

建议使用不带末尾斜杠的 Base URL，例如 `https://openlist.example.com`。命令行设置 URL 时不会自动移除末尾斜杠，客户端会直接拼接 API 路径。

## 安装

### 本地构建

```bash
mkdir -p dist
go build -o dist/op-cli .
```

Windows 下可以使用 `.exe` 后缀：

```powershell
New-Item -ItemType Directory -Force dist | Out-Null
go build -o dist/op-cli.exe .
```

### 构建全部目标平台

```bash
make all-platform
```

Makefile 还提供以下构建目标：

| 目标 | 说明 |
| --- | --- |
| `make current` | 构建当前主机平台，输出 `dist/op-cli-current`（Windows 下为 `.exe`），并写入版本元数据。 |
| `make all-platform` | 构建 Makefile 中声明的全部平台组。 |
| `make all` 或 `make` | 同时构建当前平台和全部平台目标。 |
| `make clean` | 删除 `dist/` 输出目录。 |

跨平台文件使用 `op-cli-<goos>-<goarch>` 命名。当前目标组包括：

- Linux：`amd64`、`386`、`arm64`、`arm`、`riscv64`、`ppc64`、`ppc64le`、`s390x` 和 `loong64`。
- macOS：`amd64` 和 `arm64`。
- Windows：`amd64`、`386` 和 `arm64`。
- FreeBSD、OpenBSD 和 NetBSD：以 `Makefile` 中声明的架构为准。
- DragonFly BSD、Solaris 和 illumos：以 `Makefile` 中声明的架构为准。
- Makefile 中还存在 Android 名称的目标；分发前请阅读[已知限制](#已知限制)。

初始目标集会生成例如以下文件：

```text
op-cli-linux-amd64
op-cli-darwin-amd64
op-cli-windows-amd64.exe
```

所有跨平台配方均使用 `CGO_ENABLED=0`。构建全部目标会比构建当前平台耗时更长。

## 快速开始

以下示例假设二进制文件名为 `./op-cli`。

```bash
# 设置服务器地址
./op-cli url https://openlist.example.com

# 登录；账号启用 TOTP 时，在末尾追加验证码
./op-cli login <用户名> <密码> [totp-code]

# 检查登录状态并查看根目录
./op-cli info
./op-cli ls /

# 查看并下载文件
./op-cli get /documents/report.pdf
./op-cli download /documents/report.pdf
```

不带参数启动可以进入交互式菜单：

```bash
./op-cli
```

交互式菜单提供以下功能：

- **Server setting**：查看并保存 OpenList Base URL。
- **Auth setting**：登录、登出和查询当前用户；登录表单支持可选 TOTP 验证码。
- **File browser**：浏览目录、返回上级目录、查看文件元数据并下载文件。
- **About program**：查看版本号、Git 提交和构建时间。

当前源码中的 CLI 登录参数下标检查仍有误，文档中的 CLI 登录命令可能触发 index-out-of-range；在依赖 CLI 登录前请使用交互式登录或先修复 `cmd/root.go`。

## 命令参考

| 命令 | 说明 |
| --- | --- |
| `url <base-url>` | 设置 OpenList 服务器 Base URL。 |
| `url get` 或 `url info` | 输出当前配置的 Base URL。 |
| `login <username> <password> [totp-code]` | 登录并保存服务器返回的令牌。 |
| `logout` | 登出并清除本地保存的令牌。 |
| `info` | 获取当前用户的用户名、基础路径、角色和权限。 |
| `ls [path]` | 列出目录内容；省略路径时默认为 `/`。 |
| `get [path]` | 查看文件或目录信息以及原始 URL；省略路径时默认为 `/`。 |
| `download [path]` | 下载文件到当前目录，并使用服务器返回的文件名保存；`dl`、`down` 是别名。 |
| `version` | 显示版本号、Git 提交和构建时间；`-v` 是别名。 |
| `help` | 显示内置用法说明。 |

客户端对已实现的命令也接受对应的 `-command` 和 `--command` 形式。

## 配置与认证

配置以 TOML 格式保存在：

```text
<用户主目录>/.config/op-cli/config.toml
```

Windows 中 `<用户主目录>` 是当前用户的 profile 目录，路径分隔符会按平台使用。文件内容示例：

```toml
base_url = "https://openlist.example.com"
token = "<access-token>"
```

首次运行时会自动创建配置文件。访问令牌以明文保存在本地，请妥善保护该文件，不要将其提交到版本库。

## 下载行为

`download` 会先请求文件信息，再下载服务端返回的 `raw_url`。文件保存到当前工作目录，文件名取自服务端返回值；如果同名文件已存在，会被覆盖。下载采用流式传输并显示进度条，目前不支持断点续传、校验和验证或自定义输出路径。

## 开发

可以使用以下命令格式化并测试项目：

```bash
gofmt -w main.go cmd/root.go model/*.go tui/*.go utils/*.go
go test ./...
make current
```

项目当前没有测试用例。Makefile 构建会通过链接参数写入 `Version`、`GitCommit` 和 `BuildTime`；直接使用 `go build` 不会写入这些值。

## 已知限制

- 客户端依赖上文列出的 OpenList 接口路径和响应结构。
- 底层请求模型包含密码、分页和刷新字段，但 CLI 尚未提供对应参数。
- 当前源码中的 CLI 登录参数下标检查有误，按文档执行 `login <username> <password>` 可能触发 index-out-of-range；请先使用交互式登录或修复 `cmd/root.go`。
- Makefile 中的 `android-*` 配方当前实际设置的是 `GOOS=linux`，生成的是带 Android 文件名的 Linux 二进制，并非原生 Android 二进制。
- 构建全部平台需要兼容的 Go 工具链，部分目标可能不受本地 Go 版本支持。

## 许可证

本项目采用 GNU General Public License v3.0（GPL-3.0）授权，完整许可证文本请参阅 [LICENSE](LICENSE)。
