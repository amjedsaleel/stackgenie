pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/amjedsaleel/stackgenie.git'
            }
        }

        stage('Build') {
            steps {
                sh 'docker build -t  amjedsaleel/stackgenie:${BUILD_NUMBER} .'
            }
        }

        stage('Push docker image') {
            steps {
               withCredentials([usernamePassword(credentialsId: 'docker-hub', passwordVariable: 'password', usernameVariable: 'username')]) {
                    sh 'docker login -u $username -p $password'
                }

                sh 'docker push amjedsaleel/stackgenie:${BUILD_NUMBER}'
                sh 'docker logout'
            }
        }        
    }
}
