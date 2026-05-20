package main

import "github.com/hashicorp/terraform-plugin-framework/resource"


// Other Resource methods are omitted in this example
var _ resource.Resource = &ThingResource{}
var _ resource.ResourceWithUpgradeState = &ThingResource{}

type ThingResource struct{/* ... */}

type ThingResourceModelV0 struct {
    Id                string `tfsdk:"id"`
    OptionalAttribute *bool  `tfsdk:"optional_attribute"`
    RequiredAttribute bool   `tfsdk:"required_attribute"`
}

type ThingResourceModelV1 struct {
    Id                string  `tfsdk:"id"`
    OptionalAttribute *string `tfsdk:"optional_attribute"`
    RequiredAttribute string  `tfsdk:"required_attribute"`
}

func (r *ThingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "id": schema.StringAttribute{
                Computed: true,
            },
            "optional_attribute": schema.StringAttribute{
                // As compared to prior schema.BoolAttribute below
                Optional: true,
            },
            "required_attribute": schema.StringAttribute{
                // As compared to prior schema.BoolAttribute below
                Required: true,
            },
        },
        // The resource has a prior state version of 0, which had the attribute
        // types of types.BoolType as shown below.
        Version: 1,
    }
}

