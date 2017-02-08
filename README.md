# Anonymijson

Anonymise JSON data structures.

## Usage

```
$ cat test.json
{
    "hello": [1, 2, 3],
    "bla": "lol",
    "1": 1.02
}
$ anonymijson < test.json | python -m json.tool
{
    "1": 0.6120456642749681,
    "bla": "BzgbaiCMRAjWwhTH",
    "hello": [
        0.5872982159059681,
        0.5372820936538049,
        1.2310533463860203
    ]
}
```

Will override all values in JSON objects, while leaving they key/array/hashmap
structure. The format will be kept, so strings will be replaced by random
strings and numbers by random numbers.

## Installation

Provided you have Go installed, you can simply run

```
go get github.com/justuswilhelm/anonymijson
```

and as long as Go has executables in your `$PATH`, you can immediately start
using the command `anonymijson`.
