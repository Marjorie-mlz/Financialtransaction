const FinancialTransaction = artifacts.require("FinancialTransaction");

module.exports = function (deployer) {
  deployer.deploy(FinancialTransaction, { gas: 6721974 });
}; 