#!/usr/bin/env node

var gulp = require('gulp');
var wiredep = require('wiredep').stream;
var connect = require('gulp-connect');

//-- Wire installed bower components to the conduit file.
gulp.task('wiredep', function () {
  gulp.src('./index.html')
    .pipe(wiredep())
    .pipe(gulp.dest('.'));
});

//-- Start a static file server.
gulp.task('serve', [ 'watch' ], function() {
    connect.server({
        root: '.',
        livereload: true
    });
});

//-- Reload the static file server.
gulp.task('serve:reload', function() {
    gulp.src([ './index.html' ])
        .pipe(connect.reload());
});

//-- Watch for changes and reload the static file server.
gulp.task('watch', function() {
    gulp.watch([
        'index.html',
        'css/**/*.css',
        'js/**/*.js',
        'partials/**/*.html'
    ], [ 'serve:reload' ]);
});