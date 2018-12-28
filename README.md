# fabric-common-chaincode-golang

# Notes

- full document of couchdb query syntax ```http://docs.couchdb.org/en/stable/api/database/find.html?highlight=find.html#post--db-_find```
- ParseCreator is deprecated now, please use ClientIdentityLibrary
# TODO

- Yacov M introduce about
    chaincodeStub.GetDecorations: As for Decorations, a peer may add additional input to the chaincode input via custom endorsement handlers.
    You need to specify a plugin file in the core.yaml section that implements a decorator.

    chaincodeStub.GetBinding: it's just hash over nonce || creator || epoch

-