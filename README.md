# Anonymijson

![](doc/logo.png)

Anonymise JSON data structures. The tool helps ensure that test data is
stripped of identifiable personal data and unlicensed content. This is useful
as it allows you to save time when creating new test cases or other test data
for your software. You can get started using Anonymijson provided you have Go
installed.

Scroll down for an explanation on how to install, and how to use.

## Installation

Provided you have Go installed, you can simply run

```
go get github.com/justuswilhelm/anonymijson
```

and as long as Go has executables in your `$PATH`, you can immediately start
using the command `anonymijson`.

## Usage

```bash
$ cat test.json test2.json
```

```json
{
  "1": 1.02,
  "bla": "lol",
  "hello": [
    1,
    2,
    3
  ]
}
[
  [
    1,
    2
  ],
  1,
  2,
  3
]
```

```bash
$ anonymijson test.json test2.json
```

```json
=== test.json ===
{
  "1": 0.9122551027038107,
  "bla": "scaberulous-honeysweet",
  "hello": [
    0.6288449008092282,
    0.9233224572170322,
    0.5275968199385908
  ]
}

=== test2.json ===
[
  [
    0.35278562141995556,
    0.5841195295045782
  ],
  0.9718421636089548,
  0.7314389548879902,
  0.2812210260450312
]
```

Will override all values in JSON objects, while leaving they key/array/hashmap
structure. The format will be kept, so strings will be replaced by random
strings and numbers by random numbers.

Furthermore, anonymijson can be run in-place. This means that files will be
replaced directly with the randomized values, instead of outputting the result
to stdout. This can be done by running

```bash
$ anonymijson test.json
```

The output is stored within the file provided.

```bash
$ cat test.json
```

```json
{
  "1": 0.09056946229454671,
  "bla": "Nathaniel-bonded",
  "hello": [
    0.8885474570694423,
    0.8161380166567574,
    0.4648524575247153
  ]
}
```
