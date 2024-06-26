# ConfigLoader: Dynamic Configuration Management in Go

ConfigLoader is a flexible and powerful library for managing configurations in Go applications. It allows loading, merging, and dynamically overriding configurations from various sources with ease. This makes it an ideal choice for applications that require a high level of flexibility in their configuration management.

## Features

- Load configuration from files with support for JSON, YAML, TOML, and environment variables.
- Override configuration values programmatically to cater to different environments or runtime requirements.
- Support for deserializing nested structures and arrays/slices in configurations.
- Easy integration with existing Go projects with minimal setup.

## Getting Started

To use ConfigLoader, first ensure you have it installed:

```go
import "github.com/snippetaccumulator/configloader"
```

### Basic Usage

The following example demonstrates how to load a configuration file:

```go
package main

import (
    "github.com/snippetaccumulator/configloader"
    "log"
)

type AppConfig struct {
    Field1 string
    Field2 int
    Nested struct {
        Field3 bool
    }
}

func main() {
    var config AppConfig
    loader := configloader.NewConfigLoader("config.yaml", configloader.WithPath("/path/to/config/"), configloader.WithDeserializer(new(configloader.YAMLDeserializer)))
    if err := loader.Load(&config); err != nil {
        log.Fatalf("Failed to load configuration: %s", err)
    }
    log.Println("Configuration loaded successfully:", config)
}
```

### Paths

When using deserializers, the data must follow the path of the corresponding tags used by the deserializer. This is obvious for most structures, except for structs with embedded members. Here is an example of how to handle this:

```go
type GenericConfig struct {
    Field1 string
    Field2 int
}

type WebConfig struct {
    GenericConfig
    Port int
}
```

With the `YAML` deserializer, the configuration file would need to look like this:

```yaml
GenericConfig:
  Field1: "value1"
  Field2: 42
Port: 8080
```

This does not apply for overrides or the mock loader (they use the same code behind the scenes). With overrides one can either use the path with or without the embedded struct name. Meaning both 

```go
loader.Override("Field1", "new value")
// and
loader.Override("GenericConfig.Field1", "new value")
```

are valid. Overrides/MockLoader also needs to match the field names exactly, they do not account for any tags.

### Applying Overrides

You can override specific fields of the configuration, useful for setting dynamic values or secrets:

```go
loader.Override("Nested.Field3", false)
if err := loader.Load(&config); err != nil {
    log.Fatalf("Failed to apply overrides: %s", err)
}
```

### Working with Overrides and Deserializers

ConfigLoader supports multiple deserializers out of the box. Here's how you can use an override file with a custom deserializer:

```go
loader := configloader.NewConfigLoader("main_config.yaml",
    configloader.WithPath("/path/to/config/"),
    configloader.WithDeserializer(new(configloader.JSONDeserializer)),
    configloader.WithOverrideFile("/path/to/overrides/", "override_config.yaml"),
    configloader.WithOverrideDeserializer(new(configloader.YAMLDeserializer)),
)
```

## Testing with MockLoader

For testing purposes, ConfigLoader provides a `MockLoader` to simulate loading configurations without file dependencies:

```go
mockData := map[string]interface{}{
    "Field1": "mockValue",
    "Nested.Field3": true,
}
loader := configloader.NewMockLoader(mockData)
var config AppConfig
if err := loader.Load(&config); err != nil {
    log.Fatalf("Failed to load mock configuration: %s", err)
}
```
