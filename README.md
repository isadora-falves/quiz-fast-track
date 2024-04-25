# Quiz-Fast-Track API
<p align="center">
  <img src="https://github.com/isadora-falves/quiz-fast-track/assets/77645495/dc696716-0b3c-41d9-beac-8660f2621573" width="50%" height="auto" alt="Quiz-Fast-Track API">
</p>


Welcome to the Quiz-Fast-Track API, a Go application designed for taking quizzes and tracking scores. This project was developed as a solution for a test proposed by Fast-Track Engineering.

## Features

- **In-memory data store**: Fast and efficient handling of quizzes and questions.
- **RESTful API**: Easy interaction through well-defined HTTP endpoints.
- **Swagger Documentation**: Detailed API documentation and interactive endpoint testing.

## Requirements

- Docker

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Running the Application

To start the application, execute the following command:

```bash
docker-compose up --build
```

## Running the Application

After running the docker compose, access the application by visiting:

[http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)

This URL will take you to the Swagger UI where you can interact with the API endpoints.

## Architecture

The application is structured using a clean architecture approach with four main layers:

- **Entities Layer**: Core business objects of the application.
- **Use Cases Layer**: Business rules can be executed without external interaction.
- **Controllers Layer**: Interface adapters that bridge the gap between use cases and the web framework.
- **Infrastructure Layer**: Implements external interfaces and frameworks such as data storage and web frameworks.

## Key Decisions

- **Layered Approach**: Ensures separation of concerns among components.
- **Dependency Injection**: Used extensively to manage dependencies cleanly throughout the application.
- **Development Approach**: Developed "outside-in", starting with entities, followed by use cases, infrastructure, and finally integrating everything through controllers and the main application layer.

## Built With

- **Go**: The programming language used.

## Authors

- **Isadora Alves** - Initial work - [IsadoraAlves](https://github.com/IsadoraAlves)

See also the list of [contributors](CONTRIBUTORS.md) who participated in this project.
