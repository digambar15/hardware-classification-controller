package validate

import (
	hwcc "hardware-classification-controller/api/v1alpha1"

	bmh "github.com/metal3-io/baremetal-operator/pkg/apis/metal3/v1alpha1"
)

//Comparison function compare the host against the profile and filter the valid host
func Comparison(hosts []bmh.BareMetalHost, profiles []hwcc.ExpectedHardwareConfiguration) map[interface{}][]hwcc.ExpectedHardwareConfiguration {

	validHost := make(map[interface{}][]hwcc.ExpectedHardwareConfiguration)
	for _, host := range hosts {
		for _, profile := range profiles {
			if host.Status.HardwareDetails.CPU.Count >= profile.MinimumCPU.Count &&
				int64(host.Status.HardwareDetails.Storage[0].SizeBytes) >= (profile.MinimumDisk.SizeBytesGB*1024*1024) &&
				len(host.Status.HardwareDetails.NIC) >= profile.MinimumNICS.NumberOfNICS &&
				host.Status.HardwareDetails.RAMMebibytes >= (profile.MinimumRAM*1024) {
				newHost, ok := validHost[host.Status.HardwareDetails]
				if ok {
					validHost[host.Status.HardwareDetails] = append(newHost, profile)
				} else {
					var validProfile []hwcc.ExpectedHardwareConfiguration
					validHost[host.Status.HardwareDetails] = append(validProfile, profile)
				}
			}
		}
	}

	return validHost

}
