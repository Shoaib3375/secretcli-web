// Function to handle the secret creation form submission
async function submitSecretForm(event) {
  event.preventDefault(); // Prevent the default form submission

  // Get the form data
  const formData = new FormData(event.target);
  const data = Object.fromEntries(formData);

  // Retrieve the token from localStorage
  const token = localStorage.getItem('token');

  try {
    // Send a POST request to the create secret API
    const response = await fetch('http://localhost:8080/secret/api/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`, // Include the bearer token in the headers
      },
      body: JSON.stringify(data),
    });

    // Store the response body as text
    const responseBodyText = await response.text();

    if (response.ok) {
      // Parse the JSON response
      const jsonResponse = JSON.parse(responseBodyText);

      // Show a success message
      alert('Secret created successfully!');

      // Redirect to the secrets listing page or any other desired page
      window.location.href = '/secret/web/list'; // Change to your desired success page
    } else {
      // Attempt to parse the error response
      let errorData;
      try {
        errorData = JSON.parse(responseBodyText); // Try to parse it as JSON
      } catch (jsonError) {
        errorData = { message: responseBodyText }; // Use the plain text response as the error message
      }

      // Display the error message in a popup alert
      alert(`Error: ${errorData.message || 'Failed to create secret'}`);
      console.error('Server response:', errorData.message);
    }
  } catch (error) {
    console.error('Error:', error);
    alert(`Error: ${error.message}`);
  }
}

// Attach the submitSecretForm function to the form submit event
document.getElementById('secret-form').addEventListener('submit', submitSecretForm);
