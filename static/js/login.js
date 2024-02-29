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
const signin = document.getElementById("signInBtn")
const signup = document.getElementById("signUpBtn")

form.addEventListener('submit', function (event) {
  console.log(form.getAttribute("action"));
  if (form.getAttribute("action") === "/signin") {
    if (form.getElementsByClassName("is-valid").length != 2) {
        alert("Please enter both email and password for Sign In.");
        event.preventDefault(); // Prevent form submission
    } 
} else {
  if (form.getElementsByClassName("is-valid").length != 3) {
        alert("Name: must be capitalized, email, and password for Sign Up.");
        event.preventDefault(); // Prevent form submission
    }
}
});

signin.addEventListener("click", function() {
  form.action="/signin";
});

signup.addEventListener("click", function() {
  form.action="/signup";
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
  let val = element.target.value.trim();
  el.setAttribute("value", val);
  let namee = element.target.value.trim();
  const regex = new RegExp('^[A-Z][a-z]+');
  console.log(regex.test(namee), namee);
  if (regex.test(namee) && (namee.length > 2 && namee.length < 50)) {
    setValid(el);
    return true;
  } else {
    setInvalid(el);
    return false;
  }
}

function isValidPassword(el, element) {
  let val = element.target.value.trim();
  el.setAttribute("value", val);
  let password = element.target.value;
  console.log(password);
  if (password.length < 8) {
    setInvalid(el);
    return false;
  }

  var hasLetter = /[a-zA-Z]/.test(password);
  var hasNumber = /\d/.test(password);
  var hasSpecialChar = /[!@#$%^&*()_+{}\[\]:;<>,.?~\\/-]/.test(password);

  if (hasLetter && hasNumber && hasSpecialChar) {
    setValid(el);
    return true;
  } else {
    setInvalid(el);
    return false;
  }
}


function isValidEmail(el, element) {
  const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  let val = element.target.value.trim();
  el.setAttribute("value", val);

  if (val === '') {
    setInvalid(el);
    return false;
  } else if (!re.test(String(val).toLowerCase())) {
    setInvalid(el);
    return false;
  } else {
    setValid(el);
  }
  return true;
}

function setInvalid(element) {
  element.classList.add('is-invalid');
  element.classList.remove('is-valid')
}

function setValid(element) {
  element.classList.add('is-valid');
  element.classList.remove('is-invalid')
};
