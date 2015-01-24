'use strict';

angular.module('justRunIt').controller('SnippetController', [ '$scope', '$log', '$window',
    '$stateParams', '$mdToast', 'snippet', 'LocalSnippetService', 'RemoteSnippetService', 
    function($scope, $log, $window, $stateParams, $mdToast,
        snippet, LocalSnippetService, RemoteSnippetService) {

    var supportedLanguages = LocalSnippetService.getSupportedLanguages();

    // Add language info.
    if (snippet.lang) {
        snippet.langInfo = supportedLanguages[snippet.lang];
    }

    // Add default title.
    snippet.title = snippet.title || snippet.langInfo.name + ' Snippet';

    $scope.ui = {
        snippet: snippet,
        editorConfig: {
            lineNumbers: true,
            mode: snippet.langInfo.mimeType,
            autofocus: true
        },
        state: {
            isSaving: false,
            isRunning: false,
            isForking: false
        }
    };

    function getContentHeight() {
        var viewportHeight = jQuery($window).height();
        return (viewportHeight - 128);
    }

    $scope.onEditorLoad = function(instance) {
        var editorHeight = getContentHeight();
        instance.setSize('100%', 0.9 * editorHeight);
    };

    $scope.runSnippet = function() {
        
        function onError(response) {
            $scope.ui.state.isRunning = false;
            LocalSnippetService.toastError(response.message);
        }

        $scope.ui.state.isRunning = true;
        RemoteSnippetService.runSnippet()
            .then(function(response) {
                $scope.ui.state.isRunning = false;
                if (!response.status) {
                    onError(response);
                }
            })
            .catch(onError);
    };

    $scope.saveSnippet = function() {

        function onError(response) {
            $scope.ui.state.isSaving = false;
            LocalSnippetService.toastError(response.message);
        }

        $scope.ui.state.isSaving = true;
        RemoteSnippetService.saveSnippet(snippet)
            .then(function(response) {
                $scope.ui.state.isSaving = false;
                if (response.status) {
                    var message = 'Your ' + snippet.langInfo.name + ' snippet has been saved.';
                    LocalSnippetService.toast(message);
                }
                else {
                    onError(response);
                }
            })
            .catch(onError);
    };

    var contentHeight = 0.9 * getContentHeight();
    var term = new Terminal({
      rows: Math.floor( contentHeight / 18),
      cols: 1,
      screenKeys: true
    });
    
    term.open(document.getElementById('terminal'));

} ]);