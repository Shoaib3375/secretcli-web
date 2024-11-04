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

    const result = await response.json(); // Parse the response JSON
    const secrets = result.data.secrets; // Access the secrets array
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
    row.classList.add('secret-row'); // Add class for styling

    row.innerHTML = `
      <td class="secret-name" onclick="toggleDetails(this)">${secret.title}</td>
      <td class="secret-details">
        <div>
          <strong>Username:</strong> ${secret.username} <br>
          <strong>Password:</strong>
          <span class="masked-password">********
            <span class="reveal-password" onclick="togglePassword(event, '${secret.password}')">Show</span>
          </span>
          <br>
          <strong>Note:</strong> ${secret.note} <br>
          <strong>Email:</strong> ${secret.email} <br>
          <strong>Website:</strong> ${secret.website} <br>
          <strong>Created At:</strong> ${new Date(secret.created_at).toLocaleString()} <br>
          <strong>Updated At:</strong> ${secret.updated_at ? new Date(secret.updated_at).toLocaleString() : 'N/A'}
        </div>
      </td>
    `;
    
    tableBody.appendChild(row); // Append the new row to the table body
  });
}

// Function to toggle secret details visibility
function toggleDetails(element) {
  const row = element.closest('tr');
  row.classList.toggle('active'); // Toggle the active class to show/hide details
}

// Function to toggle password visibility
function togglePassword(event, password) {
  event.stopPropagation(); // Prevent the event from bubbling up to the row click
  const passwordSpan = event.target.parentElement; // Get the password span
  
  // Check if the password is currently masked
  if (passwordSpan.innerHTML.includes("********")) {
    // Show the password
    passwordSpan.innerHTML = `${password} <span class="reveal-password" onclick="togglePassword(event, '${password}')">Hide</span>`;
  } else {
    // Hide the password
    passwordSpan.innerHTML = `******** <span class="reveal-password" onclick="togglePassword(event, '${password}')">Show</span>`;
  }
}

// Check login status before fetching secrets
async function init() {
  if (checkAuth()) { // Call the checkAuth function from auth.js
    await fetchSecrets(); // Fetch secrets only if logged in
  }
}

// Run the init function on page load
window.onload = init;
