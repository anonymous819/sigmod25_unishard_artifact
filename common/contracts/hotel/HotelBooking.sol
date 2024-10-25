// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

contract HotelBooking {
  uint256 public constant MAX_ROOMS = 300;
  uint256 public constant HOTEL_PRICE = 1;
  address public immutable paymentService;
  uint256 public roomReserved;
  address[MAX_ROOMS] public rooms;

  constructor(address _paymentService) {
    paymentService = _paymentService;
  }

  function checkRoomAvailability(address account, uint payment) public view{
    require(roomReserved < MAX_ROOMS, "No more room available");
    require(payment >= HOTEL_PRICE, "Insufficient payment");
    (bool success, ) = paymentService.staticcall(abi.encodeWithSignature("verifyDeposit(address,uint256)", account, payment));
    require(success, "Payment failed");

  }

  function bookHotel(address account) public {
    (bool success, ) = paymentService.call(abi.encodeWithSignature("receivePayment(address,uint256)", account, HOTEL_PRICE));
    require(success, "Fail to receive payment");
    rooms[roomReserved++] = account;
  }
}
