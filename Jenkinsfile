pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'todo_app'
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/arunnasarain/todo_application'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o todo_app .'
            }
        }

        stage('Package') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE .'
            }
        }

        stage('Deploy') {
            steps {
                 sh 'docker run -d -p 8080:8080 todo_app'
            }
        }
    }
}
