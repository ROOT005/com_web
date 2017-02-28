{{template "nav_banner" .}}
{{template "product_model" .}}
    <div class="container choice">
      <ul>
          <li class="rangecount" >融资金额:
             <a class="btn btn-info 0" href="/products?rangec=0" role="button" id="0">全部</a>
          {{range .rangecount}}
            <a class="btn btn-info {{.ID}}" href="/products?rangec={{.ID}}" role="button">{{.Name}}</a>
          {{end}}
          </li>
          <li class="classficiation">产品类型:
            <a class="btn btn-info 0" href="/products?id=0" role="button">全部</a>
            {{range .classfication}}
            <a class="btn btn-info {{.ID}} " href="/products?id={{.ID}}" role="button" id="{{.ID}}"> {{.Name}}</a>
            {{end}}
          </li>
      </ul>
    </div><!--choice_end-->
{{template "product_show" .}}
{{template "footer" .}}