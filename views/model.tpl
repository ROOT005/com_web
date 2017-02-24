{{define "model"}}
<div class="modal fade" id="product_info" tabindex="-1" role="dialog" aria-labelledby="gridSystemModalLabel">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
        {{range .products}}
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h2 class="modal-title" id="gridSystemModalLabel">{{.Name}}</h2>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-md-6">
                  <ul>
                      <li>资金主体：{{.SourceCompany}}</li>
                      <li>投资行业：{{.Industry}}</li>
                      <li>投资金额：{{.Count}}</li>
                      <li>还款方式：{{.RepayWay}}</li>
                  </ul>
              </div>
              <div class="col-md-6">
                  <ul>
                      <li>还款期限：{{.RepayTime}}</li>
                      <li>利率：{{.Rate}}起</li>
                      <li>放款时间：{{.LoanTime}}</li>
                  </ul>
              </div>
            </div>
            <hr style="height: 2px; border-top: 1px solid #5bc0de;">
            {{str2html .Description}}
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <a href="c_info.html" class="btn btn-info" role="button">咨询</a>
          </div>
          {{end}}
        </div><!-- /.modal-content -->
      </div><!-- /.modal-dialog -->
    </div><!-- /.modal -->
{{end}}