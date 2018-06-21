# go-aws-tools

There are many AWS tools. These are ours. 

## Install

You will need to have both `Go` (specifically a version of Go more recent than 1.6 so let's just assume you need [Go 1.8](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

Wet paint. Move along...

## Tools

### aws-mfa-session

```
> ./bin/aws-mfa-session -profile {PROFILE}
Enter your MFA token code: 123456
2018/06/21 12:00:49 Updated session credentials for 'session' profile (expires 2018-06-21 20:04:15 +0000 UTC)
```

## See also:

* https://github.com/DEEP-IMPACT-AG/skuld
* https://docs.aws.amazon.com/cli/latest/reference/sts/get-session-token.html
