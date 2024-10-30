async function submitRegistration(event) {
  event.preventDefault(); // Prevent the default form submission

  // Get the form data
  const formData = new FormData(event.target);
  const data = Object.fromEntries(formData);

  console.log(JSON.stringify(data));

  try {
    // Send a POST request to the registration API
    const response = await fetch('http://localhost:8080/auth/api/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    // Store the response body as text
    const responseBodyText = await response.text();

    if (response.ok) {
      // Handle successful registration
      alert('Registration successful! Please log in.');
      window.location.href = '/auth/web/login';
    } else {
      // Attempt to parse the error response
      let errorData;
      try {
        errorData = JSON.parse(responseBodyText);
      } catch (jsonError) {
        errorData = { message: responseBodyText };
      }

      // Display the server error message in an alert popup
      alert(errorData.message || 'Registration failed');
      console.error('Server response:', errorData.message);
    }
  } catch (error) {
    console.error('Error:', error);
    alert(`Error: ${error.message}`);
  }
}

// Attach the submitRegistration function to the form submit event
document.getElementById('registration-form').addEventListener('submit', submitRegistration);
