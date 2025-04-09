This is the project about enterprise level of user management

User-Management-Go-React/  
│── Echo/  
│   ├── cmd/                  # Application startup entry point  
│   │   └── main.go           # Main program entry  
│   ├── config/               # Configuration files (e.g., database configuration, environment variables)  
│   │   └── config.go         # Configuration loading logic  
│   ├── internal/             # Internal application logic  
│   │   ├── handler/          # Route handler functions  
│   │   │   └── user_handler.go  
│   │   ├── middleware/       # Middleware definitions  
│   │   │   └── auth.go       # Authentication-related middleware  
│   │   ├── model/            # Data model definitions  
│   │   │   └── user_model.go  
│   │   ├── repository/       # Database operation logic (encapsulating GORM queries)  
│   │   │   └── user_repository.go  
│   │   ├── service/          # Business logic layer (calls repository)  
│   │   │   └── user_service.go  
│   │   ├── router/           # Route definitions  
│   │   │   └── router.go  
│   │   ├── util/             # Utility functions (e.g., encryption, logging tools)  
│   │       └── hash_util.go  # Password encryption tool  
│   ├── migrations/           # Database migration files (SQL or GORM auto-migration)  
│       └── init.sql          # Initial migration file  
│   ├── .env                  # Environment variable file (e.g., database connection information)  
│   ├── go.mod                # Go module definition file  
│   ├── go.sum                # Go dependency lock file  
│  
│── React/  
│  
├── docker-compose.yml         # Docker Compose file for containerized deployment of backend and frontend services  
├── README.md                  # Project documentation  
