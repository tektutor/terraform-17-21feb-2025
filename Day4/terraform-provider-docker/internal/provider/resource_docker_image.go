package provider

import (
	"context"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/image"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDockerImage() *schema.Resource {

	return &schema.Resource {
		Description: "This resource will support create(pull), read(get), update and delete docker image via Terraform.",

		CreateContext: resourcePullDockerImage,
		ReadContext:   resourceGetDockerImageDetails,
		UpdateContext: resourceUpdateDockerImage,
		DeleteContext: resourceDeleteDockerImage,

		Schema: map[string]*schema.Schema {
			"docker_image": {
				Description: "Name of the docker image.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func inspectDockerImage( ctx context.Context, imageName string ) string {
	ctx = context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.47"))
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	out, err := cli.ImageInspect(ctx, imageName, nil)
	if err != nil {
		panic(err)
	}

	return out.ID
}

func resourcePullDockerImage( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	//Retrieve the inputs user provided in the terraform script .tf file
	imageName := d.Get("docker_image").(string)

	ctx = context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.47"))
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	imageId := inspectDockerImage ( ctx, imageName )
	d.SetId(imageId)

	return nil
}

func resourceGetDockerImageDetails( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	return nil
}

func resourceUpdateDockerImage( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	return nil
}

func resourceDeleteDockerImage( ctx context.Context, d *schema.ResourceData, meta any ) diag.Diagnostics {
	return nil
}
