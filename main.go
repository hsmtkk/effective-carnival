package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-google-go/google/v4/artifactregistryrepository"
	"github.com/cdktf/cdktf-provider-google-go/google/v4/cloudbuildtrigger"
	"github.com/cdktf/cdktf-provider-google-go/google/v4/cloudrunservice"
	"github.com/cdktf/cdktf-provider-google-go/google/v4/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const (
	owner      = "hsmtkk"
	project    = "effective-carnival"
	region     = "asia-northeast1"
	repository = "effective-carnival"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	provider.NewGoogleProvider(stack, jsii.String("google"), &provider.GoogleProviderConfig{
		Project: jsii.String(project),
	})

	artifactregistryrepository.NewArtifactRegistryRepository(stack, jsii.String("artifact_registry"), &artifactregistryrepository.ArtifactRegistryRepositoryConfig{
		Format:       jsii.String("docker"),
		Location:     jsii.String(region),
		RepositoryId: jsii.String("registry"),
	})

	cloudbuildtrigger.NewCloudbuildTrigger(stack, jsii.String("cloud_build_trigger"), &cloudbuildtrigger.CloudbuildTriggerConfig{
		Filename: jsii.String("cloudbuild.yaml"),
		Github: &cloudbuildtrigger.CloudbuildTriggerGithub{
			Name:  jsii.String(repository),
			Owner: jsii.String(owner),
			Push: &cloudbuildtrigger.CloudbuildTriggerGithubPush{
				Branch: jsii.String("main"),
			},
		},
	})

	cloudrunservice.NewCloudRunService(stack, jsii.String("cloud_run_service"), &cloudrunservice.CloudRunServiceConfig{
		AutogenerateRevisionName: true,
		Location:                 jsii.String(region),
		Name:                     jsii.String("example"),
		Template: &cloudrunservice.CloudRunServiceTemplate{
			Spec: &cloudrunservice.CloudRunServiceTemplateSpec{
				Containers: []*cloudrunservice.CloudRunServiceTemplateSpecContainers{{
					Image: jsii.String("us-docker.pkg.dev/cloudrun/container/hello"),
				}},
			},
		},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "effective-carnival")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendProps{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("hsmtkkdefault"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("effective-carnival")),
	})

	app.Synth()
}
