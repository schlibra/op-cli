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
- 显示版本号、Git 提交和构建时间。
- 为 Linux、macOS 和 Windows 构建静态二进制文件。

## 环境要求

- Go 1.26.6 或兼容版本；具体版本声明在 `go.mod` 中。
- 一个提供本客户端所需 API 的 OpenList 服务器。
- 使用交互式菜单时，需要支持终端输入。

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
go build -o dist/op-cli .
```

Windows 下可以使用 `.exe` 后缀：

```powershell
go build -o dist/op-cli.exe .
```

### 构建全部目标平台

```bash
make build-all
```

二进制文件会写入 `dist/`，文件名如下：

```text
op-cli-linux-amd64
op-cli-linux-arm64
op-cli-darwin-amd64
op-cli-darwin-arm64
op-cli-windows-amd64.exe
```

当前 `make build` 配方仍是占位实现；构建单个平台请使用 `go build`，构建上述全部平台请使用 `make build-all`。

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

当前菜单可以配置服务器地址、进行认证和查看版本信息。首页虽然显示了 `File` 选项，但该选项尚未接入具体操作；文件浏览和下载请使用命令行命令。

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
```

项目当前没有测试用例。使用 `make build-all` 时，会通过链接参数写入 `Version`、`GitCommit` 和 `BuildTime`。

## 已知限制

- 客户端依赖上文列出的 OpenList 接口路径和响应结构。
- 底层请求模型包含密码、分页和刷新字段，但 CLI 尚未提供对应参数。
- 当前源码中的 CLI 登录参数下标检查有误，按文档执行 `login <username> <password>` 可能触发 index-out-of-range；在依赖 CLI 登录前请先修复 `cmd/root.go`。

## 许可证

本项目采用 GNU General Public License v3.0（GPL-3.0）授权，完整许可证文本请参阅 [LICENSE](LICENSE)。
