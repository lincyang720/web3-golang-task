require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");
require("@openzeppelin/hardhat-upgrades");
// 添加这一行来引入solidity-coverage插件
require("solidity-coverage");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks:{
    sepolia: {
      url: "https://sepolia.infura.io/v3/023fb439d0cb4bf784b8827497e71da9",
      accounts: ["5fb0d527911392bd021a6cad58c90bc5794b1a7354f5840f38170c758f95c6a7"],
    },
  },
  namedAccounts: {
    deployer: 0,
    user1: 1,
    user2: 2,
  }
};
