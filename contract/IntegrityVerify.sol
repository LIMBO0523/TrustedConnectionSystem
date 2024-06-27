// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntegrityVerify {
    // 定义一个mapping，key是字符串，value是字节数组（hash值通常表示为字节数组）
    mapping(string => bytes32) public fileToHashMap;
    // 定义一个mapping，key是address，value是公钥（这里假设公钥为字节数组）
    mapping(address => bytes) public addressToPublicKeyMap;
    // 定义一个mapping，key是address，value是bool，验证address的设备是否完整
    mapping(address => bool) public addressToBoolMap;

    // 定义一个事件，用于保存hash数组，交易发起者的公钥，以及它们的哈希值
    event HashArrayStored(bytes32[] sigArray, bytes senderPublicKey, bytes32[] hashArray, address sender);
    // 定义事件，用于通知监听的程序
    event measureFinished();


    // 设置字符串和对应的hash值
    function setHash(string memory key, bytes32 value) public {
        fileToHashMap[key] = value;
    }

    // 设置字符串和对应的hash值（多组一起设置）
    function setMultipleHashes(string[] memory keys, bytes32[] memory values) public {
        require(keys.length == values.length, "Number of keys must match number of values");

        for (uint i = 0; i < keys.length; i++) {
            fileToHashMap[keys[i]] = values[i];
        }
    }

    // 获取字符串对应的hash值
    function getHash(string memory key) public view returns (bytes32) {
        return fileToHashMap[key];
    }

    // 设置地址和对应的公钥
    function setAddressToPublicKey(address addr, bytes memory publicKey) public {
        addressToPublicKeyMap[addr] = publicKey;
    }

    // 获取地址对应的公钥
    function getAddressToPublicKey(address addr) public view returns (bytes memory) {
        return addressToPublicKeyMap[addr];
    }

    // 检查单个字符串和哈希值是否匹配
    function verifyHash(string memory key, bytes32 value) public view returns (bool) {
        bytes32 storedHash = fileToHashMap[key];
        return storedHash == value;
    }

    // 检查map中的string和byte是否一一对应，如果一一对应，则保存第二个字节数组和交易发起者的公钥
    function verifyAndStoreHashes(string[] memory keys, bytes32[] memory values, bytes32[] memory hashArrayToStore) public returns (bool) {
        require(keys.length == values.length, "Keys and values array length must match");


        delete addressToBoolMap[msg.sender];
        addressToBoolMap[msg.sender] = true;

        for (uint i = 0; i < keys.length; i++) {
            if (fileToHashMap[keys[i]] != values[i]) {
                revert("Key and value mismatch");
            }
        }

        bytes memory senderPublicKey = addressToPublicKeyMap[msg.sender];
        require(senderPublicKey.length > 0, "Sender public key not set");

        // 参数依次为 文件的签名值，公钥，文件的hash值, 调用者的地址
        emit HashArrayStored(hashArrayToStore, senderPublicKey, values, msg.sender);

        return true;
    }

    // 设置address和对应的bool值
    function setAddressToBool(address addr, bool value) public {
        addressToBoolMap[addr] = value;
        emit measureFinished();
    }

    // 检查address是否存在于mapping中且bool值是否为true
    function isAddressTrue(address addr) public view returns (bool) {

        return addressToBoolMap[addr];
    }
}