// SPDX-License-Identifier: MIT
pragma solidity ^ 0.8;

//整数转罗马数字
contract IntToRoman {
    //     七个不同的符号代表罗马数字，其值如下：

    // 符号	值
    // I	1
    // V	5
    // X	10
    // L	50
    // C	100
    // D	500
    // M	1000

    function intToRoman(uint  num) public pure returns (string memory) {
         uint256[] memory values = new uint256[](13);
        values[0] = 1000;
        values[1] = 900;
        values[2] = 500;
        values[3] = 400;
        values[4] = 100;
        values[5] = 90;
        values[6] = 50;
        values[7] = 40;
        values[8] = 10;
        values[9] = 9;
        values[10] = 5;
        values[11] = 4;
        values[12] = 1;
        
        string[] memory symbols = new string[](13);
        symbols[0] = "M";
        symbols[1] = "CM";
        symbols[2] = "D";
        symbols[3] = "CD";
        symbols[4] = "C";
        symbols[5] = "XC";
        symbols[6] = "L";
        symbols[7] = "XL";
        symbols[8] = "X";
        symbols[9] = "IX";
        symbols[10] = "V";
        symbols[11] = "IV";
        symbols[12] = "I";

        string memory result="";
        uint256 i=0;

        while(num>0 && i<values.length){
            if(num>=values[i]){
                num -= values[i];
                result=concat(result,symbols[i]);
            }else{
                i++;
            }
        }
        return result;
    }

    function concat(string memory a,string memory b) internal pure returns (string memory) {
        bytes memory aBytes = bytes(a);
        bytes memory bBytes = bytes(b);

        bytes memory resultBytes = new bytes(aBytes.length+bBytes.length);
        for(uint i=0;i< aBytes.length;i++){
            resultBytes[i] = aBytes[i];
        }
        for(uint i=0;i<bBytes.length;i++){
            resultBytes[aBytes.length + i] = bBytes[i];
        }
        return string(resultBytes);
    }

}