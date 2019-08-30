# gochimp

This is a fork of an unmainted library for interacting with the [Mandrill API](https://github.com/mattbaird/gochimp). I want to thank Matt Baird for all the original work he did (especially around the massive Messages part of the API)

Several packages have been restructured as to help make testing more easy and isolate concerns.

## Design Goals

### Separation of request/response parsing from end-user api

One of the biggest changes is breaking out http api representation from go api representation.
Having written many clients for various apis, this is the biggest pain point in maintaining upgrades.
Basically directly adding your json struct tags to the structs that are used by end users of your go api creates a lot of pain.

To that end, there is a dedicate package that defines the requests and responses that are used by the higher level client.

## Minimal function signatures

Where appropriate, functions should only take the required information.
Optional arguments should variadic functional options.

This ensures that function signatures don't require multiple empty placeholders and can be modified with new functionality without breaking the existing function signature.

## Configurability

As part of the functional options, it allows you to make bits of the client configurable without needing private APIs for testing. This also benefits end-users in that using an alternate http.Client (perhaps with ochttp already wired up to it). Allowing users to use their own logging library (by accepting a stdlib-compatible logger such as zap's `NewStdLog`/`NewStdLogAt`)

## Optional simplified api

In addition to operating directly on an instance of a `mandrill.Client` there are some additional "sugared" interfaces (such as the `MessageBuilder` )that can be invoked after calling `mandrill.Connect` (which doesn't require you to throw away the created client variable which is just ugly to read)

## Examples

There are two relevant examples right now in `examples/send`

The `message` example shows using a fully instantiated client and calling relevant methods on that client to form and send a message via mandrill.

The `message-fluent` example shows the alternate fluent interface that doesn't require creating multiple structs to build out the `Message` instance to be sent.

One benefit of the fluent interface is particular is being able to specify recipients all at once with not only the email and name of the recipient, but and recipient specific metadata and merge variables.

I feel like, especially with something as flexible as message sending, it has the potential to reduce errors to an end user. The chance of missing a recipient's metadata or merge variables is reduced.

## TODO

- [ ] godoc comments finished
- [ ] redo tests
- [ ] finish remaining types
- [ ] more examples

[![GoDoc](https://godoc.org/github.com/lusis/gochimp?status.svg)](https://godoc.org/github.com/lusis/gochimp)