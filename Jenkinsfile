pipeline {
  agent {
    docker {
      image 'golang:1.8'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh '''go get -d -v ./...
go install -v ./...'''
      }
    }

  }
}