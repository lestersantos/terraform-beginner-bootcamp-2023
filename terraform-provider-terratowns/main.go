//package main: Declares the package name. The main package is special in Go it's
// where the execution of the program starts.
package main

// "fmt": Is short for format, which provides I/O formatting functions.
import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
// func main(): The main function is the entry point of the program. 
// When the program is run, execution starts from this function.
func main() {
	// Format.Printline
	// Prints to standard output
	fmt.Println("Hello, World!")

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}

// in golang, a titlecase function will be exported
func Provider() *schema.Provider {
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type: schema.TypeString,
				Required: true,
				Description: "The endpoint for the external service",
			},
			"token": {
				Type: schema.TypeString,
				Sensitive: true, // Mark the token as sensitive to hide it in logs
				Required: true,
				Description: "Bearer token for authorization",
			},
			"user_uuid":{
				Type: schema.TypeString,
				Required: true,
				Description: "UUID for configuration",
				// ValidateFunc: validateUUID,
			},
		},
	}
	// p.ConfigureContextFunc = providerConfigure(p)
	return p
}

// func validateUUID(v interface{}, k string) (ws []string, errors []error) {
// 	log.Print("validateUUID:start")
// 	value := v.(string)
// 	if _, err := uuid.Parse(value); err != nil {
// 		errors = append(error, fmt.Errorf("Invalid UUID format"))
// 	}
// 	log.Print("validateUUID:end")
// }