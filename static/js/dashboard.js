$(document).ready(function() {
    var message =0
    var id = 1;
     $.ajax({
            url: '/users',
            type: 'GET',
            dataType: 'json',
            success: function(data){
                $.each(data, function(index, val) {
                     $("tbody").append('<tr><td>'+val.ID+'</td><td>'+val.Name+'</td><td>'+val.Count+'</td><td>'+val.Phone+'</td><td>'+val.CreatedAt+'</td></tr>');
                });
            }
        })
    setInterval(function(){
        id = $('tbody tr:first td:first').text();
         $.ajax({
            url: '/users?id='+id,
            type: 'GET',
            dataType: 'json',
            success: function(data){
                if (data.length > 0) {
                $('.message').attr('class', 'message label label-info');
                $('.message a').html(data.length);
                var audio = $("#audio")[0];
                audio.play();
                $.each(data, function(index, val) {
                     $("tbody").append('<tr><td>'+val.ID+'</td><td>'+val.Name+'</td><td>'+val.Count+'</td><td>'+val.Phone+'</td><td>'+val.CreatedAt+'</td></tr>');
                });
                }  
            }
        })
    }, 60000);
});
 