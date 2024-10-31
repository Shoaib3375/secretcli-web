console.log("secrets.js is loaded");

// Function to fetch secrets from the API
async function fetchSecrets() {
  const token = localStorage.getItem('token');

  try {
    const response = await fetch('http://localhost:8080/secret/api/list', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`, // Include the token in the Authorization header
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch secrets');
    }

    const secrets = await response.json(); // Parse the response JSON
    populateSecretsTable(secrets); // Populate the table with fetched secrets
  } catch (error) {
    console.error('Error fetching secrets:', error);
    alert('Error fetching secrets: ' + error.message);
  }
}

// Function to populate the secrets table with data
function populateSecretsTable(secrets) {
  const tableBody = document.querySelector('tbody'); // Get the table body element
  tableBody.innerHTML = ''; // Clear existing rows

  // Loop through the secrets and create rows
  secrets.forEach(secret => {
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>${secret.id}</td>
      <td>${secret.title}</td>
      <td>${secret.username}</td>
      <td>${secret.password}</td>
      <td>${secret.note}</td>
      <td>${secret.email}</td>
      <td>${secret.website}</td>
      <td>${secret.user_id}</td>
      <td>${new Date(secret.created_at).toLocaleString()}</td>
      <td>${secret.updated_at ? new Date(secret.updated_at).toLocaleString() : 'N/A'}</td>
    `;
    tableBody.appendChild(row); // Append the new row to the table body
  });
}

// Check login status before fetching secrets
async function init() {
  if (checkAuth()) { // Call the checkAuth function from auth.js
    await fetchSecrets(); // Fetch secrets only if logged in
  }
}

// Run the init function on page load
window.onload = init;
