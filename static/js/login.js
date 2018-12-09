String.prototype.hashCode = function() {
    var hash = 0, i, chr;
    if (this.length === 0) return hash;
    for (i = 0; i < this.length; i++) {
        chr   = this.charCodeAt(i);
        hash  = ((hash << 5) - hash) + chr;
        hash |= 0; // Convert to 32bit integer
    }
    return hash;
  };

function check(form) {
    userName = form.userName.value;
    userHash = form.password.value.hashCode();
    method = "post";

    var form = document.createElement("form");
    form.setAttribute("method", method);
    form.setAttribute("action", "api/v1/login/");

    var fieldUserName = document.createElement("input");
    fieldUserName.setAttribute("type", "hidden");
    fieldUserName.setAttribute("name", "userName");
    fieldUserName.setAttribute("value", userName);
    form.appendChild(fieldUserName);

    var fieldUserHash = document.createElement("input");
    fieldUserHash.setAttribute("type", "hidden");
    fieldUserHash.setAttribute("name", "userHash");
    fieldUserHash.setAttribute("value", userHash);
    form.appendChild(fieldUserHash);

    console.log(userName+' '+userHash);

    document.body.appendChild(form);
    form.submit();
}

function isLoginOK(data) {
    var jsonData = JSON.parse(data);
    console.log(jsonData)
}