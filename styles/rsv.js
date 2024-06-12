// Toggle class active for the navbar
const navbarNav = document.querySelector(".navbar-nav");

document.addEventListener("DOMContentLoaded", () => {
  const tabButtons = document.querySelectorAll(".tab-button");
  const indoorContents = document.querySelectorAll(".indoor");
  const outdoorContents = document.querySelectorAll(".outdoor");

  // TAB
  tabButtons.forEach((button) => {
    button.addEventListener("click", () => {
      const tab = button.getAttribute("data-tab");

      if (tab === "indoor") {
        indoorContents.forEach((content) => content.classList.add("active"));
        outdoorContents.forEach((content) =>
          content.classList.remove("active")
        );
      } else if (tab === "outdoor") {
        indoorContents.forEach((content) => content.classList.remove("active"));
        outdoorContents.forEach((content) => content.classList.add("active"));
      }

      tabButtons.forEach((btn) => btn.classList.remove("active"));
      button.classList.add("active");
    });
  });
});

// Toggle class active for the hamburger menu
document.querySelector("#hamburger-menu").onclick = () => {
  navbarNav.classList.toggle("active");
};