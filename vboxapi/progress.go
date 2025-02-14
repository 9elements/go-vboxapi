package vboxapi

import "github.com/9elements/go-vboxapi/vboxweb"

type Progress struct {
	virtualbox *VirtualBox

	managedObjectId string
}

func (p *Progress) WaitForCompletion(timeout int32) error {
	request := vboxweb.IProgresswaitForCompletion{This: p.managedObjectId}
	request.Timeout = timeout

	_, err := p.virtualbox.IProgresswaitForCompletion(&request)
	if err != nil {
		return err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return nil
}

func (p *Progress) GetPercent() (uint32, error) {
	request := vboxweb.IProgressgetPercent{This: p.managedObjectId}

	response, err := p.virtualbox.IProgressgetPercent(&request)
	if err != nil {
		return 0, err // TODO: Wrap the error
	}

	// TODO: See if we need to do anything with the response
	return response.Returnval, nil
}

func (p *Progress) Release() error {
	return p.virtualbox.Release(p.managedObjectId)
}
