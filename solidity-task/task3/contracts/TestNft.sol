pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TestNft is ERC721, Ownable {
    string private _tokenURI;

    constructor() ERC721("TestNFT", "TNFT") Ownable(msg.sender) {}

    function mint(address to,uint256 tokenId) public  onlyOwner {
        _mint(to, tokenId);
    }

    function getTokenURI() public  view returns (string memory) {
        return _tokenURI;
    }

    function setTokenURI(string memory tokenURI) external onlyOwner {
        _tokenURI = tokenURI;
    }
}