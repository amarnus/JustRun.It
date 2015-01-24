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

        saveSnippet: function() {

        },

        runSnippet: function() {

        },

        installSnippetDeps: function() {

        },

        lintSnippet: function() {

        },

        removeSnippet: function() {
            
        }

    };

} ]);