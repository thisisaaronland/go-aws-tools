tools:
	go build -mod vendor -o bin/aws-cloudfront-invalidate cmd/aws-cloudfront-invalidate/main.go
	go build -mod vendor -o bin/aws-mfa-session cmd/aws-mfa-session/main.go