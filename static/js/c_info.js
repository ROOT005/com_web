$(document).ready(function() {

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
    //验证电话
    //验证码
    $('input.verify').change(function(){
        var id = $('.verify').val();
        var captcha_id = $('.cap_id input').val();
        $.ajax({
            type: "get",
            url:"/verify?captcha_id="+ captcha_id+"&id="+id,
            async:true,
            success:function(data){
                console.log(data)
                if (data == "false"){
                    alert("请输入正确的验证码");
                    $(".captcha-img").trigger('click');
                }else{
                    $(".verify").addClass('true');
                }
            }
        });
    });
    $('input.phone').change(function(){
        var phonenum = $(".phone").val();
        $.ajax({
            type:"get",
            url:"/submit/phone?phone="+phonenum,
            async:true,
            success:function(data){
                if (data == "false") {
                    alert("请输入正确的电话");
                }else{
                    $('input.phone').addClass('true');
                }
            }
        });
    });
    $('.submit').click(function(){
        var result =  $('.verify').hasClass('true') && $('.phone').hasClass('true');
        if (!result) {
            alert("请按要求填写资料");
        }
        alert("提交成功");
    });
});
