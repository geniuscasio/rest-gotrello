$(document).ready(function() {
    $.ajax({
        url:"http://rest-trello-server.herokuapp.com/income/1"
    }).then(function(data) {
    $('.income-id').append(data.id);
    $('.income-hint').append(data.hint);
    $('.income-amount').append(data.amount);
    });
})
