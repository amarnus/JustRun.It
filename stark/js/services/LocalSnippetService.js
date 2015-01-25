'use strict';

angular.module('justRunIt').factory('LocalSnippetService', [ '$mdToast', '$rootScope', '$cookies',
    function($mdToast, $rootScope, $cookies) {

    var toastPosition = 'top right';

    return {

        isCurrentUser: function(sessionId) {
            if (!$cookies['anonymous_session_id']) {
                return false;
            }
            var currentUserSessionId = $cookies['anonymous_session_id'];
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
                    icon: 'fa fa-code'
                },
                'clojure': {
                    name: 'Clojure',
                    mimeType: 'text/x-clojure',
                    icon: 'fa fa-code'
                }

            };
        }

    };

} ]);