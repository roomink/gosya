# Configure go app

The library allows you to organize access to configuration data in yaml format, taking into account the environment.

## Install

    go get github.com/roomink/gosya

## Usage

In the folder example complete example of use.

```go
import "github.com/roomink/gosya"

settings := &Settings{}
err := gosya.Merge(settings, path, env)
```
Organize files in the config directory

    path_to_config
    ├── settings
    │   └── development.yml
    ├── settings.local.yml
    └── settings.yml


`Merge` function takes several arguments: 
* `settings` A pointer to a structure that will project data from yaml. 
* `path` Path where the configuration files. 
* `env` Environment, for which the configuration is formed.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
