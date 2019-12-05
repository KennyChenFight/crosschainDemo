pragma solidity ^0.5.0;

contract Coordinate {

    event TransactionSet(string key, string value);

    mapping (string => string) public transactions;


    function addTransaction(string calldata key, string calldata value) external {
        transactions[key] = value;
        emit TransactionSet(key, value);
    }
}