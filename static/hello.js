$(document).ready(function() {
    $.ajax({
        url:"http://rest-trello-server.herokuapp.com/income/"
    }).then(function(data) {
        $('.row').append();
        var jsonData = JSON.parse(data);
        for (var i = 0; i < jsonData.length; i++) {
            var income = jsonData[i];
            console.log(income.hint);
            var id = income.id;
            var hint = income.hint;
            var amount = income.amount;
            var html = '<div class="col-sm - 3">'
                + '< a href = "#" class="list-group-item list-group-item-action flex-column align-items-start active" >'
                + '<div class="d-flex w-30 justify-content-between">'
                + '        <h5 class="mb-1">' + amount + '$</h5>'
                + '        <small>3 days ago</small>'
                + '    </div>'
                + '    <p class="mb-1">' + hint + '</p>'
                + '    <small>Donec id elit non mi porta.</small>'
                + '    </a >'
                + ' </div >'
            $('.row').append(html);
        }
        
    });
})