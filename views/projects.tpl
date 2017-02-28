{{template "nav_banner" .}}
{{template "project_model" .}}
 <div class="container choice">
      <ul>
          <li class="projectrange">融资金额:<a class="btn btn-info 0" href="/projects?projectrange=0" role="button">全部</a>
          {{range .projectranges}}
            <a class="btn btn-info {{.ID}}" href="/projects?projectrange={{.ID}}" role="button">{{.Name}}</a>
          {{end}}
          </li>
          <li class="category">项目类型:
            <a class="btn btn-info classficiation" href="/projects?category=0" role="button">所有</a>
            {{range .categories}}
            <a class="btn btn-info {{.ID}}" href="/projects?category={{.ID}}"  role="button"> {{.Name}}</a>
            {{end}}
          </li>
      </ul>
    </div><!--choice_end-->
{{template "project_show" .}}
{{template "footer" .}}