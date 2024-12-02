pipeline {
    agent any

    environment {
        REPO_URL = 'https://github.com/quocbao64/motel-booking-api.git'
        IMAGE_NAME = 'quocbao64/motel-booking-api'
    }

    stages {
        stage('Build Docker Image') {
            steps {
                script {
                    echo 'Building Docker image...'
                    sh 'docker build -t ${IMAGE_NAME} .'
                }
            }
        }

        stage('Run Unit Tests') {
            steps {
                script {
                    echo 'Running unit tests...'
                    sh 'docker run --rm ${IMAGE_NAME} go test ./...'
                }
            }
        }



        stage('Push Docker Image') {
            steps {
                script {
                    echo 'Pushing Docker image to Docker Hub...'
                    withCredentials([usernamePassword(credentialsId: 'docker-hub-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                        sh 'docker login -u $DOCKER_USER -p $DOCKER_PASS'
                    }
                    sh 'docker tag ${IMAGE_NAME} ${IMAGE_NAME}:latest'
                    sh 'docker push ${IMAGE_NAME}:latest'
                }
            }
        }

        stage('Deploy to Production') {
            steps {
                script {
                    echo 'Deploying application...'
                    withCredentials([usernamePassword(credentialsId: 'docker-hub-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                        sh 'docker login -u $DOCKER_USER -p $DOCKER_PASS'
                    }
                    sh 'docker pull ${IMAGE_NAME}:latest'
                    sh 'docker stop motel_booking_api || true && docker rm motel_booking_api || true'
                    sh 'nohup docker run -d -p 3006:3006 --name motel_booking_api ${IMAGE_NAME}:latest >/dev/null 2>&1 &'
                }
            }
        }

        stage('Verify Deployment') {
            steps {
                script {
                    sleep(5) // Wait a few seconds to allow the container to start
                    sh 'curl -f http://localhost:3006 || exit 1'
                }
            }
        }
    }

    post {
        success {
            echo 'Deployment completed successfully.'
        }
        failure {
            echo 'Deployment failed. Please check the logs.'
        }
    }
}