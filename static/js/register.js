// Function to toggle password visibility
function togglePasswordVisibility(fieldId) {
  const passwordField = document.getElementById(fieldId);
  const eyeIcon = document.getElementById('eye-icon-' + fieldId);
  if (passwordField.type === "password") {
      passwordField.type = "text";
      eyeIcon.textContent = '🙈';
  } else {
      passwordField.type = "password";
      eyeIcon.textContent = '👁️';
  }
}

// Function to handle the registration form submission
async function submitRegister(event) {
  event.preventDefault(); // Prevent the default form submission

  // Get the form data
  const formData = new FormData(document.getElementById('register-form'));
  const data = Object.fromEntries(formData);

  // Clear previous error messages
  const errorMessageDiv = document.getElementById('error-message');
  errorMessageDiv.textContent = '';

  try {
      // Send a POST request to the registration API
      const response = await fetch('http://localhost:8080/auth/api/register', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
      });

      // Check if response is ok
      if (response.ok) {
          // Parse JSON response
          const jsonResponse = await response.json();

          // Redirect to login page after successful registration
          window.location.href = '/auth/web/login'; // Redirect to login page
      } else {
          // If response is not ok, parse error message and display it
          const errorData = await response.json();
          errorMessageDiv.textContent = errorData.message || 'Registration failed. Please try again.';
      }
  } catch (error) {
      console.error('Error:', error);
      const errorMessageDiv = document.getElementById('error-message');
      errorMessageDiv.textContent = `Error: ${error.message}`;
  }
}
