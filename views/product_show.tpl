{{define "product_show"}}
<div class="product_show">
    <div class="container">
        {{range .paginator.products}}
        <div class="col-md-2 product">
            <div class="panel">
                <div class="scroll panel-heading product_title"><h3 class="text-center">{{.Name}}</h3></div>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item"><span>额度:</span><strong>{{.Count}}</strong></li>
                     <li class="list-group-item"><span >利率:</span><span class="rate_n">{{.Rate}}</span>起</li>
                      <li class="list-group-item"><span>期限:</span>{{.RepayTime}}</li>
                     <li class="list-group-item"><span>还款方式:</span>{{str2html .RepayWay}}</li> 
                    <li class="list-group-item"><a class="scroll btn btn-lg btn-block more"  id="{{.ID}}"data-toggle="modal" data-target="#product_info">更多</a></li>
                </ul>
            </div>
        </div>
        {{end}}
    </div>
</div>
<div class="container-fluid">
    <nav class="pag">
      <ul class="pagination pagination-lg">
        <li class="" ><a href="/products/?p={{.paginator.firstpage}}">&laquo;</a></li>
        <li> <a href="/products/?p=1">首页</a></li>
        {{range $index,$page := .paginator.pages}}
        <li  {{if eq $.paginator.currpage $page }}class="active"{{end}}>
            <a href="/products/?p={{$page}}">{{$page}}</a></li>  
        {{end}}
        <li><a href="/products/?p={{.paginator.totalpages}}">尾页</a></li>
        <li><a href="/products/?p={{.paginator.lastpage}}">&raquo;</a></li>
      </ul>
    </nav>
</div>
{{end}}