(function() {
  'use strict';

  function StationsCtrl($scope, StationService) {
    $scope.StationService = StationService;

    var map;
    map = new GMaps({
      div: '#map',
      lat: -12.0447259,
      lng: -77.0414517,
      zoom: 6
    });

    map.addMarker({
      lat: -12.0447259,
      lng: -77.0414517,
      title: 'Lima'
    });

    map.addMarker({
      lat: -16.4087455,
      lng: -71.5407495,
      title: 'Arequipa'
    });
  }

  StationsCtrl.$inject = ['$scope', 'StationService'];

  angular.module('TrainNetwork')
    .controller('StationsCtrl', StationsCtrl);
})();