$(document).ready(function() {
    $.ajax({        
        url:"api/v1/income/"
    }).then(function(data) {
        var t = document.getElementById("username");
        console.log(t);
        $('.row').append();
        var jsonData = JSON.parse(data);
        console.log(jsonData.session.login);
        t.innerText = jsonData.session.login;
        if (jsonData["session"]["status"] == false) {
            window.alert("Ошибка сессии");
            window.location.href = "/";
            return;
        }
        var fallSum = 0;
        var htmlTable = ''
        if (jsonData == null) {
            htmlTable = `
                <tr>
                <th scope="row">No</th>
                <td>No</td>
                <td>No</td>
                <td>No</td>
                <td>No</td>
                <td>No</td>
            </tr>`;
            $('.incomeTable').append(htmlTable);
        }
        jsonData = jsonData["content"]
        console.log(jsonData)
        for (var i = 0; i < jsonData.length; i++) {
            var income = jsonData[i];
            var id = income.id;
            var date = new Date(income.date);
            var hint = income.hint;
            var amount = income.amount;

            fallSum += amount;
            tags = `<a class="badge badge-warning income-tags">no tags</a>`
            if (!(typeof(income.tags) === "undefined")) {
                tags = ''
                for (var tag = 0; tag < income.tags.length; tag++){
                    tags += `<a class="badge badge-info income-tags">${income.tags[tag].name}</a>`
                }
            }
            dateOptions = {year: 'numeric', month: 'numeric', day: 'numeric' };
            htmlTable += `
            <tr>
                <th scope="row">${id}</th>
                <td>${amount}$</td>
                <td>${date.toLocaleString('ru-RU', dateOptions)}</td>
                <td>${hint}</td>
                <td>${tags}</td>
                <td>${fallSum}$</td>
            </tr>`;
        }
        $('.incomeTable').append(htmlTable);
        
    });
})
function newIncome(form) {
    var id = form.id.value;
    var amount = form.amount.value;
    var date = form.date.value;
    var hint = form.hint.value;
    var tags_t = form.tags.value.split(",");
    var tags = null;
    console.log(tags_t.length)
    if (tags_t[0] != "") {
        var pack;
        var tags = new Array();
        for(var i = 0; i < tags_t.length; i++) {
            pack = new Object();
            pack.name = tags_t[i];
            tags.push(pack);
        }
    }
    console.log(tags);

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "api/v1/income/", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        id: parseInt(id, 10),
        amount: parseInt(amount, 10),
        date: date+"T15:04:05.999999-07:00",
        hint: hint,
        tags: tags
    }));
}

function isLoginOK(data) {
    console.log('data'+data)
}