// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

/**
 * @title memeShib
 * @dev SHIB风格的Meme代币合约，具有交易税、流动性池集成和交易限制功能
 * 合约实现了以下核心功能：
 * 1. 代币税功能：对每笔交易征收一定比例的税费
 * 2. 流动性池集成：与Uniswap V2集成，支持添加和移除流动性
 * 3. 交易限制功能：设置单笔交易最大额度和每日交易次数限制
 */
contract memeShib is ERC20, Ownable {

    // 交易税率，设置为0.5%，对每笔交易征收税费
    uint256 public taxRate = 0.5; 
    
    // 税费接收地址，所有收取的税费将发送到此地址
    address public taxAddress;

    // 单笔交易最大额度限制，防止大额交易操纵市场
    uint256 public maxTxAmount = 1000 * 10 ** 9; 
    
    // 记录每个地址的每日交易次数，用于实现交易次数限制
    mapping(address => uint8) public dailyTxCount; 
    
    // 每日交易次数限制，每个地址每天最多进行5次交易
    uint8 public maxDailyTxCount = 5; 

    // Uniswap V2 路由器接口实例，用于与Uniswap交互
    IUniswapV2Router02 public uniswapV2Router; 

    // 代币销毁比例，设置为0.01%，每笔交易都会销毁一定数量的代币
    uint256 public burnRate = 0.01; 

    // 代币转账事件，记录转账的发送方、接收方和转账金额
    event Transfer(address indexed from, address indexed to, uint256 value);
    
    // 税费支付事件，记录税费的支付方、接收方和税费金额
    event TaxPaid(address indexed from, address indexed to, uint256 value);

    /**
     * @dev 合约构造函数，初始化代币名称、符号以及其他参数
     * 在构造函数中完成以下初始化操作：
     * 1. 调用父合约ERC20的构造函数设置代币名称和符号
     * 2. 铸造初始代币供应量并发送给合约创建者
     * 3. 设置税费接收地址为合约创建者
     * 4. 初始化Uniswap V2路由器接口
     */
    constructor() ERC20("Meme Shib", "MSHIB") {
        // 铸造100万枚代币并发送给合约创建者，作为初始供应量
        _mint(msg.sender, 1000000 * 10 ** decimals()); 
        
        // 设置税费地址为合约创建者
        taxAddress = msg.sender; 
        
        // 初始化Uniswap V2路由器地址
        uniswapV2Router = IUniswapV2Router02(0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D); 
    }

    /**
     * @dev 授权转账函数，允许第三方代表代币持有者转账
     * @param from 代币发送方地址
     * @param to 代币接收方地址
     * @param value 转账金额
     * @return bool 转账是否成功
     * 
     * 该函数实现了以下功能：
     * 1. 检查发送方余额是否充足
     * 2. 检查授权额度是否足够
     * 3. 执行交易限制检查
     * 4. 处理代币转账（含税费和销毁）
     * 5. 更新授权额度
     */
    function transferFrom(address from, address to, uint256 value) public returns (bool) {
        // 检查发送方余额是否充足
        require(balanceOf(from) >= value, "Insufficient balance");
        
        // 检查调用者的授权额度是否足够
        require(allowance(from, msg.sender) >= value, "Allowance exceeded");
        
        // 交易限制功能：检查转账金额是否超过单笔交易最大额度
        require(value <= maxTxAmount, "Exceeds max transaction amount");
        
        // 交易限制功能：检查发送方当日交易次数是否超过限制
        require(dailyTxCount[from] < maxDailyTxCount, "Exceeds daily transaction limit");

        // 调用内部转账函数处理税费和实际转账
        _transfer(from, to, value);
        
        // 触发转账事件
        emit Transfer(from, to, value);

        // 计算并销毁一定数量的代币（根据burnRate比例）
        uint256 burnAmount = (value * burnRate) / 100;
        _burn(from, burnAmount);

        // 增加发送方的当日交易次数计数
        dailyTxCount[from]++;

        // 更新授权额度，减少已使用的额度
        _approve(from, msg.sender, allowance(from, msg.sender) - value);
        return true;
    }

    /**
     * @dev 代币转账函数，允许用户直接转账代币
     * @param to 代币接收方地址
     * @param value 转账金额
     * @return bool 转账是否成功
     * 
     * 该函数实现了以下功能：
     * 1. 检查发送方余额是否充足
     * 2. 执行交易限制检查
     * 3. 处理代币转账（含税费和销毁）
     */
    function transfer(address to, uint256 value) public returns (bool) {
        // 检查发送方余额是否充足
        require(balanceOf(msg.sender) >= value, "Insufficient balance");
        
        // 交易限制功能：检查转账金额是否超过单笔交易最大额度
        require(value <= maxTxAmount, "Exceeds max transaction amount");
        
        // 交易限制功能：检查发送方当日交易次数是否超过限制
        require(dailyTxCount[msg.sender] < maxDailyTxCount, "Exceeds daily transaction limit");
        
        // 调用内部转账函数处理税费和实际转账
        _transfer(msg.sender, to, value);
        
        // 触发转账事件
        emit Transfer(msg.sender, to, value);

        // 计算并销毁一定数量的代币（根据burnRate比例）
        uint256 burnAmount = (value * burnRate) / 100;
        _burn(msg.sender, burnAmount);

        // 增加发送方的当日交易次数计数
        dailyTxCount[msg.sender]++;

        return true;
    }

    /**
     * @dev 内部转账函数，处理实际的代币转账逻辑
     * @param from 代币发送方地址
     * @param to 代币接收方地址
     * @param amount 转账金额
     * 
     * 该函数实现了以下功能：
     * 1. 计算交易税费
     * 2. 将扣除税费后的金额转账给接收方
     * 3. 将税费转账给税费地址
     * 4. 触发税费支付事件
     */
    function _transfer(address from, address to, uint256 amount) internal override {
        // 计算交易税费（根据taxRate比例）
        uint256 tax = (amount * taxRate) / 100;
        
        // 计算扣除税费后的实际转账金额
        uint256 amountAfterTax = amount - tax;

        // 调用父合约的转账函数，将扣除税费后的金额转账给接收方
        super._transfer(from, to, amountAfterTax);
        
        // 调用父合约的转账函数，将税费转账给税费地址
        super._transfer(from, taxAddress, tax);
        
        // 触发税费支付事件
        emit TaxPaid(from, taxAddress, tax);
    }

    /**
     * @dev 重置指定账户的每日交易次数
     * @param account 需要重置交易次数的账户地址
     * 
     * 该函数只能由合约所有者调用，用于在特殊情况下重置某个账户的交易次数
     */
    function resetDailyTxCount(address account) external onlyOwner {
        // 将指定账户的每日交易次数重置为0
        dailyTxCount[account] = 0;
    }

    /**
     * @dev 添加流动性到Uniswap池
     * @param tokenAmount 添加的代币数量
     * @param ethAmount 添加的ETH数量
     * 
     * 该函数只能由合约所有者调用，用于向Uniswap池添加流动性
     * 添加流动性可以增加代币的交易深度和流动性
     */
    function addLiquidity(uint256 tokenAmount, uint256 ethAmount) external onlyOwner {
        // 授权Uniswap路由器可以使用指定数量的代币
        _approve(address(this), address(uniswapV2Router), tokenAmount);
        
        // 调用Uniswap路由器的添加流动性函数
        uniswapV2Router.addLiquidityETH{value: ethAmount}(
            address(this),     // 代币地址
            tokenAmount,       // 代币数量
            0,                 // 最小代币数量（滑点保护）
            0,                 // 最小ETH数量（滑点保护）
            owner(),           // 流动性代币接收方
            block.timestamp    // 交易截止时间
        );
    }

    /**
     * @dev 从Uniswap池移除流动性
     * @param liquidity 要移除的流动性数量
     * 
     * 该函数只能由合约所有者调用，用于从Uniswap池移除流动性
     * 移除流动性可以获得相应的代币和ETH
     */
    function removeLiquidity(uint256 liquidity) external onlyOwner {
        // 调用Uniswap路由器的移除流动性函数
        uniswapV2Router.removeLiquidityETH(
            address(this),     // 代币地址
            liquidity,         // 流动性数量
            0,                 // 最小代币数量（滑点保护）
            0,                 // 最小ETH数量（滑点保护）
            owner(),           // 代币和ETH接收方
            block.timestamp    // 交易截止时间
        );
    }
}