// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package acmpca provides the client and types for making API
// requests to AWS Certificate Manager Private Certificate Authority.
//
// You can use the ACM PCA API to create a private certificate authority (CA).
// You must first call the CreateCertificateAuthority function. If successful,
// the function returns an Amazon Resource Name (ARN) for your private CA. Use
// this ARN as input to the GetCertificateAuthorityCsr function to retrieve
// the certificate signing request (CSR) for your private CA certificate. Sign
// the CSR using the root or an intermediate CA in your on-premises PKI hierarchy,
// and call the ImportCertificateAuthorityCertificate to import your signed
// private CA certificate into ACM PCA.
//
// Use your private CA to issue and revoke certificates. These are private certificates
// that identify and secure client computers, servers, applications, services,
// devices, and users over SSLS/TLS connections within your organization. Call
// the IssueCertificate function to issue a certificate. Call the RevokeCertificate
// function to revoke a certificate.
//
// Certificates issued by your private CA can be trusted only within your organization,
// not publicly.
//
// Your private CA can optionally create a certificate revocation list (CRL)
// to track the certificates you revoke. To create a CRL, you must specify a
// RevocationConfiguration object when you call the CreateCertificateAuthority
// function. ACM PCA writes the CRL to an S3 bucket that you specify. You must
// specify a bucket policy that grants ACM PCA write permission.
//
// You can also call the CreateCertificateAuthorityAuditReport to create an
// optional audit report that lists every time the CA private key is used. The
// private key is used for signing when the IssueCertificate or RevokeCertificate
// function is called.
//
// See https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22 for more information on this service.
//
// See acmpca package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/acmpca/
//
// Using the Client
//
// To AWS Certificate Manager Private Certificate Authority with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Certificate Manager Private Certificate Authority client ACMPCA for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/acmpca/#New
package acmpca
