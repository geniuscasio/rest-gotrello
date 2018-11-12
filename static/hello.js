$(document).ready(function() {
    $.ajax({
        url:"http://rest-trello-server.herokuapp.com/income/1"
    }).then(function(data) {
        $('.income-id').append(JSON.parse(data).id);
        $('.income-hint').append(JSON.parse(data).hint);
        $('.income-amount').append(JSON.parse(data).amount);
    });
})
