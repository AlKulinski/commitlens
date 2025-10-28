# CommitLens

A lightweight CLI tool that tracks file changes and helps you understand what's worth committing. CommitLens takes snapshots of your directories and provides detailed diffs with AI-powered summaries to give you clarity on your changes.

## Features

- **Snapshot Tracking**: Capture the state of your files at any point
- **Smart Diffing**: Compare current state against snapshots to see what changed
- **AI-Powered Summaries**: Get intelligent summaries of your changes using Groq API
- **Concurrent Processing**: Optimized performance with worker pools for large file sets
- **SQLite Storage**: Lightweight local database for snapshot metadata

## Installation

### Quick Install

**macOS / Linux (Homebrew):**
```bash
brew tap alkowskey/tap
brew install commitlens
```

**Windows (Scoop):**
```powershell
scoop bucket add alkowskey https://github.com/alkowskey/scoop-bucket
scoop install commitlens
```

**Go Install:**
```bash
go install github.com/alkowskey/commitlens@latest
```

**Download Binary:**

Download the latest release for your platform from [GitHub Releases](https://github.com/alkowskey/commitlens/releases).

See [INSTALL.md](INSTALL.md) for detailed installation instructions.

## Configuration

CommitLens supports AI-powered diff summaries via Groq API. Create a `.env` file in the project root:

```bash
cp .env.example .env
```

Configure your environment variables:

```env
GROQ_API_KEY=your_api_key_here
GROQ_MODEL=openai/gpt-oss-20b
GROQ_API_URL=https://api.groq.com/openai/v1/chat/completions
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
commitlens diff --from file1.txt --to file2.txt
# or using short flags
commitlens diff -f file1.txt -t file2.txt
```

### AI-Powered Diff Summaries

When configured with Groq API, CommitLens can generate intelligent summaries of your changes, helping you understand the impact and context of modifications across your codebase.

## How It Works

1. **Snapshot**: CommitLens creates a cached copy of your files and stores metadata (hash, size, mtime) in SQLite
2. **Track**: File changes are detected by comparing current state against stored snapshots using content hashing (xxhash)
3. **Diff**: Uses concurrent processing with configurable diff algorithms (base or patience) to efficiently compare large batches of files
4. **Analyze**: Optional AI-powered analysis via Groq API provides intelligent summaries of changes
5. **Report**: Shows added and removed lines for each changed file with detailed context

## Architecture

CommitLens follows clean architecture principles with clear separation of concerns:

```
commitlens/
├── cmd/                    # CLI commands (root, track, diff, version)
├── internal/
│   ├── snapshot/           # Snapshot tracking domain
│   │   ├── domain/         # Core snapshot entities
│   │   ├── repository/     # Data persistence layer
│   │   ├── services/       # Business logic (with benchmarks)
│   │   └── usecases/       # Application use cases (start, compare, flush)
│   ├── diff/               # File comparison engine
│   │   ├── domain/         # Diff entities and interfaces
│   │   ├── infra/          # Infrastructure (Groq API, diff algorithms)
│   │   ├── services/       # Diff business logic
│   │   ├── config/         # Configuration management
│   │   ├── prompts/        # AI prompt templates
│   │   ├── factories/      # Object creation patterns
│   │   └── usecases/       # Diff execution logic
│   ├── db/                 # SQLite database
│   │   └── migrations/     # Database schema migrations
│   └── common/             # Shared utilities
│       ├── flags/          # CLI flag definitions
│       └── utils/          # Common utility functions (87.8% test coverage)
```

### Key Components

- **Diff Algorithms**: Base differ and patience differ implementations
- **AI Integration**: Groq API integration for intelligent diff summaries
- **Repository Pattern**: Clean data access layer with SQLite backend
- **Worker Pools**: Concurrent processing for optimal performance
- **Domain-Driven Design**: Clear separation between domain, infrastructure, and application layers

## Performance

CommitLens uses an optimized concurrent worker pool pattern:

- **10 files**: ~21µs per operation
- **100 files**: ~116µs per operation
- **1000 files**: ~1ms per operation

Benchmarks show 3x faster performance compared to synchronous processing on large file sets.

## Tech Stack

- **Language**: Go 1.21+
- **CLI Framework**: urfave/cli v3
- **Database**: SQLite with mattn/go-sqlite3
- **Hashing**: xxhash (cespare/xxhash/v2) for fast content hashing
- **AI Integration**: Groq API for diff summarization
- **Testing**: Go's built-in testing framework with httptest for HTTP mocking

## Dependencies

Key dependencies include:

- `github.com/urfave/cli/v3` - Modern CLI framework
- `github.com/mattn/go-sqlite3` - SQLite driver
- `github.com/cespare/xxhash/v2` - Fast hashing algorithm
- `github.com/joho/godotenv` - Environment variable management

## Testing

Run the test suite:

```bash
# Run all tests
go test ./...
```

## License

MIT

## Contributing

Contributions welcome! Feel free to open issues or submit PRs.
