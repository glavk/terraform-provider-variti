package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOrigin() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Origin resource in the Terraform provider Variti.",

		CreateContext: resourceOriginCreate,
		ReadContext:   resourceOriginRead,
		UpdateContext: resourceOriginUpdate,
		DeleteContext: resourceOriginDelete,

		Schema: map[string]*schema.Schema{
			"ip_attribute": {
				// This description is used by the documentation generator and the language server.
				Description: "Origin IP.",
				Type:        schema.TypeString,
				Optional:    false,
			},
			"mode_attribute": {
				// This description is used by the documentation generator and the language server.
				Description: "Origin Mode: primary or backup.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"weight_attribute": {
				// This description is used by the documentation generator and the language server.
				Description: "Origin Weight: 0 - 99.",
				Type:        schema.TypeInt,
				Optional:    false,
			},
		},
	}
}

func resourceOriginCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}

func resourceOriginRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	token := d.Get("token").(string)
	if token == "" {
		return diag.Errorf("empty token")
	}

	return diag.Errorf("not implemented")
}

func resourceOriginUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceOriginDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
