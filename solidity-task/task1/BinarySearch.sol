// SPDX-License-Identifier: MIT
pragma solidity  ^0.8;

contract BinarySearch {
    // 二分查找,
    function binarySearch(int[] memory arr, int target) public pure returns (int) {
        uint left = 0;
        uint right = arr.length - 1;

        while (left <= right) {
            //有溢出的风险，重构
            uint mid = left + (right - left) / 2;

            if (arr[mid] == target) {
                return int(mid); // 返回找到的索引
            } else if (arr[mid] < target) {
                left = mid + 1; // 在右半部分继续查找
            } else {
                right = mid - 1; // 在左半部分继续查找
            }
        }

        return -1; // 如果未找到，返回 -1
    }
}
