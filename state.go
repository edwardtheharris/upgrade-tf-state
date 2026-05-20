package main

func (r *ThingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
    return map[int64]resource.StateUpgrader{
        // State upgrade implementation from 0 (prior state version) to 1 (Schema.Version)
        0: {
            PriorSchema: &schema.Schema{
                Attributes: map[string]schema.Attribute{
                    "id": schema.StringAttribute{
                        Computed: true,
                    },
                    "optional_attribute": schema.BoolAttribute{
                        // As compared to current schema.StringAttribute above
                        Optional: true,
                    },
                    "required_attribute": schema.BoolAttribute{
                        // As compared to current schema.StringAttribute above
                        Required: true,
                    },
                },
            },
            StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
                var priorStateData ThingResourceModelV0

                resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)

                if resp.Diagnostics.HasError() {
                    return
                }

                upgradedStateData := ThingResourceModelV1{
                    Id:                priorStateData.Id,
                    RequiredAttribute: fmt.Sprintf("%t", priorStateData.RequiredAttribute),
                }

                if priorStateData.OptionalAttribute != nil {
                    v := fmt.Sprintf("%t", *priorStateData.OptionalAttribute)
                    upgradedStateData.OptionalAttribute = &v
                }

                resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
            },
        },
    }
}

