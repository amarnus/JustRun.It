'use strict';

var app = angular.module('justRunIt', [
    'ngMaterial',
    'ui.router',
    'ui.codemirror'
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

    $stateProvider.state('snippetView', {
        url: '/snippet/:snippet_id',
        controller: 'SnippetController',
        templateUrl: 'partials/snippet.html',
        resolve: {
            snippet: [ 'RemoteSnippetService', '$stateParams', function(RemoteSnippetService, $stateParams) {
                return RemoteSnippetService.getSnippet($stateParams.snippet_id);
            } ]
        }
    });

    $stateProvider.state('snippetEdit', {
        url: '/snippet/:snippet_id/edit',
        controller: 'SnippetController',
        templateUrl: 'partials/snippet.html',
        resolve: {
            snippet: [ 'RemoteSnippetService', '$stateParams', function(RemoteSnippetService, $stateParams) {
                return RemoteSnippetService.getSnippet($stateParams.snippet_id);
            } ]
        }
    });

    $stateProvider.state('snippetEmbed', {
        url: '/snippet/:snippet_id/embed',
        controller: 'SnippetController',
        templateUrl: 'partials/snippet-embed.html',
        resolve: {
            snippet: [ 'RemoteSnippetService', '$stateParams', function(RemoteSnippetService, $stateParams) {
                return RemoteSnippetService.getSnippet($stateParams.snippet_id);
            } ]
        }
    });

    $urlRouterProvider.otherwise('/snippet/add');

} ]);

app.run([ '$log', function($log) {
    $log.debug('JustRun.It client has initialized...');
} ]);