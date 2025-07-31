// SPDX-License-Identifier: MIT
pragma solidity ^ 0.8;

contract MergeSortedArr {
    //合并两个有序数组

    function mergeSortedArr(int[] memory nums1, int[] memory nums2) public pure returns(int[] memory) {
        int[] memory result = new int[](nums1.length + nums2.length);
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while (i < nums1.length && j < nums2.length) {
            if (nums1[i] <= nums2[j]) {
                result[k] = nums1[i];
                i++;
            } else {
                result[k] = nums2[j];
                j++;
            }
            k++;
        }
        // 处理剩余元素
        while (i < nums1.length) {
            result[k] = nums1[i];
            i++;
            k++;
        }
        while (j < nums2.length) {
            result[k] = nums2[j];
            j++;
            k++;
        }
        return result;
    }
        
}