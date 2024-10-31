// auth.js
// Common function to check login status and redirect if not logged in

function checkAuth() {
  const token = localStorage.getItem('token');
  const expiry = localStorage.getItem('expiry');

  if (!token || !expiry) {
    redirectToLogin();
    return false;
  }

  const expiryDate = new Date(expiry);
  const now = new Date();
  if (now >= expiryDate) {
    redirectToLogin();
    return false;
  }

  showLoggedInMessage(); // Show message if logged in
  return true;
}

function redirectToLogin() {
  alert("Your session has expired or you are not logged in. Redirecting to login page.");
  window.location.href = '/auth/web/login';
}

function showLoggedInMessage() {
  const welcomeMessage = document.getElementById('welcome-message');
  if (welcomeMessage) {
    welcomeMessage.textContent = "You are logged in!";
  }
}

// Run checkAuth on every page load where this script is included
window.onload = checkAuth;
