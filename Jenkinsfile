node {
  stage('SCM') {
    checkout scm
  }
  stage('Test & Coverage') {
    sh 'go test ./... -coverprofile=coverage.out'
  }
  stage('SonarQube Analysis') {
    def scannerHome = tool 'SonarScanner';
    withSonarQubeEnv() {
      sh "${scannerHome}/bin/sonar-scanner"
    }
  }
}
