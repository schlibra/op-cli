# op-cli

[English](README.md) | [中文](README_zh.md)

`op-cli` 是一个运行在终端中的 [OpenList](https://github.com/OpenListTeam/OpenList) 客户端，同时支持交互式菜单和命令行操作，可用于配置服务器、用户认证、文件浏览、元数据查询和文件下载。

## 功能

- 配置和查看 OpenList 服务器 Base URL。
- 使用用户名和密码登录，支持可选的 TOTP 验证码。
- 查询当前用户信息并登出。
- 按路径列出文件和目录。
- 查看文件或目录的元数据。
- 使用实时进度条下载文件。
- 在交互式文件浏览器中浏览目录、查看文件信息并发起下载。
- 显示程序版本、Git 提交和构建时间。
- 根据 Makefile 中声明的平台目标构建二进制文件，包括通过 Android NDK 构建原生 Android 二进制。

## 环境要求

- Go 1.26.6 或兼容的 Go 工具链；具体版本声明在 `go.mod` 中。
- 一个提供本客户端所需 API 的 OpenList 服务器。
- 使用交互式菜单时，需要支持终端输入。
- 只有使用 Makefile 构建目标时才需要 GNU Make。
- 构建 `android-*` 目标时，需要将 Android NDK LLVM 工具链加入 `PATH`。

当前客户端调用以下 OpenList 接口：

| 操作 | 接口 |
| --- | --- |
| 登录 | `POST /api/auth/login` |
| 当前用户 | `GET /api/me` |
| 登出 | `POST /api/auth/logout` |
| 文件列表 | `POST /api/fs/list` |
| 文件信息 | `POST /api/fs/get` |

建议使用不带末尾斜杠的 Base URL，例如 `https://openlist.example.com`。客户端会直接拼接 API 路径，不会自动移除末尾斜杠。

## 安装

### 本地构建

Linux、macOS、BSD 或其他类 Unix 系统：

```bash
mkdir -p dist
go build -o dist/op-cli .
```

Windows：

```powershell
New-Item -ItemType Directory -Force dist | Out-Null
go build -o dist/op-cli.exe .
```

### 使用 Make 跨平台构建

```bash
make current       # 构建当前主机平台
make all-platform  # 构建所有平台组
make               # 等同于 make all
make clean         # 删除 dist/
```

`make current` 会将版本号、Git 提交和构建时间写入 `dist/op-cli-current`（Windows 下为 `.exe`）。跨平台文件使用 `op-cli-<goos>-<goarch>` 命名。非 Android 跨平台配方使用 `CGO_ENABLED=0`；Android 配方使用 `GOOS=android` 和 `CGO_ENABLED=1`。

Makefile 当前声明的平台组如下：

| 平台组 | 目标架构 |
| --- | --- |
| Linux | `amd64`、`386`、`arm64`、`arm`、`riscv64`、`ppc64`、`ppc64le`、`s390x`、`loong64` |
| macOS | `amd64`、`arm64` |
| Windows | `amd64`、`386`、`arm64` |
| FreeBSD | `amd64`、`386`、`arm64`、`arm`、`riscv64` |
| OpenBSD | `amd64`、`386`、`arm64`、`arm` |
| NetBSD | `amd64`、`386`、`arm64`、`arm` |
| Android | `amd64`、`386`、`arm64`、`arm` |
| 其他 Unix | DragonFly BSD `amd64`、Solaris `amd64`、illumos `amd64` |

构建全部目标会比构建当前平台花费更长时间。`android-*` 配方使用各架构对应的 Android NDK Clang 编译器，默认 API 级别为 `21`。将 NDK LLVM 的 `bin` 目录加入 `PATH` 后，可以按需覆盖 API 级别：

```bash
make android ANDROID_API_LEVEL=24
```

`make all-platform` 包含 Android 目标，因此同样需要 Android NDK 编译器。生成的原生 Android 二进制使用 `op-cli-android-<goarch>` 命名。

### Android NDK 配置

Android 配方会根据目标架构查找对应的 NDK Clang 启动器：

| 架构 | 编译器启动器 |
| --- | --- |
| `amd64` | `x86_64-linux-android<api>-clang` |
| `386` | `i686-linux-android<api>-clang` |
| `arm64` | `aarch64-linux-android<api>-clang` |
| `arm` | `armv7a-linux-androideabi<api>-clang` |

配置步骤如下：

1. 在 Android Studio 的 SDK Manager 中安装 **NDK (Side by side)**，也可以使用 Android SDK 命令行工具安装 NDK。
2. 找到已安装的 NDK 目录，例如 `<android-sdk>/ndk/<version>`。
3. 将对应主机的 LLVM `bin` 目录加入 `PATH`。该目录位于 `<ndk-root>/toolchains/llvm/prebuilt/` 下，常见主机目录名包括 `windows-x86_64`、`linux-x86_64` 和 `darwin-x86_64`（请以实际 NDK 包中的目录为准）。

PowerShell 示例：

```powershell
$env:ANDROID_NDK_HOME = "$env:LOCALAPPDATA\Android\Sdk\ndk\<version>"
$env:Path = "$env:ANDROID_NDK_HOME\toolchains\llvm\prebuilt\windows-x86_64\bin;$env:Path"
Get-Command aarch64-linux-android21-clang
```

Linux 或 macOS 示例：

```bash
export ANDROID_NDK_HOME="$ANDROID_SDK_ROOT/ndk/<version>"
export PATH="$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/<host-name>/bin:$PATH"
command -v aarch64-linux-android21-clang
```

将验证命令中的 `21` 替换为计划使用的 API 级别。确认编译器可用后，可以构建单个架构或全部 Android 目标：

```bash
make android-arm64
make android ANDROID_API_LEVEL=24
```

Windows 下请使用 Git Bash、MSYS2、WSL 或其他提供 Makefile 所需 POSIX Shell 工具的环境执行这些配方。请在运行 `make` 的同一个 Shell 中设置 `ANDROID_NDK_HOME` 和 `PATH`；在另一个 PowerShell 窗口中设置的环境变量不会自动继承。如果找不到编译器，请检查 NDK 主机目录是否已加入 `PATH`、API 级别后缀是否匹配，以及架构对应的启动器名称是否正确。

## 快速开始

以下示例假设二进制文件名为 `./op-cli`（`.\op-cli.exe`）。

1. 设置 OpenList 服务器地址：

   ```bash
   ./op-cli url https://openlist.example.com
   ```

2. 使用命令行登录；如果账户启用了 TOTP，请在末尾添加验证码：

   ```bash
   ./op-cli login <username> <password> [totp-code]
   ```

   也可以运行 `./op-cli`，然后选择 **Auth setting** > **Login user** 使用交互式登录表单。

3. 列出根目录、查看文件信息或下载文件：

   ```bash
   ./op-cli ls /
   ./op-cli get /documents/report.pdf
   ./op-cli download /documents/report.pdf
   ```

运行 `./op-cli help` 可查看内置用法说明。不带参数运行程序时始终进入交互式菜单。

## 命令参考

| 命令 | 说明 |
| --- | --- |
| `url <base-url>` | 设置 OpenList 服务器 Base URL。 |
| `url get` 或 `url info` | 输出当前配置的 Base URL。 |
| `login <username> <password> [totp-code]` | 登录并保存服务器返回的令牌。 |
| `logout` | 登出并清除本地保存的令牌。 |
| `info` | 显示当前用户的用户名、基础路径、角色和权限。 |
| `ls [path]` | 列出目录内容；路径省略时默认为 `/`。 |
| `get [path]` | 显示文件或目录元数据及原始 URL；路径省略时默认为 `/`。 |
| `download [path]` | 将文件下载到当前目录，并使用服务器返回的文件名；`dl` 和 `down` 是别名。 |
| `version` | 显示版本号、Git 提交和构建时间；`-v` 是别名。 |
| `help` | 显示内置用法说明。 |

命令分发器也接受对应的 `-command` 和 `--command` 写法，包括下载命令的别名。

## 配置与认证

配置以 TOML 格式保存在：

```text
<用户主目录>/.config/op-cli/config.toml
```

Windows 下，`<用户主目录>` 是当前用户的 profile 目录，路径分隔符会按平台变化。首次使用时会自动创建该文件：

```toml
base_url = "https://openlist.example.com"
token = "<access-token>"
```

访问令牌以明文保存在本地。请妥善保护该文件，切勿将其提交到版本库。

## 下载行为

`download` 会先请求 `/api/fs/get` 获取文件元数据，再将返回的 `raw_url` 下载到当前工作目录。输出文件名来自服务器；如果同名文件已存在，`os.Create` 会覆盖它。下载采用流式传输并显示进度条，目前不支持断点续传、校验和或自定义输出路径。

## 网络行为

程序启动时会将 Go 默认 DNS 解析器设置为公共 DNS `223.5.5.5`，并通过 UDP 使用该服务器。若当前网络屏蔽此 DNS 或要求使用内网 DNS，OpenList 主机名可能无法解析。需要其他 DNS 行为时，请调整 `utils/dns.go`。

## 开发

可以使用以下命令格式化、测试和构建项目：

```bash
gofmt -w main.go cmd/root.go model/*.go tui/*.go utils/*.go
go test ./...
make current
```

仓库当前没有自动化测试文件。Makefile 构建会通过链接参数写入 `Version`、`GitCommit` 和 `BuildTime`；直接执行 `go build` 时这些值会保持为空。

## 已知限制

- 虽然底层请求模型包含密码、分页和刷新字段，但 CLI 尚未提供对应参数。
- 下载不支持断点续传、校验和或自定义目标路径，并且会直接使用服务器返回的文件名。
- Android 交叉编译要求每个目标架构对应的 NDK Clang 可执行文件均可通过 `PATH` 访问。
- 构建全部平台需要 Go 版本和外部工具链支持每一组 `GOOS`/`GOARCH` 组合。

## 许可证

本项目采用 GNU General Public License v3.0（GPL-3.0）授权，完整许可证文本请参阅 [LICENSE](LICENSE)。
