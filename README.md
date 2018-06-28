# go-aws-tools

There are many AWS tools. These are ours. 

## Install

You will need to have both `Go` (specifically a version of Go more recent than
1.6 so let's just assume you need [Go 1.10](https://golang.org/dl/) or higher)
and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

Wet paint. Move along...

## Tools

Many (all?) of these tools are available or can be replicated with the offical
`awscli` tool. The point of these tools is that they are specific to a need and,
importantly, they can be precompiled and not require a "hilarious"
install-all-the-dependencies dances which isn't always an option. Computers, amirite?

### aws-mfa-session

```
$> ./bin/aws-mfa-session -h
Usage of ./bin/aws-mfa-session:
  -duration int
    	    The time (in seconds) that your session should last (default 3600)
  -profile string
    	   A valid AWS credentials profile (default "default")
  -session-profile string
    		   The name of the AWS credentials profile to update with session credentials (default "session")
```

For example:

```
$> ./bin/aws-mfa-session -profile {PROFILE}
Enter your MFA token code: 123456
2018/06/21 12:00:49 Updated session credentials for 'session' profile (expires 2018-06-21 20:04:15 +0000 UTC)
```

### aws-cloudfront-invalidate

```
$> ./bin/aws-cloudfront-invalidate -h
Usage of ./bin/aws-cloudfront-invalidate:
  -distribution-id string
    		   A valid CloudFront distribution ID
  -path value
    	A path to invalidate (you can pass multiple -path flags for multiple paths)
  -profile string
    	   A valid AWS credentials profile (default "session")
```

For example:

```
$> ./bin/aws-cloudfront-invalidate -distribution-id {DISTRIBUTION} -path '/*'
{
  Invalidation: {
    CreateTime: 2018-06-28 00:41:43.093 +0000 UTC,
    Id: "I3KC8TSZPCCCKC",
    InvalidationBatch: {
      CallerReference: "08c0235a70e271838a2faf42bccb578c9672e01c",
      Paths: {
        Items: ["/*"],
        Quantity: 1
      }
    },
    Status: "InProgress"
  },
  Location: "https://cloudfront.amazonaws.com/2017-10-30/distribution/{DISTRIBUTION}/invalidation/I3KC8TSZPCCCKC"
}
```

## See also:

* https://github.com/DEEP-IMPACT-AG/skuld
* https://docs.aws.amazon.com/cli/latest/reference/sts/get-session-token.html
