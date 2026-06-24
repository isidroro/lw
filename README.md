# LW

### Description:
Small CLI utility to list files and their sizes recursively in a target directory. Written in Go.

### Use:
From the command line: `lw <target>` to list files and weights from the `<target>` directory.

### Sample output:
lw mydir

mydir/
main.go        1.2K
  lib/
    walk.go    800B
    files.go   400B
