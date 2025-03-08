# TARDIS

Tar Dis and Untar Dat

## Features

- Support for Windows, MacOS, and Linux
- Create an archive of one or multiple files or directory
- Support to delete source files/directories after a successful archive
- Recursively extract an archive
- Support for compressing archives using gzip
- Support for decompressing archives using gzip and bzip2
- Delete protection for `/` and `C:\` directories

## Usage

### Creating an archive

For compression, end the archive name parameter with a `.gz`, `.tgz`, or `.gzip` file extension.

Creating an archive of one file or directory
```
$ tardis create -a <archive.tar> -p <file-to-archive>
```

Create an archive of multiple files or directories
```
$ tardis create -a <path-of-archive-to-create> -p <file-to-archive> -p <directory-to-archive>
```

Create an archive of multiple files or directories and delete source files if archive was successful
```
$ tardis create -a <path of archive to create> -p <file-to-archive> -p <directory-to-archive> --delete
```

### Extract an archive

`tardis` uses the file extension to detect whether the archive is uncompressed or uses `gzip` or `bzip2` compression.
```
$ tardis extract -f <tar-file-to-extract> -d <destination-directory>
```

## Building from Code

```
$ git clone
$ cd tardis
$ go build -o tardis
```