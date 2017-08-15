package resources

// AWSKinesisFirehoseDeliveryStream_BufferingHints AWS CloudFormation Resource (AWS::KinesisFirehose::DeliveryStream.BufferingHints)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html
type AWSKinesisFirehoseDeliveryStream_BufferingHints struct {

	// IntervalInSeconds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints-intervalinseconds
	IntervalInSeconds int64 `json:"IntervalInSeconds"`

	// SizeInMBs AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints.html#cfn-kinesisfirehose-kinesisdeliverystream-s3destinationconfiguration-bufferinghints-sizeinmbs
	SizeInMBs int64 `json:"SizeInMBs"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisFirehoseDeliveryStream_BufferingHints) AWSCloudFormationType() string {
	return "AWS::KinesisFirehose::DeliveryStream.BufferingHints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisFirehoseDeliveryStream_BufferingHints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
