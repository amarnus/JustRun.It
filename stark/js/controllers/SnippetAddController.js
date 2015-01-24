'use strict';

angular.module('justRunIt').controller('SnippetAddController', [ '$scope', '$log',
    'LocalSnippetService', 'RemoteSnippetService', function($scope, $log, LocalSnippetService,
    RemoteSnippetService) {

    $scope.languages = LocalSnippetService.getSupportedLanguages();

    $scope.pickLanguage = function(l) {
        $log.log('You picked ' + l.name + '...');
        // Create a snippet.
        // Redirect to the edit page.
    };

} ]);