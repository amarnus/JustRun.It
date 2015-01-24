'use strict';

angular.module('justRunIt').factory('RemoteSnippetService', [ '$http', '$q', '$timeout',
    function($http, $q, $timeout) {

    var baseUrl = 'http://gophergala.justrun.it';

    return {

        createSnippet: function(languageCode) {
            return $http({
                method: 'POST',
                url: baseUrl + '/snippets',
                data: {
                    language_code: languageCode
                }
            });
        },

        getSnippets: function() {
            return [];
        },

        getSnippet: function(snippetId) {
            return $http({
                method: 'GET',
                url: baseUrl + '/snippet/' + snippetId
            });
            // return {
            //     language_code: 'php',
            //     title: 'Exchange Selection Sort',
            //     description: 'Simple algorithm to sort a list of numbers.',
            //     tags: [ 'algorithm', 'web' ],
            //     code: '<?php\n\necho "Hello World";\n',
            //     deps: []
            // };
        },

        saveSnippet: function(snippet) {
            delete snippet.langInfo;
            return $http({
                method: 'PUT',
                url: baseUrl + '/snippet/' + snippet._id,
                data: snippet
            });
        },

        forkSnippet: function(snippetId) {
            var deferred = $q.defer();
            $timeout(function() {
                deferred.resolve({
                    _id: 'bar'
                });
            }, 4000);
            return deferred.promise;
        },

        runSnippet: function(snippetId) {
            var deferred = $q.defer();
            deferred.resolve({ status: 0, message: 'Your snippet cannot be run now as our backend isn\'t ready yet' });
            return deferred.promise;
        },

        lintSnippet: function() {
            var deferred = $q.defer();
            deferred.resolve({ status: 1, messages: [] });
            return deferred.promise;
        }

    };

} ]);