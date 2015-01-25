'use strict';

angular.module('justRunIt').controller('SnippetAddController', [ '$scope', '$log', '$state', '$mdToast',
    'LocalSnippetService', 'RemoteSnippetService', function($scope, $log, $state, $mdToast, LocalSnippetService,
    RemoteSnippetService) {

    $scope.ui = { shouldProgressBarShow: false };

    $scope.languages = LocalSnippetService.getSupportedLanguages();

    $scope.pickLanguage = function(lang) {
        $log.log('You picked ' + lang + '...');
        var languagePicked = $scope.languages[lang];
        $scope.ui.shouldProgressBarShow = true;
        RemoteSnippetService.createSnippet(lang)
            .success(function(response) {
                $scope.ui.shouldProgressBarShow = false;
                var message = 'Your ' + languagePicked.name + ' snippet has been created.';
                LocalSnippetService.toast(message);
                $state.go('snippetDetail', { snippet_id: response.result.snippet_id });
            })
            .error(function(response) {
                $scope.ui.shouldProgressBarShow = false;
                if (response && response.message) {
                    LocalSnippetService.toastError(response.message);
                }
                else {
                    LocalSnippetService.toastError('You cannot create a new snippet right now. Please try again later.');
                }
            });
    };

} ]);