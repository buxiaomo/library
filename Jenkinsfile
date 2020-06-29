pipeline {
    agent {
        label 'jnlp-slave'
    }

    environment {
        PROJECT_NAME="buxiaomo"

        REPOSITORY_URL = "https://github.com/buxiaomo/library.git"

        REGISTRY_PROTOCOL="https"
        REGISTRY_URL = "hub.xiaomo.io"

        Dockerfile="Dockerfile"
    }

    options {
        buildDiscarder(logRotator(numToKeepStr: '15'))
    }

    stages {
        stage('checkout') {
            steps {
                git url: "${env.REPOSITORY_URL}"
            }
        }

        stage('build') {
            steps {
                container('dockerd') {
                    sh label: 'Docker Image', script: "docker build -t ${env.REGISTRY_URL}/${env.PROJECT_NAME}/${JOB_NAME}:${BUILD_ID} -f ${env.Dockerfile} ."
                }
            }
        }

        stage('push') {
            steps {
                container('dockerd') {
                    sh label: 'Docker', script: "docker push ${env.REGISTRY_URL}/${env.PROJECT_NAME}/${JOB_NAME}:${BUILD_ID}"
                }
            }
        }

        stage('deploy dev') {
            when {
                not { branch 'master' }
            }
            steps {
                container('kubectl') {
                    sh label: 'namespace', script: "kubectl get ns ${env.PROJECT_NAME}-dev || kubectl create ns ${env.PROJECT_NAME}-dev"
                    sh label: 'change tag', script: "sed -i \"s#APPLICATION_NAME#${JOB_NAME}#\" ./manifests/deployment.yaml"
                    sh label: 'change tag', script: "sed -i \"s#PROJECT_NAME#${env.PROJECT_NAME}#\" ./manifests/deployment.yaml"
                    sh label: 'change tag', script: "sed -i \"s#REGISTRY_URL#${env.REGISTRY_URL}#\" ./manifests/deployment.yaml"
                    sh label: 'change tag', script: "sed -i \"s#IMAGE_TAG#${BUILD_ID}#\" ./manifests/deployment.yaml"
                    sh label: 'deploy', script: "kubectl apply -n ${env.PROJECT_NAME}-dev -f ./manifests/deployment.yaml"
                }
            }
        }

        stage('deploy prod') {
            when { branch 'master' }
            steps {
                container('kubectl') {
                    sh label: 'namespace', script: "kubectl get ns ${env.PROJECT_NAME}-uat || kubectl create ns ${env.PROJECT_NAME}-uat"
                    sh label: 'change tag', script: "sed -i \"s#IMAGE_TAG#${BUILD_ID}#\" ./manifests/deployment.yaml"
                    sh label: 'deploy', script: "kubectl apply -n ${env.PROJECT_NAME}-prod -f ./manifests/deployment.yaml"
                }
            }
        }
    }
}
