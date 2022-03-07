# Map Headers for Traefik

Allows the mapping of values from one header to another:
  - Full value
  - Partial value
  - Replacement value

## Configuration:

### Static
```
experimental:
  plugins:
    mapheaders:
      moduleName: "github.com/fdufault/mapheaders"
      version: "v0.0.7"
```
### Dynamic
```
http:
  middlewares:
    mapSomeHeader:                                                                                                          
       plugin:                                                                                                               
          mapheaders:                                                                                                         
            FromHeader: "Some-Header"                                                                                       
            ToHeader: "Some-Other-Header"                                                                                           
            Mappings:                                                                                                     
              - "value1"                                                                                                    
              - "value2=>valueB"
              - "default=>somedef
 ```             
 - FromHeader: The name of the header to use as the source.
 - ToHeader: The name of the header to set.
 - Mappings (optional):
 1) No mappings provided or no mappings match: the full value of the `FromHeader` header is copied to the `ToHeader` header
 2) Simple mapping: if the mapping value is found in the `FromHeader` header value, it will be set as the `ToHeader` header value.
 3) Remapping: if the value before `=>` is found in the `FromHeader` header value, the value after the `=>` will be set as the `ToHeader` header value. The mapping for `default` will be used if no other mapping matches.

### NOTE: the mappings are processed in order and once a mapping matches, processing stops at that mapping (including the default mapping).

Examples based on the configuration above:

- A request contains `Some-Header: value1`: `Some-Other-Header` will be set to `value1`
- A request contains `Some-Header: value2`: `Some-Other-Header` will be set to `valueB`
- A request contains `Some-Header: value1,value2`: `Some-Other-Header` will be set to `value1` (the first match stops the parsing)
- A request contains `Some-Header: somevalue`: the default value `somedef` will be used.
