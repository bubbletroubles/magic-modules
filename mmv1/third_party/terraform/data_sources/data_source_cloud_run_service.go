package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleCloudRunService() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceCloudRunService().Schema)
	addRequiredFieldsToSchema(dsSchema, "name", "location")
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceGoogleCloudRunServiceRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleCloudRunServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	id, err := ReplaceVars(d, config, "locations/{{location}}/namespaces/{{project}}/services/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceCloudRunServiceRead(d, meta)
}