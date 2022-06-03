# Connection Test Exporter

A quick utility I wrote to test my internet connection, and record a log of when it was up or down.

For now it just connects by default to bbc.co.uk, and writes a log to a directory in your home folder. It should work on all platforms supported by the go compiler, but has only been tested on Windows 10 for now, as that's the PC I have that's connected by ethernet directly to my home router.

# Build

```bash
go build
```