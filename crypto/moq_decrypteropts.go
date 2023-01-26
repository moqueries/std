// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package crypto

import (
	"crypto"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that crypto.DecrypterOpts is mocked
// completely
var _ crypto.DecrypterOpts = (*MoqDecrypterOpts_mock)(nil)

// MoqDecrypterOpts holds the state of a moq of the DecrypterOpts type
type MoqDecrypterOpts struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDecrypterOpts_mock

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqDecrypterOpts_mock isolates the mock interface of the DecrypterOpts type
type MoqDecrypterOpts_mock struct {
	Moq *MoqDecrypterOpts
}

// MoqDecrypterOpts_recorder isolates the recorder interface of the
// DecrypterOpts type
type MoqDecrypterOpts_recorder struct {
	Moq *MoqDecrypterOpts
}

// NewMoqDecrypterOpts creates a new moq of the DecrypterOpts type
func NewMoqDecrypterOpts(scene *moq.Scene, config *moq.Config) *MoqDecrypterOpts {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDecrypterOpts{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDecrypterOpts_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the DecrypterOpts type
func (m *MoqDecrypterOpts) Mock() *MoqDecrypterOpts_mock { return m.Moq }

// OnCall returns the recorder implementation of the DecrypterOpts type
func (m *MoqDecrypterOpts) OnCall() *MoqDecrypterOpts_recorder {
	return &MoqDecrypterOpts_recorder{
		Moq: m,
	}
}

// Reset resets the state of the moq
func (m *MoqDecrypterOpts) Reset() {}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDecrypterOpts) AssertExpectationsMet() { m.Scene.T.Helper() }