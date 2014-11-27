(function() {
  'use strict';

  function run(StationService) {
    StationService.init();
  }

  run.$inject = ['StationService'];

  angular.module('TrainNetwork')
    .run(run);
})();