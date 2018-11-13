$(document).ready(function() {
    $.ajax({
        url:"http://rest-trello-server.herokuapp.com/income/"
    }).then(function(data) {
        $('.row').append();
        var jsonData = JSON.parse(data);
        for (var i = 0; i < jsonData.length; i++) {
            var income = jsonData[i];
            console.log(income);
            var id = income.id;
            var date = income.date;
            var hint = income.hint;
            var amount = income.amount;
            var htmlTable =`
            <tr>
                <th scope="row">${id}</th>
                <td>${amount}$</td>
                <td>${date}</td>
                <td>${hint}</td>
                <td>Tags</td>
                <td>235$</td>
            </tr>`;
            $('.incomeTable').append(htmlTable);
        }
        
    });
})