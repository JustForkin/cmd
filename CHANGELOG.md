# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]
### Added
- `Graceful` for network servers to implement graceful restart.
- `SystemdListeners` returns `[]net.Listener` for [systemd socket activation][activation].

## [1.1.0] - 2016-08-24
### Added
- `IDGenerator` generates UUID-like ID string for request tracking.
- `Go` issues new request tracking ID and store it in the derived context.
- `HTTPClient`, a wrapper for `http.Client` that exports request tracking ID and logs results.
- `LogCmd`, a wrapper for `exec.Cmd` that records command execution results together with request tracking ID.

### Changed
- `HTTPServer` adds or imports request tracking ID for every request.
- `Server` adds request tracking ID for each new connection.
- Install signal handler only for the global environment.

### Removed
- `Context` method of `Environment` is removed.  It was a design flaw.

## [1.0.1] - 2016-08-22
### Changed
- Update docs.
- Use [cybozu-go/netutil](https://github.com/cybozu-go/netutil).
- Conform to cybozu-go/log v1.1.0 spec.

[activation]: http://0pointer.de/blog/projects/socket-activation.html
[Unreleased]: https://github.com/cybozu-go/cmd/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/cybozu-go/cmd/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/cybozu-go/cmd/compare/v1.0.0...v1.0.1
