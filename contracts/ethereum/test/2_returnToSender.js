const sha256 = require("sha256")
const converter = require("./helper/converter.js")
const HashTimeLock = artifacts.require("HashTimeLock");

contract("Returns HTLC tests", async accounts => {
    let htlc = null;
    before(async () => {
        htlc = await HashTimeLock.deployed();
    });
    const secret = [ 0x34, 0x31, 0x32, 0x33]
    const hexSecret = "0x" + converter.toHexString(secret)
    const hash = "0x" + sha256(secret)
    
    it("Lock", async () => {
        const amount = "0.01"
        const expireHeight = (await web3.eth.getBlockNumber()) - 10
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

    it("Unlock execption", async () => {
        try {
            var result = await htlc.unlock(requestHash, hexSecret, { from: accounts[1], value: 0 })
            console.error(result)
        } catch(ex) {
            return
        }

        throw("unlock success")
    });

    it("Return to sender", async () => {
        const result = await htlc.returnToSender(requestHash, { from: accounts[0], value: 0 })
        const height = result.receipt.blockNumber
        const swapRequest = await htlc.swapRequests.call(requestHash);
        assert.equal(swapRequest.amount, 0);
        assert.equal(height > swapRequest.expireHeight, true);
    });
});