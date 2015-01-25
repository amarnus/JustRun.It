'use strict';

angular.module('justRunIt').factory('LocalSnippetService', [ '$mdToast', '$rootScope', '$cookies',
    function($mdToast, $rootScope, $cookies) {

    var toastPosition = 'top right';
    var sessionKey = 'session_id';

    function getUserSessionId() {
        return $cookies[sessionKey];
    }

    return {

        getUserSessionId: getUserSessionId,

        isCurrentUser: function(sessionId) {
            var currentUserSessionId = getUserSessionId();
            if (!currentUserSessionId) {
                return false;
            }
            return (sessionId === currentUserSessionId);
        },

        showGlobalProgressBar: function() {
            $rootScope.shouldShowGlobalProgressBar = true;
        },

        hideGlobalProgressBar: function() {
            $rootScope.shouldShowGlobalProgressBar = false;
        },

        toast: function(message) {
            var snippetConfig = $mdToast.simple().content(message).position(toastPosition);
            $mdToast.show(snippetConfig);
        },

        toastError: function(message) {
            $mdToast.show({
                controller: [ '$scope', function($scope) {
                    $scope.message = message;
                } ],
                templateUrl: 'partials/toast-error.html',
                position: toastPosition
            });
        },

        getSupportedLanguages: function() {
            return {
                'php': {
                    name: 'PHP',
                    mimeType: 'application/x-httpd-php-open',
                    icon: 'devicon-php-plain'
                },
                'python': {
                    name: 'Python',
                    mimeType: 'text/x-python',
                    icon: 'devicon-python-plain'
                },
                'ruby': {
                    name: 'Ruby',
                    mimeType: 'text/x-ruby',
                    icon: 'devicon-ruby-plain'
                },
                'javascript': {
                    name: 'Javascript',
                    mimeType: 'text/javascript',
                    icon: 'devicon-javascript-plain'
                },
                'go': {
                    name: 'Go',
                    mimeType: 'text/x-go',
                    icon: 'fa fa-code' // TODO: Find a better icon for this guy.
                },
                // 'clojure': {
                //     name: 'Clojure',
                //     mimeType: 'text/x-clojure',
                //     icon: 'fa fa-code'
                // }

            };
        }

    };

} ]);