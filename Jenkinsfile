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

        stage('Docker Build & Push') {
            agent {
                docker {
                    image 'docker:24.0.5-dind'
                    args '--privileged -v /var/run/docker.sock:/var/run/docker.sock'
                }
            }
            steps {
                withDockerRegistry([credentialsId: 'docker-hub-creds', url: '']) {
                    sh "docker build -t ${DOCKER_IMAGE} ."
                    sh "docker push ${DOCKER_IMAGE}"
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
