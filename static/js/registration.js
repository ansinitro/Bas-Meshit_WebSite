const form = document.getElementById('form');
const namee = document.getElementById('name');
const surname = document.getElementById('surname');
const phoneNumber = document.getElementById('phone-number');
const email = document.getElementById('email');
const course = document.getElementById('course');
const age = document.getElementById('age');
const check = document.getElementById('check');

form.addEventListener('submit', function (event) {
  console.log("Iamhere");
  if (course.value === 'choose') {
    setInvalid(course);
    event.preventDefault();
  }
  if (form.getElementsByClassName("is-valid").length != 7) {
    event.preventDefault();
  }
});

surname.addEventListener("input", function (element) {
  isValidName(surname, element);
});
namee.addEventListener("input", function (element) {
  isValidName(namee, element);
});
phoneNumber.addEventListener("input", function (element) {
  isValidNumber(phoneNumber, element);
});
email.addEventListener("input", function (element) {
  isValidEmail(email, element);
});
course.addEventListener("input", function (element) {
  isCourseChose(course, element);
});
age.addEventListener("input", function (element) {
  isValidAge(age, element);
});
check.addEventListener("change", function (element) {
  if (!check.checked) {
    setInvalid(check);
  } else {
    setValid(check);
  }
});

function isCourseChose(el, element) {
  if (course.value == "choose") {
    setInvalid(el);
  } else {
    setValid(el);
  }
}

function isValidAge(el, element) {
  let val = element.target.value.trim();

  el.setAttribute("value", val);
  if (val == '') {
    setInvalid(el);
  } else if (val < 100) {
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

function isValidName(el, element) {
  let val = element.target.value.trim();
  el.setAttribute("value", val)
  const regex = new RegExp('^[A-Z][a-z]+');
  if (regex.test(val) && (val.length >= 2 && val.length < 50)) {
    setValid(el);
  } else {
    setInvalid(el);
  }
}

function isValidNumber(el, element) {
  let val = element.target.value.trim();

  el.setAttribute("value", val);
  if (val === '') {
    setInvalid(el);
  } else if (val.length != 10) {
    setInvalid(el);
  } else if (!(/^\d+$/.test(val)) || val[0] != '7') {
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