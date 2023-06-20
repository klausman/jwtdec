# Ultra-simple JWT decoder in Go

Simple commandline tool so people don't have to use other people's websites to
decode JWTs. Those websites may all be aboveboard and do everything clientside,
but you can't audit every site out there. Better to run things locally, and
know where your data is.

## Usage

From a checkout of this tool:

```
$ ACCESSTOKEN="eyJ0...." go run .
```

If the tool has already been compiled and is in your path:

```
ACCESSTOKEN="eyJ0...." jwtdec
```

You cna also use `export ACCESSTOKEN="eyJ0...."` and then run the tool, but
then all subsequent programs run from the same shell will se the environment
variable as well, unless you use `unset ACCESSTOKEN`.
