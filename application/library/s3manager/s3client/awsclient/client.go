package awsclient

import (
	"github.com/admpub/nging/application/dbschema"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Connect(m *dbschema.NgingCloudStorage) (client *s3.S3, err error) {
	isSecure := m.Secure == `Y`
	config := &aws.Config{
		DisableSSL:  aws.Bool(!isSecure),
		Endpoint:    aws.String(m.Endpoint),
		Credentials: credentials.NewStaticCredentials(m.Key, m.Secret, ""),
	}
	if len(m.Region) > 0 {
		config.Region = aws.String(m.Region)
	}
	var sess *session.Session
	sess, err = session.NewSession(config)
	if err != nil {
		return
	}
	client = s3.New(sess)
	return client, nil
}
