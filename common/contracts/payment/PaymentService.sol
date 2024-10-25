// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

contract PaymentService {
  mapping (address => uint) public balances;

  function depositpayment() public payable {
    balances[msg.sender] += msg.value;
  }

  function verifyDeposit(address account, uint payment) public view {
    require(balances[account] >= payment, "Insufficient balance");
  }

  function receivePayment(address account, uint payment) public {
    require(balances[account] >= payment, "Insufficient balance");
    balances[account] -= payment;
    balances[msg.sender] += payment;  
  }
}