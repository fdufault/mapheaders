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

  if len(config.FromHeader) == 0 {
    fmt.Println("FromHeader is not defined!")
    return nil, fmt.Errorf("FromHeader is not defined")
  }

  if len(config.ToHeader) == 0 {
    fmt.Println("ToHeader is not defined!")
    return nil, fmt.Errorf("ToHeader is not defined")
  }

  useMappings := true
  splitMappings := []Mapping{}
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
    fromHeaderValue := req.Header.Get(config.FromHeader)
    if len(fromHeaderValue) > 0 {
      if useMappings {
        for _, mapping := range splitMappings {
          if strings.Contains(fromHeaderValue, mapping.From) {
            req.Header.Set(config.ToHeader, mapping.To)
            break
          }
        }
      } else {
        req.Header.Set(config.ToHeader, fromHeaderValue)
      }
    } else {
      requestPath := req.Header.Get("RequestPath")
      fmt.Printf("Header '%s' has no value for path: %s\n", config.FromHeader, requestPath)
    }
    next.ServeHTTP(rw, req)
  }), nil
}
