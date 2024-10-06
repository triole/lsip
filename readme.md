# lsip ![build](https://github.com/triole/lsip/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/lsip/actions/workflows/test.yaml/badge.svg)

<!-- toc -->

- [Synopsis](#synopsis)
- [Help](#help)

<!-- /toc -->

## Synopsis

Quickly get and print your external IPv4 or IPv6 Address.

## Help

```go mdox-exec="r -h"

list your ip

Flags:
  -h, --help                      Show context-sensitive help.
  -p, --print="any"               print certain ip version, can be: 4, 6, both,
                                  any or all
  -t, --threads=16                threads, max of parallel requests
  -n, --dry-run                   dry run, just print don't do
      --log-file="/dev/stdout"    log file
      --log-level="info"          log level
      --log-no-colors             disable output colours, print plain text
      --log-json                  enable json log, instead of text one
  -V, --version-flag              display version
```
