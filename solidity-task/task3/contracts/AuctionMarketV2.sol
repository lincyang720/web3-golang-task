// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
//可升级的（先安装npm install hardhat-deploy,@openzeepelin/contracts-upgradeable）
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "hardhat/console.sol";


contract AuctionMarketV2 is Initializable {

    struct Auction {
        address seller;
        uint256 duration;
        uint256 startingPrice;
        uint256 startTime;
        bool ended;
        // 最高出价者
        address highestBidder;
        // 最高出价
        uint256 highestBid;
        // nft合约地址
        address nftContract;
        uint256 tokenId;
    }

    mapping(uint256 => Auction) public auctions;
    uint256 public nextAuctionId;
    //合约所有者
    address public owner;

    // 初始化函数，使用 Initializable 代替构造函数
    function initialize() initializer public {
        owner = msg.sender;
    }


    function createAuction(
        uint256 duration,
        uint256 startingPrice,
        address nftAddress,
        uint256 tokenId
    ) public {
        require(msg.sender == owner, "Only the owner can create an auction");
        require(duration > 1000 * 60, "Duration must be greater than 10 seconds");
        require(startingPrice > 0, "Starting price must be greater than 0");

        auctions[nextAuctionId] = Auction({
            tokenId: tokenId,
            seller: msg.sender,
            duration: duration,
            startingPrice: startingPrice,
            startTime: block.timestamp,
            ended: false,
            nftContract: nftAddress,
            highestBidder: address(0),
            highestBid: 0
        });

        nextAuctionId++;
    }

    function  placeBid(uint256 auctionId,uint256 amount) external payable {
        Auction storage auction = auctions[auctionId];
    
        require(!auction.ended && auction.startTime+auction.duration>block.timestamp, "Auction has already ended");

        require(amount >=auction.startingPrice && amount > auction.highestBid, "Bid must be higher than the starting price and the highest bid");
        

        if (auction.highestBidder != address(0)) {
            // Transfer the previous highest bid back to the bidder
            payable(auction.highestBidder).transfer(auction.highestBid);
        }

        auction.highestBidder = msg.sender;
        auction.highestBid = amount;
    }

    function endAuction(uint256 auctionId) public {
        Auction storage  auction = auctions[auctionId];

        console.log(
            "endAuction",
            auction.startTime,
            auction.duration,
            block.timestamp
        );
        

        require(!auction.ended && auction.startTime+auction.duration<=block.timestamp, "Auction has not ended");

        // IERC721(auction.nftContract).safeTransferFrom(address(this),auction.highestBidder,auction.tokenId);

        // payable(address(this)).transfer(address(this).balance);

        auction.ended=true;
    }


    function testUpgrade() public pure returns (string memory) {
        // 测试升级后的合约功能
        return "Upgrade successful";
    }
}