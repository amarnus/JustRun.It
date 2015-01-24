#!/usr/bin/env node

var gulp = require('gulp');
var wiredep = require('wiredep').stream;

//-- Wire installed bower components to my conduit file.
gulp.task('wiredep', function () {
  gulp.src('./index.html')
    .pipe(wiredep())
    .pipe(gulp.dest('.'));
});