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

    $(".upvote-link").click(function (event) {
        var $self = $(this);
        var id = $self.data('id');
        $.ajax({
            type: "POST",
            url: "/products/"+ id + "/vote/",
            dataType: 'json',
        }).done(function (response) {
            if (response.code === 1) {
                console.log(response);
                var voteCount = response.data.vote_count;
                var $vote = $self.find(".vote-count");
                $vote.text(voteCount);
                $self.addClass("upvote-active");
                $self.addClass("disabled");
            } else {
                alert(response.message);
            }
        }).fail(function (data) {
            console.log(data);
            alert(data.statusText);
        });
    });
});
