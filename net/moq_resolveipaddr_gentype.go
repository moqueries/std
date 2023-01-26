// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"net"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ResolveIPAddr_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ResolveIPAddr_genType func(network, address string) (*net.IPAddr, error)

// MoqResolveIPAddr_genType holds the state of a moq of the
// ResolveIPAddr_genType type
type MoqResolveIPAddr_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqResolveIPAddr_genType_mock

	ResultsByParams []MoqResolveIPAddr_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Network moq.ParamIndexing
			Address moq.ParamIndexing
		}
	}
}

// MoqResolveIPAddr_genType_mock isolates the mock interface of the
// ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_mock struct {
	Moq *MoqResolveIPAddr_genType
}

// MoqResolveIPAddr_genType_params holds the params of the
// ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_params struct{ Network, Address string }

// MoqResolveIPAddr_genType_paramsKey holds the map key params of the
// ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_paramsKey struct {
	Params struct{ Network, Address string }
	Hashes struct{ Network, Address hash.Hash }
}

// MoqResolveIPAddr_genType_resultsByParams contains the results for a given
// set of parameters for the ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqResolveIPAddr_genType_paramsKey]*MoqResolveIPAddr_genType_results
}

// MoqResolveIPAddr_genType_doFn defines the type of function needed when
// calling AndDo for the ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_doFn func(network, address string)

// MoqResolveIPAddr_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_doReturnFn func(network, address string) (*net.IPAddr, error)

// MoqResolveIPAddr_genType_results holds the results of the
// ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_results struct {
	Params  MoqResolveIPAddr_genType_params
	Results []struct {
		Values *struct {
			Result1 *net.IPAddr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResolveIPAddr_genType_doFn
		DoReturnFn MoqResolveIPAddr_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqResolveIPAddr_genType_fnRecorder routes recorded function calls to the
// MoqResolveIPAddr_genType moq
type MoqResolveIPAddr_genType_fnRecorder struct {
	Params    MoqResolveIPAddr_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqResolveIPAddr_genType_results
	Moq       *MoqResolveIPAddr_genType
}

// MoqResolveIPAddr_genType_anyParams isolates the any params functions of the
// ResolveIPAddr_genType type
type MoqResolveIPAddr_genType_anyParams struct {
	Recorder *MoqResolveIPAddr_genType_fnRecorder
}

// NewMoqResolveIPAddr_genType creates a new moq of the ResolveIPAddr_genType
// type
func NewMoqResolveIPAddr_genType(scene *moq.Scene, config *moq.Config) *MoqResolveIPAddr_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqResolveIPAddr_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqResolveIPAddr_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Network moq.ParamIndexing
				Address moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Network moq.ParamIndexing
			Address moq.ParamIndexing
		}{
			Network: moq.ParamIndexByValue,
			Address: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ResolveIPAddr_genType type
func (m *MoqResolveIPAddr_genType) Mock() ResolveIPAddr_genType {
	return func(network, address string) (*net.IPAddr, error) {
		m.Scene.T.Helper()
		moq := &MoqResolveIPAddr_genType_mock{Moq: m}
		return moq.Fn(network, address)
	}
}

func (m *MoqResolveIPAddr_genType_mock) Fn(network, address string) (result1 *net.IPAddr, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqResolveIPAddr_genType_params{
		Network: network,
		Address: address,
	}
	var results *MoqResolveIPAddr_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(network, address)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(network, address)
	}
	return
}

func (m *MoqResolveIPAddr_genType) OnCall(network, address string) *MoqResolveIPAddr_genType_fnRecorder {
	return &MoqResolveIPAddr_genType_fnRecorder{
		Params: MoqResolveIPAddr_genType_params{
			Network: network,
			Address: address,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqResolveIPAddr_genType_fnRecorder) Any() *MoqResolveIPAddr_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqResolveIPAddr_genType_anyParams{Recorder: r}
}

func (a *MoqResolveIPAddr_genType_anyParams) Network() *MoqResolveIPAddr_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqResolveIPAddr_genType_anyParams) Address() *MoqResolveIPAddr_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqResolveIPAddr_genType_fnRecorder) Seq() *MoqResolveIPAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqResolveIPAddr_genType_fnRecorder) NoSeq() *MoqResolveIPAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqResolveIPAddr_genType_fnRecorder) ReturnResults(result1 *net.IPAddr, result2 error) *MoqResolveIPAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.IPAddr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResolveIPAddr_genType_doFn
		DoReturnFn MoqResolveIPAddr_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *net.IPAddr
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqResolveIPAddr_genType_fnRecorder) AndDo(fn MoqResolveIPAddr_genType_doFn) *MoqResolveIPAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqResolveIPAddr_genType_fnRecorder) DoReturnResults(fn MoqResolveIPAddr_genType_doReturnFn) *MoqResolveIPAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.IPAddr
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResolveIPAddr_genType_doFn
		DoReturnFn MoqResolveIPAddr_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqResolveIPAddr_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqResolveIPAddr_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqResolveIPAddr_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqResolveIPAddr_genType_paramsKey]*MoqResolveIPAddr_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqResolveIPAddr_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqResolveIPAddr_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqResolveIPAddr_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values *struct {
					Result1 *net.IPAddr
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqResolveIPAddr_genType_doFn
				DoReturnFn MoqResolveIPAddr_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqResolveIPAddr_genType) PrettyParams(params MoqResolveIPAddr_genType_params) string {
	return fmt.Sprintf("ResolveIPAddr_genType(%#v, %#v)", params.Network, params.Address)
}

func (m *MoqResolveIPAddr_genType) ParamsKey(params MoqResolveIPAddr_genType_params, anyParams uint64) MoqResolveIPAddr_genType_paramsKey {
	m.Scene.T.Helper()
	var networkUsed string
	var networkUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Network == moq.ParamIndexByValue {
			networkUsed = params.Network
		} else {
			networkUsedHash = hash.DeepHash(params.Network)
		}
	}
	var addressUsed string
	var addressUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Address == moq.ParamIndexByValue {
			addressUsed = params.Address
		} else {
			addressUsedHash = hash.DeepHash(params.Address)
		}
	}
	return MoqResolveIPAddr_genType_paramsKey{
		Params: struct{ Network, Address string }{
			Network: networkUsed,
			Address: addressUsed,
		},
		Hashes: struct{ Network, Address hash.Hash }{
			Network: networkUsedHash,
			Address: addressUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqResolveIPAddr_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqResolveIPAddr_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}