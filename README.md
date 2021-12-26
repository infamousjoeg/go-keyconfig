# go-keyconfig <!-- omit in toc -->

[![Go Reference](https://pkg.go.dev/badge/github.com/infamousjoeg/go-keyconfig.svg)](https://pkg.go.dev/github.com/infamousjoeg/go-keyconfig) [![keyconfig CI](https://github.com/infamousjoeg/go-keyconfig/actions/workflows/ci.yml/badge.svg)](https://github.com/infamousjoeg/go-keyconfig/actions/workflows/ci.yml)

A Golang package for simplifying storing configuration in the OS-provided secret manager.

- [Operating System Support](#operating-system-support)
- [Installation](#installation)
- [Usage](#usage)
  - [Example](#example)
  - [func SetConfig](#func-setconfig)
  - [func GetConfig](#func-getconfig)
  - [func DeleteConfig](#func-deleteconfig)
- [Testing](#testing)
- [Security](#security)
- [License](#license)
- [Maintainer](#maintainer)
- [Contributions](#contributions)

## Operating System Support

|OS|Secret Manager|
|--|--|
|MacOS|OSX Keychain|
|Linux|Keyring|
|Windows|Windows Credential Manager|

## Installation

```shell
go get -u github.com/infamousjoeg/go-keyconfig
```

## Usage

### Example

For a full example of usage, please see [example/example.go]().

```golang
import "github.com/infamousjoeg/go-keyconfig"
```

### func SetConfig

```golang
func SetConfig (configID string, config interface{}) error
```

`keyconfig.SetConfig` sets a struct type containing config key/values to the secret store of the current OS using the `configID` value as the identifier. It first encodes the struct to JSON and then Base64 encodes it for prior to setting in the secret store. Only an error will be returned, if one occurs.

### func GetConfig

```golang
func GetConfig (configID string, config interface{}) error
```

`keyconfig.GetConfig` gets a Base64-encoded key/value config from the current OS secret store using the `configID` value as the identifier. It then Base64 decodes to JSON and then will unmarshal the JSON data into the struct provided as the `config interface{}` value. Since the struct is provided when the function is called, only an error will be returned otherwise, if one occurs.

### func DeleteConfig

```golang
func DeleteConfig (configID string) error
```

`keyconfig.DeleteConfig` removes a Base64-encoded key/value config from the current OS secret store using the `configID` value as the identifier. Only an error will be returned, if one occurs.

## Testing

```shell
go test -v ./...
```

## Security

Please responsibly disclose any security issues or concerns as outlined in the [Security Guidelines](SECURITY.md).

## License

MIT

## Maintainer

[@infamousjoeg](https://github.com/infamousjoeg)

[![Buy me a coffee][buymeacoffee-shield]][buymeacoffee]

[buymeacoffee]: https://www.buymeacoffee.com/infamousjoeg
[buymeacoffee-shield]: https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png

## Contributions

Pull Requests are currently being accepted.  Please read and follow the guidelines laid out in [CONTRIBUTING.md]().
