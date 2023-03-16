# Archives

[![Tests](https://github.com/rymdport/archives/actions/workflows/tests.yml/badge.svg)](https://github.com/rymdport/archives/actions/workflows/tests.yml)
[![Analysis](https://github.com/rymdport/archives/actions/workflows/analysis.yml/badge.svg)](https://github.com/rymdport/archives/actions/workflows/analysis.yml)

A Go package for easy handling of archive formats like `zip` and `tar` (both compressed and compressed).

## Supported formats

### Tar

Tar archives can be created using the `tar` package.

**The following compression formats are supported:**
- bzip2 (only decompression)
- gzip
- xz
- zstd

### Zip

Zip archives can be created using the `zip` package.

**The following compression formats are supported:**
- gzip

## Security

This library takes care to make sure that file paths are cleaned to try and avoid path traversals.
All close methods that returns errors are also handled accordingly when defered. 