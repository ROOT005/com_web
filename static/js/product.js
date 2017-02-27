$(document).ready(function() {
   $('.more').click(function(){
    var id = $(this).attr('id');
    $.ajax({
        url: '/products/product_info?id='+id,
        type: 'get',
        dataType: 'json',
        success:function(data){
            $(".description").html(data.Description)
            console.log(data.Description)
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

});
