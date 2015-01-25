'use strict';

angular.module('justRunIt').factory('RemoteSnippetService', [ '$http', '$q', '$timeout', 'LocalSnippetService',
    function($http, $q, $timeout, LocalSnippetService) {

    var baseUrl = 'http://gophergala.justrun.it';

    return {

        createSnippet: function(languageCode) {
            return $http({
                method: 'POST',
                url: baseUrl + '/snippets',
                data: JSON.stringify({
                    language_code: languageCode
                })
            });
        },

        getSnippets: function(opts) {
            var deferred = $q.defer();
            $http({
                method: 'GET',
                url: baseUrl + '/snippets',
                params: opts
            })
            .success(function(data) {
                deferred.resolve(data);
            })
            .error(function(data) {
                deferred.reject(data);
            });
            return deferred.promise;
        },

        getSnippet: function(snippetId) {
            var deferred = $q.defer();
            $http({
                method: 'GET',
                url: baseUrl + '/snippet/' + snippetId
            })
            .success(function(data) {
                // TODO: Check data.status.
                deferred.resolve(data.result);
            })
            .error(function(data) {
                deferred.reject(data);
            });
            return deferred.promise;
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
            var snippetClone = _.cloneDeep(snippet);
            delete snippetClone.langInfo;
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

        runSnippet: function(snippetId, code) {
            return $http({
                method: 'POST',
                url: baseUrl + '/run',
                data: {
                    uid: snippetId,
                    sid: LocalSnippetService.getUserSessionId() || 'foo', // TODO: Remove this.
                    snippet: code
                } 
            });
        },

        lintSnippet: function() {
            var deferred = $q.defer();
            deferred.resolve({ status: 1, messages: [] });
            return deferred.promise;
        }

    };

} ]);