'use strict';

angular.module('justRunIt').controller('SnippetAddController', [ '$scope', '$log', '$state', '$mdToast',
    'LocalSnippetService', 'RemoteSnippetService', function($scope, $log, $state, $mdToast, LocalSnippetService,
    RemoteSnippetService) {

    $scope.languages = LocalSnippetService.getSupportedLanguages();

    $scope.pickLanguage = function(lang) {
        $log.log('You picked ' + lang + '...');
        var languagePicked = $scope.languages[lang];
        RemoteSnippetService.createSnippet({ lang: lang })
            .then(function(snippet) {
                $mdToast.show($mdToast.simple().content('Your ' + languagePicked.name + ' snippet has been created.'));
                $state.go('snippetEdit', { snippet_id: snippet.snippet_id });
            })
            .catch(function() {
                // Respond to error.
            });
    };

} ]);