[![Build Status](https://travis-ci.org/yowcow/goircparser.svg?branch=master)](https://travis-ci.org/yowcow/goircparser)

IRC Parser
==========

Go-port of Perl module [Parse::IRC](https://metacpan.org/pod/Parse::IRC) a.k.a. regular expression IRC command parser in Go.

HOW TO USE
----------

### Import

Do:

```
go get github.com/yowcow/goircparser
```

and import like:

```go
import "github.com/yowcow/goircparser"
```

### Parse

Pass a string of IRC command:

```go
row := goircparser.Parse(line) // *goircparser.Row
```

SEE ALSO
--------

[RFC 2812](https://tools.ietf.org/html/rfc2812)
[Parse::IRC](https://metacpan.org/pod/Parse::IRC)
