// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract Counter{
    uint256 private count;

    function increment() external returns (uint256) {
        count += 1;
        return count;
    }

    function decrement() external returns (uint256) {
        require(count > 0, "Counter: count is already zero");
        count -= 1;
        return count;
    }
}