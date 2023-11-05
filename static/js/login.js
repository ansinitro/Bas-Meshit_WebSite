let signUpBtn = document.getElementById("signUpBtn");
let signInBtn = document.getElementById("signInBtn");
let nameField = document.getElementById("nameField");
let title = document.getElementById("title");
let forgotPasswd = document.getElementById("forgotPasswd");

signInBtn.onclick = function() {
  nameField.style.maxHeight = "0";
  title.innerHTML = "Кіру"
  signUpBtn.classList.add("bg-secondary");
  signUpBtn.classList.remove("bg-info");
  signInBtn.classList.remove("bg-secondary");
  signInBtn.classList.add("bg-info");
  forgotPasswd.classList.remove("visually-hidden")
}

signUpBtn.onclick = function() {
  nameField.style.maxHeight = "60px";
  title.innerHTML = "Тіркелу"
  signUpBtn.classList.add("bg-info");
  signUpBtn.classList.remove("bg-secondary");
  signInBtn.classList.add("bg-secondary");
  signInBtn.classList.remove("bg-info");
  forgotPasswd.classList.add("visually-hidden") 
}

// Validation Check

const form = document.getElementById('form');
const namee = document.getElementById('name');
const email = document.getElementById('email');
const passwd = document.getElementById('passwd');

form.addEventListener('submit', function (event) {
  if (form.getElementsByClassName("is-valid").length == 7) {
    alert("Сіздің форманыз қабылданды.");
  } else {
    alert("Қатені дұрыстаңыз.");
  }
});

namee.addEventListener("input", function (element) {
  isValidName(namee, element);
});

email.addEventListener("input", function (element) {
  isValidEmail(email, element);
});

passwd.addEventListener("input", function (element) {
  isValidPassword(passwd, element);
});

function isValidName(el, element) {
  let namee = element.target.value.trim();
  const regex = new RegExp('^[A-Z][a-z]+');
  console.log(regex.test(namee), namee);
  if (regex.test(namee) && (namee.length > 2 && namee.length < 50)) {
    setValid(el);
  } else {
    setInvalid(el);
  }
}

function isValidPassword(el, element) {
  let password = element.target.value;
  console.log(password);
  if (password.length < 8) {
    return false;
  }

  var hasLetter = /[a-zA-Z]/.test(password);
  var hasNumber = /\d/.test(password);
  var hasSpecialChar = /[!@#$%^&*()_+{}\[\]:;<>,.?~\\/-]/.test(password);

  if (hasLetter && hasNumber && hasSpecialChar) {
    setValid(el);
  } else {
    setInvalid(el);
  }
}


function isValidEmail(el, element) {
  const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  let val = element.target.value.trim();

  el.setAttribute("value", val);

  if (val === '') {
    setInvalid(el);
  } else if (!re.test(String(val).toLowerCase())) {
    setInvalid(el);
  } else {
    setValid(el);
  }
}

function setInvalid(element) {
  element.classList.add('is-invalid');
  element.classList.remove('is-valid')
}

function setValid(element) {
  element.classList.add('is-valid');
  element.classList.remove('is-invalid')
};
