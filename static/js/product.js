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
    
    $(document).on("click", ".upvote-link", function (event) {
    // $(".upvote-link").click(function (event) { // 这种方式不能用于后面append的元素点击
        var $self = $(this);
        var id = $self.data('id');
        $.ajax({
            type: "POST",
            url: "/products/"+ id + "/vote/",
            dataType: 'json',
        }).done(function (response) {
            if (response.code === 1) {
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

    $(".btn-more").click(function (event) {
        var lastDate = $(".products-content .date").last().text();
        console.log(lastDate);
        $.ajax({
            type: "GET",
            url: "/?last_dt="+ lastDate
        }).done(function(response) {
            $(".products-content").append(response);
        }).fail(function (data) {
            console.log(data);
            alert(data.statusText);
        });
    });
});
