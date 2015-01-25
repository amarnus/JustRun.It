'use strict';

angular.module('justRunIt').factory('RemoteSnippetService', [ '$http', '$q', '$timeout', '$log', 'LocalSnippetService',
    function($http, $q, $timeout, $log, LocalSnippetService) {

    var baseUrl = 'http://gophergala.justrun.it';
    
    var ws = new ReconnectingWebSocket( 'ws://gophergala.justrun.it/ws/io' );

    ws.onopen = function() {
        $log.debug('WebSocket connection was initiated successfully...');
    };

    ws.onclose = function() {
        $log.debug('WebSocket connection was closed...');
    };

    ws.onerror = function(error) {
        $log.error('Error with the WebSocket connection...');
        $log.error(error);
    };

    return {

        getWebSocket: function() {
            return ws;
        },

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
            return $http({
                method: 'POST',
                url: baseUrl + '/snippet/' + snippetId + '/fork'
            });
        },

        runSnippet: function(language, snippetId, code) {
            var sessionId = 'session_' + LocalSnippetService.getUserSessionId() + '_snippet_' + snippetId;
            if (language === 'javascript') {
                language = 'nodejs';
            }
            if (ws) {
                ws.send(JSON.stringify({
                    id: sessionId
                }));
            }
            return $http({
                method: 'POST',
                url: baseUrl + '/run',
                data: {
                    language: language,
                    uid: snippetId,
                    sid: sessionId,
                    snippet: code
                } 
            });
        },

        installDeps: function(language, snippetId, code) {
            var sessionId = 'session_' + LocalSnippetService.getUserSessionId() + '_snippet_' + snippetId;
            if (language === 'javascript') {
                language = 'nodejs';
            }
            if (ws) {
                ws.send(JSON.stringify({
                    id: sessionId
                }));
            }
            return $http({
                method: 'POST',
                url: baseUrl + '/install/deps',
                data: {
                    language: language,
                    uid: snippetId,
                    sid: sessionId,
                    snippet: code
                } 
            });
        },

        lintSnippet: function() {
            var sessionId = 'session_' + LocalSnippetService.getUserSessionId() + '_snippet_' + snippetId;
            if (language === 'javascript') {
                language = 'nodejs';
            }
            return $http({
                method: 'POST',
                url: baseUrl + '/lint/complete',
                data: {
                    language: language,
                    uid: snippetId,
                    sid: sessionId,
                    snippet: code
                } 
            });
        }

    };

} ]);