# Changelog

## 0.9.5: Added extended metrics

This release adds a `WithLabel` method to create metrics primed with certain labels. This can be used when passing labels between modules.

## 0.9.4: Add `Must*` methods

This release adds methods starting with `Must` that panic instead of throwing an error.

## 0.9.3: Custom label support

Each of the metric methods now allow adding extra labels:

```go
testCounter.Increment(
    net.ParseIP("127.0.0.1"),
    metrics.Label("foo", "bar"),
    metrics.Label("somelabel","somevalue")
)
```

The following rules apply and will cause a `panic` if violated:

- Label names and values cannot be empty.
- The `country` label name is reserved for GeoIP usage.

## 0.9.2: Fixed JSON and YAML marshalling

In the previous version the JSON and YAML configuration marshalling / unmarshalling created an unnecessary sub-map, which was incompatible to ContainerSSH 0.3. This release fixes that and restores compatibility.

## 0.9.1: Updating GeoIP to 0.9.3

This release updates the [GeoIP dependency](https://github.com/containerssh/geoip) to version 0.9.3 for a cleaner API.

## 0.9.0: Initial release

This is the initial port from ContainerSSH 0.3