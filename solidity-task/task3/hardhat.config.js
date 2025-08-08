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
      url: "https://sepolia.infura.io/v3/",
      accounts: [""],
    },
  },
  namedAccounts: {
    deployer: 0,
    user1: 1,
    user2: 2,
  }
};
