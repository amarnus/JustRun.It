'use strict';

angular.module('justRunIt').controller('SnippetGalleryController', [ '$scope', 'publicSnippets', 'mySnippets',
    function($scope, publicSnippets, mySnippets) {

    $scope.publicSnippets = publicSnippets;
    $scope.mySnippets = mySnippets;

} ]);