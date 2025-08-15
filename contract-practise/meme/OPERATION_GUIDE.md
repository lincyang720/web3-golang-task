# MemeShib 代币合约操作指南

本文档详细说明了如何部署和使用 MemeShib 代币合约，包括合约部署、代币交易、添加和移除流动性等操作。

## 目录
1. [合约功能概述](#合约功能概述)
2. [部署前准备](#部署前准备)
3. [合约部署](#合约部署)
4. [代币交易](#代币交易)
5. [流动性管理](#流动性管理)
6. [交易限制机制](#交易限制机制)
7. [常见操作示例](#常见操作示例)

## 合约功能概述

MemeShib 是一个基于以太坊的 SHIB 风格 Meme 代币合约，具有以下核心功能：

1. **代币税机制**：对每笔交易征收 0.5% 的税费，税费将发送到指定的税费地址
2. **代币销毁机制**：每笔交易将销毁 0.01% 的代币，减少流通供应量
3. **流动性池集成**：与 Uniswap V2 集成，支持添加和移除流动性
4. **交易限制机制**：设置单笔交易最大额度和每日交易次数限制，防止市场操纵

## 部署前准备

在部署合约之前，请确保完成以下准备工作：

1. 安装 Node.js 和 npm
2. 安装 Hardhat 开发环境：
   ```bash
   npm init -y
   npm install --save-dev hardhat
   ```
3. 安装 OpenZeppelin 合约库：
   ```bash
   npm install @openzeppelin/contracts
   ```
4. 安装 Uniswap V2 接口：
   ```bash
   npm install @uniswap/v2-periphery
   ```

## 合约部署

1. 将 [memeShib.sol](memeShib.sol) 文件放置在 `contracts` 目录中
2. 编写部署脚本（如 `deploy.js`）：
   ```javascript
   const { ethers } = require("hardhat");
   
   async function main() {
     const MemeShib = await ethers.getContractFactory("memeShib");
     const memeShib = await MemeShib.deploy();
     
     await memeShib.deployed();
     
     console.log("MemeShib deployed to:", memeShib.address);
   }
   
   main()
     .then(() => process.exit(0))
     .catch((error) => {
       console.error(error);
       process.exit(1);
     });
   ```
3. 编译合约：
   ```bash
   npx hardhat compile
   ```
4. 部署合约：
   ```bash
   npx hardhat run scripts/deploy.js --network <network-name>
   ```

## 代币交易

### 标准转账 (transfer)

任何持有代币的用户都可以使用 `transfer` 函数向其他地址发送代币：

```javascript
// JavaScript 示例
const memeShib = new ethers.Contract(contractAddress, abi, signer);
await memeShib.transfer(recipientAddress, amount);
```

转账过程将自动执行以下操作：
1. 检查发送方余额是否充足
2. 检查是否超过单笔交易最大额度（1000 * 10^9）
3. 检查当日交易次数是否超过限制（5次）
4. 从转账金额中扣除 0.5% 作为税费发送到税费地址
5. 销毁转账金额的 0.01%
6. 增加发送方当日交易次数

### 授权转账 (transferFrom)

代币持有者可以授权第三方代表其转账代币：

1. 首先授权额度：
   ```javascript
   await memeShib.approve(spenderAddress, amount);
   ```
2. 被授权方执行转账：
   ```javascript
   await memeShib.transferFrom(ownerAddress, recipientAddress, amount);
   ```

## 流动性管理

流动性操作只能由合约所有者执行。

### 添加流动性

```javascript
// 添加流动性示例
await memeShib.addLiquidity(tokenAmount, ethAmount, {
  value: ethAmount // 发送 ETH 到合约
});
```

该函数将：
1. 授权 Uniswap 路由器使用指定数量的代币
2. 调用 Uniswap V2 的 `addLiquidityETH` 函数添加流动性

### 移除流动性

```javascript
// 移除流动性示例
await memeShib.removeLiquidity(liquidityAmount);
```

该函数将调用 Uniswap V2 的 `removeLiquidityETH` 函数移除流动性。

## 交易限制机制

合约实现了两层交易限制机制：

1. **单笔交易最大额度**：默认为 1000 * 10^9 代币
2. **每日交易次数限制**：默认为 5 次

### 限制重置

合约所有者可以手动重置指定账户的每日交易次数：

```javascript
await memeShib.resetDailyTxCount(userAddress);
```

## 常见操作示例

### 查看账户余额

```javascript
const balance = await memeShib.balanceOf(address);
```

### 查看交易次数

```javascript
const txCount = await memeShib.dailyTxCount(address);
```

### 修改税费地址（仅限所有者）

由于税费地址在构造函数中设置为合约创建者，如果需要修改，需要通过合约升级或添加新函数实现。

### 查看合约参数

```javascript
const taxRate = await memeShib.taxRate();
const taxAddress = await memeShib.taxAddress();
const maxTxAmount = await memeShib.maxTxAmount();
const maxDailyTxCount = await memeShib.maxDailyTxCount();
const burnRate = await memeShib.burnRate();
```

## 注意事项

1. 合约使用 Uniswap V2 路由器地址 `0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D`，该地址在主网和大多数测试网上有效
2. 交易税费默认为 0.5%，销毁比例默认为 0.01%，这些参数可以在合约中修改
3. 每日交易次数限制默认为 5 次，需要手动重置或通过额外开发实现自动重置
4. 流动性操作仅限合约所有者执行，以确保资金安全

## 故障排除

1. **编译错误**：确保所有依赖库已正确安装
2. **部署失败**：检查网络连接和账户余额
3. **交易失败**：检查余额、交易限制和授权额度
4. **流动性操作失败**：确认调用者为合约所有者

通过遵循本指南，您可以成功部署和使用 MemeShib 代币合约。如有任何问题，请参考 Solidity 文档和相关库的官方文档。