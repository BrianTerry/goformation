package resources

// AWSCognitoUserPool_Policies AWS CloudFormation Resource (AWS::Cognito::UserPool.Policies)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html
type AWSCognitoUserPool_Policies struct {

	// PasswordPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-policies.html#cfn-cognito-userpool-policies-passwordpolicy
	PasswordPolicy AWSCognitoUserPool_PasswordPolicy `json:"PasswordPolicy"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_Policies) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.Policies"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_Policies) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
