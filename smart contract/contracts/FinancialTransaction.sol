// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract FinancialTransaction is Ownable, ReentrancyGuard {
    using Strings for uint256;
    
    enum Currency { USD, EUR }
    
    struct Transaction {
        string transactionId;
        string senderAccount;
        string receiverAccount;
        uint256 amount;
        Currency currency;
        uint256 timestamp;
        uint256 blockTimestamp;
        string note;
    }

    mapping(bytes32 => Transaction) private transactions;
    mapping(string => bytes32[]) private accountTransactions;
    mapping(address => bool) private authorizedUsers;

    event TransactionRecorded(
        bytes32 indexed hash,
        string transactionId,
        string senderAccount,
        string receiverAccount,
        uint256 amount,
        Currency currency,
        uint256 timestamp,
        uint256 blockTimestamp,
        string note
    );

    event TransactionNoteUpdated(bytes32 indexed hash, string note);
    event UserAuthorized(address user);
    event UserUnauthorized(address user);

    constructor() Ownable(msg.sender) {
        authorizedUsers[msg.sender] = true;
    }

    modifier onlyAuthorized() {
        require(authorizedUsers[msg.sender], "Not authorized");
        _;
    }

    function authorizeUser(address user) public onlyOwner {
        authorizedUsers[user] = true;
        emit UserAuthorized(user);
    }

    function unauthorizeUser(address user) public onlyOwner {
        authorizedUsers[user] = false;
        emit UserUnauthorized(user);
    }

    function recordTransaction(
        string memory _transactionId,
        string memory _senderAccount,
        string memory _receiverAccount,
        uint256 _amount,
        Currency _currency,
        uint256 _timestamp,
        string memory _note
    ) public onlyAuthorized nonReentrant returns (bytes32) {
        require(bytes(_transactionId).length > 0, "Transaction ID cannot be empty");
        require(bytes(_senderAccount).length > 0, "Sender account cannot be empty");
        require(bytes(_receiverAccount).length > 0, "Receiver account cannot be empty");
        require(_amount > 0, "Amount must be greater than 0");
        
        bytes32 transactionHash = keccak256(abi.encodePacked(_transactionId, _senderAccount, _receiverAccount, _amount, _currency, _timestamp, block.timestamp));
        
        require(transactions[transactionHash].timestamp == 0, "Transaction already exists");

        transactions[transactionHash] = Transaction({
            transactionId: _transactionId,
            senderAccount: _senderAccount,
            receiverAccount: _receiverAccount,
            amount: _amount,
            currency: _currency,
            timestamp: _timestamp,
            blockTimestamp: block.timestamp,
            note: _note
        });

        accountTransactions[_senderAccount].push(transactionHash);
        accountTransactions[_receiverAccount].push(transactionHash);

        emit TransactionRecorded(
            transactionHash,
            _transactionId,
            _senderAccount,
            _receiverAccount,
            _amount,
            _currency,
            _timestamp,
            block.timestamp,
            _note
        );

        return transactionHash;
    }

    function getTransactionByHash(bytes32 _transactionHash) public view onlyAuthorized returns (Transaction memory) {
        require(transactions[_transactionHash].timestamp != 0, "Transaction does not exist");
        return transactions[_transactionHash];
    }

    function getAccountTransactions(string memory _account) public view onlyAuthorized returns (bytes32[] memory) {
        return accountTransactions[_account];
    }

    function getCurrencyName(Currency _currency) public pure returns (string memory) {
        if (_currency == Currency.USD) {
            return "USD";
        } else if (_currency == Currency.EUR) {
            return "EUR";
        } else {
            revert("Invalid currency");
        }
    }

    function updateTransactionNote(bytes32 _transactionHash, string memory _note) public onlyAuthorized {
        require(transactions[_transactionHash].timestamp != 0, "Transaction does not exist");
        transactions[_transactionHash].note = _note;
        emit TransactionNoteUpdated(_transactionHash, _note);
    }

    function isAuthorized(address user) public view returns (bool) {
        return authorizedUsers[user];
    }
}