# reposnip

Download files and folders from any GitHub repo without cloning the whole thing.

## The problem

You find a useful folder in a GitHub repo. To get it you have to clone the entire repo — sometimes hundreds of MBs — just to use one folder.

reposnip fixes that.

## Demo
```bash
reposnip https://github.com/golang/example/tree/master/hello
```
```
Found 5 files:
  hello/go.mod
  hello/hello.go
  hello/reverse/example_test.go
  hello/reverse/reverse.go
  hello/reverse/reverse_test.go

✅ Done! 5 file(s) saved.
```

## Install

Download the binary for your platform from [Releases](https://github.com/Vydeh07/reposnip/releases)

**Linux**
```bash
chmod +x reposnip-linux
sudo cp reposnip-linux /usr/local/bin/reposnip
```

**Mac**
```bash
chmod +x reposnip-mac
sudo cp reposnip-mac /usr/local/bin/reposnip
```

**Windows**
```
add reposnip-windows.exe to your PATH
```

## Usage
```bash
# download a folder — saves where you are
reposnip https://github.com/owner/repo/tree/main/foldername

# download a single file
reposnip https://github.com/owner/repo/blob/main/file.go

# save to a specific folder
reposnip https://github.com/owner/repo/tree/main/src --output ~/myproject

# private repo
reposnip https://github.com/owner/private-repo/tree/main/src --token ghp_xxx
```

## Why reposnip over git clone?

| | git clone | reposnip |
|---|---|---|
| Downloads entire repo | ✅ always | ❌ never |
| Gets just one folder | ❌ not possible | ✅ always |
| Private repo support | ✅ | ✅ via --token |
| Preserves folder structure | ✅ | ✅ |
| Single binary, no runtime | ❌ needs git | ✅ |

## Built with

- Go — compiles to a single binary, no runtime needed
- GitHub Contents API — fetches only what you ask for
- Zero external dependencies except cobra for CLI flags