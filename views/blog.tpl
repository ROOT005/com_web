{{template "nav_banner" .}}
  <div class="container choice">
    <ul>
        <li>
        <a class="btn btn-info" href="/news?blogcategory=0" title="" role="button">全部</a>
        {{range .blogcategory}}
		<a class="btn btn-info" href="/news?blogcategory={{.ID}}" title="" role="button">
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
					<div class="head-section text-center">
						<h2>行业动态</h2>
						<span> </span>
					</div>
					<!-- /head-section -->
				<div class="blog-post-grids blog_grids">
					<div class="col-md-4 blog-posts">
						<div class="blog-post">
							<a href="single.html"><img src="/static/img/b1.jpg"></a>
							<a href="single.html" class="blog-title">{{.title}}</a>
							<p class="sub_head">由<a href="#">{{.author}}</a>于{{.time}}发表</p>
							<span></span>
							<p class="content">{{str2html .abstract}}</p>
						</div>
					</div>
					<div class="clearfix"></div>
				</div>
			<div class="pagination text-center">
				<ul>
					<li><a class="prev" href="#">上一页</a></li>
					<li><a href="#">1</a></li>
					<li><a href="#">2</a></li>
					<li><a href="#">3</a></li>
					<li><a href="#">4</a></li>
					<li><a href="#">5</a></li>
					<li><a href="#">6</a></li>
					<li><a href="#">7</a></li>
					<li><a href="#">8</a></li>
					<li><span>.....</span></li>
					<li><a href="#">12</a></li>
					<li><a href="#" class="next">下一页</a></li>
				</ul>
			</div>
			</div>
		</div>
	</div>
    {{template "footer" .}}
