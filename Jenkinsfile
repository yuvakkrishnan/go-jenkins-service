pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/yourusername/go-jenkins-service.git'
            }
        }

        stage('Build Go App') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o main .'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t go-jenkins-service .'
            }
        }

        stage('Docker Run') {
            steps {
                sh 'docker run -d -p 8081:8081 go-jenkins-service'
            }
        }
    }
}
