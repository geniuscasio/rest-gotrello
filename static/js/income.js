var sortType = 0
var sortMode = 1
var incomes = []

const incomeStyle = "income";
const outcomeStyle = "outcome";

$(document).ready(function() {
    $.ajax({        
        url:"api/v1/income/"
    }).then(function(data) {
        var userNamePlaceholder = document.getElementById("username");
        var userName = getCookie("userName")
        $('.row').append();
        userNamePlaceholder.innerText = userName
        // data = '{}'
        incomes = JSON.parse(data);
        // incomes = [{"id": 1, "date": "2021-05-20T15:04:05.999999-07:00", "amount": 10}, 
        // {"id": 1, "date": "2021-05-20T15:04:05.999999-07:00", "amount": -7},
        // {"id": 1, "date": "2019-05-20T15:04:05.999999-07:00", "amount": 2}, 
        // {"id": 1, "date": "2020-05-20T15:04:05.999999-07:00", "amount": 3}]
        console.log(incomes)
        incomes = sortIncomes(incomes);
        console.log(incomes)
        updateIncomeView()
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

function sortIncomes(what) {
    console.log(sortType, sortMode)
    var result = what.sort(function (a, b) {
        var valueA = 0;
        var valueB = 0;
        if(sortType == 1) {
            valueA = new Date(a.date).getTime();
            valueB = new Date(b.date).getTime();
        } else if (sortType == 0) {
            valueA = a.amount;
            valueB = b.amount;
        }
        if (valueA < valueB) {
            return -1 * sortMode;
        }
        if (valueA > valueB) {
            return 1 * sortMode;
        }
        return 0;
    });
    return result;
}

function updateIncomeView() {
    var balance = 0;
    var totalIncome = 0;
    var totalOutcome = 0;
    var htmlTable = ''
    $('.incomeTable tr').remove();
    if (incomes == null) {
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
    for (var i = 0; i < incomes.length; i++) {
        var income = incomes[i];
        var id = i + 1;
        var date = new Date(income.date);
        var hint = income.hint;
        var amount = income.amount;
        var rowStyle;

        balance += amount;
        if (amount < 0) { 
            rowStyle = outcomeStyle;
            totalOutcome += amount;
        } else if (amount > 0) {
            rowStyle = incomeStyle;
            totalIncome += amount;
        }
        if (hint == undefined) {
            hint = "Немає"; 
        }
        
        tags = ""
        if(typeof(income.tags) === "undefined") {
            tags = `<a class="badge badge-warning income-tags">no tags</a>`
        } else {
            tag_list = income.tags.split(",")
            for (var k = 0; k < tag_list.length; k++){
                tags += `<a class="badge badge-info income-tags">${tag_list[k]}</a>`
            }
        }
        dateOptions = {year: 'numeric', month: 'numeric', day: 'numeric' };
        htmlTable += `
        <tr class="amount ${rowStyle}">
            <th scope="row">${id}</th>
            <td><a>${amount}</a></td>
            <td>${date.toLocaleString('ru-RU', dateOptions)}</td>
            <td>${hint}</td>
            <td>${tags}</td>
            <td>↓${balance}</td>
            <td><button onClick='deleteIncome(${id})'>❌</button></td>
        </tr>`;
    }
    $('.incomeTable').append(htmlTable);
    document.getElementById("totalIncome").innerText = "Доходи: " + totalIncome;
    document.getElementById("totalOutcome").innerText = "Витрати: " + totalOutcome;
    document.getElementById("balance").innerText = "Баланс: " + balance;
}

function deleteIncome(id){
    var isApproved = confirm("Вы - администратор?");
    if(isApproved) {
        var xhr = new XMLHttpRequest();
        xhr.open("DELETE", "api/v1/incomeDelete/", true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.send(JSON.stringify({
            id: parseInt(id, 10)
        }));
    }
}

function refreshSort() {
    document.getElementById('sortAmount').innerText = "Сума 💲";
    document.getElementById('sortDate').innerText = "Дата 📅";
}

function sortByAmount() {
    sortType = 0;
    var sortHint = document.getElementById('sortAmount');
    var icon = sortReverse();
    refreshSort()
    sortHint.innerText = sortHint.innerText + icon;
    incomes = sortIncomes(incomes);
    updateIncomeView();
    console.log("sort by amount");
}

function sortReverse() {
    var icon = ""
    if(sortMode == 1) {
        icon = "⬇️";
        sortMode = -1;
    } else {
        icon = "⬆️";
        sortMode = 1;
    }
    return icon;
}

function sortByDate() {
    sortType = 1
    var sortHint = document.getElementById('sortDate');
    var icon = sortReverse();
    refreshSort()
    sortHint.innerText = sortHint.innerText + icon;
    incomes = sortIncomes(incomes);
    updateIncomeView();
    console.log("sort by date");
}

function newIncome(form) {
    var id = form.id.value;
    var amount = form.amount.value;
    var date = form.date.value;
    var hint = form.hint.value;
    var tags = form.tags.value;
    if(amount == "") {
        alert("Поле сумма обов'язкове до заповнення!");
    }
    if(date == "") {
        var today = new Date();
        var dd = String(today.getDate()).padStart(2, '0');
        var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
        var yyyy = today.getFullYear();

        date = yyyy + '-' + mm + '-' + dd;
    }
    console.log(date)
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
