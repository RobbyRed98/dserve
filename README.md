# dserve
A small web server written in go that hosts the content of a directory and its subdirectories.

## Building & Packaging
Requires: `goreleaser` https://github.com/goreleaser/goreleaser
```bash
$ goreleaser release --skip-publish --snapshot --rm-dist
```

Creates the following packages:
* `deb`: arm6 arm7 arm8 arm64 i386 amd64
* `linux.tar.xz`: arm6 arm7 arm8 arm64 i386 amd64
* `freebsd.tar.xz`: arm6 arm7 arm8 i386 amd64
* `windows.zip`: i386 amd64

## Usage

Host the content of the directory in which the application is located.
```bash
$ dserve 
```
or
```powershell
dserve.exe
```

Open the following URL: http://localhost:8080/file_name

### Options

#### Host address
Default: `localhost`

```bash
$ dserve --ip 123.45.67.81
```
or
```powershell
dserve.exe --ip 123.45.67.81
```

#### Port
Default: `8080`
```bash
$ dserve --ip 1234
```
or
```powershell
dserve.exe --ip 1234
```

#### Directory to host
Defaults: `.`
```bash
$ dserve --dir path/to/directory
```
or
```powershell
dserve.exe --dir path\to\directory
```