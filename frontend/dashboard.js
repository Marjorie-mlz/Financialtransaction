// 从 localStorage 获取 JWT 令牌
const token = localStorage.getItem("token");

if (!token) {
    window.location.href = "index.html"; // 如果没有 token，跳转到登录页面
}

// 辅助函数：处理 fetch 请求的通用错误
function handleError(response) {
    if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return response.json();
}

// 获取用户信息并显示
async function loadUserProfile() {
    try {
        const response = await fetch("http://localhost:8080/api/profile", {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }
        });
        const data = await handleError(response);
        document.getElementById("username").textContent = data.username || "未知用户";
        document.getElementById("email").textContent = data.email || "无邮箱信息";
    } catch (error) {
        console.error("Error fetching profile:", error);
        alert("无法获取用户信息，请重试。");
    }
}

// 获取账户余额并显示
async function loadAccountBalance() {
    try {
        const response = await fetch("http://localhost:8080/api/account-balance", {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }
        });
        const data = await handleError(response);
        document.getElementById('balance').innerText =
            data.balance ? `${data.balance} ETH` : '0 ETH';
    } catch (error) {
        console.error("Error fetching account balance:", error);
        alert("无法获取账户余额，请稍后再试。");
    }
}


// 获取最近交易记录并显示
async function loadRecentTransactions() {
    try {
        const response = await fetch("http://localhost:8080/api/transactions", {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json"
            }
        });
        const data = await handleError(response);

        const transactionList = document.getElementById("recent-transactions");
        transactionList.innerHTML = ''; // 确保之前内容清空

        if (data.transactions && data.transactions.length > 0) {
            data.transactions.forEach(tx => {
                const li = document.createElement("li");
                li.textContent =
                    `ID: ${tx.tx_id}, Amount: ${tx.amount} ETH, Status: ${tx.status}`;
                transactionList.appendChild(li);
            });
        } else {
            transactionList.innerHTML = '<li>暂无交易记录</li>';
        }
    } catch (error) {
        console.error("Error fetching transactions:", error);
        document.getElementById("recent-transactions").innerHTML =
            '<li>无法获取交易记录</li>';
    }
}

// 登出按钮点击事件
document.getElementById("logout-btn").addEventListener("click", () => {
    localStorage.removeItem("token");
    window.location.href = "index.html";
});
// 页面加载时调用各个数据加载函数
window.addEventListener('load', () => {
    loadUserProfile();
    loadAccountBalance();
    loadRecentTransactions();
});