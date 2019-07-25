function alert(msg, title) {
    if (!msg) return;
    title = title || "提示信息";
    var htm = '';
    htm += '<div class="modal fade" tabindex="-1" role="dialog" aria-labelledby="exampleModalCenterTitle" aria-hidden="true">';
    htm += '  <div class="modal-dialog modal-dialog-centered" role="document">';
    htm += '    <div class="modal-content">';
    htm += '      <div class="modal-header">';
    htm += '        <h5 class="modal-title" id="exampleModalLabel">' + title + '</h5>';
    htm += '        <button type="button" class="close" data-dismiss="modal" aria-label="Close">';
    htm += '          <span aria-hidden="true">&times;</span>';
    htm += '        </button>';
    htm += '      </div>';
    htm += '      <div class="modal-body">' + msg + '</div>';
    htm += '      <div class="modal-footer">';
    htm += '        <button type="button" class="btn btn-primary" data-dismiss="modal">关闭</button>';
    htm += '      </div>';
    htm += '    </div>';
    htm += '  </div>';
    htm += '</div>';
    $(htm).modal();
}