'use strict';

var app = angular.module('justRunIt', [
    'ngMaterial',
    'ui.router'
]);

app.config([ function() {

} ]);

app.run([ '$log', function($log) {
    $log.debug('JustRun.It client has initialized...');
} ]);