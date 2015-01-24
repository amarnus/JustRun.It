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
                var message = 'Your ' + languagePicked.name + ' snippet has been created.';
                LocalSnippetService.toast(message);
                $state.go('snippetEdit', { snippet_id: snippet.snippet_id });
            })
            .catch(function(response) {
                LocalSnippetService.toastError(response.message);
            });
    };

} ]);