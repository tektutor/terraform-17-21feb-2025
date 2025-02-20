package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

type DockerConfig struct {
   imageName string
   containerName string
   hostName string
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider {
			DataSourcesMap: map[string]*schema.Resource {

			},
			ResourcesMap: map[string]*schema.Resource {
			     "tektutor_docker_container": resourceDockerContainer(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure( version string, p *schema.Provider) func(context.Context, *schema.ResourceData) ( any, diag.Diagnostics) {
  	return func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
		dockerConfig := DockerConfig{}

		return &dockerConfig, nil
	}
}
