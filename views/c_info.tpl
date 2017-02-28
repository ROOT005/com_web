
<!DOCTYPE html>
<html lang="zh_CN">
    <head>
        <meta charset="utf-8">
        <title>信息录入</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="description" content="">
        <meta name="author" content="">
        <script src="/static/js/jquery.js"></script>
        <script src="/static/js/bootstrap.min.js"></script>
        <script src="/static/js/jquery.backstretch.min.js"></script>
        <script src="/static/js/c_info.js"></script>

        <link rel="stylesheet" href="/static/css/bootstrap.css">
        <link rel="stylesheet" href="/static/css/c_info.css">
    </head>
    <body>

        <div class="register-container container">
            <div class="row">
                <div class="iphone col-md-5">
                    <img src="/static/img/iphone.png" alt="">
                </div>
                <div class="register col-md-6">
                    <form action="/home" method="get">
                        <h2><span class="red"><strong>信息录入</strong></span></h2>
                        <label for="firstname">姓名：</label>
                        <input type="text" id="firstname" name="firstname" placeholder="您的称呼....">
                        <label for="username">资金需求：/万元</label>
                        <input type="text" id="username" name="username" placeholder="输入您的资金需求额度...">
                        <label for="email">联系电话：</label>
                        <input type="text" id="email" name="email" placeholder="请输入您的联系电话...">
                        <label for="password">其他</label>
                        <textarea class="form-control" rows="3" placeholder="可以输入您的其他需求..."></textarea>
                        <button type="submit">提交</button>
                    </form>
              </div>
          </div>
      </div>
  </body>
</html>
