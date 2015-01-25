'use strict';

angular.module('justRunIt').controller('SnippetController', [ '$scope', '$log', '$window', '$state',
    '$stateParams', '$mdToast', '$mdDialog', 'snippet', 'LocalSnippetService', 'RemoteSnippetService', 
    function($scope, $log, $window, $state, $stateParams, $mdToast, $mdDialog,
        snippet, LocalSnippetService, RemoteSnippetService) {

    var supportedLanguages = LocalSnippetService.getSupportedLanguages();
    var editorTheme = localStorage.getItem('editorTheme') || 'default';
    var editor;

    // Edit mode or Run-only mode.
    var isAuthor = false; // LocalSnippetService.isCurrentUser(snippet.session_id)

    // Add language info.
    if (snippet.language_code) {
        snippet.langInfo = supportedLanguages[snippet.language_code];
    }

    // Add default title.
    snippet.title = snippet.title || snippet.langInfo.name + ' Snippet';

    $scope.ui = {
        snippet: snippet,
        editorConfig: {
            lineNumbers: true,
            mode: snippet.langInfo.mimeType,
            theme: editorTheme,
            readOnly: !isAuthor ? 'nocursor' : false
        },
        state: {
            runOnly: !isAuthor,
            isSaving: false,
            isRunning: false,
            isForking: false,
            isLinting: false
        }
    };

    function getContentHeight() {
        var viewportHeight = jQuery($window).height();
        return (viewportHeight - 128);
    }

    function setEditorHeight() {
        var editorHeight = getContentHeight();
        editor.setSize('100%', 0.9 * editorHeight);
    }

    $scope.onEditorLoad = function(instance) {
        editor = instance;
        setEditorHeight();
    };

    $(window).resize(setEditorHeight);

    $scope.openThemeChooser = function() {
        $mdDialog.show({
            controller: [ '$scope', function($scope) {
                $scope.editorTheme = editorTheme;
                $scope.themes = [ 'default', 'solarized', '3024-night' ];
                $scope.$watch('editorTheme', function(newValue, oldValue) {
                    if (typeof(oldValue) === 'undefined') {
                        return;
                    }
                    if (newValue === oldValue) {
                        return;
                    }
                    editorTheme = newValue;
                    editor.setOption('theme', newValue);
                    localStorage.setItem('editorTheme', newValue);
                });
            } ],
            templateUrl: 'partials/theme-chooser.html',
            clickOutsideToClose: true,
            escapeToClose: true,
            hasBackdrop: false
        });
    };

    $scope.tagSnippet = function() {
        $mdDialog.show({
            controller: [ '$scope', function($scope) {
                $scope.tagIndices = [ 1, 2, 3, 4, 5 ];
                $scope.tags = snippet.tags;
                $scope.runOnly = !isAuthor;
                $scope.saveTags = function() {
                    var tags = [];
                    for(var i = 0; i < $scope.tags.length; i++) {
                        if ($scope.tags[i]) {
                            tags.push($scope.tags[i]);
                        }
                    }
                    snippet.tags = tags;
                    $mdDialog.hide();
                }
            } ],
            templateUrl: 'partials/tag-list.html',
            clickOutsideToClose: true,
            escapeToClose: true,
            hasBackdrop: true
        });
    };

    $scope.forkSnippet = function() {
        
        function onError(response) {
            $scope.ui.state.isForking = false;
            LocalSnippetService.hideGlobalProgressBar();
            LocalSnippetService.toastError(response.message);
        }

        $scope.ui.state.isForking = true;
        LocalSnippetService.showGlobalProgressBar();
        RemoteSnippetService.forkSnippet(snippet._id)
            .then(function(response) {
                $scope.ui.state.isForking = false;
                LocalSnippetService.hideGlobalProgressBar();
                $state.go('snippetDetail', { snippet_id: response._id });
                LocalSnippetService.toast('You have successfully forked a ' + snippet.langInfo.name + ' snippet.');
            })
            .catch(onError);
    };

    $scope.runSnippet = function() {
        
        function onError(response) {
            $scope.ui.state.isRunning = false;
            LocalSnippetService.hideGlobalProgressBar();
            LocalSnippetService.toastError(response.message);
        }

        $scope.ui.state.isRunning = true;
        LocalSnippetService.showGlobalProgressBar();
        RemoteSnippetService.runSnippet(snippet._id)
            .then(function(response) {
                $scope.ui.state.isRunning = false;
                LocalSnippetService.hideGlobalProgressBar();
                if (!response.status) {
                    onError(response);
                }
            })
            .catch(onError);
    };

    $scope.saveSnippet = function() {

        function onError(response) {
            $scope.ui.state.isSaving = false;
            LocalSnippetService.hideGlobalProgressBar();
            LocalSnippetService.toastError(response.message);
        }

        $scope.ui.state.isSaving = true;
        LocalSnippetService.showGlobalProgressBar();
        RemoteSnippetService.saveSnippet(snippet)
            .success(function(response) {
                $scope.ui.state.isSaving = false;
                LocalSnippetService.hideGlobalProgressBar();
                if (response.status) {
                    var message = 'Your ' + snippet.langInfo.name + ' snippet has been saved.';
                    LocalSnippetService.toast(message);
                }
                else {
                    onError(response);
                }
            })
            .error(onError);
    };

    var contentHeight = 0.9 * getContentHeight();
    var term = new Terminal({
      rows: Math.floor(contentHeight / 14),
      cols: 80,
      screenKeys: true,
      cursorBlink: true 
    });
    term.open(document.getElementById('terminal'));

    // var socket = io.connect();
    // socket.on('connect', function() {
        
    //     term.on('data', function(data) {
    //       socket.emit('data', data);
    //     });

    //     term.on('title', function(title) {
    //       document.title = title;
    //     });

    //     socket.on('data', function(data) {
    //       term.write(data);
    //     });

    //     socket.on('disconnect', function() {
    //       term.destroy();
    //     });
    // });

    Mousetrap.bind([ 'command+s', 'ctrl+s' ], function() {
        $scope.$apply(function() {
            if (!$scope.ui.state.isSaving && isAuthor) {
                $scope.saveSnippet();
            }
            else {
                $log.debug('Ignoring "Save" command...');
            }
        });
        return false;
    });

    Mousetrap.bind([ 'command+r', 'ctrl+r' ], function() {
        $scope.$apply(function() {
            if (!$scope.ui.state.isRunning) {
                $scope.runSnippet();
            }
        });
        return false;
    });

    Mousetrap.bind([ 'command+l', 'ctrl+l' ], function() {
        $scope.$apply(function() {
            if (!$scope.ui.state.isLinting) {
                // TODO: Lint
            }
        });
        return false;
    });

    Mousetrap.bind([ 'command+k', 'ctrl+k' ], function() {
        $scope.$apply(function() {
            if (!$scope.ui.state.isForking) {
                $scope.forkSnippet();
            }
        });
        return false;
    });

} ]);