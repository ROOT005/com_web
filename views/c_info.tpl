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
                <form action="/submit" method="POST">
                    <h2><span class="red"><strong>信息录入</strong></span></h2>
                    <label for="name">姓名:</label>
                    <input type="text" id="name" name="name" placeholder="您的称呼....">
                    <label for="demend">需求:</label>
                    <select name="demend" class="form-control" id="demend">
                        <option value="1">找资金</option>
                        <option value="2">资方入驻</option>
                    </select>
                    <label for="count">额度：/万元</label>
                    <input type="text" id="count" name="count" placeholder="输入您的资金具体额度或者范围...">
                    <label for="phone">联系电话：</label>
                    <input class="phone" type="text" id="phone" name="phone" placeholder="请输入您的联系电话...">
                    <label for="others">其他:</label>
                    <textarea id="others" name="others" class="form-control" rows="3" placeholder="可以输入您的其他需求..."></textarea>
                    <label class="cap_id control-label" for="captcha">{{create_captcha}}</label>
                    <input class="verify form-control" type="text" id="captcha" name="captcha" placeholder="请输入验证码 点击图片刷新">
                    <button class="submit" type="submit">提交</button>
                </form>
          </div>
      </div>
  </div>
</body>
</html>
