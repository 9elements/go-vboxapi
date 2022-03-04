package vboxapi

import (
	"github.com/9elements/go-vboxapi/vboxweb"
)

type SystemProperties struct {
	virtualbox      *VirtualBox
	managedObjectId string
}

func (sp *SystemProperties) GetMaxNetworkAdapters(chipset *vboxweb.ChipsetType) (uint32, error) {
	request := vboxweb.ISystemPropertiesgetMaxNetworkAdapters{This: sp.managedObjectId, Chipset: chipset}

	response, err := sp.virtualbox.ISystemPropertiesgetMaxNetworkAdapters(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sp *SystemProperties) GetMaxDevicesPerPortForStorageBus(bus vboxweb.StorageBus) (uint32, error) {
	request := vboxweb.ISystemPropertiesgetMaxDevicesPerPortForStorageBus{This: sp.managedObjectId, Bus: &bus}
	response, err := sp.virtualbox.ISystemPropertiesgetMaxDevicesPerPortForStorageBus(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sp *SystemProperties) GetMinPortCountForStorageBus(bus vboxweb.StorageBus) (uint32, error) {
	request := vboxweb.ISystemPropertiesgetMinPortCountForStorageBus{This: sp.managedObjectId, Bus: &bus}
	response, err := sp.virtualbox.ISystemPropertiesgetMinPortCountForStorageBus(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	return response.Returnval, nil
}

func (sp *SystemProperties) Release() error {
	return sp.virtualbox.Release(sp.managedObjectId)
}
