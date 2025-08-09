# Go BMI (IMT) Calculator

A simple interactive Go program that calculates Body Mass Index (BMI, also called IMT) from user input. It prompts for height (in meters) and weight (in kilograms) and prints the computed BMI.

## Prerequisites

### Installing Go

Before running this application, you need to install Go on your system.

#### Linux/macOS
1. Download Go from the official website: https://golang.org/dl/
2. Extract the archive to `/usr/local`:
   ```bash
   sudo tar -C /usr/local -xzf go<version>.linux-amd64.tar.gz
   ```
3. Add Go to your PATH by adding this line to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.):
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
4. Reload your shell profile:
   ```bash
   source ~/.bashrc  # or ~/.zshrc
   ```
5. Verify the installation:
   ```bash
   go version
   ```

#### Windows
1. Download the Windows installer from https://golang.org/dl/
2. Run the installer and follow the installation wizard
3. Open Command Prompt and verify the installation:
   ```cmd
   go version
   ```

## Project Setup

### Creating a Go Module

If you're starting a new Go project, you can create a module using:

```bash
# Create a new directory for your project
mkdir my-go-project
cd my-go-project

# Initialize a new Go module
go mod init my-project-name

# This creates a go.mod file that tracks your dependencies
```

### Understanding go.mod

The `go.mod` file defines your module and its dependencies. In this project, it contains:

```
module demo/app-1

go 1.24.6
```

- `module demo/app-1`: The module name
- `go 1.24.6`: The Go version used

## Running the Application

### Method 1: Direct execution
```bash
go run .
```

Example interactive session:

```text
$ go run .
Enter your height:
1.78
Enter your weight:
85.5
27.0...
```

### Method 2: Build and run
```bash
# Build the application (outputs a binary named "go-demo" by default)
go build -o go-demo

# Run the built binary
./go-demo
```

### Method 3: Install locally (GOBIN/GOPATH)
```bash
# Install to $(go env GOPATH)/bin (ensure it's on your PATH)
go install .

# Then run (path may vary by setup)
go-demo
```

## Project Structure

```
go-demo/
├── go.mod          # Module definition and dependencies
├── main.go         # Main application file
└── README.md       # This file
```

## Development

### Behavior and Assumptions

- Expects metric units: height in meters, weight in kilograms
- Reads from standard input; no flags/arguments are currently supported
- Outputs the raw floating-point BMI value using default formatting

### Adding Dependencies

To add external packages to your project:

```bash
# Add a dependency
go get github.com/example/package

# This automatically updates go.mod and creates go.sum
```

### Running Tests

If you have test files (e.g., `main_test.go`):

```bash
go test
```

### Formatting Code

```bash
go fmt
```

### Running Linter

```bash
# Install golangci-lint (if not already installed)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run the linter
golangci-lint run
```

## Troubleshooting

### Common Issues

1. **"go: command not found"**
   - Make sure Go is installed and added to your PATH
   - Restart your terminal after installation

2. **"module not found"**
   - Run `go mod tidy` to clean up dependencies
   - Check that your `go.mod` file exists and is properly formatted

3. **Permission denied when running built binary**
   - Make the binary executable: `chmod +x go-demo`

## Next Steps

- Add input validation and friendly error messages
- Format BMI output (e.g., to two decimals) and add categorization
- Optional: add CLI flags for non-interactive usage (e.g., `--height`, `--weight`)
- Add tests in `*_test.go` files

## Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Go Modules Reference](https://golang.org/ref/mod)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
