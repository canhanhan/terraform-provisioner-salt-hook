package main

import (
	"context"

	salt "github.com/finarfin/go-salt-netapi-client/cherrypy"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provisioner() terraform.ResourceProvisioner {
	return &schema.Provisioner{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},

		ApplyFunc:    apply,
		ValidateFunc: validate,
	}
}

func apply(ctx context.Context) error {
	data := ctx.Value(schema.ProvConfigDataKey).(*schema.ResourceData)

	cli := salt.NewClient(
		data.Get("address").(string),
		data.Get("username").(string),
		data.Get("password").(string),
		data.Get("backend").(string),
	)

	if err := cli.Login(); err != nil {
		return err
	}

	defer cli.Logout()

	tag := data.Get("id").(string)
	eventData := data.Get("data").(map[string]interface{})

	if err := cli.Hook(tag, eventData); err != nil {
		return err
	}

	return nil
}

func validate(c *terraform.ResourceConfig) (ws []string, es []error) {
	return ws, es
}
