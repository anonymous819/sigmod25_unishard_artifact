// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

contract TrainBooking {
  uint256 public constant MAX_SEATS = 300;
  uint256 public constant TRAIN_PRICE = 1;
  address public immutable paymentService;
  uint256 public ticketSold;
  address[MAX_SEATS] public tickets;

  constructor(address _paymentService) {
    paymentService = _paymentService;
  }

  function checkSeatAvailability(address account, uint payment) public view {
    require(ticketSold < MAX_SEATS, "No more ticket available");
    require(payment >= TRAIN_PRICE, "Insufficient payment");
    (bool success, ) = paymentService.staticcall(abi.encodeWithSignature("verifyDeposit(address,uint256)", account, payment));
    require(success, "Payment failed");
  }

  function bookTrain(address account) public {
    (bool success, ) = paymentService.call(abi.encodeWithSignature("receivePayment(address,uint256)", account, TRAIN_PRICE));
    require(success, "Fail to receive payment");
    tickets[ticketSold++] = account;
  }
}