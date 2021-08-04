package dependents

import (
	"fmt"
	"x-msa-core/helper"
	"x-msa-core/service/client"
)

func NewSDM() ServiceDependentsManager {
	return &serviceDependents{
		dep: map[helper.GroupsType]map[string]client.ServiceClient{},
	}
}

func (d *serviceDependents) Add(key, addr string, group helper.GroupsType) error {
	c, err := client.NewServiceClient(addr, group)
	if err != nil {
		return fmt.Errorf("cannot craete service client %v", err)
	}

	d.rw.Lock()
	defer d.rw.Unlock()

	if _, ok := d.dep[group]; ok {
		if _, ok := d.dep[group][key]; ok {
			return fmt.Errorf("service exist by key %v", key)
		}
		d.dep[group][key] = c
	} else {
		d.dep[group] = map[string]client.ServiceClient{key: c}
	}

	return nil
}

func (d *serviceDependents) Get(group helper.GroupsType, key string) (client.ServiceClient, error) {
	d.rw.Lock()
	defer d.rw.Unlock()

	if _, ok := d.dep[group]; ok {
		if c, ok := d.dep[group][key]; ok {
			return c, nil
		}
		return nil, fmt.Errorf("service not found by key %v in group %v", key, group)
	} else {
		return nil, fmt.Errorf("service not found by key %v in group %v", key, group)
	}
}

func (d *serviceDependents) GetFirstByGroup(group helper.GroupsType) (client.ServiceClient, error) {
	d.rw.Lock()
	defer d.rw.Unlock()

	if list, ok := d.dep[group]; ok {
		for _, c := range list {
			return c, nil
		}
		return nil, fmt.Errorf("service not found by group %v", group)
	} else {
		return nil, fmt.Errorf("service not found by group %v", group)
	}
}

func (d *serviceDependents) Delete(group helper.GroupsType, key string) error {
	d.rw.Lock()
	defer d.rw.Unlock()

	if _, ok := d.dep[group]; ok {
		if c, ok := d.dep[group][key]; ok {
			c.Close()
			delete(d.dep[group], key)
			if len(d.dep[group]) == 0 {
				delete(d.dep, group)
			}
			return nil
		}
		return fmt.Errorf("service not found by key %v in group %v", key, group)
	} else {
		return fmt.Errorf("service not found by key %v in group %v", key, group)
	}
}
