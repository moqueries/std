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

// DialIP_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type DialIP_genType func(network string, laddr, raddr *net.IPAddr) (*net.IPConn, error)

// MoqDialIP_genType holds the state of a moq of the DialIP_genType type
type MoqDialIP_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDialIP_genType_mock

	ResultsByParams []MoqDialIP_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Network moq.ParamIndexing
			Laddr   moq.ParamIndexing
			Raddr   moq.ParamIndexing
		}
	}
}

// MoqDialIP_genType_mock isolates the mock interface of the DialIP_genType
// type
type MoqDialIP_genType_mock struct {
	Moq *MoqDialIP_genType
}

// MoqDialIP_genType_params holds the params of the DialIP_genType type
type MoqDialIP_genType_params struct {
	Network      string
	Laddr, Raddr *net.IPAddr
}

// MoqDialIP_genType_paramsKey holds the map key params of the DialIP_genType
// type
type MoqDialIP_genType_paramsKey struct {
	Params struct {
		Network      string
		Laddr, Raddr *net.IPAddr
	}
	Hashes struct {
		Network      hash.Hash
		Laddr, Raddr hash.Hash
	}
}

// MoqDialIP_genType_resultsByParams contains the results for a given set of
// parameters for the DialIP_genType type
type MoqDialIP_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDialIP_genType_paramsKey]*MoqDialIP_genType_results
}

// MoqDialIP_genType_doFn defines the type of function needed when calling
// AndDo for the DialIP_genType type
type MoqDialIP_genType_doFn func(network string, laddr, raddr *net.IPAddr)

// MoqDialIP_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the DialIP_genType type
type MoqDialIP_genType_doReturnFn func(network string, laddr, raddr *net.IPAddr) (*net.IPConn, error)

// MoqDialIP_genType_results holds the results of the DialIP_genType type
type MoqDialIP_genType_results struct {
	Params  MoqDialIP_genType_params
	Results []struct {
		Values *struct {
			Result1 *net.IPConn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialIP_genType_doFn
		DoReturnFn MoqDialIP_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDialIP_genType_fnRecorder routes recorded function calls to the
// MoqDialIP_genType moq
type MoqDialIP_genType_fnRecorder struct {
	Params    MoqDialIP_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDialIP_genType_results
	Moq       *MoqDialIP_genType
}

// MoqDialIP_genType_anyParams isolates the any params functions of the
// DialIP_genType type
type MoqDialIP_genType_anyParams struct {
	Recorder *MoqDialIP_genType_fnRecorder
}

// NewMoqDialIP_genType creates a new moq of the DialIP_genType type
func NewMoqDialIP_genType(scene *moq.Scene, config *moq.Config) *MoqDialIP_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDialIP_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDialIP_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Network moq.ParamIndexing
				Laddr   moq.ParamIndexing
				Raddr   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Network moq.ParamIndexing
			Laddr   moq.ParamIndexing
			Raddr   moq.ParamIndexing
		}{
			Network: moq.ParamIndexByValue,
			Laddr:   moq.ParamIndexByHash,
			Raddr:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the DialIP_genType type
func (m *MoqDialIP_genType) Mock() DialIP_genType {
	return func(network string, laddr, raddr *net.IPAddr) (*net.IPConn, error) {
		m.Scene.T.Helper()
		moq := &MoqDialIP_genType_mock{Moq: m}
		return moq.Fn(network, laddr, raddr)
	}
}

func (m *MoqDialIP_genType_mock) Fn(network string, laddr, raddr *net.IPAddr) (result1 *net.IPConn, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqDialIP_genType_params{
		Network: network,
		Laddr:   laddr,
		Raddr:   raddr,
	}
	var results *MoqDialIP_genType_results
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
		result.DoFn(network, laddr, raddr)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(network, laddr, raddr)
	}
	return
}

func (m *MoqDialIP_genType) OnCall(network string, laddr, raddr *net.IPAddr) *MoqDialIP_genType_fnRecorder {
	return &MoqDialIP_genType_fnRecorder{
		Params: MoqDialIP_genType_params{
			Network: network,
			Laddr:   laddr,
			Raddr:   raddr,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDialIP_genType_fnRecorder) Any() *MoqDialIP_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDialIP_genType_anyParams{Recorder: r}
}

func (a *MoqDialIP_genType_anyParams) Network() *MoqDialIP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqDialIP_genType_anyParams) Laddr() *MoqDialIP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqDialIP_genType_anyParams) Raddr() *MoqDialIP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqDialIP_genType_fnRecorder) Seq() *MoqDialIP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDialIP_genType_fnRecorder) NoSeq() *MoqDialIP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDialIP_genType_fnRecorder) ReturnResults(result1 *net.IPConn, result2 error) *MoqDialIP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.IPConn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialIP_genType_doFn
		DoReturnFn MoqDialIP_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *net.IPConn
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDialIP_genType_fnRecorder) AndDo(fn MoqDialIP_genType_doFn) *MoqDialIP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDialIP_genType_fnRecorder) DoReturnResults(fn MoqDialIP_genType_doReturnFn) *MoqDialIP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.IPConn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialIP_genType_doFn
		DoReturnFn MoqDialIP_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDialIP_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDialIP_genType_resultsByParams
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
		results = &MoqDialIP_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDialIP_genType_paramsKey]*MoqDialIP_genType_results{},
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
		r.Results = &MoqDialIP_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDialIP_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDialIP_genType_fnRecorder {
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
					Result1 *net.IPConn
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqDialIP_genType_doFn
				DoReturnFn MoqDialIP_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDialIP_genType) PrettyParams(params MoqDialIP_genType_params) string {
	return fmt.Sprintf("DialIP_genType(%#v, %#v, %#v)", params.Network, params.Laddr, params.Raddr)
}

func (m *MoqDialIP_genType) ParamsKey(params MoqDialIP_genType_params, anyParams uint64) MoqDialIP_genType_paramsKey {
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
	var laddrUsed *net.IPAddr
	var laddrUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Laddr == moq.ParamIndexByValue {
			laddrUsed = params.Laddr
		} else {
			laddrUsedHash = hash.DeepHash(params.Laddr)
		}
	}
	var raddrUsed *net.IPAddr
	var raddrUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Raddr == moq.ParamIndexByValue {
			raddrUsed = params.Raddr
		} else {
			raddrUsedHash = hash.DeepHash(params.Raddr)
		}
	}
	return MoqDialIP_genType_paramsKey{
		Params: struct {
			Network      string
			Laddr, Raddr *net.IPAddr
		}{
			Network: networkUsed,
			Laddr:   laddrUsed,
			Raddr:   raddrUsed,
		},
		Hashes: struct {
			Network      hash.Hash
			Laddr, Raddr hash.Hash
		}{
			Network: networkUsedHash,
			Laddr:   laddrUsedHash,
			Raddr:   raddrUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDialIP_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDialIP_genType) AssertExpectationsMet() {
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
