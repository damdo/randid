## randid

**generate** (crypto-secure) **random** (alphanumeric string) **ids** - in Go


### FEATURES:
- **secure** :lock: crypto random ids thanks to `crypto/rand`
- **fast** :zap:
- both fixed and arbitrary long generation methods
- **tested** :white_check_mark:


### USAGE:
download the package
```sh
go get github.com/damdo/randid
```

import and use the package
```go
import "github.com/damdo/randid"

func foo(){

    // generate a default len (32 char) id
    your32CharLongID := randid.ID()
    // -> 26e99c11f0a31ec57924c8e2a0712cd3

    // or
    desiredLen := 64
    your64CharLongID := randid.SizedID(desiredLen)
    // -> 6cc77090ff64c232613574bb562510f78d88a84e8351f3a68ac1caa902750bb7


    // change randid default len
    // every id generated from now on with .ID() will be 33 char long
    randid.DefaultLen = 33

    your33CharLongID := randid.ID()
    // -> 091015d5146b6d2e3f77ff50d805cee50

}
```
### CREDITS:
Credits goes to https://github.com/moby/moby/blob/master/pkg/stringid <br>
upon which this mini package is heavily based.