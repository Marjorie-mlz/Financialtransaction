// 确认 app.js 文件已成功加载
console.log('app.js loaded successfully');

// 等待页面加载完成后再绑定事件
document.addEventListener('DOMContentLoaded', () => {
    const transactionForm = document.getElementById('transactionForm');
    const loginForm = document.getElementById('login-form');
    const registerForm = document.getElementById('register-form');
    const logoutBtn = document.getElementById('logout-btn');
    const transactionStatus = document.getElementById('transactionStatus');
    const transactionHistoryTable = document.querySelector('#transactionTable tbody');  // 用于交易历史展示
    const uploadForm = document.getElementById('uploadForm');

    // 登录逻辑
    if (loginForm) {
        loginForm.addEventListener('submit', async function (event) {
            event.preventDefault();
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;

            try {
                const response = await fetch('http://localhost:8080/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ email, password }),
                });

                const data = await response.json();
                if (response.ok) {
                    console.log("Token from server:", data.token);
                    localStorage.setItem("token", data.token);
                    window.location.href = 'dashboard.html';
                } else {
                    alert(data.error);
                }
            } catch (error) {
                console.error('Error:', error);
                alert('Something went wrong. Please try again later.');
            }
        });
    }

    // 注册逻辑
    if (registerForm) {
        registerForm.addEventListener('submit', async function (event) {
            event.preventDefault();
            const username = document.getElementById('username').value;
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;

            try {
                const response = await fetch('http://localhost:8080/register', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ username, email, password }),
                });

                const data = await response.json();
                if (response.ok) {
                    alert('Registration successful! Please log in.');
                    window.location.href = 'index.html';
                } else {
                    alert(data.error);
                }
            } catch (error) {
                console.error('Error:', error);
                alert('Something went wrong. Please try again later.');
            }
        });
    }

    // 退出功能
    if (logoutBtn) {
        logoutBtn.addEventListener('click', () => {
            console.log('Logout button clicked');
            localStorage.removeItem('token');
            window.location.href = 'index.html';
        });
    }
    // 交易记录展示逻辑
    async function loadTransactionHistory() {
        try {
            const response = await fetch('http://localhost:8080/api/transactions', {
                headers: {
                    'Authorization': 'Bearer ' + localStorage.getItem('token'),
                },
            });
            const transactions = await response.json();

            transactionHistoryTable.innerHTML = '';  // 清空表格内容
            transactions.forEach(transaction => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${transaction.id}</td>
                    <td>${transaction.sender}</td>
                    <td>${transaction.receiver}</td>
                    <td>${transaction.amount}</td>
                    <td class="${transaction.status.toLowerCase()}">${transaction.status}</td>
                    <td>${new Date(transaction.timestamp).toLocaleString()}</td>
                `;
                transactionHistoryTable.appendChild(row);
            });
        } catch (error) {
            console.error('Failed to load transactions:', error);
        }
    }

    // 如果在交易历史页面，则加载交易记录
    if (transactionHistoryTable) {
        loadTransactionHistory();
    }

    
    // 上传数据表单的逻辑 - 新增功能
    if (uploadForm) {
        uploadForm.addEventListener('submit', async function (event) {
            event.preventDefault();

            const sender = document.getElementById('sender').value.trim();
            const receiver = document.getElementById('receiver').value.trim();
            const amount = parseFloat(document.getElementById('amount').value.trim());
            const note = document.getElementById('note').value.trim();

            if (sender.length !== 42 || !sender.startsWith('0x') || receiver.length !== 42 || !receiver.startsWith('0x')) {
                alert('Invalid Ethereum address');
                return;
            }
            if (amount <= 0 || isNaN(amount)) {
                alert('Amount must be greater than 0');
                return;
            }

            try {
                const response = await fetch('http://localhost:8080/api/upload-data', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': 'Bearer ' + localStorage.getItem('token'),
                    },
                    body: JSON.stringify({
                        sender: sender,
                        receiver: receiver,
                        amount: amount,
                        note: note,
                    }),
                });

                const result = await response.json();
                if (response.ok) {
                    document.getElementById('transactionStatus').innerHTML = `
                        <div class="alert alert-success" role="alert">
                            Data uploaded successfully! <br>
                            Transaction Hash: <strong>${result.transactionHash}</strong>
                        </div>`;
                } else {
                    document.getElementById('transactionStatus').innerHTML = `
                        <div class="alert alert-danger" role="alert">
                            Upload failed: ${result.error}
                        </div>`;
                }
            } catch (error) {
                console.error('Upload failed', error);
                document.getElementById('transactionStatus').innerHTML = `
                    <div class="alert alert-danger" role="alert">
                        An unexpected error occurred. Please check the console logs.
                    </div>`;
            }
        });
    }
    
    // 修改密码功能
    async function changePassword() {
        const oldPassword = document.getElementById('oldPassword').value;
        const newPassword = document.getElementById('newPassword').value;

        try {
            const response = await fetch('/change-password', {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem('token'),
                },
                body: JSON.stringify({ old_password: oldPassword, new_password: newPassword }),
            });

            if (response.ok) {
                localStorage.removeItem('token');
                document.getElementById('message').innerHTML = `
                    <p>Password changed successfully. Please log in again.</p>
                    <button id="loginBtn" class="btn btn-primary mt-2">Go to Login</button>`;
                document.getElementById('loginBtn').addEventListener('click', () => {
                    window.location.href = 'index.html';
                });
            } else {
                const error = await response.json();
                document.getElementById('message').innerHTML = `<p style="color:red">Failed to change password: ${error.error}</p>`;
            }
        } catch (error) {
            console.error('Error:', error);
            document.getElementById('message').innerHTML = `<p style="color:red">An error occurred: ${error.message}</p>`;
        }
    }
    // 绑定返回仪表板按钮事件
    if (returnDashboardBtn) {
        returnDashboardBtn.addEventListener('click', () => {
            window.location.href = 'dashboard.html'; // 跳转到仪表板页面
        });
    }
});
