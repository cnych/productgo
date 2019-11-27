$(document).ready(function () {
    $("#login-form").submit(function (event) {
        event.preventDefault();
        var formData = {
            'username': $('input[name=username]').val(),
            'password': $('input[name=password]').val()
        };
        $.ajax({
            type: 'POST',
            url: '/login',
            data: formData,
            dataType: 'json',
        }).done(function (data) {
            console.log(data);
            if (data.code === 1) {
                window.location.href = '/';
            } else {
                alert(data.message);
            }
        }).fail(function (data) {
            console.log(data);
            alert(data.statusText);
        });
    });
});
