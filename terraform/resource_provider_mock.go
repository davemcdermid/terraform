package terraform

import "sync"

// MockResourceProvider implements ResourceProvider but mocks out all the
// calls for testing purposes.
type MockResourceProvider struct {
	sync.Mutex

	// Anything you want, in case you need to store extra data with the mock.
	Meta interface{}

	CloseCalled                  bool
	CloseError                   error
	InputCalled                  bool
	InputInput                   UIInput
	InputConfig                  *ResourceConfig
	InputReturnConfig            *ResourceConfig
	InputReturnError             error
	InputFn                      func(UIInput, *ResourceConfig) (*ResourceConfig, error)
	ApplyCalled                  bool
	ApplyInfo                    *InstanceInfo
	ApplyState                   *InstanceState
	ApplyDiff                    *InstanceDiff
	ApplyFn                      func(*InstanceInfo, *InstanceState, *InstanceDiff) (*InstanceState, error)
	ApplyReturn                  *InstanceState
	ApplyReturnError             error
	ConfigureCalled              bool
	ConfigureConfig              *ResourceConfig
	ConfigureFn                  func(*ResourceConfig) error
	ConfigureReturnError         error
	DiffCalled                   bool
	DiffInfo                     *InstanceInfo
	DiffState                    *InstanceState
	DiffDesired                  *ResourceConfig
	DiffFn                       func(*InstanceInfo, *InstanceState, *ResourceConfig) (*InstanceDiff, error)
	DiffReturn                   *InstanceDiff
	DiffReturnError              error
	RefreshCalled                bool
	RefreshInfo                  *InstanceInfo
	RefreshState                 *InstanceState
	RefreshFn                    func(*InstanceInfo, *InstanceState) (*InstanceState, error)
	RefreshReturn                *InstanceState
	RefreshReturnError           error
	ResourcesCalled              bool
	ResourcesReturn              []ResourceType
	ValidateCalled               bool
	ValidateConfig               *ResourceConfig
	ValidateFn                   func(*ResourceConfig) ([]string, []error)
	ValidateReturnWarns          []string
	ValidateReturnErrors         []error
	ValidateResourceFn           func(string, *ResourceConfig) ([]string, []error)
	ValidateResourceCalled       bool
	ValidateResourceType         string
	ValidateResourceConfig       *ResourceConfig
	ValidateResourceReturnWarns  []string
	ValidateResourceReturnErrors []error
}

func (p *MockResourceProvider) Close() error {
	p.CloseCalled = true
	return p.CloseError
}

func (p *MockResourceProvider) Input(
	input UIInput, c *ResourceConfig) (*ResourceConfig, error) {
	p.InputCalled = true
	p.InputInput = input
	p.InputConfig = c
	if p.InputFn != nil {
		return p.InputFn(input, c)
	}
	return p.InputReturnConfig, p.InputReturnError
}

func (p *MockResourceProvider) Validate(c *ResourceConfig) ([]string, []error) {
	p.Lock()
	defer p.Unlock()

	p.ValidateCalled = true
	p.ValidateConfig = c
	if p.ValidateFn != nil {
		return p.ValidateFn(c)
	}
	return p.ValidateReturnWarns, p.ValidateReturnErrors
}

func (p *MockResourceProvider) ValidateResource(t string, c *ResourceConfig) ([]string, []error) {
	p.Lock()
	defer p.Unlock()

	p.ValidateResourceCalled = true
	p.ValidateResourceType = t
	p.ValidateResourceConfig = c

	if p.ValidateResourceFn != nil {
		return p.ValidateResourceFn(t, c)
	}

	return p.ValidateResourceReturnWarns, p.ValidateResourceReturnErrors
}

func (p *MockResourceProvider) Configure(c *ResourceConfig) error {
	p.Lock()
	defer p.Unlock()

	p.ConfigureCalled = true
	p.ConfigureConfig = c

	if p.ConfigureFn != nil {
		return p.ConfigureFn(c)
	}

	return p.ConfigureReturnError
}

func (p *MockResourceProvider) Apply(
	info *InstanceInfo,
	state *InstanceState,
	diff *InstanceDiff) (*InstanceState, error) {
	p.Lock()
	defer p.Unlock()

	p.ApplyCalled = true
	p.ApplyInfo = info
	p.ApplyState = state
	p.ApplyDiff = diff
	if p.ApplyFn != nil {
		return p.ApplyFn(info, state, diff)
	}

	return p.ApplyReturn, p.ApplyReturnError
}

func (p *MockResourceProvider) Diff(
	info *InstanceInfo,
	state *InstanceState,
	desired *ResourceConfig) (*InstanceDiff, error) {
	p.Lock()
	defer p.Unlock()

	p.DiffCalled = true
	p.DiffInfo = info
	p.DiffState = state
	p.DiffDesired = desired
	if p.DiffFn != nil {
		return p.DiffFn(info, state, desired)
	}

	return p.DiffReturn, p.DiffReturnError
}

func (p *MockResourceProvider) Refresh(
	info *InstanceInfo,
	s *InstanceState) (*InstanceState, error) {
	p.Lock()
	defer p.Unlock()

	p.RefreshCalled = true
	p.RefreshInfo = info
	p.RefreshState = s

	if p.RefreshFn != nil {
		return p.RefreshFn(info, s)
	}

	return p.RefreshReturn, p.RefreshReturnError
}

func (p *MockResourceProvider) Resources() []ResourceType {
	p.Lock()
	defer p.Unlock()

	p.ResourcesCalled = true
	return p.ResourcesReturn
}
