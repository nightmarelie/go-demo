# Go Demo Application

A simple Go application that prints "Hello, World!" to demonstrate basic Go programming concepts.

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
go run main.go
```

### Method 2: Build and run
```bash
# Build the application
go build

# Run the built binary
./app-1
```

### Method 3: Install globally
```bash
# Install the application globally
go install

# Run from anywhere (if your GOPATH/bin is in your PATH)
app-1
```

## Project Structure

```
go-demo/
├── go.mod          # Module definition and dependencies
├── main.go         # Main application file
└── README.md       # This file
```

## Development

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
   - Make the binary executable: `chmod +x app`

## Next Steps

- Add more functionality to `main.go`
- Create additional packages in subdirectories
- Add tests in `*_test.go` files
- Set up a CI/CD pipeline
- Add configuration management
- Implement logging and error handling

## Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Go Modules Reference](https://golang.org/ref/mod)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
