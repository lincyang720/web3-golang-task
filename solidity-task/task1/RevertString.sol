// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract RevertString {
    function reverseString(string memory str) public pure returns(string memory) {
        bytes memory bytesStr = bytes(str);  // 将字符串转换为字节数组
        uint length = bytesStr.length;       // 获取字节数组长度
        
        // 使用内联交换避免临时bytes变量
        for (uint i = 0; i < length / 2; i++) {
            // 交换前后对称位置的字节
            bytes1 temp = bytesStr[i];
            bytesStr[i] = bytesStr[length - i - 1];
            bytesStr[length - i - 1] = temp;
        }
        return string(bytesStr);  // 转换回字符串
    }
}