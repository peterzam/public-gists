pragma solidity ^0.8.17;

contract ERC20 {
    string public constant name = "";
    string public constant symbol = "";
    uint8 public constant decimals = 0;

     mapping(address => uint256) private _balances;

    function balanceOf(address account) public view virtual returns (uint256) {
        return _balances[account];
    }
}
