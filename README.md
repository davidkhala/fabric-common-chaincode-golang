# fabric-common-chaincode-golang

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
- when iterate >100 times from queryResult iterator, it shows as:

  ```
  QUERY_STATE_NEXT failed: transaction ID: c84cb2a85b169e4f515c304199c38b22c8c818d55ed3cbf2c4cb91dfb67b0250: query iterator not found
  ```
    - This apply to [StateIterator][HistoryIterator] in golang chaincode only
    - [totalQueryLimit](https://github.com/hyperledger/fabric/blob/release-1.4/sampleconfig/core.yaml) is not the cause
- `failed to invoke chaincode name:"lscc" , error: API error (400): OCI runtime create failed: container_linux.go:348: starting container process caused "exec: \"chaincode\": executable file not found in $PATH": unknown`
    - means package name for golang-chaincode entrance is not `main`
- Yacov M introduce about
    chaincodeStub.GetDecorations: As for Decorations, a peer may add additional input to the chaincode input via custom endorsement handlers.
    You need to specify a plugin file in the core.yaml section that implements a decorator.

    chaincodeStub.GetBinding: it's just hash over nonce || creator || epoch    
## TODO

    
