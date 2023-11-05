let i = 0;

document.getElementsByClassName('carousel-control-next')[0].addEventListener('click', function() {
  if (i >= 4) {
    i = 0;
  }
  url = '../images/mosque-background-' + i + '.jpg';
  document.getElementById("header").style.setProperty("background-image", "linear-gradient(rgba(4, 9, 3, 0.7), rgba(4, 9, 3, 0.7)), url('" + url + "')");
  // document.getElementById("sub-header").style.setProperty("background-image", "linear-gradient(rgba(4, 9, 3, 0.7), rgba(4, 9, 3, 0.7)), url('" + url + "')");
  i++;
});

document.getElementsByClassName('carousel-control-prev')[0].addEventListener('click', function() {
  if (i >= 4) {
    i = 0;
  }
  url = '../images/mosque-background-' + i + '.jpg';
  document.getElementById("header").style.setProperty("background-image", "linear-gradient(rgba(4, 9, 3, 0.7), rgba(4, 9, 3, 0.7)), url('" + url + "')");
  // document.getElementById("sub-header").style.setProperty("background-image", "linear-gradient(rgba(4, 9, 3, 0.7), rgba(4, 9, 3, 0.7)), url('" + url + "')");
  i++;
});