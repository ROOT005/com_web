<!DOCTYPE html>
<html>
<head>
<title>新闻动态页</title>
<link href="/static/css/bootstrap.css" rel='stylesheet' type='text/css' />
<!-- Custom Theme files -->
<link href="/static/css/blog.css" rel="stylesheet" type="text/css" media="all" />
<!-- Custom Theme files -->
<script src="/static/js/jquery.js"></script>
<script src="/static/js/bootstrap.min.js" charset="utf-8"></script>
<!-- Custom Theme files -->
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="keywords" content="" />
<script type="application/x-javascript"> addEventListener("load", function() { setTimeout(hideURLbar, 0); }, false); function hideURLbar(){ window.scrollTo(0,1); } </script>
<!--webfont-->
 <script type="text/javascript">
		jQuery(document).ready(function($) {
			$(".scroll").click(function(event){
				event.preventDefault();
				$('html,body').animate({scrollTop:$(this.hash).offset().top},1200);
			});
		});
	</script>
</head>
<body>
	<!-- header-section-starts -->
  <div class="container nav" id="header">
      <div class="navigation">
          <nav class="navbar navbar-default" role="navigation">
              <div class="navbar-header">
                  <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                      <span class="sr-only">Toggle navigation</span>
                      <span class="icon-bar"></span>
                      <span class="icon-bar"></span>
                      <span class="icon-bar"></span>
                  </button>
              </div>
              <div class="collapse navbar-collapse nav-wil" id="bs-example-navbar-collapse-1">
                  <ul class="nav navbar-nav navbar-left link-effect">
                      <li><a href="home.html" class="scroll">主页</a></li>
                      <li><a href="product.html" class="scroll">找资金</a></li>
                      <li><a href="project.html" class="scroll">找项目</a></li>
                      <li><a href="blog.html" class="scroll">行业动态</a></li>
                      <li><a href="about.html" class="scroll">关于我们</a></li>
                      <div class="clearfix"></div>
                  </ul>
              </div><!--/.navbar-collapse-->
          </nav>
      </div>
  </div><!--navgination_end-->
  <div class="container-fluid banner">
      <div  id="carousel-example-generic" class="carousel slide" data-ride="carousel">
          <!-- 下边的点 -->
          <ol class="carousel-indicators visible-lg-inline">
              <li data-target="#carousel-example-generic" data-slide-to="0" class="active"></li>
              <li data-target="#carousel-example-generic" data-slide-to="1"></li>
              <li data-target="#carousel-example-generic" data-slide-to="2"></li>
          </ol>
          <!-- 内容 -->
          <div class="carousel-inner" role="listbox">
              <div class="item active">
                  <img class="img-responsive" src="/static/img/product_banner.jpg" alt="">
                  <div class="carousel-caption">
                  </div>
              </div>
              <div class="item">
                  <img class="img-responsive" src="/static/img/product_banner.jpg" alt="">
                  <div class="carousel-caption">
                  </div>
              </div>
              <div class="item">
                  <img class="img-responsive" src="/static/img/product_banner.jpg" alt="">
                  <div class="carousel-caption">
                  </div>
              </div>
          </div>
          <!-- 控制 -->
          <a class="left carousel-control" href="#carousel-example-generic" role="button" data-slide="prev">
          </a>
          <a class="right carousel-control" href="#carousel-example-generic" role="button" data-slide="next">
          </a>
      </div>
  </div>
  <div class="container choice">
    <ul>
        <li><a class="btn btn-info" href="" title="" role="button">行业动态</a><a class="btn btn-info" href="" title="" role="button">融资攻略</a></li>
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
							<p>{{str2html .content}}</p>
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
</body>
</html>
