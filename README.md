# fabric-common-chaincode-golang

## Notes

- full document of couchdb query syntax ```http://docs.couchdb.org/en/stable/api/database/find.html?highlight=find.html#post--db-_find```
- https://hlf.readthedocs.io/en/latest/endorsement-policies.html?highlight=endorse#validation

  "the key-level endorsement policy overrides the chaincode-level endorsement policy." not just a new layer of restriction.
  If a key’s endorsement policy is removed (set to nil), the chaincode-level endorsement policy becomes the default again.
- https://jira.hyperledger.org/browse/FAB-5094 GetHistoryForPrivateKey
  
      The workaround simply make a shadow copy of privateData in public scope. 
      And how do we implement that copy depends on requirements.
- best practice for errors in golang chaincode: https://hyperledger-fabric.readthedocs.io/en/release-1.4/error-handling.html
- [2.0] dependency on fabric-core is now split to 
    ```
      "github.com/hyperledger/fabric-chaincode-go"
      "github.com/hyperledger/fabric-protos-go"
    ```
- [2.0] align `shim.NewLogger` removal [FAB-15366](https://jira.hyperledger.org/browse/FAB-15366)
- `github.com/hyperledger/fabric/core/chaincode/shim/ext/entities` is removed in [FAB-16213](https://jira.hyperledger.org/browse/FAB-16213)
- `lscc.go` still depends on fabric core, but it is deprecated and giving way to `_lifecycle`
    ```http request
    github.com/hyperledger/fabric/core/common/ccprovider
    ```
 

    
