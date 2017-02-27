{{template "nav_banner" .}}
{{template "model" .}}
    <div class="container choice">
      <ul>
          <li>融资金额:<a class="btn btn-info" href="" title="" role="button">0-50万</a><a class="btn btn-info" href="" title="" role="button">50-100万</a></li>
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