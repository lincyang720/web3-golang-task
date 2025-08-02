// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

event Donation(address from, uint256 amount);

contract BeggingContract {
    mapping(address => uint256) public balance;
    mapping(address => uint256) private donateMapping;
    address owner;

    constructor() {
        owner = msg.sender;
    }

    function donate() public payable returns (bool) {
        donateMapping[msg.sender] += msg.value;
        balance[owner] += msg.value;
        emit Donation(msg.sender, msg.value);
        return true;
    }

    modifier onlyOwner {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    function withdraw(uint256 amount) public onlyOwner {
        require(address(this).balance >= amount, "Insufficient contract balance");
        payable(owner).transfer(amount);
        balance[owner] -= amount;
    }

    function getDonation(address from) public view returns (uint256) {
        return donateMapping[from];
    }
}