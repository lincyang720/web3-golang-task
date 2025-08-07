const {deployments, upgrades, ethers} = require("hardhat");

const path = require("path");
const fs = require("fs");

module.exports = async ({getNamedAccounts, deployments}) =>
{
  const {save} = deployments;
  const {deployer} = await getNamedAccounts();
  console.log("部署用户地址：",   deployer);

  //读取.cache/proxyAuctionMarket.json文件
  const storePath = path.resolve(__dirname, "./.cache/proxyAuctionMarket.json");


  const {proxyAddress} = JSON.parse(fs.readFileSync(storePath, "utf8"));

  //升级版的业务合约
  const AuctionMarketV2 = await ethers.getContractFactory("AuctionMarketV2");

  //升级代理合约
  const auctionMarketProxyV2 = await upgrades.upgradeProxy(proxyAddress, AuctionMarketV2);
  await auctionMarketProxyV2.waitForDeployment();

  const proxyAddressV2 = await auctionMarketProxyV2.getAddress();


    await save("AuctionMarketProxyV2", {
        address: proxyAddressV2,
        // abi: JSON.parse(auctionMarketProxyV2.interface.format("json")),
        // args: [],
        // log: true
    });

}

module.exports.tags = ["upgradeAuctionMarket"];
