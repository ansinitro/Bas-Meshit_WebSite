let comeUsBtn = document.getElementById("comeUsBtn");
comeUsBtn.addEventListener("click", function () {
  alert("This feature will be added on the next release.")
});

buttons = document.getElementsByClassName("registration");
  for (i = 0; i < buttons.length; i++) {
    buttons[i].onclick = (function() {
      window.location.assign("registration.html");
    });
  }