$(document).ready(function() {
   var blogcate =  GetUrlParameter("blogcategory");
   //变色
   $(".blogcategory ."+blogcate).css('background-color', '#2b90d9');
   $(".content img").attr('class', 'img-responsive');
});
function GetUrlParameter(name){
     var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
     var r = window.location.search.substr(1).match(reg);
     if(r!=null)return  unescape(r[2]); return null;
}