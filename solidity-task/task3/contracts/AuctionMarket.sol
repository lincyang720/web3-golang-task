// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
//可升级的（先安装npm install hardhat-deploy,@openzeepelin/contracts-upgradeable）
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "hardhat/console.sol";


contract AuctionMarket is Initializable {

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

        //参与竞价的资产类型
        //0 表示ETH
        //其他表示ERC20
        address tokenAddress;
    }

    mapping(uint256 => Auction) public auctions;
    uint256 public nextAuctionId;
    //合约所有者
    address public owner;

    //代币到喂价合约的映射
    // AggregatorV3Interface internal priceETHFeed;
    mapping(address => AggregatorV3Interface) public priceFeeds;

    //设置兑换eth到usd的地址
    function setPriceETHFeed(address tokenAddress,address _priceFeed) public {
        priceFeeds[tokenAddress] = AggregatorV3Interface(_priceFeed);
    }

    //ETH=>USD  371253943274 3712.53943274
    //USDC=>USD 99985036  0.99985036
    function getChainlinkDataFeedLatestAnswer(address tokenAddress) public view returns (int) {
        if(tokenAddress == address(0)) {
            // ETH
            return 371253943274; // 3712.53943274      }
        }
        AggregatorV3Interface priceFeed = priceFeeds[tokenAddress];
        (
            /*uint80 roundID*/,
            int256 price,
            /*uint startedAt*/,
            /*uint timeStamp*/,
            /*uint80 answeredInRound*/
        ) = priceFeed.latestRoundData();
        return price;
    }

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
        require(duration >= 10, "Duration must be greater than 10 seconds");
        require(startingPrice > 0, "Starting price must be greater than 0");

        //转移nft到合约
        IERC721(nftAddress).approve(address(this), tokenId);
        // IERC721(nftAddress).safeTransferFrom(msg.sender, address(this), tokenId);

        auctions[nextAuctionId] = Auction({
            tokenId: tokenId,
            seller: msg.sender,
            duration: duration,
            startingPrice: startingPrice,
            startTime: block.timestamp,
            ended: false,
            nftContract: nftAddress,
            highestBidder: address(0),
            highestBid: 0,
            tokenAddress: address(0)
        });

        nextAuctionId++;
    }

    function  placeBid(uint256 auctionId,uint256 amount,address tokenAddress) external payable {
        Auction storage auction = auctions[auctionId];
    
        require(!auction.ended && auction.startTime+auction.duration>block.timestamp, "Auction has already ended");

        uint payValue;
        if(tokenAddress!=address(0)){
            //处理ERC20
            payValue = amount * uint(getChainlinkDataFeedLatestAnswer(tokenAddress));

        }else{
            //处理ETH
            console.log("msg.value",msg.value);
            amount = msg.value;
            payValue = amount * uint(getChainlinkDataFeedLatestAnswer(address(0)));
            console.log("amount",amount);
            console.log("payValue",payValue);
        }
        
        uint256 startPriceValue = auction.startingPrice * uint(getChainlinkDataFeedLatestAnswer(auction.tokenAddress));
        uint256 highestBidValue = auction.highestBid * uint(getChainlinkDataFeedLatestAnswer(auction.tokenAddress));

        require(payValue >= startPriceValue && payValue > highestBidValue, "Bid must be higher than the starting price and the highest bid");

        if (tokenAddress != address(0)) { 
            IERC20(tokenAddress).transferFrom(msg.sender,address(this),amount);
        }
        
        if(auction.highestBid>0){
            if(auction.tokenAddress == address(0)){
                payable(auction.highestBidder).transfer(auction.highestBid);
            }else{
                //转移之前的最高出价者的资金
                IERC20(auction.tokenAddress).transfer(auction.highestBidder,auction.highestBid);
            }
        }
        auction.tokenAddress = tokenAddress;
        auction.highestBid = amount;
        auction.highestBidder = msg.sender;
    }

    function endAuction(uint256 auctionId) public {
        Auction storage  auction = auctions[auctionId];

        console.log(
            "endAuction",
            auction.startTime,
            auction.duration,
            block.timestamp
        );
        
        console.log(auction.startTime+auction.duration,block.timestamp);

        require(!auction.ended && (auction.startTime+auction.duration)<=block.timestamp, "Auction has not ended");

        //转移nft到最高出价者
        IERC721(auction.nftContract).safeTransferFrom(address(this),auction.highestBidder,auction.tokenId);
        //转移剩余的资金到卖家
        // payable(auction.seller).transfer(auction.highestBid);

        auction.ended=true;
    }
}