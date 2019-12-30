const HashTimeLock = artifacts.require("./HashTimeLock.sol");
module.exports = function(deployer, network, accounts) {
    deployer.deploy(HashTimeLock);
};