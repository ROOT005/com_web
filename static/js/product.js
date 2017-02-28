$(document).ready(function() {
   $('.more').click(function(){
        var id = $(this).attr('id');
        $.ajax({
            url: '/products/product_info?id='+id,
            type: 'get',
            dataType: 'json',
            success:function(data){
                $(".description").html(data.Description)
                $(".modal-title").html(data.Name)
                $(".company").html(data.SourceCompany)
                $(".industry").html(data.Industry)
                $(".count").html(data.Count)
                $(".repayway").html(data.RepayWay)
                $(".repaytime").html(data.RepayTime)
                $(".rate").html(data.Rate)
                $(".loantime").html(data.LoanTime)
            }      
        });
    });
    var urlid = GetUrlParameter("id");
    var urlrangec = GetUrlParameter("rangec");
    var len = $(".pagination li").length;
    //修改分页url参数
    for (var i = 0; i < len; i++) {
        var currenturl = $(".pagination a").eq(i).attr('href');
        $('.pagination a').eq(i).attr('href', currenturl+"&id="+urlid+"&rangec="+urlrangec);
    }
    //修改范围url参数
    for (var i = 0; i < $('.rangecount a').length; i++) {
        var currenturl = $(".rangecount a").eq(i).attr('href');
        $('.rangecount a').eq(i).attr('href', currenturl+"&id="+urlid);
        console.log($(".rangecount a").eq(i).attr('href'));
    }
    //修改分类url
    for (var i = 0; i < $('.classficiation a').length; i++){
        var currenturl = $('.classficiation a').eq(i).attr('href');
        $('.classficiation a').eq(i).attr('href', currenturl+"&rangec="+urlrangec);
    }
});

function GetUrlParameter(name)
{
     var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
     var r = window.location.search.substr(1).match(reg);
     if(r!=null)return  unescape(r[2]); return null;
}