# Docker - Multiple Traefik

Register different configurations in different traefiks

This plugin is useful when you have multiple instances of traefik, but each instance has its own configuration

### Configuration

For each plugin, the Traefik static configuration must define the module name (as is usual for Go packages).

The following declaration (given here in YAML) defines a plugin:

```yaml
# Static configuration

experimental:
  plugins:
    example:
      moduleName: github.com/leandrormd/docker-multiple-traefik
      version: v0.1.0

providers:
  plugin:
    example:
      pollInterval: 2s
```