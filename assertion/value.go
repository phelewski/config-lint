package assertion

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/url"
)

// StandardValueSource can fetch values from external sources
type StandardValueSource struct {
	Log LoggingFunction
}

// GetValue looks up external values when an Assertion includes a ValueFrom attribute
func (v StandardValueSource) GetValue(assertion Assertion) (string, error) {
	if assertion.ValueFrom.URL != "" {
		v.Log(fmt.Sprintf("Getting value_from %s", assertion.ValueFrom.URL))
		parsedURL, err := url.Parse(assertion.ValueFrom.URL)
		if err != nil {
			return "", err
		}
		if parsedURL.Scheme != "s3" && parsedURL.Scheme != "S3" {
			return "", fmt.Errorf("Unsupported protocol for value_from: %s", parsedURL.Scheme)
		}
		content, err := v.GetValueFromS3(parsedURL.Host, parsedURL.Path)
		if err != nil {
			return "", err
		}
		v.Log(content)
		return content, nil
	}
	return assertion.Value, nil
}

// GetValueFromS3 looks up external values for an Assertion when the S3 protocol is specified
func (v StandardValueSource) GetValueFromS3(bucket string, key string) (string, error) {
	region := &aws.Config{Region: aws.String("us-east-1")}
	awsSession := session.New()
	s3Client := s3.New(awsSession, region)
	response, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.String(), nil
}
