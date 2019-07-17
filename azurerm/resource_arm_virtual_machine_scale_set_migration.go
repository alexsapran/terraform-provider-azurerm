package azurerm

import (
	"fmt"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

func resourceVirtualMachineScaleSetMigrateState(v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	switch v {
	case 0:
		log.Println("[INFO] Found AzureRM Scale Set State v0; migrating to v1")
		return resourceVirtualMachineScaleSetStateV0toV1(is, meta)
	case 1:
		log.Printf("[INFO] Found AzureRM Scale Set State v0; migrating to v2")
		return resourceVirtualMachineScaleSetStateV1toV2(is, meta)
	default:
		return is, fmt.Errorf("Unexpected schema version: %d", v)
	}
}

func resourceVirtualMachineScaleSetStateV0toV1(is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Println("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}

	log.Printf("[DEBUG] ARM Virtual Machine Scale Set Attributes before Migration: %#v", is.Attributes)

	client := meta.(*ArmClient).vmScaleSetClient
	ctx := meta.(*ArmClient).StopContext

	resGroup := is.Attributes["resource_group_name"]
	name := is.Attributes["name"]

	read, err := client.Get(ctx, resGroup, name)
	if err != nil {
		return is, err
	}

	is.ID = *read.ID

	log.Printf("[DEBUG] ARM Virtual Machine Scale Set Attributes after State Migration: %#v", is.Attributes)

	return is, nil
}

func resourceVirtualMachineScaleSetStateV1toV2(is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Println("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}
	log.Printf("[INFO] ARM Virtual Machine Scale Set Attributes before migration: %#v", is.Attributes)
	d := resourceArmVirtualMachineScaleSet().Data(is)
	log.Printf("[INFO] ARM Virtual Machine Scale Set Attributes after migration: %#v", d.State())
	return d.State(), nil
}
