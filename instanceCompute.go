package main

import (
	"os"
)

func instanceCompute() map[string]string{
	instanceCompute := map[string]string{"location": os.Getenv("LOCATION"),
        "name": os.Getenv("VM_NAME"),
        "offer": "",
        "osType": "Linux",
        "placementGroupId": "",
        "platformFaultDomain": "",
        "platformUpdateDomain": "",
        "publisher": "",
        "resourceGroupName": os.Getenv("RG_NAME"),
        "sku": "",
        "subscriptionId": os.Getenv("ID"),
        "tags": "",
        "vmId": "",
        "vmScaleSetName": "",
        "vmSize": "",
        "zone": ""}
	return instanceCompute
}