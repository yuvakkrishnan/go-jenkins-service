pipeline {
    agent any

    stages {
<<<<<<< HEAD
        stage('Checkout') {
=======
        
        stage('Checkout') {
            steps {
                git credentialsId: 'github-token', url: 'https://github.com/yuvakkrishnan/go-jenkins-service.git', branch: 'main'
            }
        }
        stage('Clone Repo') {
>>>>>>> 562e7156a22ba4c89edb91bcfca11893664a4778
            steps {
                git credentialsId: 'github-token', url: 'https://github.com/yuvakkrishnan/go-jenkins-service.git', branch: 'main'
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
