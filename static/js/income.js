$(document).ready(function() {
    $.ajax({        
        url:"api/v1/income/"
    }).then(function(data) {
        var t = document.getElementById("username");
        var userName = getCookie("userName")

        $('.row').append();
        var jsonData = JSON.parse(data);
        t.innerText = userName
        var fallSum = 0;
        var htmlTable = ''
        console.log(jsonData)
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
        console.log()
        const incomeStyle = "income";
        const outcomeStyle = "outcome";
        for (var i = 0; i < jsonData.length; i++) {
            var income = jsonData[i];
            var id = income.id;
            var date = new Date(income.date);
            var hint = income.hint;
            var amount = income.amount;
            var rowStyle = incomeStyle;
            fallSum += amount;
            console.log(id, date, hint, amount, rowStyle)
            if (amount < 0) { rowStyle = outcomeStyle; }

            tags = `<a class="badge badge-warning income-tags">no tags</a>`
            if (!(typeof(income.tags) === "undefined")) {
                tags = ''
                for (var tag = 0; tag < income.tags.length; tag++){
                    tags += `<a class="badge badge-info income-tags">${income.tags[tag].name}</a>`
                }
            }
            dateOptions = {year: 'numeric', month: 'numeric', day: 'numeric' };
            htmlTable += `
            <tr class="amount ${rowStyle}">
                <th scope="row">${id}</th>
                <td><a>${amount}$</a></td>
                <td>${date.toLocaleString('ru-RU', dateOptions)}</td>
                <td>${hint}</td>
                <td>${tags}</td>
                <td>↓${fallSum}$</td>
                <td><button>Delete</button></td>
            </tr>`;
        }
        $('.incomeTable').append(htmlTable);
        
    });
})

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++) {
        var c = ca[i].trim();
           if (c.indexOf(name)==0) return c.substring(name.length,c.length);
    }
    return "";
}   

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
    location.reload();
}

function isLoginOK(data) {
    console.log('data'+data)
}