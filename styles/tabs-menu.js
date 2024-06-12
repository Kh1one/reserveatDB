// Toggle class active
const navbarNav = document.querySelector(".navbar-nav");

document.addEventListener("DOMContentLoaded", function () {
  const tabButtons = document.querySelectorAll(".tab-button");
  const tabContents = document.querySelectorAll(".menu-content");

  tabButtons.forEach((tab) => {
    tab.addEventListener("click", () => {
      // Remove active class from all buttons
      tabButtons.forEach((btn) => btn.classList.remove("active"));

      // Add active class to the clicked button
      tab.classList.add("active");

      // Hide all tab contents
      tabContents.forEach((content) => content.classList.remove("active"));

      // Show the corresponding tab content
      const target = document.getElementById(tab.getAttribute("data-tab"));
      target.classList.add("active");
    });
  });
});

//Ketika hamburger menu di klik
document.querySelector("#hamburger-menu").onclick = () => {
  navbarNav.classList.toggle("active");
};
