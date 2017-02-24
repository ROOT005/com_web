{{define "product_show"}}
<div class="product_show">
    <div class="container">
        {{range .products}}
        <div class="col-md-2 product">
            <div class="panel">
                <div class="scroll panel-heading product_title"><h3 class="text-center">{{.Name}}</h3></div>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item"><span>额度:</span><strong>{{.Count}}</strong></li>
                     <li class="list-group-item"><span>利率:</span>{{.Rate}}起</li>
                      <li class="list-group-item"><span>期限:</span>{{.RepayTime}}</li>
                     <li class="list-group-item"><span>还款方式:</span>{{.RepayWay}}</li>
                    <li class="list-group-item"><a class="scroll btn btn-lg more btn-block" href="" data-toggle="modal" data-target="#product_info">更多</a></li>
                </ul>
            </div>
        </div>
        {{end}}
    </div>
</div>
{{end}}