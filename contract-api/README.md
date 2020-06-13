# fabric-common-chaincode-golang: contract API


## High light
- it bundles one or more "contracts" into a running chaincode. 
- All contracts for use in chaincode must implement the contractapi.ContractInterface. 
    - The easiest way to do this is to embed the contractapi.Contract struct within your own contract which will provide default functionality for meeting this interface.
- By default all public functions of a struct are assumed to be callable via the final chaincode. They must match a set of rules
- About function parameters
    - First parameter as type `contractapi.TransactionContextInterface`
    - For the rest only following parameter types are accepted: 
    
      ```
      string
      bool
      int (including int8, int16, int32 and int64)
      uint (including uint8, uint16, uint32 and uint64)
      float32
      float64
      time.Time
      Arrays/slices of any allowable type
      Structs (whose public fields are all of the allowable types or another struct)
      Pointers to structs
      Maps with a key of type string and values of any of the allowable types
      interface{} (Only allowed when directly taken in, will receive a string type when called via a transaction)  
      ```
- About function return values
    - If the function is defined to return zero values then a success response will be returned for all calls to that contract function
    - If the function is defined to return one value then that value may be any of the allowable types listed for parameters (except interface{}) or error.
    - If the function is defined to return two values then the first may be any of the allowable types listed for parameters (except interface{}) and the second must be error
    - [**FASHION WE ARE USING**] return one value and panic all error    
       