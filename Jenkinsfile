pipeline {

    agent any

    environment {
        GOCACHE = "${WORKSPACE}/.cache/go-build"
        GOPATH  = "${WORKSPACE}/.go"
        CGO_ENABLED = '0'

        IMAGE_NAME = 'dathahi/golang-demo'
        IMAGE_TAG  = 'latest'
    }

    stages {

        stage('Prepare') {
            steps {
                sh '''
                    mkdir -p $GOCACHE
                    mkdir -p $GOPATH

                    go version
                    docker version
                '''
            }
        }

        stage('Test') {
            steps {
                sh '''
                    go test ./...
                '''
            }
        }

        stage('Build Binary') {
            steps {
                sh '''
                    go build -o go-demo .
                '''
            }
        }

        stage('Build Image') {
            steps {
                sh '''
                    docker build \
                      -t $IMAGE_NAME:$IMAGE_TAG \
                      .
                '''
            }
        }

        stage('Login DockerHub') {
            steps {

                withCredentials([
                    usernamePassword(
                        credentialsId: 'dockerhub-creds',
                        usernameVariable: 'DOCKER_USER',
                        passwordVariable: 'DOCKER_PASS'
                    )
                ]) {

                    sh '''
                        echo "$DOCKER_PASS" | docker login \
                            -u "$DOCKER_USER" \
                            --password-stdin
                    '''
                }
            }
        }

        stage('Push Image') {
            steps {
                sh '''
                    docker push $IMAGE_NAME:$IMAGE_TAG
                '''
            }
        }
    }

    post {

        always {
            cleanWs()
        }

        success {
            echo 'Pipeline success'
        }

        failure {
            echo 'Pipeline failed'
        }
    }
}