'use strict';

angular.module('justRunIt').controller('SnippetGalleryController', [ '$scope', 'publicSnippets',
    'mySnippets', 'LocalSnippetService',
    function($scope, publicSnippets, mySnippets, LocalSnippetService) {

        var supportedLanguages = LocalSnippetService.getSupportedLanguages();

        for (var i = 0; i < publicSnippets.length; i++) {
            publicSnippets[i].langInfo = supportedLanguages[publicSnippets[i].language_code];
        }

        for (var i = 0; i < mySnippets.length; i++) {
            mySnippets[i].langInfo = supportedLanguages[mySnippets[i].language_code];
        }

        $scope.publicSnippets = publicSnippets;
        $scope.publicSnippetsCount = publicSnippets.length;
        $scope.mySnippets = mySnippets;
        $scope.mySnippetsCount = mySnippets.length;

} ]);