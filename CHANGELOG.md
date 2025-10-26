# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-10-26

### Added
- Go module support with `go.mod`
- CHANGELOG.md for tracking version changes
- Semantic versioning tags

### Changed
- **BREAKING**: Made library thread-safe by removing global mutable variables
  - Removed global `multiGSM7` and `multiUnicode` variables
  - Converted to local variables in `Split()` method
  - Updated `appendSms()` signature to accept these values as parameters
- Improved code documentation with thread-safety comments

### Fixed
- Race condition issues when using the library in concurrent goroutines
- Thread-safety now verified with `go test -race`

## [0.0.0] - Previous releases

### Features
- SMS splitting for GSM 7-bit charset
- SMS splitting for Unicode charset
- Support for basic and extended GSM 7 table
- Configurable UDH size (6 or 7 bytes)
- Automatic charset detection
- Force charset option
- High surrogate (emoji) support
