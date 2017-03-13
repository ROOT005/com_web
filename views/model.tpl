{{define "product_model"}}
<div class="modal fade" id="product_info" tabindex="-1" role="dialog" aria-labelledby="gridSystemModalLabel">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h2 class="modal-title" id="gridSystemModalLabel"></h2>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-md-6">
                  <ul>
                      <li>资金主体：<span class="company"></span> </li>
                      <li>额度：<span class="count"></span></li>
                      <li>还款方式：<span class="repayway"></span></li>
                  </ul>
              </div>
              <div class="col-md-6">
                  <ul>
                      <li>还款期限：<span class="repaytime"></span></li>
                      <li>利率：<span class="rate"></span>起</li>
                      <li>放款时间：<span class="loantime"></span></li>
                  </ul>
              </div>
            </div>
            <div class="description">
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <a href="/consult" class="btn btn-info" role="button">咨询</a>
          </div>
        </div><!-- /.modal-content -->
      </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->
{{end}}

{{define "project_model"}}
<div class="modal fade" id="product_info" tabindex="-1" role="dialog" aria-labelledby="gridSystemModalLabel">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h2 class="modal-title" id="gridSystemModalLabel"></h2>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-md-6">
                  <ul>
                      <li>额度：<span class="count"></span> </li>
                      <li>融资方式：<span class="financeway"></span></li>
                      <li>位置：<span class="location"></span></li>
                  </ul>
              </div>
              <div class="col-md-6">
                  <ul>
                      <li>融资主体：<span class="financecompany"></span></li>
                       <li>所在行业：<span class="industry"></span></li>
                  </ul>
              </div>
            </div>
            <div class="description">
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <a href="/consult" class="btn btn-info" role="button">咨询</a>
          </div>
        </div><!-- /.modal-content -->
      </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->
{{end}}
