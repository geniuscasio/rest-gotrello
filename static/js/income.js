$(document).ready(function() {
    $.ajax({
        //http://rest-trello-server.herokuapp.com/
        
        url:"api/v1/income/"
    }).then(function(data) {
        $('.row').append();
        var jsonData = JSON.parse(data);
        var fallSum = 0;
        for (var i = 0; i < jsonData.length; i++) {
            var income = jsonData[i];
            var id = income.id;
            var date = income.date;
            var hint = income.hint;
            var amount = income.amount;

            fallSum += amount;
            var htmlTable =`
            <tr>
                <th scope="row">${id}</th>
                <td>${amount}$</td>
                <td>${date}</td>
                <td>${hint}</td>
                <td>Tags</td>
                <td>${fallSum}$</td>
            </tr>`;
            $('.incomeTable').append(htmlTable);
        }
        
    });
})
function newIncome(form) {
    var id = form.id.value;
    var amount = form.amount.value;
    var date = form.date.value;
    var hint = form.hint.value;

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "api/v1/income/", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        id: id,
        amount: amount,
        date: date,
        hint: hint
    }));
}