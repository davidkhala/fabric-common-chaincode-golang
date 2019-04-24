# fabric-common-chaincode-golang

## Before git clone...
Wait, you should use `go get github.com/davidkhala/fabric-common-chaincode-golang` instead. 
And then `dep ensure`

## Notes

- full document of couchdb query syntax ```http://docs.couchdb.org/en/stable/api/database/find.html?highlight=find.html#post--db-_find```
- ParseCreator is deprecated now, please use ClientIdentityLibrary
- https://hlf.readthedocs.io/en/latest/endorsement-policies.html?highlight=endorse#validation

  "the key-level endorsement policy overrides the chaincode-level endorsement policy." not just a new layer of restriction.
  If a key’s endorsement policy is removed (set to nil), the chaincode-level endorsement policy becomes the default again.
- https://jira.hyperledger.org/browse/FAB-5094 GetHistoryForPrivateKey
  
      The workaround simply make a shadow copy of privateData in public scope. 
      And how do we implement that copy depends on requirements.
- best practice for errors in golang chaincode: https://hyperledger-fabric.readthedocs.io/en/release-1.4/error-handling.html
## TODO

- Yacov M introduce about
    chaincodeStub.GetDecorations: As for Decorations, a peer may add additional input to the chaincode input via custom endorsement handlers.
    You need to specify a plugin file in the core.yaml section that implements a decorator.

    chaincodeStub.GetBinding: it's just hash over nonce || creator || epoch
- when iterate >100 queryResult:  QUERY_STATE_NEXT failed: transaction ID: c84cb2a85b169e4f515c304199c38b22c8c818d55ed3cbf2c4cb91dfb67b0250: query iterator not found

    
