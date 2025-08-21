const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");
const { time } = require("@nomicfoundation/hardhat-network-helpers");

describe("MetaNodeStake test", function() {
    let metaNodeStakeContract;
    let metaNodeStakeAddress;
    let metaNodeContract;
    let metaNodeAddress;
    let owner;
    let user1;
    let user2;
    let startBlock;
    let endBlock;
    let metaNodePerBlock;
    let poolAdded = false;
    
    // 区块管理常量
    const START_BLOCK_DELAY = 5;
    const END_BLOCK_DELAY = 1000;

    // 获取最新区块号并转换为BigInt
    async function getLatestBlockNumber() {
        return Number(await ethers.provider.getBlockNumber());
    }

    before(async function() {
        this.timeout(60000);
        
        [owner, user1, user2] = await ethers.getSigners();
        
        // 部署 MetaNode 合约
        const MetaNode = await ethers.getContractFactory("MetaNode");
        metaNodeContract = await MetaNode.deploy(owner.address);
        await metaNodeContract.waitForDeployment();
        metaNodeAddress = await metaNodeContract.getAddress();
        console.log("MetaNode deployed to:", metaNodeAddress);
        
        // 添加额外的地址有效性检查
        expect(ethers.isAddress(metaNodeAddress)).to.be.true;
        expect(metaNodeAddress).to.not.be.null;
        expect(metaNodeAddress).to.not.be.undefined;
        
        // deploy MetaNodeStake contract
        const MetaNodeStake = await ethers.getContractFactory("MetaNodeStake");
        const currentBlock = await getLatestBlockNumber();
        startBlock = currentBlock + START_BLOCK_DELAY;
        endBlock = startBlock + END_BLOCK_DELAY;
        metaNodePerBlock = ethers.parseEther("1"); // 1 MetaNode per block
        
        // 添加参数验证
        expect(ethers.isAddress(metaNodeAddress)).to.be.true;
        expect(metaNodeAddress).to.not.be.null;
        expect(metaNodeAddress).to.not.be.undefined;
        expect(startBlock).to.be.lt(endBlock);
        
        // 添加日志输出调试信息
        console.log("Deploying MetaNodeStake with parameters:");
        console.log("MetaNode Address:", metaNodeAddress);
        console.log("startBlock:", startBlock);
        console.log("endBlock:", endBlock);
        console.log("metaNodePerBlock:", metaNodePerBlock.toString());
        
        metaNodeStakeContract = await upgrades.deployProxy(
            MetaNodeStake, [metaNodeAddress, startBlock, endBlock, metaNodePerBlock],
            {
                initializer: 'initialize',
                kind: 'uups',
             });
        await metaNodeStakeContract.waitForDeployment();
        metaNodeStakeAddress = await metaNodeStakeContract.getAddress();
        console.log("MetaNodeStake deployed to:", metaNodeStakeAddress);
        
        // 添加部署后验证
        expect(ethers.isAddress(metaNodeStakeAddress)).to.be.true;
        expect(metaNodeStakeAddress).to.not.be.null;
        expect(metaNodeStakeAddress).to.not.be.undefined;
        
        const deployedMetaNodeAddress = await metaNodeStakeContract.MetaNode(); // 使用合约中定义的变量名
        console.log("Deployed MetaNode Address:", deployedMetaNodeAddress);
        expect(deployedMetaNodeAddress).to.equal(metaNodeAddress);
        
        // 给MetaNodeStake合约铸造大量MetaNode代币用于测试
        const mintAmount = ethers.parseEther("1000000");
        await metaNodeContract.mint(metaNodeStakeAddress, mintAmount);
        console.log("Minted", mintAmount.toString(), "MetaNode tokens to MetaNodeStake contract");
        
        // 验证MetaNodeStake合约的MetaNode余额
        const metaNodeStakeBalance = await metaNodeContract.balanceOf(metaNodeStakeAddress);
        console.log("MetaNodeStake contract MetaNode balance:", metaNodeStakeBalance.toString());
        expect(metaNodeStakeBalance).to.be.gte(mintAmount);
        
        // 等待到startBlock
        await advanceToStartBlock();
    });

    // 等待到startBlock的专用函数
    async function advanceToStartBlock() {
        const latestBlock = await getLatestBlockNumber();
        if (latestBlock < startBlock) {
            await time.advanceBlockTo(startBlock);
        }
    }

    // 确保池子已添加的函数
    async function ensurePoolAdded() {
        if (!poolAdded) {
            // 双重验证区块状态
            await advanceToStartBlock();
            
            await metaNodeStakeContract.addPool(
                ethers.ZeroAddress, 
                500, 
                ethers.parseEther("0.1"),
                20, 
                false
            );
            poolAdded = true;
        }
    }

    // 动态更新endBlock的专用函数
    async function updateEndBlockIfNeeded() {
        const currentBlock = await getLatestBlockNumber();
        const contractEndBlock = Number(await metaNodeStakeContract.endBlock());
        
        if (currentBlock >= contractEndBlock) {
            const newEndBlock = currentBlock + END_BLOCK_DELAY;
            await metaNodeStakeContract.setEndBlock(newEndBlock);
            endBlock = newEndBlock;
        }
    }

    // 动态更新startBlock的专用函数
    async function updateStartBlockIfNeeded() {
        const currentBlock = await getLatestBlockNumber();
        const contractStartBlock = Number(await metaNodeStakeContract.startBlock());
        const contractEndBlock = Number(await metaNodeStakeContract.endBlock());
        
        // 如果startBlock不安全或过期
        if (contractStartBlock >= contractEndBlock || 
            contractStartBlock + 20 < currentBlock || 
            contractStartBlock > currentBlock + 50) {
            const newStartBlock = currentBlock + START_BLOCK_DELAY;
            const newEndBlock = newStartBlock + END_BLOCK_DELAY;
            
            await metaNodeStakeContract.setStartBlock(newStartBlock);
            await metaNodeStakeContract.setEndBlock(newEndBlock);
            startBlock = newStartBlock;
            endBlock = newEndBlock;
        }
    }

    // 统一的区块管理钩子
    beforeEach(async function () {
        // 动态更新endBlock
        await updateEndBlockIfNeeded();
        // 动态更新startBlock
        await updateStartBlockIfNeeded();
    });

    describe("管理员功能测试", function () {
        it("should successfully add pool", async () => {
            expect(poolAdded).to.be.false;
            
            // 确保达到startBlock
            await advanceToStartBlock();
            
            await metaNodeStakeContract.addPool(
                ethers.ZeroAddress, 
                500, 
                ethers.parseEther("0.1"),
                20, 
                false
            );
            poolAdded = true;
            
            const poolInfo = await metaNodeStakeContract.pool(0);
            expect(poolInfo.stTokenAddress).to.equal(ethers.ZeroAddress);
            expect(poolInfo.poolWeight).to.equal(500);
        });

        it("应该能暂停和恢复提现", async function () {
            await ensurePoolAdded();
            await expect(metaNodeStakeContract.pauseWithdraw())
                .to.emit(metaNodeStakeContract, "PauseWithdraw");
            
            expect(await metaNodeStakeContract.withdrawPaused()).to.be.true;

            await expect(metaNodeStakeContract.unpauseWithdraw())
                .to.emit(metaNodeStakeContract, "UnpauseWithdraw");
            
            expect(await metaNodeStakeContract.withdrawPaused()).to.be.false;
        });

        it("应该能暂停和恢复领取奖励", async function () {
            await ensurePoolAdded();
            await expect(metaNodeStakeContract.pauseClaim())
                .to.emit(metaNodeStakeContract, "PauseClaim");
            
            expect(await metaNodeStakeContract.claimPaused()).to.be.true;

            await expect(metaNodeStakeContract.unpauseClaim())
                .to.emit(metaNodeStakeContract, "UnpauseClaim");
            
            expect(await metaNodeStakeContract.claimPaused()).to.be.false;
        });

        it("应该能设置开始和结束区块", async function () {
            await ensurePoolAdded();
            const currentBlock = await getLatestBlockNumber();
            const newStartBlock = currentBlock + 100;
            const newEndBlock = currentBlock + 2000;

            await expect(metaNodeStakeContract.setStartBlock(newStartBlock))
                .to.emit(metaNodeStakeContract, "SetStartBlock")
                .withArgs(newStartBlock);
            
            await expect(metaNodeStakeContract.setEndBlock(newEndBlock))
                .to.emit(metaNodeStakeContract, "SetEndBlock")
                .withArgs(newEndBlock);

            expect(await metaNodeStakeContract.startBlock()).to.equal(newStartBlock);
            expect(await metaNodeStakeContract.endBlock()).to.equal(newEndBlock);
        });

        it("应该能设置每区块奖励数量", async function () {
            await ensurePoolAdded();
            const newMetaNodePerBlock = ethers.parseEther("2");

            await expect(metaNodeStakeContract.setMetaNodePerBlock(newMetaNodePerBlock))
                .to.emit(metaNodeStakeContract, "SetMetaNodePerBlock")
                .withArgs(newMetaNodePerBlock);

            expect(await metaNodeStakeContract.MetaNodePerBlock()).to.equal(newMetaNodePerBlock);
        });

        it("应该能更新池子信息", async function () {
            await ensurePoolAdded();
            const newMinDepositAmount = ethers.parseEther("0.5");
            const newUnstakeLockedBlocks = 30;

            await expect(metaNodeStakeContract.updatePool(0, newMinDepositAmount, newUnstakeLockedBlocks))
                .to.emit(metaNodeStakeContract, "UpdatePoolInfo")
                .withArgs(0, newMinDepositAmount, newUnstakeLockedBlocks);

            const poolInfo = await metaNodeStakeContract.pool(0);
            expect(poolInfo.minDepositAmount).to.equal(newMinDepositAmount);
            expect(poolInfo.unstakeLockedBlocks).to.equal(newUnstakeLockedBlocks);
        });
     
        // 为质押功能测试添加专用区块推进
    describe("质押功能测试", function () {
        it("应该能质押 ETH", async function () {
            // 强制等待到startBlock
            await advanceToStartBlock();
            
            // 确保池子已添加
            await ensurePoolAdded();
            
            // 推进区块到startBlock+1，确保startBlock < current < endBlock
            const currentBlock = await getLatestBlockNumber();
            if (currentBlock < startBlock + 1) {
                await time.advanceBlockTo(startBlock + 1);
            }
            
            const stakeAmount = ethers.parseEther("1");
            await expect(metaNodeStakeContract.connect(user1).depositETH({
                value: stakeAmount
            }))
                .to.emit(metaNodeStakeContract, "Deposit")
                .withArgs(user1.address, 0, stakeAmount);

            const userInfo = await metaNodeStakeContract.user(0, user1.address);
            expect(userInfo.stAmount).to.equal(stakeAmount);
        });
    });

    // 解质押功能测试
    describe("解质押功能测试", function () {
        it("应该能申请解质押 ETH", async function () {
            // 专用测试前准备
            await advanceToStartBlock();
            await ensurePoolAdded();
            
            // 提前推进区块
            const currentBlock = await getLatestBlockNumber();
            if (currentBlock < startBlock + 1) {
                await time.advanceBlockTo(startBlock + 1);
            }
            
            // 质押并解质押
            const stakeAmount = ethers.parseEther("1");
            await metaNodeStakeContract.connect(user2).depositETH({
                value: stakeAmount
            });

            await expect(metaNodeStakeContract.connect(user2).unstake(0, stakeAmount))
                .to.emit(metaNodeStakeContract, "RequestUnstake")
                .withArgs(user2.address, 0, stakeAmount);
        });
    });

    // 奖励功能测试
    describe("奖励功能测试", function () {
        it("应该能正确计算和领取奖励", async function () {
            await advanceToStartBlock();
            await ensurePoolAdded();
            
            // 推进到startBlock+1
            const currentBlock = await getLatestBlockNumber();
            if (currentBlock < startBlock + 1) {
                await time.advanceBlockTo(startBlock + 1);
            }
            
            // 质押并获取奖励
            const stakeAmount = ethers.parseEther("100"); // 增加质押金额
            await metaNodeStakeContract.connect(user1).depositETH({
                value: stakeAmount
            });

            // 推进100个区块
            for (let i = 0; i < 100; i++) {
                await time.advanceBlock();
            }

            // 先更新池子状态再获取pending值，确保准确性
            await metaNodeStakeContract.updatePool(0);
            const pending = await metaNodeStakeContract.pendingMetaNode(0, user1.address);
            expect(pending).to.be.gt(0);

            // 领取奖励
            const balanceBefore = await metaNodeContract.balanceOf(user1.address);
            const contractMetaNodeBalanceBefore = await metaNodeContract.balanceOf(metaNodeStakeAddress);
            
            await metaNodeStakeContract.connect(user1).claim(0);
            
            const balanceAfter = await metaNodeContract.balanceOf(user1.address);
            const contractMetaNodeBalanceAfter = await metaNodeContract.balanceOf(metaNodeStakeAddress);

            // 添加详细日志输出
            console.log("质押金额:", stakeAmount.toString());
            console.log("待领取奖励:", pending.toString());
            console.log("MetaNodePerBlock:", metaNodePerBlock.toString());
            console.log("用户奖励前余额:", balanceBefore.toString());
            console.log("用户奖励后余额:", balanceAfter.toString());
            console.log("合约MetaNode余额变化:", (contractMetaNodeBalanceBefore - contractMetaNodeBalanceAfter).toString());
            console.log("余额实际变化:", (balanceAfter - balanceBefore).toString());

            // 验证用户确实收到了奖励
            expect(balanceAfter - balanceBefore).to.be.gt(0);
            
            // 验证收到的奖励数量是否合理（基于质押量和推进的区块数）
            // 我们质押了100 ETH并推进了100个区块，每区块奖励1个MetaNode
            const minimumExpected = ethers.parseEther("50");   // 至少应收到50个奖励
            const maximumExpected = ethers.parseEther("300");  // 最多应收到300个奖励
            
            expect(balanceAfter - balanceBefore).to.be.gte(minimumExpected);
            expect(balanceAfter - balanceBefore).to.be.lte(maximumExpected);
        });
    });

    // 池子参数更新测试
    describe("池子参数更新测试", function () {
        it("应该能更新池子权重", async function () {
            await advanceToStartBlock();
            await ensurePoolAdded();
            
            const currentBlock = await getLatestBlockNumber();
            if (currentBlock < startBlock + 1) {
                await time.advanceBlockTo(startBlock + 1);
            }
            
            const newWeight = 300;
            await expect(metaNodeStakeContract.setPoolWeight(0, newWeight, true))
                .to.emit(metaNodeStakeContract, "SetPoolWeight")
                .withArgs(0, newWeight, newWeight);

            const pool = await metaNodeStakeContract.pool(0);
            expect(pool.poolWeight).to.equal(newWeight);
        });
    });

    // 查询功能测试
    describe("查询功能测试", function () {
        it("应该能正确返回池子数量", async function () {
            await ensurePoolAdded();
            const poolLength = await metaNodeStakeContract.poolLength();
            expect(poolLength).to.be.gte(1);
        });

        it("应该能正确计算奖励乘数", async function () {
            await ensurePoolAdded();
            
            // 获取当前区块号
            const currentBlock = await getLatestBlockNumber();
            
            // 使用固定的区块范围进行测试，避免测试过程中区块变化
            const fromBlock = currentBlock;
            const toBlock = currentBlock + 10;
            
            // 获取合约中的参数用于调试
            const contractMetaNodePerBlock = await metaNodeStakeContract.MetaNodePerBlock();
            
            // 获取实际的乘数
            const multiplier = await metaNodeStakeContract.getMultiplier(BigInt(fromBlock), BigInt(toBlock));
            
            // 根据合约实现计算预期乘数: (to - from) * MetaNodePerBlock
            const blockDiff = toBlock - fromBlock;
            const expectedMultiplier = BigInt(blockDiff) * BigInt(contractMetaNodePerBlock);
            
            // 添加日志输出帮助调试
            console.log("fromBlock:", fromBlock);
            console.log("toBlock:", toBlock);
            console.log("blockDiff:", blockDiff);
            console.log("contractMetaNodePerBlock:", contractMetaNodePerBlock.toString());
            console.log("multiplier:", multiplier.toString());
            console.log("expectedMultiplier:", expectedMultiplier.toString());
            
            expect(multiplier).to.equal(expectedMultiplier);
        });

        it("应该能正确查询用户质押余额", async function () {
            await ensurePoolAdded();
            // 使用一个新的用户地址以避免之前测试的影响
            const stakeAmount = ethers.parseEther("1");
            await metaNodeStakeContract.connect(user2).depositETH({
                value: stakeAmount
            });

            const stakingBalance = await metaNodeStakeContract.stakingBalance(0, user2.address);
            expect(stakingBalance).to.equal(stakeAmount);
        });
    });

    // 提现功能测试
    describe("提现功能测试", function () {
        it("应该能提现解质押的ETH", async function () {
            await ensurePoolAdded();
            
            // 质押并解质押
            const stakeAmount = ethers.parseEther("1");
            await metaNodeStakeContract.connect(user2).depositETH({
                value: stakeAmount
            });

            await metaNodeStakeContract.connect(user2).unstake(0, stakeAmount);

            // 推进足够的区块以解锁资金
            for (let i = 0; i < 20; i++) {
                await time.advanceBlock();
            }

            // 提现
            await expect(metaNodeStakeContract.connect(user2).withdraw(0))
                .to.emit(metaNodeStakeContract, "Withdraw");
        });
    });
});  // 补全外层describe的闭合括号
});