# Golang API structure


### Pre-requisites
- Go installed on your machine
- Taskfile installed on your machine

### Getting Started

1. Clone this repository
```bash
git clone https://github.com/hello/go-struct.git
```

2. Rename the project to your own repository name
```bash
# Replace 'github.com/your-username/your-project' with your actual repository path
task utils:rename-project -- github.com/your-username/your-project
```

3. Initialize your Go module
```bash
go mod init
go mod tidy
```

4. Run the project
```bash
task dev
```

### Project Structure
```
.
├── cmd/              # Application entry points
├── hacks/            # Scripts and tools
├── internal/         # Private application and library code
│   ├── api/         # API controllers and routes
│   ├── utils/       # Utility functions
│   └── core/        # Core application logic
│      └── domain/   # Business entities and value objects
│      ├── ports/    # Business entities and value objects
│      └── services/ # Business services
├── pkg/              # Library code that's ok to use by external applications
│   ├── db/          # Database connection and repositories
│   ├── gateways/    # External services and clients
│   └── utils/       # Utility classes
├── taskfiles/        # Task definitions
└── README.md         # Project documentation
```
