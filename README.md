# Connection Test Exporter

A quick utility I wrote to test my internet connection, and record a log of when it was up or down.

*** Now with prometheus metrics support! ***

By default it attempts to connect by default to bbc.co.uk, and will export a number of metrics to prometheus as it does so, along with writing out some cool output to your terminal. It can also write a log to a directory in your home folder. It should work on all platforms supported by the go compiler, but has only been tested on Windows 10 for now, as that's the PC I have that's connected by ethernet directly to my home router.

## Build

```bash
go build
```

## Requires

Some kind of emoji font installed because I like being able to spot the dropped connections at a glance in the output.

## Run

```bash
PS C:\Users\point\Documents\Development\connection-test> connection-test 
[info] 2022/06/03 10:08:18.456512 main.go:108: script invoked using log file: C:\Users\point\connection-tests\Fri-3-Jun-2022.log
[info] 2022/06/03 10:08:18.666002 main.go:55: 👍 Connection to https://bbc.co.uk was a success, status code: 200
[info] 2022/06/03 10:08:20.795040 main.go:55: 👍 Connection to https://bbc.co.uk was a success, status code: 200
[info] 2022/06/03 10:08:22.844687 main.go:55: 👍 Connection to https://bbc.co.uk was a success, status code: 200
[info] 2022/06/03 10:08:24.896820 main.go:55: 👍 Connection to https://bbc.co.uk was a success, status code: 200
[info] 2022/06/03 10:08:26.949276 main.go:55: 👍 Connection to https://bbc.co.uk was a success, status code: 200
```

Example log output:
```Text
[error] 2022/06/03 00:48:00.811478 main.go:46: Get "https://bbc.co.uk": dial tcp: lookup bbc.co.uk: no such host
[error] 2022/06/03 00:48:00.811478 main.go:55: 😡 Connection to https://bbc.co.uk was a fail, response nil
[error] 2022/06/03 00:48:09.872935 main.go:46: Get "https://bbc.co.uk": dial tcp: lookup bbc.co.uk: no such host
[error] 2022/06/03 00:48:09.872935 main.go:55: 😡 Connection to https://bbc.co.uk was a fail, response nil
[error] 2022/06/03 00:48:11.879708 main.go:46: Get "https://bbc.co.uk": dial tcp: lookup bbc.co.uk: no such host
[error] 2022/06/03 00:48:11.879708 main.go:55: 😡 Connection to https://bbc.co.uk was a fail, response nil
[info] 2022/06/03 10:08:18.456512 main.go:108: script invoked using log file: C:\Users\point\connection-tests\Fri-3-Jun-2022.log
[info] 2022/06/03 10:08:18.666002 main.go:55: 👍 Connection to https://bbc.co.uk was a success, status code: 200
[info] 2022/06/03 10:08:20.795040 main.go:55: 👍 Connection to https://bbc.co.uk was a success, status code: 200
```

## FAQs

### Why not just use the blackbox exporter?

Good question - because I wanted to write something myself.