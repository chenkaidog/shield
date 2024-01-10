document.addEventListener("DOMContentLoaded", function() {
  document.getElementById("loginBut").addEventListener("click",  function() {
    var loginData = {
      username: document.getElementById("username").value,
      password: document.getElementById("password").value
    };
  
    fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(loginData)
    })
    .then(function(response) {
      if (response.ok) {
        
      }
    })
    .catch(function(error) {
      alert(error)
    });
  });
})