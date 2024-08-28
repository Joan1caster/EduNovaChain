### Build
```shell
$ forge build
```

### deploy
anvil

forge script script/nft.s.sol --broadcast (可选)--rpc-url (可选)--private-key (可选)--interactive


### interact

```send
cast send <contract_address> "<funtcion_name>(<parameter_classes>)" <parameter> (可选)--rpc-url (可选)--private-key (可选)--interactive
```

```call
cast call <contract_address> "<funtcion_name>(<parameter_classes>)" <parameter> (可选)--rpc-url (可选)--private-key (可选)--interactive
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
