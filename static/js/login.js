// Function to toggle password visibility
function togglePasswordVisibility() {
  const passwordField = document.getElementById('password');
  const eyeIcon = document.getElementById('eye-icon');
  if (passwordField.type === "password") {
    passwordField.type = "text";
    eyeIcon.textContent = '🙈'; // Hide icon
  } else {
    passwordField.type = "password";
    eyeIcon.textContent = '👁️'; // Show icon
  }
}

// Function to handle login form submission
async function submitLogin(event) {
  event.preventDefault(); // Prevent the default form submission

  // Clear previous error message
  const errorMessageElement = document.getElementById('error-message');
  errorMessageElement.style.display = 'none';  // Hide error message initially
  errorMessageElement.textContent = '';

  // Get the form data
  const formData = new FormData(document.getElementById('login-form'));
  const data = Object.fromEntries(formData);

  try {
    // Send a POST request to the login API
    const response = await fetch('http://localhost:8080/auth/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    // Check if response is ok
    if (response.ok) {
      // Parse JSON and save the token
      const jsonResponse = await response.json();

      // Store the token and expiry in localStorage
      localStorage.setItem('token', jsonResponse.data.token);
      localStorage.setItem('expiry', jsonResponse.data.expiry);

      // Redirect to the home page or desired page
      window.location.href = '/'; // Adjust as needed
    } else {
      // If response is not ok, parse error message and display it
      const errorData = await response.json();
      errorMessageElement.textContent = errorData.message || 'Login failed';
      errorMessageElement.style.display = 'block'; // Show error message
    }
  } catch (error) {
    console.error('Error:', error);
    errorMessageElement.textContent = `Error: ${error.message}`;
    errorMessageElement.style.display = 'block'; // Show error message
  }
}
