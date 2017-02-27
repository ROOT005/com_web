{{template "nav_banner" .}}
{{template "product_model" .}}
    <div class="container choice">
      <ul>
          <li>融资金额:<a class="btn btn-info" href="/products?rangec = 0" title="" role="button">全部</a>
          {{range .rangecount}}
            <a class="btn btn-info " href="/products?rangec={{.ID}}" role="button" id="{{.ID}}">{{.Name}}</a>
          {{end}}
          </li>
          <li>产品类型:
            <a class="btn btn-info classficiation" href="/products?id=0" title="" role="button" id="0">所有</a>
            {{range .classfication}}
            <a class="btn btn-info classficiation" href="/products?id={{.ID}}" title="" role="button" id="{{.ID}}"> {{.Name}}</a>
            {{end}}
          </li>
      </ul>
    </div><!--choice_end-->
{{template "product_show" .}}
{{template "footer" .}}