function toggle() {
  var element = document.getElementById("mobile-menu"); 

  if (element.style.display != "none") {
    element.style.display = 'none';
  } else if (element.style.display == 'none') {
    element.style.display = 'block';
  }

};
