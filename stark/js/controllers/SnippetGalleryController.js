'use strict';

angular.module('justRunIt').controller('SnippetGalleryController', [ 'publicSnippets', 'mySnippets',
    function(publicSnippets, mySnippets) {

    $scope.publicSnippets = publicSnippets;
    $scope.mySnippets = mySnippets;

} ]);