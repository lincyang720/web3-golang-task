// const { ethers, deployments,upgrades } = require("hardhat");
// const {expect} = require("chai");


// // describe("start", async ()=>{
// //     it("should be able to  deploy ", async ()=>{
// //         const AuctionMarket = await ethers.getContractFactory("AuctionMarket");
// //         const auctionMarket = await AuctionMarket.deploy();
// //         await auctionMarket.waitForDeployment();

// //         await auctionMarket.createAuction(
// //             100 * 1000, 
// //             ethers.parseEther("0.1"), 
// //             ethers.ZeroAddress,
// //             1);

// //             const auction = await auctionMarket.auctions(0);
// //             console.log(auction);
// //     }); 

// // }



// describe("Test Upgrade", async function () {
//     it("should be able to upgrade the contract", async function () {
//         //1.部署业务合约(tags的名字)
//         await deployments.fixture(["deployAuctionMarket"]);
//         //导出的save文件的名字
//         const auctionMarketProxy = await deployments.get("AuctionMarketProxy");

//         //2.调用createAuction方法创建拍卖
//         //拿到合约客户端
//         const auctionMarket = await ethers.getContractAt("AuctionMarket", auctionMarketProxy.address);
//         await auctionMarket.createAuction(100 * 1000,
//             ethers.parseEther("0.01"),
//             ethers.ZeroAddress,
//             1);

//         const auction = await auctionMarket.auctions(0);
//         console.log("创建拍卖成功", auction);
//         const implementationAddress = await upgrades.erc1967.getImplementationAddress(auctionMarketProxy.address);
//         console.log("实现合约地址：", implementationAddress);

//         //3.升级合约
//         await deployments.fixture(["upgradeAuctionMarket"]);
//         const implementationAddressV2 = await upgrades.erc1967.getImplementationAddress(auctionMarketProxy.address);
//         console.log("升级后实现合约地址：", implementationAddressV2);
        
//         //4.读取合约的auctions[0]
//         const auction2 = await auctionMarket.auctions(0);
//         console.log("升级后拍卖信息：", auction2);

//         const auctionMarketV2 = await ethers.getContractAt("AuctionMarketV2", auctionMarketProxy.address);
//         const upgrade = await auctionMarketV2.testUpgrade();
//         console.log("升级后的值：", upgrade);

//         expect(auction2.startTime).to.equal(auction.startTime, "升级后拍卖信息不一致");
//         // expect(implementationAddress).to.equal(implementationAddressV2, "升级后实现合约地址不一致");
//     });
// });