(function() {
  'use strict';

  function config($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'stations.html',
        controller: 'StationsCtrl'
      })
  }

  config.$inject = ['$routeProvider'];

  angular.module('TrainNetwork')
    .config(config)
})();