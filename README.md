# op-cli

[English](README.md) | [中文](README_zh.md)

`op-cli` is a terminal client for an [OpenList](https://github.com/OpenListTeam/OpenList) server. It supports both an interactive menu and command-line operations for server configuration, authentication, file browsing, metadata inspection, and downloads.

## Features

- Configure and inspect the OpenList server base URL.
- Log in with a username and password, with optional TOTP support.
- Query the current user's profile and log out.
- List files and directories under a path.
- Inspect file or directory metadata.
- Download files with a live progress bar.
- Browse directories, inspect files, and start downloads from the interactive file browser.
- Show the program version, Git commit, and build time.
- Build binaries for the platform targets declared in the Makefile, including native Android builds through the Android NDK.

## Requirements

- Go 1.26.6 or a compatible Go toolchain. The required version is declared in `go.mod`.
- An OpenList server exposing the API endpoints used by this client.
- A terminal that supports interactive input when using menu mode.
- GNU Make only if you use the Makefile build targets.
- The Android NDK LLVM toolchain on `PATH` when building `android-*` targets.

The client currently calls these OpenList endpoints:

| Operation | Endpoint |
| --- | --- |
| Login | `POST /api/auth/login` |
| Current user | `GET /api/me` |
| Logout | `POST /api/auth/logout` |
| List files | `POST /api/fs/list` |
| File metadata | `POST /api/fs/get` |

Use a base URL such as `https://openlist.example.com` without a trailing slash. The client concatenates API paths directly and does not remove a trailing slash for you.

## Installation

### Build locally

On Linux, macOS, BSD, or another Unix-like system:

```bash
mkdir -p dist
go build -o dist/op-cli .
```

On Windows:

```powershell
New-Item -ItemType Directory -Force dist | Out-Null
go build -o dist/op-cli.exe .
```

### Cross-compile with Make

```bash
make current       # Build the current host platform
make all-platform  # Build every platform group
make                # Equivalent to: make all
make clean         # Remove dist/
```

`make current` embeds the version, Git commit, and build time into `dist/op-cli-current` (or `.exe` on Windows). Cross-compiled files use the `op-cli-<goos>-<goarch>` naming pattern. Non-Android cross-platform recipes set `CGO_ENABLED=0`; Android recipes use `GOOS=android` and `CGO_ENABLED=1`.

The platform groups currently declared in the Makefile are:

| Group | Targets |
| --- | --- |
| Linux | `amd64`, `386`, `arm64`, `arm`, `riscv64`, `ppc64`, `ppc64le`, `s390x`, `loong64` |
| macOS | `amd64`, `arm64` |
| Windows | `amd64`, `386`, `arm64` |
| FreeBSD | `amd64`, `386`, `arm64`, `arm`, `riscv64` |
| OpenBSD | `amd64`, `386`, `arm64`, `arm` |
| NetBSD | `amd64`, `386`, `arm64`, `arm` |
| Android | `amd64`, `386`, `arm64`, `arm` |
| Other Unix | DragonFly BSD `amd64`, Solaris `amd64`, illumos `amd64` |

Building every target can take considerably longer than a native build. The `android-*` recipes use architecture-specific Android NDK Clang compilers and default to API level `21`. Add the NDK LLVM `bin` directory to `PATH`, then override the API level when needed:

```bash
make android ANDROID_API_LEVEL=24
```

`make all-platform` includes the Android targets and therefore also requires the Android NDK compilers. The resulting files are native Android binaries named `op-cli-android-<goarch>`.

### Android NDK setup

The Android recipes resolve the NDK Clang launcher for each target architecture:

| Architecture | Compiler launcher |
| --- | --- |
| `amd64` | `x86_64-linux-android<api>-clang` |
| `386` | `i686-linux-android<api>-clang` |
| `arm64` | `aarch64-linux-android<api>-clang` |
| `arm` | `armv7a-linux-androideabi<api>-clang` |

Configure the toolchain as follows:

1. Install **NDK (Side by side)** from Android Studio's SDK Manager, or install an NDK package with the Android SDK command-line tools.
2. Locate the installed NDK directory, for example `<android-sdk>/ndk/<version>`.
3. Add its host-specific LLVM `bin` directory to `PATH`. The directory is under `<ndk-root>/toolchains/llvm/prebuilt/`; common host names are `windows-x86_64`, `linux-x86_64`, and `darwin-x86_64` (use the directory that exists in your NDK package).

On PowerShell:

```powershell
$env:ANDROID_NDK_HOME = "$env:LOCALAPPDATA\Android\Sdk\ndk\<version>"
$env:Path = "$env:ANDROID_NDK_HOME\toolchains\llvm\prebuilt\windows-x86_64\bin;$env:Path"
Get-Command aarch64-linux-android21-clang
```

On Linux or macOS:

```bash
export ANDROID_NDK_HOME="$ANDROID_SDK_ROOT/ndk/<version>"
export PATH="$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/<host-name>/bin:$PATH"
command -v aarch64-linux-android21-clang
```

Replace `21` in the verification command with the API level you plan to build. After the compiler check succeeds, build one architecture or all Android targets:

```bash
make android-arm64
make android ANDROID_API_LEVEL=24
```

On Windows, run these Makefile recipes from Git Bash, MSYS2, WSL, or another environment that provides the POSIX shell utilities used by the Makefile. Set `ANDROID_NDK_HOME` and `PATH` in the same shell that runs `make`; environment variables set in a separate PowerShell window are not inherited. If a compiler is not found, check the NDK host directory in `PATH`, the API level suffix, and the architecture-specific launcher name.

## Quick start

The examples below assume the binary is available as `./op-cli` (`.\op-cli.exe` on Windows).

1. Set the OpenList server URL:

   ```bash
   ./op-cli url https://openlist.example.com
   ```

2. Log in from the command line. Append a TOTP code when the account requires it:

   ```bash
   ./op-cli login <username> <password> [totp-code]
   ```

   Alternatively, run `./op-cli` and choose **Auth setting** > **Login user** to use the interactive login form.

3. Browse the root directory or inspect a file from the command line:

   ```bash
   ./op-cli ls /
   ./op-cli get /documents/report.pdf
   ./op-cli download /documents/report.pdf
   ```

Run `./op-cli help` for the built-in usage text. Running the program without arguments always opens the interactive menu.

## Command reference

| Command | Description |
| --- | --- |
| `url <base-url>` | Set the OpenList server base URL. |
| `url get` or `url info` | Print the configured base URL. |
| `login <username> <password> [totp-code]` | Authenticate and save the returned token. |
| `logout` | Log out and clear the locally stored token. |
| `info` | Show the current user's username, base path, role, and permission. |
| `ls [path]` | List a directory. The default path is `/`. |
| `get [path]` | Show file or directory metadata and its raw URL. The default path is `/`. |
| `download [path]` | Download a file to the current directory using the server-provided name. `dl` and `down` are aliases. |
| `version` | Show version, Git commit, and build time. `-v` is an alias. |
| `help` | Show the built-in usage text. |

The command dispatcher also accepts the corresponding `-command` and `--command` spellings, including the download aliases.

## Configuration and authentication

Configuration is stored as TOML at:

```text
<user-home>/.config/op-cli/config.toml
```

On Windows, `<user-home>` is the current user's profile directory and path separators follow the platform. The file is created automatically on first use:

```toml
base_url = "https://openlist.example.com"
token = "<access-token>"
```

The access token is stored locally in plain text. Protect this file and never commit it to source control.

## Downloads

`download` first requests file metadata from `/api/fs/get`, then downloads the returned `raw_url` into the current working directory. The output filename comes from the server. `os.Create` overwrites an existing file with the same name. Downloads are streamed and show a progress bar; resume, checksum verification, and a custom output path are not currently supported.

## Network behavior

At startup the program configures Go's default resolver to use the public DNS server `223.5.5.5` over UDP. Environments that block this server or require an internal DNS resolver may therefore be unable to resolve the OpenList host. Adjust `utils/dns.go` if your deployment needs different DNS behavior.

## Development

Format, test, and build the project with:

```bash
gofmt -w main.go cmd/root.go model/*.go tui/*.go utils/*.go
go test ./...
make current
```

The repository currently has no automated test files. Makefile builds embed `Version`, `GitCommit`, and `BuildTime` through linker flags; a direct `go build` leaves those values empty.

## Known limitations

- File listing and metadata requests do not expose password, pagination, or refresh options through the CLI, even though the underlying request models contain those fields.
- Downloads do not support resume, checksums, or a custom destination path, and the server-provided filename is used as-is.
- Android cross-compilation requires the matching NDK Clang executable for each requested architecture to be available on `PATH`.
- Cross-compiling every target requires a Go release and external toolchains that support each requested `GOOS`/`GOARCH` pair.

## License

This project is licensed under the GNU General Public License v3.0. See [LICENSE](LICENSE) for the full license text.
