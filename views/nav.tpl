{{define "nav_banner"}}
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="description" content="">
<script src="/static/js/jquery.js" type="text/javascript"></script>
<script src="/static/js/bootstrap.min.js"></script>
<link href="/static/css/{{.style_name}}.css" rel="stylesheet" type="text/css" media="all" />
<script src="/static/js/headroom.js" type="text/javascript"></script>
<link rel="icon" href="">
<link href="/static/css/bootstrap.css" rel="stylesheet">
<link href="/static/css/product.css" rel="stylesheet">
    <title>{{.head_title}}</title>
</head>
<body>
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
                     <li><a href="/home" class="scroll">主页</a></li>
                        <li><a href="/products" class="scroll">找资金</a></li>
                        <li><a href="/projects" class="scroll">找项目</a></li>
                        <li><a href="/news" class="scroll">行业动态</a></li>
                        <li><a href="/about" class="scroll">关于我们</a></li>
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
{{end}}