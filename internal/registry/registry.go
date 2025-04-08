package registry

func (r *Registry) ResolveWithDependencies(names []string) ([]string, error) {
    resolved := []string{}
    visited := make(map[string]bool)

    var resolve func(name string)
    resolve = func(name string) {
        if visited[name] {
            return
        }
        visited[name] = true

        var module *IndexEntry
        for i := range r.Modules {
            if r.Modules[i].Name == name {
                module = &r.Modules[i]
                break
            }
        }
        if module == nil {
            return
        }

        resolved = append(resolved, module.Name)
        for _, dep := range module.RegistryDeps {
            resolve(dep)
        }
    }

    for _, name := range names {
        resolve(name)
    }

    return resolved, nil
}