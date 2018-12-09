# go-monkey [![Build Status](https://travis-ci.org/hangingman/go-monkey.svg?branch=master)](https://travis-ci.org/hangingman/go-monkey)

Golang lexer/parser learning scripts

## develop

- `$GOPATH` is usually `~/.go`

```
$ go get github.com/hangingman/go-monkey
$ cd ~/.go/src/github.com/hangingman/go-monkey
```

## test

```
$ make test
```

## build

```
$ make
```

# Hobby-programming-languages stored in this repo

## Monkey language

- Main dish of this repository

## Min language

- ALGOL like programming language

### Code Sample

```
var x, i;
  input x;
  if x=0 then x:=1 fi;
  i:=x-1;
  while i>=2
    begin
      x:=x*i;
      i:=i-1
    end;

  output x;
```

## RCF language

- RCF(Recursively Calling Functions)

### Code Sample

```
def fact(x) = if x>=2 then fact(x-1)*x else 1 fi
in fact(3)
```
