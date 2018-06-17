# videotoogo
A container focused microservice based video server in Go, AngularJS and Docker.

# Build

Building the project requires go to be installed and configured as well as Docker. 

The build expects to be able to call Docker to build images so you must be a user in the Docker group or running as root (which, of course, is highly unadvisable).

Once go and Docker are configured pull source:
  go get github.com/dlockamy/videotogo

Or:
  git clone https://github.com/dlockamy/videotoogo.git   

Enter directory source directory.
(For go get the source will be located at $GOPATH/github.com/dlockamy/videotogo)

Run the build script:

  chmod +x ./script/build.sh
  ./script/build.sh
  
This will build four local docker images:

  local/listen
  local/processor
  local/upload
  local/web
  
  
