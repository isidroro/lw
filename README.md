# lw

CLI tool to list files and directories recursively with their sizes. Written in Go.

Directories show the sum of all their contents. Entries are sorted with directories first, then files alphabetically.

## Usage

```
lw [dir]
```

- No argument: lists the current directory.
- `.` or `..` or any path: resolves and lists that directory.

## Sample output

```
lw src/

src/                 12.4K
  lib/               8.1K
    files.go         4.2K
    walk.go          3.9K
  main.go            4.3K
```
