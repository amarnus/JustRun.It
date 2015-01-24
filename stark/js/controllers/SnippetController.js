'use strict';

angular.module('justRunIt').controller('SnippetController', [ '$scope', '$log', '$window',
    '$stateParams', 'snippet', 'LocalSnippetService', function($scope, $log, $window, $stateParams, 
        snippet, LocalSnippetService) {

    var supportedLanguages = LocalSnippetService.getSupportedLanguages();

    // Add language info.
    if (snippet.lang) {
        snippet.langInfo = supportedLanguages[snippet.lang];
    }

    $scope.ui = {
        snippet: snippet,
        editorConfig: {
            lineNumbers: true,
            mode: snippet.langInfo.mimeType,
            autofocus: true
        }
    };

    function getContentHeight() {
        var viewportHeight = jQuery($window).height();
        return (viewportHeight - 2 * 64 - 1);
    }

    $scope.onEditorLoad = function(instance) {
        var editorHeight = getContentHeight();
        instance.setSize('100%', editorHeight);
    };

    var term = new Terminal({
      rows: Math.ceil(getContentHeight() / 18),
      cols: 1,
      screenKeys: true
    });
    
    term.open(document.getElementById('terminal'));

} ]);