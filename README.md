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
- Show build version, commit, and build time information.
- Build static binaries for Linux, macOS, and Windows.

## Requirements

- Go 1.26.6 or a compatible Go toolchain. The required version is declared in `go.mod`.
- An OpenList server exposing the API endpoints used by this client.
- A terminal that supports interactive input for menu mode.

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
go build -o dist/op-cli .
```

On Windows, use an `.exe` suffix if desired:

```powershell
go build -o dist/op-cli.exe .
```

### Cross-compile all targets

```bash
make build-all
```

The binaries are written to `dist/` with these names:

```text
op-cli-linux-amd64
op-cli-linux-arm64
op-cli-darwin-amd64
op-cli-darwin-arm64
op-cli-windows-amd64.exe
```

The `make build` recipe is currently a placeholder; use `go build` for a single native binary or `make build-all` for all listed targets.

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

The menu currently exposes server URL, authentication, and about/version screens. A `File` option is displayed in the home menu but is not wired to an action yet; use the command-line file commands for browsing and downloads.

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
```

There are currently no project tests. The build embeds `Version`, `GitCommit`, and `BuildTime` through linker flags when using `make build-all`.

## Known limitations

- The client is coupled to the OpenList response shapes and endpoint paths listed above.
- File listing and metadata requests do not expose password, pagination, or refresh options through the CLI, even though the underlying request models contain those fields.
- The CLI login argument indexes are currently checked incorrectly. Running the documented `login <username> <password>` form can produce an index-out-of-range error; fix `cmd/root.go` before relying on CLI login.

## License

This project is licensed under the GNU General Public License v3.0. See [LICENSE](LICENSE) for the full license text.
