// 绑定“检查交易”按钮事件
document.getElementById('checkTransaction').addEventListener('click', async function () {
    const txID = document.getElementById('txID').value.trim();

    if (!txID) {
        alert('Please enter a valid transaction ID.');
        return;
    }

    try {
        const token = localStorage.getItem('token');
        if (!token) {
            alert('Authorization token not found. Please log in.');
            window.location.href = 'index.html';
            return;
        }

        // 请求后端API查询交易详情
        const response = await fetch(`http://localhost:8080/api/transaction/${txID}`, {
            method: 'GET',
            headers: {
                "Authorization": `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });

        if (response.ok) {
            const transaction = await response.json();
            console.log('Transaction Details:', transaction);

            // 展示交易详情
            const resultContainer = document.getElementById('transactionResult');
            const detailsElement = document.getElementById('transactionDetails');
            detailsElement.textContent = JSON.stringify(transaction, null, 2);
            resultContainer.style.display = 'block';
        } else {
            const errorMessage = await response.text();
            alert(`Failed to fetch transaction: ${errorMessage}`);
        }
    } catch (error) {
        console.error('Error fetching transaction:', error);
        alert('An error occurred while fetching the transaction. Please try again.');
    }
});

// 绑定“返回仪表板”按钮事件
document.getElementById('returnDashboard').addEventListener('click', function () {
    window.location.href = 'dashboard.html'; // 跳转到仪表板页面
});
