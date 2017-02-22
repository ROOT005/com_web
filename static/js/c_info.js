
jQuery(document).ready(function() {

    /*
        Background slideshow
    */
    $.backstretch([
      "images/1.jpg"
    , "images/2.jpg"
    , "images/3.jpg"
    ], {duration: 3000, fade: 750});
    $('.links a.home').tooltip();
    $('.links a.blog').tooltip();

});
