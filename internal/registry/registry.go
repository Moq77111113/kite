package registry

func (r *Registry) FindByName(name string) *RegistryItem {
	for i := range r.Items {
		if r.Items[i].Name == name {
			return &r.Items[i]
		}
	}
	return nil
}