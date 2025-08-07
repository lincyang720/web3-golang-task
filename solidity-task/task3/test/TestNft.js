const { ethers, deployments } = require("hardhat")
const { expect } = require("chai");

describe("Test auction market", async function () {
    it("should be ok", async function () {
        await main();
    });
        
})
async function main() {
    await deployments.fixture(["deployAuctionMarket"])
    const auctionMarketProxy = await deployments.get("AuctionMarketProxy");

    const [signer,buyer] = await ethers.getSigners();

    //部署721合约
    const TestNft = await ethers.getContractFactory("TestNft");
    const testNft = await TestNft.deploy();
    await testNft.waitForDeployment();

    const testNftAddress = await testNft.getAddress();
    console.log("NFT合约地址：", testNftAddress);

    for (let i = 0; i < 10; i++) {
        await testNft.mint(signer.address, i + 1);
    }
    //调用createAuction方法创建拍卖
    const tokenId = 1; // 假设我们要拍卖的NFT的ID是1
    const auctionMarket = await ethers.getContractAt("AuctionMarket", auctionMarketProxy.address);

    //给代理合约授权
    await testNft.connect(signer).setApprovalForAll(auctionMarketProxy.address, true);

    await auctionMarket.createAuction(10,
        ethers.parseEther("0.01"),
        testNftAddress,
        tokenId
    );


    // await auctionMarket.setPriceETHFeed(ethers.ZeroAddress, "0x694AA1769357215DE4FAC081bf1f309aDC325306");
    // ethPriceValue = await auctionMarket.getChainlinkDataFeedLatestAnswer(ethers.ZeroAddress);
    // console.log("eth价格", ethPriceValue);

    const auction = await auctionMarket.auctions(0);
    console.log("创建拍卖成功", auction); 

 
   
    //参与拍卖
    await auctionMarket.connect(buyer).placeBid(0,0, ethers.ZeroAddress,{ value: ethers.parseEther("0.01") });

    //结束拍卖
    await new Promise((r) => setTimeout(r, 10 * 1000)); // 等待10秒钟
    await auctionMarket.connect(signer).endAuction(0);

    //验证结果
    const auctionResult = await auctionMarket.auctions(0);
    console.log("拍卖结束后的结果：", auctionResult);
    expect(auctionResult.highestBidder).to.equal(buyer.address, "最高出价者不正确");
    expect(auctionResult.highestBid).to.equal(ethers.parseEther("0.01"), "最高出价不正确");

    //验证nft所有权
    const owner = await testNft.ownerOf(tokenId);
    console.log("NFT的所有者：", owner);
    expect(owner).to.equal(buyer.address, "NFT的所有者不正确");
}