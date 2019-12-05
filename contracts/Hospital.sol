pragma solidity ^0.5.0;

contract Hospital {

    event InfoSet(string key, string value);

    mapping (string => string) public infos;

    function setInfo(string calldata key, string calldata value) external {
        infos[key] = value;
        emit InfoSet(key, value);
    }

    function verifyMessage(string memory message) public pure returns(bool) {
        return (keccak256(abi.encodePacked((message))) == keccak256(abi.encodePacked(("clinic"))) );
    }
}