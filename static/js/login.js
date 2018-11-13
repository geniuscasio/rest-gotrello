function check(form)/*function to check userid & password*/ {
    /*the following code checkes whether the entered userid and password are matching*/
    if (form.userid.value == "myuserid" && form.pswrd.value == "mypswrd") {
        alert("ok")/*opens the target page while Id & password matches*/
        Console.log("ok")
    }
    else {
        alert("Error Password or Username")/*displays error message*/
        Console.log("err")
    }
}