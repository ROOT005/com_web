
jQuery(document).ready(function() {

    /*
        Background slideshow
    */
    $.backstretch([
      "/static/img/1.jpg"
    , "/static/img/2.jpg"
    , "/static/img/3.jpg"
    ], {duration: 3000, fade: 750});
    $('.links a.home').tooltip();
    $('.links a.blog').tooltip();

});
