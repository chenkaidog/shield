$(document).ready(function() {
  $('#loginBut').click(function() {
    var username = $('#username').val();
    var password = $('#password').val();

    $.post('/login', 
      { 
        username: username, 
        password: password 
      },
      function(data, status, xhr) {
        if (status=='success') {
          var code = data.code;
          var success = data.success;
          var msg = data.msg;

          if (success) {
            var csrfToken = xhr.getResponseHeader('X-Csrf-Token');
            localStorage.setItem('csrfToken', csrfToken);

            var urlParams = new URLSearchParams(window.location.search);
            var redirectUrl = decodeURIComponent(urlParams.get('redirect'));
            if (redirectUrl.length <= 0) {
              window.location.href = redirectUrl;
            } else {
              window.location.href = "index/home"
            }
          } else {
            alert("error code: " + code + "\nmsg: " + msg);
          }
        } else {
          alert("network err, please retry!");
        }
      }
    )
  });
});