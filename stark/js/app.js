'use strict';

var app = angular.module('justRunIt', [
    'ngMaterial',
    'ui.router',
    'ui.codemirror'
]);

app.config([ '$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {

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
        } ]
    });

    $stateProvider.state('snippetView', {
        url: '/snippet/:snippet_id',
        controller: 'SnippetViewController',
        templateUrl: 'partials/snippet.html'
    });

    $stateProvider.state('snippetEdit', {
        url: '/snippet/:snippet_id/edit',
        controller: 'SnippetEditController',
        templateUrl: 'partials/snippet.html'
    });

    $stateProvider.state('snippetEmbed', {
        url: '/snippet/:snippet_id/embed',
        controller: 'SnippetViewController',
        templateUrl: 'partials/snippet-embed.html'
    });

    $urlRouterProvider.otherwise('/snippet/add');

} ]);

app.run([ '$log', function($log) {
    $log.debug('JustRun.It client has initialized...');
} ]);