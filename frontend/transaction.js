const transactionForm = document.getElementById('transactionForm');

// 监听表单提交事件
transactionForm.addEventListener('submit', async function (event) {
    event.preventDefault(); // 阻止表单默认提交

    // 获取输入的值
    const transactionId = document.getElementById('transactionId').value.trim();
    const sender = document.getElementById('sender').value.trim();
    const receiver = document.getElementById('receiver').value.trim();
    const amount = parseFloat(document.getElementById('amount').value);
    const currency = document.getElementById('currency').value.trim();
    const status = document.getElementById('status').value.trim();
    const timestamp = document.getElementById('timestamp').value.trim();
    const note = document.getElementById('note').value.trim();

    // 从 localStorage 获取 JWT 令牌
    const token = localStorage.getItem("token");

    if (!token) {
        alert("Authorization token not found. Please log in.");
        window.location.href = "index.html"; // 如果没有token，跳转到登录页面
        return;
    }

    // 验证输入
    if (!transactionId || !sender || !receiver) {
        alert('Transaction ID, Sender, and Receiver cannot be empty');
        return;
    }

    if (isNaN(amount) || amount <= 0) {
        alert('Amount must be a valid number greater than 0');
        return;
    }

    if (!currency || !status || !timestamp) {
        alert('Currency, Status, and Timestamp are required');
        return;
    }

    // 准备交易数据
    const transactionData = {
        transactionId, sender, receiver, amount, currency, status, timestamp, note
    };

    console.log('Transaction Data:', transactionData); // 调试信息

    // 添加Loading提示
    const submitButton = transactionForm.querySelector('button[type="submit"]');
    submitButton.textContent = 'Submitting...';
    submitButton.disabled = true;

    try {
        // 向后端发送交易数据
        const response = await fetch('http://localhost:8080/api/submit-transaction', {
            method: 'POST', headers: {
                "Authorization": `Bearer ${token}`, 'Content-Type': 'application/json'
            }, body: JSON.stringify(transactionData)
        });

        if (response.ok) {
            const result = await response.json(); // 从响应中解析 JSON 数据
            console.log('Transaction saved:', result);

            // 提示用户交易成功，并显示交易哈希（txID）
            alert(`Transaction submitted successfully! TXID: ${result.transactionHash}`);

            // 显示完TXID之后跳转到 transactionhistory.html
            window.location.href = 'transactionhistory.html'; // 跳转到交易历史页面
        } else {
            const errorMessage = await response.text();
            console.error('Failed to submit transaction:', errorMessage);
            alert(`Failed to submit transaction: ${errorMessage}`);
        }
    } catch (error) {
        console.error('Error submitting transaction:', error);
        alert('An error occurred while submitting the transaction. Please try again.');
    } finally {
        // 恢复按钮状态
        submitButton.textContent = 'Submit Transaction';
        submitButton.disabled = false;
    }
});