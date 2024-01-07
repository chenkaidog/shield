document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Prevent default form submission
  
    // Gather form data
    var formData = new FormData(this);
    var loginData = {
      username: formData.get("username"),
      password: formData.get("password")
    };
  
    // Send login data as JSON to the server
    fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(loginData)
    })
    .then(function(response) {
     
    })
    .catch(function(error) {
      alert(error)
    });
  });