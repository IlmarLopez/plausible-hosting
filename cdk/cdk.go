package main

import (
	"os"

	"github.com/IlmarLopez/plausible-hosting/cdk/plausible-hosting/lib"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "dev"
	}

	app := awscdk.NewApp(nil)
	stackName := "plausible-stack"
	lib.NewPlausibleStack(app, *jsii.Sprintf("%s-%s", stackName, environment), &lib.PlausibleStackProps{
		StackProps: awscdk.StackProps{
			Env:       env(),
			StackName: jsii.String(stackName),
		},
		StackEnv: environment,
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
