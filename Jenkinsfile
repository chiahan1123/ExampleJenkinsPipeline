// Configuration map for each sub-project with key being the relative path
// of the sub-project and value being the master branch's Jenkins Pipeline
// job to execute, which allows for independent statistical data for each
// sub-project (eg. Test coverage trend report).
def subProjectConfigs = [
  'android/someApp/': 'master-android-someApp',
  'ios/someApp/': 'master-ios-someApp',
  'mobile/': 'master-mobile',
  'backend/': 'master-backend'
]

def masterBuilders = [:]
def pipelineBuilders = [:]

node {
  // Clean workspace before repository checkout
  cleanWs()

  // Checkout repository and obtain git commit hash
  def scmVars = checkout scm
  def commitHash = scmVars.GIT_COMMIT

  // Dispatcher stage determines the build procedures based on the branch and
  // the git diff.
  stage('Dispatcher') {
    subProjectConfigs.each { key, value ->
      def grepResult = sh(
        returnStdout: true,
        script: "git show --name-only $commitHash | grep $key | wc -m"
      ).trim()
      if (grepResult != '0') {
        if (env.BRANCH_NAME == 'master'
          || env.BRANCH_NAME == 'development'
          || env.BRANCH_NAME == 'production') {
          // TODO: Dispatch to jobs accordingly
          echo "Dispatching to $value"
          masterBuilders[key] = {
            build value
          }
        } else {
          def jenkinsfile = readFile "./${key}Jenkinsfile"
          // Add Jenkinsfile to pipeline builders
          pipelineBuilders[key] = {
            evaluate jenkinsfile
          }
        }
      }
    }
  }

  // Clean workspace at the end to make sure a fresh repository checkout
  cleanWs()
}

if (masterBuilders) {
  echo 'Executing master builders in parallel...'
  parallel masterBuilders
  echo '!!!DONE!!!'
}

if (pipelineBuilders) {
  echo 'Executing pipeline builders in parallel...'
  parallel pipelineBuilders
}
