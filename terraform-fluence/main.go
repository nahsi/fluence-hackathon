package main

import (
  "context"
  "log"

  "github.com/hashicorp/terraform-plugin-framework/providerserver"
  "github.com/cloudlesslabs/terraform-provider-fluence/internal/provider_fluence"
)

func main() {
  opts := providerserver.ServeOpts{
    Address: "registry.terraform.io/cloudlesslabs/fluence",
  }
  if err := providerserver.Serve(context.Background(), provider_fluence.New(), opts); err != nil {
    log.Fatal(err)
  }
}
