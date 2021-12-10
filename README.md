# Map Headers for Traefik

Allows the mapping of values from one header to another:
  - Full value
  - Partial value
  - Replacement value

## Configuration:

### Static

experimental:
  plugins:
    mapheaders:
      moduleName: "github.com/fdufault/mapheaders"
      version: "v0.0.1"

### Dynamic

http:
  middlewares:
    mapSomeHeader:                                                                                                          
       plugin:                                                                                                               
          mapheaders:                                                                                                         
            FromHeader: "Some-Header"                                                                                       
            ToHeader: "Some-Other-Header"                                                                                           
            Mappings:                                                                                                     
              - "admins"                                                                                                    
              - "users=>project_user" 
              
 Mappings
 
 1) No mappings: the full value of the `From` header is copied to the `To` header
 2) Simple mapping: if the mapping value is found in the `From` header value, it will be set as the `To` header value.
 3) Remapping: if the value before `=>` is found in the `From` header value, the value after the `=>` will be set as the `To` header value.
