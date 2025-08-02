// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract myNFT  is ERC721URIStorage {

    uint256 private _tokenIdCounter;
    address public owner;

    constructor() ERC721("myNFT","mNFT"){
        owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == owner,"Only owner can call this function");
        _;
    }

    function mintNFT(address recipient,string memory tokenURI) public onlyOwner returns (uint256) {
        _tokenIdCounter++;
        uint256 newTokenId = _tokenIdCounter;
        _safeMint(recipient,_tokenIdCounter);
        _setTokenURI(_tokenIdCounter,tokenURI);
        return newTokenId;
    }
}