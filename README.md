# videotoogo
A container focused microservice based video server in Go, AngularJS and Docker.

# Build

Building the project requires go to be installed and configured as well as Docker. 

The build expects to be able to call Docker to build images so you must be a user in the Docker group or running as root (which, of course, is highly unadvisable).

Once go and Docker are configured pull source:
  go get github.com/dlockamy/videotogo

Or:
```
  git clone https://github.com/dlockamy/videotoogo.git   
```

Enter directory source directory.
(For go get the source will be located at $GOPATH/github.com/dlockamy/videotogo)

Run the build script:
```
  chmod +x ./scripts/build.sh
  ./scripts/build.sh
```

This will build four local docker images:
```
  local/listen
  local/processor
  local/upload
  local/www
 ```
  # Run

To run the docker containers as a series of services:

```
  chmod +x ./scripts/services.sh
  ./scripts/services.sh
```

TODO:
To run the services as a Kubernetes Pod.

Services can be stoped with 

```
  chmod +x ./scripts/nukeservices.sh
  ./scripts/nukeservides.sh
  
```
