$(document).ready(function () {
    $("#register-form").submit(function (event) {
        event.preventDefault();
        $.ajax({
            type: 'POST',
            url: '/register',
            data: {
                'username': $('input[name=username]').val(),
                'password': $('input[name=password]').val(),
                'repassword': $('input[name=repassword]').val(),
            },
            dataType: 'json',
            encode: true
        }).done(function (data) {
        //    {code: 1, message: '注册成功'}
        //    {code: 0, message: '用户名已存在}
            console.log(data);
            if (data.code === 1) {
                window.location.href = '/login';
            }  else {
                alert(data.message);
            }

        }).fail(function (err) {
            console.log(err);
            alert(data.statusText);
        });
    });
});