$(document).ready(function () {
    $("#product-form").submit(function (event) {
        event.preventDefault();
        var formData = {
            'name': $('input[name=name]').val(),
            'link': $('input[name=link]').val(),
            'description': $('input[name=description]').val()
        };
        $.ajax({
            type: 'POST',
            url: '/products',
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
