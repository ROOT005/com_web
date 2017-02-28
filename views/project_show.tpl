{{define "project_show"}}
    <div class="project_show">
    <div class="container">
        {{range .paginator.projects}}
        <div class="col-md-2 project">
            <div class="panel">
                <div class="scroll panel-heading project_title"><h6 class="text-center">{{.Name}}</h6></div>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item"><span>额度:</span><strong>{{.Count}}</strong></li>
                     <li class="list-group-item"><span>行业:</span>{{.Industry}}</li>
                      <li class="list-group-item"><span>地区:</span>{{.Location}}</li>
                     <li class="list-group-item"><span>融资方式:</span>{{.FinanceWay}}</li>
                    <li class="list-group-item"><a class="scroll btn btn-lg btn-block more"  id="{{.ID}}"data-toggle="modal" data-target="#product_info">更多</a></li>
                </ul>
            </div>
        </div>
        {{end}}
    </div>
</div>
<div class="container">
    <nav class="pag">
      <ul class="pagination pagination-lg">
        <li class=""><a href="/projects/?p={{.paginator.firstpage}}">&laquo;</a></li>
        <li> <a href="/projects/?p=1">首页</a></li>
        {{range $index,$page := .paginator.pages}}
        <li  {{if eq $.paginator.currpage $page }}class="active"{{end}}>
            <a href="/projects/?p={{$page}}">{{$page}}</a></li>  
        {{end}}
        <li><a href="/projects/?p={{.paginator.totalpages}}">尾页</a></li>
        <li><a href="/projects/?p={{.paginator.lastpage}}">&raquo;</a></li>
      </ul>
</nav>
</div>
{{end}}