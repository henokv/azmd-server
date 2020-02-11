pipeline {
  agent {
    docker {
      image 'go:1.8'
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
