# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- /league endpoint, which returns the scores for every player in the league
- Concurrent operation support in InMemoryPlayerStore
### Changed
- PlayerStore interface moved from the server package to the store package
- Promoted InMemoryPlayerStore to proper feature
### Deprecated
### Removed
### Fixed
### Security

## [0.1.0] - 2024-10-01

### Added

- HTTP server with /player endpoint
- In-memory store to keep track of player state. Note that this feature is experimental and may be removed.
