// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.2;

contract Seeds{
    uint public count;
    uint public capacity;

    mapping (uint => address) public seeds;
    address public seedSetter;

    constructor(address setter) {
        seedSetter = setter;
        count = 0;
        capacity = 16;
    }

    function deleteSeed(uint index) external {
        require(msg.sender == seedSetter, 'Seed Delete: FORBIDDEN');
        require(index<count, 'Require: index<count');
        delete seeds[index];
    }

    function addSeed(address seed) external{
        require(msg.sender == seedSetter, 'Seed Add: FORBIDDEN');
        require(count<capacity, 'Capacity have achive warning line.');
        seeds[count]=seed;
        count=count+1;
    }

    function updateSeed(uint index, address seed) external{
        require(msg.sender == seedSetter, 'Seed Update: FORBIDDEN');
        require(index<count, 'Index must be smaller than really capacity(count value)');
        seeds[index] = seed;
    }

    function updateCapacity(uint cap) external{
        require(msg.sender == seedSetter, 'Seed Update: FORBIDDEN');
        capacity = cap;
    }

    function updateSeedSetter(address setter) external{
        require(msg.sender == seedSetter,'Owner Update: FORBIDDEN');
        seedSetter = setter;
    }
}