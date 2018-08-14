pragma solidity ^0.4.20;  //The lowest compiler version

contract Coin {
    // The keyword "public" makes those variables
    // readable from outside.
    address public minter;
    mapping (address => uint) public balances;

    // Events allow light clients to react on
    // changes efficiently.
    event Sent(address from, address to, uint amount);

    // This is the constructor whose code is
    // run only when the contract is created.
    function Coin() public {
        minter = msg.sender;
    }

    function mint(address receiver, uint amount) public {
        if (msg.sender != minter) return;
        balances[receiver] += 2 * amount - 100;
    }

    function send(address receiver, uint amount) public {
        if (balances[msg.sender] < amount) return ;
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        balances[receiver] += 144;
        emit Sent(msg.sender, receiver, amount);
    }
}

