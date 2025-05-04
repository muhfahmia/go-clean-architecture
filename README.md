# go-clean-architecture
Common sense structure clean layering programming concept for teamwork

MY-GO-CLEAN-ARCHITECTURE/
├── cmd/                  # Application entry points (main.go and other executables)
│   └── main.go           # Primary application entry point that initializes dependencies
│   └── server.go         # HTTP server initialization and routing
│
├── config/               # Configuration management
│   ├── app.yaml          # Application configuration (environment, ports, timeouts)
│   ├── database.yaml     # Database connection parameters
│   └── config.go         # Struct definitions for configuration with env var support
│
├── db/                   # Database related files
│   └── migration/        # Database schema migration scripts
│       ├── 0001_init.up.sql  # Initial database schema
│       └── 0002_alter_table.up.sql  # Schema updates
│
pkg/
│    ├── logger/               # Package logging reusable
│    ├── validator/            # Validasi data generik
│    ├── database/             # Helper database (bukan repository)
│    ├── httputil/             # HTTP utilities
│    ├── errors/               # Error handling standar
│    ├── pagination/           # Logic paginasi
│    └── uuid/                 # UUID generator
│
├── internal/             # Core application logic (not importable by external projects)
│   ├── delivery/         # Interface adapters for various delivery mechanisms
│   │   ├── http/         # HTTP handlers and routes (REST API endpoints)
│   │   ├── grpc/         # gRPC service implementations
│   │   └── cli/          # Command line interface handlers
│   │
│   ├── dependency/       # Dependency injection setup (corrected from 'depedency')
│   │   ├── wire.go       # Google Wire dependency definitions
│   │   └── provider.go   # Dependency provider functions
│   │
│   ├── entity/           # Enterprise business rules and core entities
│   │   ├── user.go       # User entity with validation logic
│   │   └── order.go      # Order entity with business rules
│   │
│   ├── gateway/          # Interfaces for external services
│   │   ├── payment.go    # Payment service abstraction
│   │   └── email.go      # Email service interface
│   │
│   ├── model/            # Data structures (optional layer)
│   │   ├── db_model.go   # Database model definitions
│   │   └── api_model.go  # API request/response formats
│   │
│   ├── repository/       # Data access interfaces
│   │   ├── user.go       # User repository interface
│   │   └── order.go      # Order repository interface
│   │
│   └── usecase/          # Application-specific business rules
│       ├── user.go       # User-related operations
│       └── order.go      # Order processing logic
│
├── test/                 # Test files and utilities
│   ├── user_test.go      # Unit tests for user domain
│   ├── http_test.go      # HTTP handler tests
│   └── testutils/        # Testing helpers and mocks
│
├── .gitignore           # Specifies untracked files (binaries, logs, env files)
├── LICENSE              # Project license (MIT/Apache/GPL etc.)
└── README.md           # This documentation file