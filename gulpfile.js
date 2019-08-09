var gulp = require('gulp'),uglify = require("gulp-uglify")
    concat = require("gulp-concat");

    gulp.task('scripts',function(done){
		gulp.src(['./node_modules/jquery/dist/jquery.js'
			,'./static/plugins/moment-v2.22.1/moment.min.js'
			,'./static/plugins/bootstrap/js/bootstrap.min.js'
			,'./static/lte/js/adminlte.min.js'
			,'./static/plugins/layer-v3.1.0/layer.js'
			,'./static/plugins/jquery-sdajax/jquery.sdajax.min.js'
			,'./static/plugins/jquery.cookie.js'
			,'./static/sdtheme/scripts/sdtheme.js'
		//	,'/static/product/scripts/rms.js'
			])
		.pipe(concat('admin.js'))
		.pipe(uglify()) 
		.pipe(gulp.dest('./static/dist/js'))
		  done();
	});

	gulp.task('uglify1',function(done){
		gulp.src('./static/product/scripts/rms.js')
		.pipe(uglify()) 
		.pipe(gulp.dest('./static/dist/js'))
	//	  done();
	});

 //    gulp.task('uglify',function(done){
	// 	gulp.src('./static/dist/js/admin.js')
	// 	.pipe(uglify()) 
	// 	.pipe(gulp.dest('./static/dist/js'))
	// //	  done();
	// });
	