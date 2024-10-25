// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

contract TravelAgency {
  address public immutable trainBooking;
  address public immutable hotelBooking;
  mapping(address=>bool) customers;

  constructor(address _trainBooking, address _hotelBooking) {
    trainBooking = _trainBooking;
    hotelBooking = _hotelBooking;
  }

  function bookTrainAndHotel(uint payment1, uint payment2) public {
    bool trainAvailable;
    bool hotelAvailable;
    bool bookingSuccess;

    (trainAvailable, ) = trainBooking.staticcall(abi.encodeWithSignature("checkSeatAvailability(address,uint256)", msg.sender, payment1));
    require(trainAvailable, "Train seat is not available.");

    (hotelAvailable, ) = hotelBooking.staticcall(abi.encodeWithSignature("checkRoomAvailability(address,uint256)", msg.sender, payment2));
    require(hotelAvailable, "Hotel room is not available.");

    (bookingSuccess, ) = trainBooking.call(abi.encodeWithSignature("bookTrain(address)", msg.sender));
    require(bookingSuccess, "Train booking failed.");

    (bookingSuccess, ) = hotelBooking.call(abi.encodeWithSignature("bookHotel(address)", msg.sender));
    require(bookingSuccess, "Hotel booking failed.");
    
    customers[msg.sender] = true;
  }
}