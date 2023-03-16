# Archives

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