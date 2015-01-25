'use strict';

var app = angular.module('justRunIt', [
    'ngMaterial',
    'ui.router',
    'ui.codemirror',
    'ngCookies'
]);

app.config([ '$stateProvider', '$urlRouterProvider', '$mdThemingProvider',
    function($stateProvider, $urlRouterProvider, $mdThemingProvider) {

    $mdThemingProvider.theme('default').primaryColor('blue-grey');

    $stateProvider.state('snippetGallery', {
        url: '/snippet?tag&me&page',
        controller: 'SnippetGalleryController',
        templateUrl: 'partials/snippet-gallery.html'
    });

    $stateProvider.state('snippetAdd', {
        url: '/snippet/add',
        onEnter: [ '$mdDialog', function($mdDialog) {
            $mdDialog.show({
                controller: 'SnippetAddController',
                templateUrl: 'partials/snippet-add.html',
                clickOutsideToClose: false,
                escapeToClose: false,
                hasBackdrop: true
            });
        } ],
        onExit: [ '$mdDialog', function($mdDialog) {
            $mdDialog.hide();
        } ]
    });

    $stateProvider.state('snippetDetail', {
        url: '/snippet/:snippet_id?mode',
        controller: 'SnippetController',
        templateUrl: 'partials/snippet.html',
        resolve: {
            snippet: [ 'RemoteSnippetService', '$stateParams', function(RemoteSnippetService, $stateParams) {
                return RemoteSnippetService.getSnippet($stateParams.snippet_id);
            } ]
        }
    });

    $urlRouterProvider.otherwise('/snippet/add');

} ]);

app.run([ '$log', '$rootScope', 'RemoteSnippetService', function($log, $rootScope, RemoteSnippetService) {
    $log.debug('JustRun.It client has initialized...');

    $rootScope.$on('$stateChangeSuccess', function(event, toState, toParams, fromState, fromParams) {
        if (toState === 'snippetDetail') {
            $log.debug('Initializing WebSocket connection...');
            ws = RemoteSnippetService.getWebSocket();
            ws.open();
        }

        if (fromState === 'snippetDetail') {
            $log.debug('Closing WebSocket connection...');
            ws = RemoteSnippetService.getWebSocket();
            ws.close(); 
        }
    });
    
} ]);