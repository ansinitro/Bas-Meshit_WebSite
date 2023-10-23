var prompted = false;

// Function to check if the user has reached the bottom of the page
function isAtBottom() {
  return window.innerHeight + window.scrollY + 100 >= document.body.scrollHeight;
}

// Function to show a prompt when the user reaches the bottom
function showBottomPrompt() {
  if (isAtBottom() && !prompted) {
    prompted = true;
    var userInput = prompt('Сізге сайт ұнады ма?\nӨзініздің ойыныз жазып кетініз:)');
    if (userInput !== null) {
      alert("Жауап бергенінізге рахмет!");
    }
  }
}

// Add a scroll event listener to trigger the prompt
window.addEventListener('scroll', showBottomPrompt);