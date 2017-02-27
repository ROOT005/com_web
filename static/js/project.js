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
});