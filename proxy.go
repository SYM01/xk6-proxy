package k6proxy

import (
	"errors"
	"net/http"
	"net/url"

	"go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/proxy", New())
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
		// proxier is the exported type
		proxier *Proxy
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu: vu,
		proxier: &Proxy{
			vu: vu,
		},
	}
}

// Compare is the type for our custom API.
type Proxy struct {
	vu           modules.VU
	oldProxyFunc func(*http.Request) (*url.URL, error)
}

func (p *Proxy) SetProxy(proxyURL string) error {
	if p.vu.State() == nil || p.vu.State().Transport == nil {
		return errors.New("unable to set proxy: unexpected nil transport")
	}

	tp, ok := p.vu.State().Transport.(*http.Transport)
	if !ok {
		return errors.New("unable to set proxy: other extensions might highjack the http already")
	}

	u, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}

	if p.oldProxyFunc == nil {
		p.oldProxyFunc = tp.Proxy
	}
	tp.Proxy = http.ProxyURL(u)
	return nil
}

func (p *Proxy) ClearProxy() {
	if p.oldProxyFunc == nil {
		return
	}

	p.vu.State().Transport.(*http.Transport).Proxy = p.oldProxyFunc
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.proxier,
	}
}
