```
/TorTorCoinSystem
|-- /api
|   |-- /handlers
|   |   |-- user.go
|   |   |-- transaction.go
|   |   |-- friendship.go
|   |   |-- request.go
|   |-- /middleware
|   |   |-- auth.go
|   |-- router.go
|-- /log
|   |-- error.log
|   |-- warning.log
|   |-- info.log
|-- /cmd
|   |-- main.go
|-- /pkg
|   |-- /config
|   |   |-- config.go
|   |-- /model
|   |   |-- user.go
|   |   |-- transaction.go
|   |   |-- friendship.go
|   |   |-- request.go
|   |-- /repository
|   |   |-- user.go
|   |   |-- transaction.go
|   |   |-- friendship.go
|   |   |-- request.go
|   |-- /service
|   |   |-- user.go
|   |   |-- transaction.go
|   |   |-- friendship.go
|   |   |-- request.go
|-- /internal
|   |-- /util
|   |   |-- hash.go
|   |   |-- validator.go
|-- /migrations
|   |-- 001_init_schema.up.sql
|   |-- 001_init_schema.down.sql
|-- go.mod
|-- go.sum
|-- README.md
```

Here's what each directory and file is for:

- `/api`: Contains all the API-related code.
    - `/handlers`: Contains handler functions for your API endpoints.
    - `/middleware`: Contains middleware functions, such as authentication.
    - `router.go`: Sets up your API routes and associates them with handler functions.

- `/cmd`: Contains the entry point of your application.
    - `main.go`: The main file where your application starts.

- `/pkg`: Contains the core logic of your application.
    - `/config`: Contains configuration-related code.
    - `/model`: Contains data models that represent your database tables.
    - `/repository`: Contains code to interact with the database.
    - `/service`: Contains business logic.

- `/internal`: Contains internal utility functions and helpers.
    - `/util`: Utility functions, such as password hashing and input validation.

- `/migrations`: Contains database migration files.
    - `001_init_schema.up.sql`: SQL script to set up the initial database schema.
    - `001_init_schema.down.sql`: SQL script to tear down the initial database schema if needed.

- `go.mod` and `go.sum`: Go module files for managing dependencies.

- `README.md`: A markdown file to describe your project, how to build, run, and test it.
