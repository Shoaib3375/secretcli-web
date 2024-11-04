async function submitLogin(event) {
  event.preventDefault(); // Prevent the default form submission

  // Get the form data
  const formData = new FormData(event.target);
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

    // Store the response body as text
    const responseBodyText = await response.text();

    if (response.ok) {
      // Parse the JSON response
      const jsonResponse = JSON.parse(responseBodyText);

      // Save the token and expiry in localStorage
      localStorage.setItem('token', jsonResponse.data.token);
      localStorage.setItem('expiry', jsonResponse.data.expiry);

      // Redirect to success page or show success message
      window.location.href = '/'; // Change to your desired success page
    } else {
      // Attempt to parse the error response
      let errorData;
      try {
        errorData = JSON.parse(responseBodyText); // Try to parse it as JSON
      } catch (jsonError) {
        errorData = { message: responseBodyText }; // Use the plain text response as the error message
      }

      // Display the error message in a popup alert
      alert(`Error: ${errorData.message || 'Login failed'}`);
      console.error('Server response:', errorData.message);
    }
  } catch (error) {
    console.error('Error:', error);
    alert(`Error: ${error.message}`);
  }
}

// Attach the submitLogin function to the form submit event
document.getElementById('login-form').addEventListener('submit', submitLogin);
