pipeline {
    agent any
    stages {
        stage('Clone Repo') {
            steps {
                git 'https://github.com/yourusername/go-jenkins-service.git'
            }
        }
        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o main .'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }
        stage('Docker Build & Run') {
            steps {
                sh 'docker build -t go-jenkins-service .'
                sh 'docker run -d -p 8081:8081 go-jenkins-service'
            }
        }
    }
}
