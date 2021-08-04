package dependents

import (
	"fmt"
	"x-msa-core/monitor/client"
)

func NewMDM() MonitorDependentsManager {
	return &monitorDependents{
		dep: map[string]client.MonitorClient{},
	}
}

func (d *monitorDependents) Add(key, addr string) error {
	client, err := client.NewMonitorClient(addr)
	if err != nil {
		return fmt.Errorf("cannot craete monitor client: %v", err)
	}

	d.rw.Lock()
	defer d.rw.Unlock()

	if _, ok := d.dep[key]; ok {
		return fmt.Errorf("monitor exist by key: %v", key)
	}
	d.dep[key] = client

	return nil
}

func (d *monitorDependents) Get(key string) (client.MonitorClient, error) {
	d.rw.Lock()
	defer d.rw.Unlock()

	if c, ok := d.dep[key]; ok {
		return c, nil
	} else {
		return nil, fmt.Errorf("monitor not found by key: %v", key)
	}
}

func (d *monitorDependents) GetFirst() (client.MonitorClient, error) {
	d.rw.Lock()
	defer d.rw.Unlock()

	for _, c := range d.dep {
		return c, nil
	}

	return nil, fmt.Errorf("monitor not found")
}

func (d *monitorDependents) Delete(key string) error {
	d.rw.Lock()
	defer d.rw.Unlock()

	if c, ok := d.dep[key]; ok {
		c.Close()
		delete(d.dep, key)
		return nil
	} else {
		return fmt.Errorf("monitor not found by key: %v", key)
	}
}
