pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'yuvakkrishnans/go-jenkins-service:latest'
    }

    stages {
        stage('Checkout') {
            steps {
                git credentialsId: 'github-token', url: 'https://github.com/yuvakkrishnan/go-jenkins-service.git', branch: 'main'
            }
        }

        stage('Build & Test in Go Container') {
            steps {
                script {
                    sh '''
                    docker run --rm -v "$PWD":/app -w /app golang:1.21 sh -c "
                        go mod tidy &&
                        go build -o main . &&
                        go test ./...
                    "
                    '''
                }
            }
        }

        stage('Docker Build & Push') {
            steps {
                script {
                    docker.withRegistry('', 'docker-hub-creds') {
                        sh '''
                        docker build -t $DOCKER_IMAGE .
                        docker push $DOCKER_IMAGE
                        '''
                    }
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh '''
                kubectl apply -f k8s/deployment.yaml
                kubectl apply -f k8s/service.yaml
                '''
            }
        }
    }

    post {
        failure {
            echo 'Pipeline failed.'
        }
        success {
            echo 'Pipeline completed successfully.'
        }
    }
}
