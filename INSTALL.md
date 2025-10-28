# Installation Guide

## Quick Install

### macOS / Linux (Homebrew)

```bash
# Add the tap
brew tap alkowskey/tap

# Install commitlens
brew install commitlens
```

### Windows (Scoop)

```powershell
# Add the bucket
scoop bucket add alkowskey https://github.com/alkowskey/scoop-bucket

# Install commitlens
scoop install commitlens
```

### Go Install (All Platforms)

If you have Go installed:

```bash
go install github.com/alkowskey/commitlens@latest
```

### Download Binary (All Platforms)

Download the latest release for your platform from:
https://github.com/alkowskey/commitlens/releases

#### Linux / macOS

```bash
# Download and extract (replace VERSION and OS/ARCH)
curl -L https://github.com/alkowskey/commitlens/releases/download/v1.0.0/commitlens_Linux_x86_64.tar.gz | tar xz

# Move to PATH
sudo mv commitlens /usr/local/bin/

# Verify installation
commitlens version
```

#### Windows

1. Download `commitlens_Windows_x86_64.zip` from releases
2. Extract the ZIP file
3. Add the directory to your PATH or move `commitlens.exe` to a directory in your PATH

## Build from Source

```bash
# Clone the repository
git clone https://github.com/alkowskey/commitlens.git
cd commitlens

# Build
go build -o commitlens

# Install to GOPATH/bin
go install
```

## Configuration

After installation, create a `.env` file in your project directory:

```bash
cp .env.example .env
```

Edit `.env` and add your API keys:

```env
# For Groq API
GROQ_API_KEY=your_groq_api_key
GROQ_MODEL=llama-3.3-70b-versatile
GROQ_API_URL=https://api.groq.com/openai/v1/chat/completions

# For OpenAI API (alternative)
OPENAI_API_KEY=your_openai_api_key
```

## Verify Installation

```bash
commitlens version
commitlens --help
```

## Uninstall

### Homebrew
```bash
brew uninstall commitlens
brew untap alkowskey/tap
```

### Scoop
```powershell
scoop uninstall commitlens
```

### Manual
```bash
# Remove binary
sudo rm /usr/local/bin/commitlens

# Remove data (optional)
rm -rf ~/.commitlens
```

## Troubleshooting

### Command not found

Make sure the binary is in your PATH:

```bash
# Check if commitlens is in PATH
which commitlens

# Add to PATH (Linux/macOS - add to ~/.bashrc or ~/.zshrc)
export PATH="$PATH:/path/to/commitlens"
```

### Permission denied

```bash
# Make binary executable
chmod +x commitlens
```

### API Key Issues

Ensure your `.env` file is in the directory where you run commitlens, or set environment variables:

```bash
export GROQ_API_KEY=your_key_here
```
