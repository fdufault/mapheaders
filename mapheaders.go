package mapheaders

import (
  "context"
  "fmt"
  "net/http"
  "strings"
)

// Config the plugin configuration.
type Config struct {
  FromHeader string   `json:"FromHeader,omitempty"` // source header to check for value
  ToHeader   string   `json:"ToHeader,omitempty"`   // destination header to set
  Mappings   []string `json:"Mappings,omitempty"`   // values to map
}

// Map values
type Mapping struct {
  From string
  To   string
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
  return &Config{}
}

// New creates and returns a plugin instance.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

  var fromHeader string
  if len(config.FromHeader) == 0 {
    fmt.Println("FromHeader is not defined!")
    return nil, fmt.Errorf("FromHeader is not defined")
  } else {
    fromHeader = config.FromHeader
  }

  var toHeader string
  if len(config.ToHeader) == 0 {
    fmt.Println("ToHeader is not defined!")
    return nil, fmt.Errorf("ToHeader is not defined")
  } else {
    toHeader = config.ToHeader
  }

  splitMappings := []Mapping{}
  useMappings := true
  if len(config.Mappings) == 0 {
    useMappings = false
  } else {
    mappings := config.Mappings
    for _, mapping := range mappings {
      m := strings.Split(mapping, "=>")
      from := strings.TrimSpace(m[0])
      to := from
      if len(m) == 2 {
        to = strings.TrimSpace(m[1])
      }
      newMapping := Mapping{
        From: from,
        To: to,
      }
      splitMappings = append(splitMappings, newMapping)
    }
  }

  return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
    headers := req.Header
    fromHeaderValues, found := headers[fromHeader]
    if found {
      fromHeaderValue := strings.Join(fromHeaderValues, ",")
      if useMappings {
        for _, mapping := range splitMappings {
          if strings.Contains(fromHeaderValue, mapping.From) {
            req.Header.Set(toHeader, mapping.To)
            break
          }
        }
      } else {
        req.Header.Set(toHeader, fromHeaderValue)
      }
    } else {
      requestPaths, found := headers["RequestPath"]
      if found {
        requestPath := strings.Join(requestPaths, ",")
        fmt.Printf("Header '%s' has no value for path: %s\n", fromHeader, requestPath)
      }
    }
    next.ServeHTTP(rw, req)
  }), nil
}
