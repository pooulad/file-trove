# 📂file-trove

Simple backup cli tool written with Go


## Flags Overview

This document provides details about the available flags for the File Trove CLI tool, their descriptions, and example usages.


| Flag             | Type    | Default            | Description                                                                 |
|-------------------|---------|--------------------|-----------------------------------------------------------------------------|
| `--operation`     | String  | `"backup"`         | Specifies the operation type. Allowed values: `backup`, `restore`.         |
| `--src`           | String  | `""` (required)    | Path to the source directory for the operation.                            |
| `--dest`          | String  | `""` (required)    | Path to the destination directory for the operation.                       |
| `--period`        | Int     | `0`                | Sets the backup interval in minutes. If `0`, the backup runs once.         |
| `--compress`      | Bool    | `false`            | Enables compression of the backup into a ZIP file.                         |
| `--detach`        | Bool    | `false`            | Runs the tool in detached mode (background) and logs to a specified file.  |
| `--log`           | String  | `"file-trove.log"` | Specifies the log file path.                                               |
| `--exclude`       | String  | `""`               | Comma-separated list of files/directories to exclude from the operation.   |

---

## Usage and Examples

This section provides examples of how to use the File Trove CLI tool with different flags and configurations.

---

## Build the CLI Tool
Compile the tool to generate the executable:

```bash
git clone https://github.com/pooulad/file-trove.git

cd file-trove
```

```bash
go build -o file-trove.exe . // in windows os
```
or
```bash
go build -o file-trove . 
```

## Basic Backup
Perform a simple backup from the source directory to the destination directory:

Windows example:
```bash
.\file-trove --operation backup --src "C:\Users\YourName\Documents" --dest "D:\Backups"
```

Linux example:
```bash
./file-trove --operation backup --src "src/folder" --dest "dest/folder"
```

## Periodic Backup

Perform a backup every 10 minutes:

Windows example:
```bash
.\file-trove --operation backup --src "C:\Projects" --dest "E:\ProjectBackups" --period 10
```

Linux example:
```bash
./file-trove --operation backup --src "src/folder" --dest "dest/folder" --period 10
```

## Full Configuration Example

Perform a periodic backup every 15 minutes, compress the result, exclude unnecessary files, and log to a file:

Windows example:
```bash
.\file-trove --operation backup --src "C:\Source" --dest "D:\Destination" --period 15 --compress --log "backup.log" --exclude "temp,cache"
```

Linux example:
```bash
./file-trove --operation backup --src "src/folder" --dest "dest/folder" --period 15 --compress --log "backup.log" --exclude "temp,cache"
```
Checkout `example` directory for see result of tool.

## TODO Checklist

This section tracks the progress of implemented features in the File Trove CLI tool.

- [x] Implement `--period` flag to allow periodic backups.
- [x] Implement `--compress` flag to enable compression of backups into a ZIP file.
- [x] Implement `--detach` flag to run the backup operation in detached mode with logging to a file.
- [x] Implement `--log` flag to define the log file path.
- [ ] **Implement `--exclude` flag to skip specific files or directories during backup.**

## ⭐️ Support the Project

If you find this project useful or interesting, please consider giving it a ⭐️ on GitHub! It means a lot and helps others discover this project. Your support keeps the project alive and evolving!

[![GitHub stars](https://img.shields.io/github/stars/pooulad/file-trove?style=social)](https://github.com/pooulad/file-trove)

---

### Here's how you can star the repo:

1. Go to the [GitHub repository](https://github.com/pooulad/file-trove).
2. Click on the ⭐️ button at the top-right corner.

---

### Why Star This Repo?

🌟 **Show Your Support**: Your star demonstrates appreciation for the effort behind the project.  
🌟 **Spread the Word**: Stars make the project more visible to others on GitHub.  
🌟 **Stay Updated**: Starring makes it easy to track updates in your list of starred repositories.

---

### A Big Thank You! 🙏

Thank you for taking the time to explore and support this project! 🚀

## License

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)


