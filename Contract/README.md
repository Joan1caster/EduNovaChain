### Build
```shell
$ forge build
```

### deploy
anvil

forge script script/nft.s.sol --broadcast --rpc-url http://127.0.0.1:8545 --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d --interactive


### interact

```create
cast send 0x71C95911E9a5D330f4D621842EC243EE1343292e "createInnovation(string,string,uint256,bool)" metadata QmSnuWmxptJZdLJpKRarxBMS2Ju2oANVrgbr2xWbie9b2D 0.77ether true --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
```

```buy
cast send 0x71C95911E9a5D330f4D621842EC243EE1343292e "purchaseInnovation(uint256)" 5 --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d --value 0.77ether
```

### Gas Snapshots

```shell
$ forge snapshot
```
## help
```shell
$ forge --help
$ anvil --help
$ cast --help
```
