# videotoogo
A container focused microservice based video server in Go, AngularJS and Docker.

# Build

Building the project requires go to be installed and configured as well as Docker. 

The build expects to be able to call Docker to build images so you must be a user in the Docker group or running as root (which, of course, is highly unadvisable).

Once go and Docker are configured pull source:
```
  go get github.com/dlockamy/videotogo
```
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
  chmod +x ./scripts/killservices.sh
  ./scripts/killservies.sh 
```

Cleanup all containers:

```
  chmod +x ./scripts/nukeservices.sh
  ./scripts/nukeservides.sh 
```


# Use

While four services are started, only three have public facing endpoints. Each service handles a single task and is intended to sit behind an ingress server that would present the group of services as a single api as well as handle authentication and https transport.

```
localhost:3000
The user interface and AngularJS application allow for the upload, browsing and playback of video files.
```
```
localhost:3001 
This is the upload service that listens POST upload.
```

```
localhost:3002 
Listens for browse requests and returns a list of available video files. File meta data is kept in a .json file that is used as an ad hoc database.
```
```
Processor: 
While it doesn't have a public interface, this process is run in a container just as the other processes. This service uses Linux fsnotify to listen for changes to an upload directory and once a new file arrives creates a hash based on file contents (to prevent duplicate video files from being stored) and move the file to the permament storage folder.
```
