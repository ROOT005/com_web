{{template "nav_banner" .}}
{{template "project_model" .}}
 <div class="container choice">
      <ul>
          <li>融资金额:<a class="btn btn-info" href="/projects?rangec = 0" title="" role="button">全部</a>
          {{range .projectranges}}
            <a class="btn btn-info " href="/projects?projectrange={{.ID}}" role="button" id="{{.ID}}">{{.Name}}</a>
          {{end}}
          </li>
          <li>项目类型:
            <a class="btn btn-info classficiation" href="/projects?id=0" title="" role="button" id="0">所有</a>
            {{range .categories}}
            <a class="btn btn-info classficiation" href="/projects?category={{.ID}}" title="" role="button" id="{{.ID}}"> {{.Name}}</a>
            {{end}}
          </li>
      </ul>
    </div><!--choice_end-->
{{template "project_show" .}}
{{template "footer" .}}