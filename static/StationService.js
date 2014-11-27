(function() {
  'use strict';

  function StationService($http) {
    this.availableStations = [];
    this.getAvailableStations = function() {
      return $http({
        method: 'GET',
        url: '/stations'
      });
    };
    this.init = function() {
      this.getAvailableStations()
        .then(_.bind(onGetAvailableStationsSuccess, this))
    };

    function onGetAvailableStationsSuccess(response) {
      this.availableStations = response.data;
    }
  }

  StationService.$inject = ['$http'];

  angular.module('TrainNetwork')
    .service('StationService', StationService);
})();