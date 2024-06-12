// Toggle class active
const navbarNav = document.querySelector(".navbar-nav");

document.addEventListener("DOMContentLoaded", function () {
  // TAB
  const tabButtons = document.querySelectorAll(".tab-button");
  const tabContents = document.querySelectorAll(".history-content, .card-list");

  // CART
  const addToCartButtons = document.querySelectorAll(".add-to-cart");
  const cartItemCount = document.querySelector(".cart-button span");
  const cartItemsList = document.querySelector(".cart-items");
  const cartTotal = document.querySelector(".cart-total");
  const cartIcon = document.querySelector(".cart-button");
  const sidebar = document.querySelector(".sidebar");

  // TAB BUTTON START
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

  // CART MENU

  cartIcon.addEventListener("click", () => {
    sidebar.classList.toggle("open");
  });

  const closeButton = document.querySelector(".sidebar-close");
  closeButton.addEventListener("click", () => {
    sidebar.classList.remove("open");
  });
});

// Ketika hamburger menu di klik
document.querySelector("#hamburger-menu").onclick = () => {
  navbarNav.classList.toggle("active");
};