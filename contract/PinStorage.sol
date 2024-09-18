// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PinStorage {
    // 存储PIN码的变量
    string private pinCode;
    address private owner;

    // 事件，用于记录PIN码的更改
    event PinChanged(string oldPin, string newPin);
    event PinSet(string pin);

    // 构造函数，设置初始PIN码和合约拥有者
    constructor(string memory _initialPin) {
        pinCode = _initialPin;
        owner = msg.sender;
        emit PinSet(_initialPin);
    }

    // 获取PIN码的方法，任何人都可以调用
    function getPin() public view returns (string memory) {
        return pinCode;
    }

    // 修改PIN码的方法，只有合约拥有者可以调用
    function setPin(string memory _newPin) public {
        require(msg.sender == owner, "Only the owner can change the PIN.");
        string memory oldPin = pinCode;
        pinCode = _newPin;
        emit PinChanged(oldPin, _newPin);
    }
}
