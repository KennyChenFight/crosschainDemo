pragma solidity ^0.5.0;

contract Clinic {

    event InfoSet(bytes32 key, bytes32 value);

    mapping (bytes32 => bytes32) public infos;

    function setInfo(bytes32 key, bytes32 value) external {
        infos[key] = value;
        emit InfoSet(key, value);
    }
}