// Simple test to ping the backend
const testRequest = {
  user_id: 1,
  latitude: 37.352395,
  longitude: -121.953076,
};

async function pingBackend() {
  console.log('Pinging backend with:', testRequest);
  
  try {
    const response = await fetch('http://localhost:8080/api/sunset-prediction', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(testRequest),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.json();
    console.log('Backend response:', result);
    return result;
  } catch (error) {
    console.error('Failed to ping backend:', error);
    throw error;
  }
}

// Run the test
pingBackend(); 