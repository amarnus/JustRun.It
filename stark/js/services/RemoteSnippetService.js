'use strict';

angular.module('justRunIt').factory('RemoteSnippetService', [ '$http', '$q', function($http, $q) {

    return {

        createSnippet: function(languageCode) {
            var deferred = $q.defer();
            deferred.resolve({
                snippet_id: 'foo',
                code: '',
                lang: languageCode
            });
            // deferred.reject({
            //     code: 'TIMEDOUT',
            //     message: 'Could not reach the server.'
            // });
            return deferred.promise;
        },

        getSnippet: function(snippetId) {
            return {
                code: '<?php\n\necho "Hello World"\n',
                lang: 'php'
            };
        },

        saveSnippet: function(snippet) {
            var deferred = $q.defer();
            deferred.resolve({ status: 1 });
            return deferred.promise;
        },

        runSnippet: function() {
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