# hmoj

I got tired of looking at hashes and trying to see if they changed. Thus: a way to turn a `[]byte` into a deterministic emoji. 

Usage: 

```go
hmoj.Encode([]byte{0x01, 0x02, 0x03, 0x04})
``` 

will return

```
🙃 01020304
```

This package follows a similar design to the standard log package, with the intention that the Encoder is configurable. (Don't have to hex encode; don't have to include hash values; can adjust the length of the "short" hash.)