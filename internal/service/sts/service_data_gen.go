// Code generated by internal/generate/servicedata/main.go; DO NOT EDIT.

package sts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/experimental/intf"
)

var sd = &serviceData{}

func registerFrameworkDataSourceFactory(factory func(context.Context) (datasource.DataSourceWithConfigure, error)) {
	sd.frameworkDataSourceFactories = append(sd.frameworkDataSourceFactories, factory)
}

func registerFrameworkResourceFactory(factory func(context.Context) (resource.ResourceWithConfigure, error)) {
	sd.frameworkResourceFactories = append(sd.frameworkResourceFactories, factory)
}

type serviceData struct {
	frameworkDataSourceFactories []func(context.Context) (datasource.DataSourceWithConfigure, error)
	frameworkResourceFactories   []func(context.Context) (resource.ResourceWithConfigure, error)
}

func (d *serviceData) Configure(ctx context.Context, meta any) error {
	return nil
}

func (d *serviceData) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return d.frameworkDataSourceFactories
}

func (d *serviceData) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return d.frameworkResourceFactories
}

var ServiceData intf.ServiceData = sd
