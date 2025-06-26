pipeline {
    agent {
        docker {
            image 'golang:1.21'
        }
    }

    environment {
        DOCKER_IMAGE = 'yuvakkrishnans/go-jenkins-service:latest'
    }

    stages {
        stage('Checkout') {
            steps {
                git credentialsId: 'github-token', url: 'https://github.com/yuvakkrishnan/go-jenkins-service.git', branch: 'main'
            }
        }

        stage('Build & Test') {
            steps {
                sh 'go version'
                sh 'go mod tidy'
                sh 'go build -o main .'
                sh 'go test ./...'
            }
        }

        stage('Docker Build & Push') {
            steps {
                withDockerRegistry([credentialsId: 'docker-hub-creds', url: '']) {
                    sh 'docker build -t ${DOCKER_IMAGE} .'
                    sh 'docker push ${DOCKER_IMAGE}'
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl apply -f k8s/service.yaml'
            }
        }
    }
}
