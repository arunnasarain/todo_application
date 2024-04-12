# Todo Application
The Todo Application is a simple yet efficient task management tool built with Golang, Gin, and PostgreSQL. It allows users to create, read, update, and delete tasks through a user-friendly web interface. The application is containerized using Docker and orchestrated with Docker-compose for easy deployment. Additionally, Jenkins is integrated for continuous integration and deployment, streamlining the development workflow.

## Features
- **Task Management**: Create, read, update, and delete tasks effortlessly.
- **User-friendly Interface**: Intuitive web interface for seamless task management.
- **Containerized Deployment**: Utilizes Docker and Docker-compose for containerized deployment, ensuring consistency across environments.
- **Continuous Integration**: Jenkins integration for automated testing and deployment, enabling continuous integration and deployment practices.
- **Scalability**: Easily scalable architecture to handle increasing workload demands.

## Technologies Used
- **Golang:** Backend development language for building robust and scalable applications.
- **Gin:** Lightweight HTTP web framework for Golang, providing essential features for building web applications.
- **PostgreSQL:** Powerful open-source relational database management system for storing task data.
- **Docker:** Containerization platform for packaging applications and their dependencies into containers.
- **Docker-compose:** Tool for defining and running multi-container Docker applications.
- **Jenkins:** Automation server for continuous integration and continuous deployment (CI/CD) pipelines.

## Installation
To run the Todo application locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/your_username/todo-application.git

2. Navigate to the project directory:

    ```bash
    cd todo-application

3. Build and start the containers using Docker-compose:

    ```bash
    docker-compose up

Access the APIs in Postman at http://localhost:8080.

