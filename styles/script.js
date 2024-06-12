// Toggle class active
const navbarNav = document.querySelector(".navbar-nav");

//Ketika hamburger menu di klik
document.querySelector("#hamburger-menu").onclick = () => {
  navbarNav.classList.toggle("active");
};

// Tombol contact us
const rsvNow = document.querySelector(".rsv-now");

document.querySelector("#cta").onclick = () => {
  rsvNow.classList.toggle("active");
};
