const sha256 = require("sha256")
const converter = require("./helper/converter.js")
const HashTimeLock = artifacts.require("HashTimeLock");

contract("Unlock HTLC tests", async accounts => {
    let htlc = null;
    before(async () => {
        htlc = await HashTimeLock.deployed();
    });
    const secret = [ 0x29, 0x29 ]
    const hexSecret = "0x" + converter.toHexString(secret)
    const hash = "0x" + sha256(secret)
    let requestHash = []
    it("Lock", async () => {
        const amount = "0.01"
        const expireHeight = (await web3.eth.getBlockNumber()) + 10
        const result = await htlc.lock(accounts[1], hash, expireHeight, {from: accounts[0], value: web3.utils.toWei(amount, 'ether')})
        requestHash = result.logs[0].args.requestHash

        const swapRequest = await htlc.swapRequests.call(requestHash);
      
        assert.equal(swapRequest.amount, web3.utils.toWei(amount, 'ether'));
        assert.equal(swapRequest.recipient, accounts[1]);
        assert.equal(swapRequest.sender, accounts[0]);
        assert.equal(swapRequest.secretHash, hash);
        assert.equal(swapRequest.expireHeight, expireHeight);

        assert.equal(result.logs[0].event, "NewSwap")
        assert.equal(result.logs[0].args.recipient, accounts[1])
        assert.equal(result.logs[0].args.sender, accounts[0])
        assert.equal(result.logs[0].args.secretHash, hash)
        assert.equal(result.logs[0].args.expireHeight, expireHeight)
    });
    it("Return to sender (execption)", async () => {
        try {
            var result = await htlc.returnToSender(requestHash, { from: accounts[0], value: 0 })
            console.error(result)
        } catch(ex) {
            return
        }

        throw("return to sender success")
    });
    
    it("Unlock", async () => {
        const result = await htlc.unlock(requestHash, hexSecret, { from: accounts[1], value: 0 })
        const height = result.receipt.blockNumber
        const swapRequest = await htlc.swapRequests.call(requestHash);
        
        assert.equal(swapRequest.amount, 0);
        assert.equal(swapRequest.secret, hexSecret)
        assert.equal(height <= swapRequest.expireHeight, true);

        assert.equal(result.logs[0].event, "SuccessSwap")
        assert.equal(result.logs[0].args.requestHash, requestHash)
        assert.equal(result.logs[0].args.secret, hexSecret)
    });
});