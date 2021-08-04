package dependents

import (
	"fmt"
	"x-msa-core/observer/client"
)

func NewODM() ObserverDependentsManager {
	return &observerDependents{
		dep: map[string]client.ObserverClient{},
	}
}

func (d *observerDependents) Add(key, addr string) error {
	client, err := client.NewObserverClient(addr)
	if err != nil {
		return fmt.Errorf("cannot craete observer client: %v", err)
	}

	d.rw.Lock()
	defer d.rw.Unlock()

	if _, ok := d.dep[key]; ok {
		return fmt.Errorf("observer exist by key: %v", key)
	}
	d.dep[key] = client

	fmt.Println("observer added")

	return nil
}

func (d *observerDependents) Get(key string) (client.ObserverClient, error) {
	d.rw.Lock()
	defer d.rw.Unlock()

	if c, ok := d.dep[key]; ok {
		return c, nil
	} else {
		return nil, fmt.Errorf("observer not found by key: %v", key)
	}
}

func (d *observerDependents) GetFirst() (client.ObserverClient, error) {
	d.rw.Lock()
	defer d.rw.Unlock()

	for _, c := range d.dep {
		return c, nil
	}

	return nil, fmt.Errorf("observer not found")
}

func (d *observerDependents) Delete(key string) error {
	d.rw.Lock()
	defer d.rw.Unlock()

	if c, ok := d.dep[key]; ok {
		c.Close()
		delete(d.dep, key)
		return nil
	} else {
		return fmt.Errorf("observer not found by key: %v", key)
	}
}
