document.getElementById("change-password-form").addEventListener("submit", function (event) {
    event.preventDefault();

    const oldPassword = document.getElementById("old-password").value;
    const newPassword = document.getElementById("new-password").value;

    const token = localStorage.getItem("token");

    fetch("http://localhost:8080/api/change-password", {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${token}`,
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            old_password: oldPassword,
            new_password: newPassword
        })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            document.getElementById("status").textContent = data.message || "Password changed successfully!";
            localStorage.removeItem('token'); // 清除 token

            // 添加返回登录按钮
            document.getElementById('status').innerHTML += `
            <button id="loginBtn" class="btn btn-primary mt-2">Go to Login</button>
        `;
            document.getElementById('loginBtn').addEventListener('click', function () {
                window.location.href = 'index.html';
            });
        })
        .catch(error => {
            document.getElementById("status").textContent = "Error changing password! Please try again.";
            console.error("Error:", error);
        });
});