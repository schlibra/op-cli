# op-cli

[English](README.md) | [中文](README_zh.md)

`op-cli` is a terminal client for an [OpenList](https://github.com/OpenListTeam/OpenList) server. It provides both an interactive menu and command-line operations for server configuration, authentication, file browsing, file metadata, and downloads.

## Features

- Configure the OpenList server base URL.
- Log in with a username and password, with optional TOTP support.
- Query the current user's profile.
- List files and directories under a path.
- Inspect file or directory metadata.
- Download a file with a progress bar.
- Browse directories, inspect files, and start downloads from the interactive file browser.
- Show build version, commit, and build time information.
- Build static binaries for Linux, macOS, Windows, BSD, and other Unix-like targets.

## Requirements

- Go 1.26.6 or a compatible Go toolchain. The required version is declared in `go.mod`.
- An OpenList server exposing the API endpoints used by this client.
- A terminal that supports interactive input for menu mode.
- GNU Make is required only for the Makefile-based build targets.

The client currently calls these OpenList endpoints:

| Operation | Endpoint |
| --- | --- |
| Login | `POST /api/auth/login` |
| Current user | `GET /api/me` |
| Logout | `POST /api/auth/logout` |
| List files | `POST /api/fs/list` |
| File metadata | `POST /api/fs/get` |

Use a base URL such as `https://openlist.example.com` without a trailing slash. The command-line URL setter does not remove a trailing slash before composing API paths.

## Installation

### Build locally

```bash
mkdir -p dist
go build -o dist/op-cli .
```

On Windows, use an `.exe` suffix if desired:

```powershell
New-Item -ItemType Directory -Force dist | Out-Null
go build -o dist/op-cli.exe .
```

### Cross-compile all targets

```bash
make all-platform
```

The Makefile also provides these build targets:

| Target | Description |
| --- | --- |
| `make current` | Build the current host platform as `dist/op-cli-current` (or `.exe` on Windows), embedding version metadata. |
| `make all-platform` | Build every platform group declared by the Makefile. |
| `make all` or `make` | Build the current platform and all declared platform targets. |
| `make clean` | Remove the `dist/` output directory. |

Cross-compiled files use the `op-cli-<goos>-<goarch>` naming pattern. The current target groups are:

- Linux: `amd64`, `386`, `arm64`, `arm`, `riscv64`, `ppc64`, `ppc64le`, `s390x`, and `loong64`.
- macOS: `amd64` and `arm64`.
- Windows: `amd64`, `386`, and `arm64`.
- FreeBSD, OpenBSD, and NetBSD: the architectures declared in `Makefile`.
- DragonFly BSD, Solaris, and illumos: the architectures declared in `Makefile`.
- Android-named targets are present in the Makefile; see [Known limitations](#known-limitations) before distributing them.

For example, the initial target set produces files such as:

```text
op-cli-linux-amd64
op-cli-darwin-amd64
op-cli-windows-amd64.exe
```

All cross-platform recipes use `CGO_ENABLED=0`. Building every target may take considerably longer than a native build.

## Quick start

The examples below assume the binary is available as `./op-cli`.

```bash
# Set the server URL
./op-cli url https://openlist.example.com

# Log in (append a TOTP code when the account requires it)
./op-cli login <username> <password> [totp-code]

# Verify the session and inspect the root directory
./op-cli info
./op-cli ls /

# Inspect and download a file
./op-cli get /documents/report.pdf
./op-cli download /documents/report.pdf
```

Run the program without arguments to open the interactive menu:

```bash
./op-cli
```

The interactive menu provides:

- **Server setting**: view and save the OpenList Base URL.
- **Auth setting**: log in, log out, or query the current user. The login form accepts an optional TOTP code.
- **File browser**: navigate directories, move to the parent directory, view file metadata, and download files.
- **About program**: view version, Git commit, and build time.

The CLI login argument validation in the current source is still incorrect. The documented CLI login form may fail with an index-out-of-range error; use the interactive login form or fix `cmd/root.go` before relying on CLI login.

## Command reference

| Command | Description |
| --- | --- |
| `url <base-url>` | Set the OpenList server base URL. |
| `url get` or `url info` | Print the configured base URL. |
| `login <username> <password> [totp-code]` | Authenticate and save the returned token. |
| `logout` | Log out and clear the locally stored token. |
| `info` | Get the current user's username, base path, role, and permission. |
| `ls [path]` | List a directory. Defaults to `/` when no path is supplied. |
| `get [path]` | Show metadata and the raw URL for a file or directory. Defaults to `/`. |
| `download [path]` | Download a file to the current directory using the server-provided name. `dl` and `down` are aliases. |
| `version` | Show version, Git commit, and build time. `-v` is an alias. |
| `help` | Show the built-in usage text. |

The commands also accept the corresponding `-command` and `--command` forms where implemented by the client.

## Configuration and authentication

Configuration is stored as TOML at:

```text
<user-home>/.config/op-cli/config.toml
```

On Windows, `<user-home>` is the current user's profile directory and path separators follow the platform. The file contains:

```toml
base_url = "https://openlist.example.com"
token = "<access-token>"
```

The file is created automatically on the first run. The access token is stored locally in plain text, so protect the file and do not commit it to source control.

## Download behavior

`download` first requests file metadata, then downloads the returned `raw_url`. The output file is created in the current working directory and uses the name returned by the server. An existing file with the same name is overwritten. Downloads are streamed and display a progress bar; resume, checksum verification, and a custom output path are not currently supported.

## Development

Format and test the project with:

```bash
gofmt -w main.go cmd/root.go model/*.go tui/*.go utils/*.go
go test ./...
make current
```

There are currently no project tests. Makefile builds embed `Version`, `GitCommit`, and `BuildTime` through linker flags; a direct `go build` does not set these values.

## Known limitations

- The client is coupled to the OpenList response shapes and endpoint paths listed above.
- File listing and metadata requests do not expose password, pagination, or refresh options through the CLI, even though the underlying request models contain those fields.
- The CLI login argument indexes are currently checked incorrectly. Running the documented `login <username> <password>` form can produce an index-out-of-range error; use interactive login or fix `cmd/root.go` first.
- The `android-*` Makefile recipes currently set `GOOS=linux`, so their output is Linux code with Android-style filenames rather than native Android binaries.
- Cross-compiling every target requires a compatible Go toolchain and can fail for targets not supported by the local Go release.

## License

This project is licensed under the GNU General Public License v3.0. See [LICENSE](LICENSE) for the full license text.
