const {deployments,upgrades, ethers} = require("hardhat")

const fs = require("fs");
const path = require("path");


// const {deploy} = deployments;

module.exports = async ({getNamedAccounts, deployments}) => {
    const {save} = deployments;
    const {deployer} = await getNamedAccounts(); 
    
    console.log("部署用户地址：",   deployer);
    const AuctionMarket = await ethers.getContractFactory("AuctionMarket");

    //通过代理部署合约
    const auctionMarketProxy = await upgrades.deployProxy(AuctionMarket, [],{initializer: "initialize"});

    await auctionMarketProxy.waitForDeployment();

    const proxyAddress = await auctionMarketProxy.getAddress();
    console.log("代理 合约地址：", proxyAddress);
    const implementationAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
    console.log("实现合约地址：", implementationAddress);

    const storePath = path.resolve(__dirname, "./.cache/proxyAuctionMarket.json");
    fs.writeFileSync(storePath, 
        JSON.stringify({proxyAddress, implementationAddress}));

    await save("AuctionMarketProxy", {
        address: proxyAddress,
        // abi: JSON.parse(auctionMarketProxy.interface.format("json")),
        args: [],
        log: true
    });
    // await deploy("AuctionMarket", {
    //     from: deployer,
    //     args: [],
    //     log: true,
    // });
};

module.exports.tags = ["deployAuctionMarket"];