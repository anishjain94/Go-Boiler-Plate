# Go Boilerplate Repository

This repository provides a boilerplate for Go (Golang) projects, designed to provide a quick start for new applications with a robust and scalable structure. It includes a set of best practices, tools, and configurations allowing developers to focus more on writing business logic rather than setting up a project.

## Features

- **Standard Project Layout**:

  - The project follows the Standard Go Project Layout which is the standard layout in the Go ecosystem.

      ```/myapp
        /cmd
            main.go
        /infra
          /database
            gorm.go
          /environment
            environment.go
          /rest
            rest.go
        /modules
          /health
        /utils
        go.mod
        go.sum
        README.md
      ```

- **Dependency Management**:

  - This boilerplate uses Go Modules for managing dependencies, which is the official dependency management solution in Go.

## Getting Started

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Run go mod tidy to download the necessary dependencies.
4. Start building your application!

## Contributing

Contributions to this boilerplate are welcome! Please feel free to open an issue or submit a pull request with your improvements or suggestions.

## License

This project is licensed under the MIT License. See the LICENSE file for details.



Folder structure
/user-service
    /cmd
        main.go        
    /common
        constants.go
        helpers.go
    /config
        config.yaml
        .env
    /docker
        /compose
        /dockerfiles
        /scripts
    /infra
        /rest
            rest.go
        /redis
            redis.go
        /middleware
            auth.middleware.go
            cors.middleware.go
            log.middleware.go
            throttle.middleware.go
        /database
            /migration
                001.sql
            gorm.go
        /environment
            environment.go
        /rest
            rest.go
        /handler
            handler.go
    /integrations
        /google
            google.go
            google.helper.go
            google.constants.go
        /aws
            aws.go
    /modules
        /health
        /users
    /utils
        string.go
        struct.go
    go.mod
    go.sum
    README.md
