// SPDX-License-Identifier: MIT
pragma solidity ^ 0.8;

contract RomanToInt {
    mapping(bytes => uint) public  romanValues;
    
     // 辅助函数，获取罗马字符对应的值
    function getRomanValue(bytes1 romanChar) internal pure returns (uint) {
        if (romanChar == 'I') return 1;
        if (romanChar == 'V') return 5;
        if (romanChar == 'X') return 10;
        if (romanChar == 'L') return 50;
        if (romanChar == 'C') return 100;
        if (romanChar == 'D') return 500;
        if (romanChar == 'M') return 1000;
        return 0;
    }

    

    //罗马数字转整数
    function romanToInt(string memory s) public pure returns (uint) {
        //罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
            // 字符          数值
            // I             1
            // V             5
            // X             10
            // L             50
            // C             100
            // D             500
            // M             1000
        // 初始化结果
        uint result = 0;
        uint prevValue = 0;

        // 遍历字符串
        for (uint i = 0; i < bytes(s).length; i++) {
            bytes1  currentChar = bytes(s)[i];
            uint currentValue = getRomanValue(currentChar);

            // 如果当前值大于前一个值，减去两倍的前一个值
            if (currentValue > prevValue) {
                result += currentValue - 2 * prevValue;
            } else {
                result += currentValue;
            }

            prevValue = currentValue;
        }

        return result;

    }
}