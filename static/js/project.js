$(document).ready(function() {
    $(".more").click(function() {
        /* Act on the event */
        var id = $(this).attr('id');
        $.ajax({
            url: '/projects/project_info?id='+id,
            type: 'get',
            dataType: 'json',
            success: function(data){
                $(".description").html(data.Description)
                $(".modal-title").html(data.Name)
                $(".count").html(data.Count)
                $(".financeway").html(data.FinanceWay)
                $(".location").html(data.Location)
                $(".financecompany").html(data.FinanceCompany)
                $(".industry").html(data.Industry)
            }
        })
    });
    var urlrangep= GetUrlParameter("projectrange");
    var urlcate = GetUrlParameter("category");
    var len = $(".pagination li").length;

    //修改分页url参数
    for (var i = 0; i < len; i++) {
        var currenturl = $(".pagination a").eq(i).attr('href');
        $('.pagination a').eq(i).attr('href', currenturl+"&projectrange="+urlrangep+"&category="+urlcate);
    }
    //修改范围url参数
    for (var i = 0; i < $('.projectrange a').length; i++) {
        var currenturl = $(".projectrange a").eq(i).attr('href');
        $('.projectrange a').eq(i).attr('href', currenturl+"&id="+urlcate);
    }
    //修改分类url
    for (var i = 0; i < $('.category a').length; i++){
        var currenturl = $('.category a').eq(i).attr('href');
        $('.category a').eq(i).attr('href', currenturl+"&projectrange="+urlrangep);
    }
    //变色
    $("li.projectrange ."+urlrangep).css('background-color', '#2b90d9');
    $("li.category ."+urlcate).css('background-color', '#2b90d9');
});

function GetUrlParameter(name){
     var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
     var r = window.location.search.substr(1).match(reg);
     if(r!=null)return  unescape(r[2]); return null;
}