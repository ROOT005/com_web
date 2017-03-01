{{template "nav_banner" .}}
 <div class="container choice">
	<ul>
	    <li class="blogcategory">
	    <a class="btn btn-info 0" href="/news?blogcategory=0" role="button">全部</a>
	    {{range .blogcategory}}
		<a class="btn btn-info {{.ID}}" href="/news?blogcategory={{.ID}}" role="button">
			{{.Name}}
		</a>
		{{end}}	
	    </li>
	</ul>
</div><!--choice_end-->
<div class="content">
	<div class="container">
		<div class="blog-posts-section">
			<!-- head-section -->
			
			<!-- /head-section -->
			
			<div class="blog-post-grids blog_grids">
				{{range .paginator.blogs}}
				<div class="col-md-3 blog-posts">
					<div class="blog-post">	
						<a href="/news/newinfo?id={{.ID}}"><img src="{{ .Mainimage.URL}}"></a>
						<a href="/news/newinfo?id={{.ID}}" class="blog-title">{{.Title}}</a>
						<p class="sub_head">由<a href="#">{{.Author}}</a> 于{{date .CreatedAt "Y-m-d"}}发表</p>
						<span></span>
						<p class="content">{{str2html .Abstract}}</p>
                        <a class="more" href="/news/newinfo?id={{.ID}}" title="">详细</a>
					</div>
				</div>
				{{end}}
			</div>
		</div>
	</div>
</div>
<div class="container">
    <nav class="pag">
      <ul class="pagination pagination-lg">
        <li class=""><a href="/news/?p={{.paginator.firstpage}}">&laquo;</a></li>
        <li> <a href="/news/?p=1">首页</a></li>
        {{range $index,$page := .paginator.pages}}
        <li  {{if eq $.paginator.currpage $page }}class="active"{{end}}>
            <a href="/news/?p={{$page}}">{{$page}}</a>
        </li>  
        {{end}}
        <li><a href="/news/?p={{.paginator.totalpages}}">尾页</a></li>
        <li><a href="/news/?p={{.paginator.lastpage}}">&raquo;</a></li>
      </ul>
    </nav>
</div>
{{template "footer" .}}