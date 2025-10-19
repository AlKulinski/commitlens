# CommitLens

A lightweight CLI tool that tracks file changes and helps you understand what's worth committing. CommitLens takes snapshots of your directories and provides detailed diffs to give you clarity on your changes.

## Features

- **Snapshot Tracking**: Capture the state of your files at any point
- **Smart Diffing**: Compare current state against snapshots to see what changed
- **Concurrent Processing**: Optimized performance with worker pools for large file sets
- **SQLite Storage**: Lightweight local database for snapshot metadata

## Installation

```bash
go install github.com/alkowskey/commitlens@latest
```

Or build from source:

```bash
git clone https://github.com/alkowskey/commitlens.git
cd commitlens
go build
```

## Usage

### Start Tracking

Begin tracking changes in a directory:

```bash
commitlens track start -d ./src
```

This creates a snapshot of all files in the specified directory.

### Compare Changes

See what's changed since your last snapshot:

```bash
commitlens track compare -d ./src
```

### Flush Snapshots

Clear all stored snapshots:

```bash
commitlens track flush
```

### Run Diff

Compare two files directly:

```bash
commitlens diff -s file1.txt -t file2.txt
```

## How It Works

1. **Snapshot**: CommitLens creates a cached copy of your files and stores metadata (hash, size, mtime) in SQLite
2. **Track**: File changes are detected by comparing current state against stored snapshots
3. **Diff**: Uses concurrent processing to efficiently compare large batches of files
4. **Report**: Shows added and removed lines for each changed file

## Architecture

```
commitlens/
├── cmd/              # CLI commands
├── internal/
│   ├── snapshot/     # Snapshot tracking logic
│   ├── diff/         # File comparison engine
│   ├── db/           # SQLite database
│   └── common/       # Shared utilities
```

## Performance

CommitLens uses an optimized concurrent worker pool pattern:

- **10 files**: ~21µs per operation
- **100 files**: ~116µs per operation  
- **1000 files**: ~1ms per operation

Benchmarks show 3x faster performance compared to synchronous processing on large file sets.

## License

MIT

## Contributing

Contributions welcome! Feel free to open issues or submit PRs.
