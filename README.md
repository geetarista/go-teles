# go-teles [![Build Status](https://drone.io/github.com/geetarista/go-teles/status.png)](https://drone.io/github.com/geetarista/go-teles/latest)

A [teles](https://github.com/armon/teles) client powered by [Go](http://golang.org).

## Installation

```bash
go get github.com/geetarista/go-teles/teles
```

## Documentation

[Read it online](http://godoc.org/github.com/geetarista/go-teles/teles)

Or read it locally:

```bash
go doc github.com/geetarista/go-bloomd/teles
```

## Testing

I use Vagrant to run the tests against a Teles server. Use the included [Vagrantfile](Vagrantfile) and make sure you use your VM's IP address in `test_helpers.go`.

## License

MIT. See `LICENSE`.
