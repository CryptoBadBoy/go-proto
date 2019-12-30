pragma solidity >=0.4.22 <0.6.0;
contract HashTimeLock {
    event NewSwap(bytes32 requestHash, uint expireHeight, address recipient, address sender, bytes32 secretHash);
    event SuccessSwap(bytes32 requestHash, bytes secret);

    struct SwapRequest {
        uint256 amount;
        uint expireHeight;
        bytes secret;
        bytes32 secretHash;
        address sender;
        address recipient;
    }
    // [recipient][sender][hash]: amount
    mapping(bytes32 => SwapRequest) public swapRequests;
    function returnToSender(bytes32 requestHash) public {
        require(swapRequests[requestHash].amount > 0, "swap is over");
        require(block.number > swapRequests[requestHash].expireHeight, "timeout is not over yet");

        msg.sender.transfer(swapRequests[requestHash].amount);
        swapRequests[requestHash].amount = 0;
    }
    function unlock(bytes32 requestHash, bytes memory secret) public {
        require(swapRequests[requestHash].amount > 0, "swap is over");
        require(block.number <= swapRequests[requestHash].expireHeight, "timeout is not over yet");
        
        msg.sender.transfer(swapRequests[requestHash].amount);
        swapRequests[requestHash].secret = secret;
        swapRequests[requestHash].amount = 0;
        emit SuccessSwap(requestHash, secret);
    }
    function lock(address recipient, bytes32 secretHash, uint expireHeight) public payable {
        bytes32 requestHash = keccak256(abi.encodePacked(this, msg.sender, recipient, secretHash));
        require(msg.value > 0, "value less or equals 0");
        require(expireHeight > 0, "expire height less or equals 0");

        require(swapRequests[requestHash].expireHeight == 0, "swap exist");

        bytes memory empty;
        swapRequests[requestHash] = SwapRequest(
            {expireHeight: expireHeight, amount: msg.value, secret: empty, secretHash: secretHash, sender: msg.sender, recipient: recipient });
        emit NewSwap(requestHash, expireHeight, recipient, msg.sender, secretHash);
    }
}