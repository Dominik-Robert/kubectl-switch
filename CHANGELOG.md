# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
- Pipeline for building the binary
## [1.0.1] - 2022-05-04
### Added
- A right readme and this changelog
- docs: Added a helper function to all subcommands

## [1.0.0] - 2022-05-02
### Added
- feat: Add root command with global parameters
  - The global parameters for all subcommands are added in the root
- feat: :sparkles: Namespace switch
  - Implemented the namespaces switch. with kubectl switch ns <NAME> you can now change the namespace
- feat: :sparkles: Context switch
  - implemented the context switch. With kubectl switch ctx <CONTEXT> you can now change the context

### Changed

### Removed