# Social Media API Gateway

A scalable API Gateway built using Golang that acts as the single entry point for a microservices-based social media platform.

The gateway is responsible for routing requests, authentication validation, service communication, and centralized request handling between client applications and backend services.

## Architecture Overview

```mermaid
flowchart TB

    Client["📱 Web / Mobile Client"]

    Gateway["🚪 API Gateway"]

    Auth["🔐 Authentication Service"]
    Post["📝 Post & Relationship Service"]

    AuthDB[("🐘 PostgreSQL")]
    PostDB[("🐘 PostgreSQL")]

    Client --> Gateway

    Gateway --> Auth
    Gateway --> Post

    Auth --> AuthDB
    Post --> PostDB
```

## Services

| Service                     | Description                                                    |
| --------------------------- | -------------------------------------------------------------- |
| Authentication Service      | User registration, login, JWT authentication, OTP verification |
| Post & Relationship Service | Post management, feeds, follow/unfollow functionality          |
| Chat Service                | Planned                                                        |
| Notification Service        | Planned                                                        |

## Engineering Concepts Demonstrated

* Microservice Architecture
* API Gateway Pattern
* gRPC Communication
* Protocol Buffers
* JWT Authentication
* Clean Architecture
* Dependency Injection
* Repository Pattern
* PostgreSQL Database Design
* Unit Testing

## Features

* JWT-based Authentication
* gRPC Service Communication
* Centralized Request Routing
* REST to gRPC Translation
* Middleware-based Authorization
* Scalable Clean Architecture
* Environment-based Configuration

## Technology Stack

* Golang
* gRPC
* Protocol Buffers
* JWT
* PostgreSQL

## Related Services

### Authentication Service

https://github.com/Anvarsha-k/SocialMediaAuthService

### Post & Relationship Service

https://github.com/Anvarsha-k/SocialMediaPostAndRelationService

## Future Enhancements

* Real-time Chat Service
* Notification Service
* Redis Caching
* Kubernetes Deployment
* Distributed Logging and Monitoring

## Project Goal

This project was developed to explore scalable backend systems, microservice communication patterns, API Gateway design, and service-oriented architecture commonly used in modern social media platforms.
