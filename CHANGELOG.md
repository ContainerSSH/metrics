# Changelog

## 0.9.2: Fixed JSON and YAML marshalling

In the previous version the JSON and YAML configuration marshalling / unmarshalling created an unnecessary sub-map, which was incompatible to ContainerSSH 0.3. This release fixes that and restores compatibility.

## 0.9.1: Updating GeoIP to 0.9.3

This release updates the [GeoIP dependency](https://github.com/containerssh/geoip) to version 0.9.3 for a cleaner API.

## 0.9.0: Initial release

This is the initial port from ContainerSSH 0.3