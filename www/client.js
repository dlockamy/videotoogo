var app = angular.module("videotogo", ["ngRoute"]);

app.config(function($routeProvider) {
    $routeProvider
    .when("/", {
        templateUrl : "main.htm"
    })
    .when("/player/:videoId", {
        templateUrl : "player.htm",
        controller : "moviePlayerCtr"
    })
    .when("/upload", {
        templateUrl : "upload.html"
    });
});

app.controller('movieListCtr', function($scope, $http) {
    $http.get("http://127.0.0.1:3002/list")
    .then(function(response) {
        var movieList = response.data
        var list = document.getElementById("movie-list")
        list.className="nav nav-pills flex-column";
        movieList.forEach(element => {
           var li = document.createElement("li");
           li.className="nav-item";
           var link = document.createElement("a");
            link.className="nav-link";
           link.href = "#!player/" + element.id;
           link.innerHTML = element.name;
 
           li.appendChild(link);

           list.appendChild(li);
        });
    });
});

app.controller('moviePlayerCtr', function($scope,$routeParams) {
    $scope.playVideo=$routeParams.videoId

    if($scope.playVideo != ""){
        var videoPlayer = document.getElementById("movieScreen");

        videoPlayer.src = "http://127.0.0.1:3002/stream/" + $scope.playVideo;
    }

});