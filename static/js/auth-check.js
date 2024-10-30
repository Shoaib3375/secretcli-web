// Function to check if the user is logged in based on token and expiry date
function isLoggedIn() {
  const token = localStorage.getItem('token');
  const expiry = localStorage.getItem('expiry');

  if (!token || !expiry) {
    return false;
  }

  // Check if the expiry date is still valid
  const expiryDate = new Date(expiry);
  const now = new Date();
  return now < expiryDate;
}

// Function to verify login status on page load
function checkLogin() {
  if (!isLoggedIn()) {
    alert("You are not logged in or your session has expired. Please log in.");
    window.location.href = '/auth/web/login'; // Redirect to login page if not logged in
  } else {
    // This will work now since the element exists
    document.getElementById('welcome-message').textContent = "You are logged in!";
  }
}

// Run the checkLogin function on page load
window.onload = checkLogin;
