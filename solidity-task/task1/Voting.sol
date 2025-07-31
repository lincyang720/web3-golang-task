// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {

    mapping(string name => uint8 num) public votesReceived;

    string[] public candidates;

    constructor(string[] memory initialCandidates) {
        candidates = initialCandidates;
    }

    function vote(string memory name , uint8  num) public{
        votesReceived[name] += num;
    }

    function getVotes(string memory name) public view  returns (uint8) {
        return votesReceived[name];
    }

    //重置所有候选人的得票数
    function resetVotes() public {
        for(uint i = 0; i < candidates.length; i++) {
            votesReceived[candidates[i]] = 0;
        }
    }

}