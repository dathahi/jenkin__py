pipeline {
    agent {docker { image 'golang:1.26.2-alpine3.23'} }
    
    stages {
        stage('Build') {
            steps {
                sh 'go build -o go-demo .'
            }
        }
        
        stage('Test') {
        	steps {
                sh 'go test ./...'
            }
        }

	    stage('Login'){
	        steps{
                withCredentials([
                    usernamePassword( credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')
                    ]) {
                            sh '''
                                echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                            '''
                        }
            }
	    }
        stage('Push'){
            steps{
                sh 'docker push dathahi/golang-demo:latest'
            }
        }

        stage('Deploy'){
            steps{
                sh './go-demo'
                
                timeout(time: 3, unit: 'MINUTES'){
                    sh './healthcheck.sh'
                }
            }
        }
        post {
            always {
                echo 'pipeline đang chạy'
            }
            success {
                echo 'thành công'
            }
            failure {
                echo 'thất bại'
            }
        }
    }
}
