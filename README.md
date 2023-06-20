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
$ ACCESSTOKEN="eyJ0...." jwtdec
```

You cna also use `export ACCESSTOKEN="eyJ0...."` and then run the tool, but
then all subsequent programs run from the same shell will se the environment
variable as well, unless you use `unset ACCESSTOKEN`.

## JSON

Since the output of this program is JSON strings, it can be piped to `jq` for easier readability:

```
$ go run . |jq .
{
  "typ": "JWT",
  "alg": "RS256"
}
{
  "aud": "...",
  "jti": "...",
  "iat": 1234567890,
  "nbf": 9876543210,
  "exp": 11223344556,
  "sub": "12332112",
  "iss": "https://api.example.org",
  "ratelimit": {
    "requests_per_unit": 1,
    "unit": "DAY"
  },
  "scopes": [
    "basic"
  ]
}
```
