'use strict';

angular.module('justRunIt').factory('LocalSnippetService', [ function() {

    return {

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