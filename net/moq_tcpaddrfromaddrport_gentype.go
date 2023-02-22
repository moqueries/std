// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"net"
	"net/netip"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// TCPAddrFromAddrPort_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type TCPAddrFromAddrPort_genType func(addr netip.AddrPort) *net.TCPAddr

// MoqTCPAddrFromAddrPort_genType holds the state of a moq of the
// TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqTCPAddrFromAddrPort_genType_mock

	ResultsByParams []MoqTCPAddrFromAddrPort_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Addr moq.ParamIndexing
		}
	}
}

// MoqTCPAddrFromAddrPort_genType_mock isolates the mock interface of the
// TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_mock struct {
	Moq *MoqTCPAddrFromAddrPort_genType
}

// MoqTCPAddrFromAddrPort_genType_params holds the params of the
// TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_params struct{ Addr netip.AddrPort }

// MoqTCPAddrFromAddrPort_genType_paramsKey holds the map key params of the
// TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_paramsKey struct {
	Params struct{ Addr netip.AddrPort }
	Hashes struct{ Addr hash.Hash }
}

// MoqTCPAddrFromAddrPort_genType_resultsByParams contains the results for a
// given set of parameters for the TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTCPAddrFromAddrPort_genType_paramsKey]*MoqTCPAddrFromAddrPort_genType_results
}

// MoqTCPAddrFromAddrPort_genType_doFn defines the type of function needed when
// calling AndDo for the TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_doFn func(addr netip.AddrPort)

// MoqTCPAddrFromAddrPort_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_doReturnFn func(addr netip.AddrPort) *net.TCPAddr

// MoqTCPAddrFromAddrPort_genType_results holds the results of the
// TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_results struct {
	Params  MoqTCPAddrFromAddrPort_genType_params
	Results []struct {
		Values *struct {
			Result1 *net.TCPAddr
		}
		Sequence   uint32
		DoFn       MoqTCPAddrFromAddrPort_genType_doFn
		DoReturnFn MoqTCPAddrFromAddrPort_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTCPAddrFromAddrPort_genType_fnRecorder routes recorded function calls to
// the MoqTCPAddrFromAddrPort_genType moq
type MoqTCPAddrFromAddrPort_genType_fnRecorder struct {
	Params    MoqTCPAddrFromAddrPort_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTCPAddrFromAddrPort_genType_results
	Moq       *MoqTCPAddrFromAddrPort_genType
}

// MoqTCPAddrFromAddrPort_genType_anyParams isolates the any params functions
// of the TCPAddrFromAddrPort_genType type
type MoqTCPAddrFromAddrPort_genType_anyParams struct {
	Recorder *MoqTCPAddrFromAddrPort_genType_fnRecorder
}

// NewMoqTCPAddrFromAddrPort_genType creates a new moq of the
// TCPAddrFromAddrPort_genType type
func NewMoqTCPAddrFromAddrPort_genType(scene *moq.Scene, config *moq.Config) *MoqTCPAddrFromAddrPort_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqTCPAddrFromAddrPort_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqTCPAddrFromAddrPort_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Addr moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Addr moq.ParamIndexing
		}{
			Addr: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the TCPAddrFromAddrPort_genType type
func (m *MoqTCPAddrFromAddrPort_genType) Mock() TCPAddrFromAddrPort_genType {
	return func(addr netip.AddrPort) *net.TCPAddr {
		m.Scene.T.Helper()
		moq := &MoqTCPAddrFromAddrPort_genType_mock{Moq: m}
		return moq.Fn(addr)
	}
}

func (m *MoqTCPAddrFromAddrPort_genType_mock) Fn(addr netip.AddrPort) (result1 *net.TCPAddr) {
	m.Moq.Scene.T.Helper()
	params := MoqTCPAddrFromAddrPort_genType_params{
		Addr: addr,
	}
	var results *MoqTCPAddrFromAddrPort_genType_results
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
		result.DoFn(addr)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(addr)
	}
	return
}

func (m *MoqTCPAddrFromAddrPort_genType) OnCall(addr netip.AddrPort) *MoqTCPAddrFromAddrPort_genType_fnRecorder {
	return &MoqTCPAddrFromAddrPort_genType_fnRecorder{
		Params: MoqTCPAddrFromAddrPort_genType_params{
			Addr: addr,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) Any() *MoqTCPAddrFromAddrPort_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqTCPAddrFromAddrPort_genType_anyParams{Recorder: r}
}

func (a *MoqTCPAddrFromAddrPort_genType_anyParams) Addr() *MoqTCPAddrFromAddrPort_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) Seq() *MoqTCPAddrFromAddrPort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) NoSeq() *MoqTCPAddrFromAddrPort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) ReturnResults(result1 *net.TCPAddr) *MoqTCPAddrFromAddrPort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.TCPAddr
		}
		Sequence   uint32
		DoFn       MoqTCPAddrFromAddrPort_genType_doFn
		DoReturnFn MoqTCPAddrFromAddrPort_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *net.TCPAddr
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) AndDo(fn MoqTCPAddrFromAddrPort_genType_doFn) *MoqTCPAddrFromAddrPort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) DoReturnResults(fn MoqTCPAddrFromAddrPort_genType_doReturnFn) *MoqTCPAddrFromAddrPort_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.TCPAddr
		}
		Sequence   uint32
		DoFn       MoqTCPAddrFromAddrPort_genType_doFn
		DoReturnFn MoqTCPAddrFromAddrPort_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTCPAddrFromAddrPort_genType_resultsByParams
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
		results = &MoqTCPAddrFromAddrPort_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTCPAddrFromAddrPort_genType_paramsKey]*MoqTCPAddrFromAddrPort_genType_results{},
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
		r.Results = &MoqTCPAddrFromAddrPort_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTCPAddrFromAddrPort_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTCPAddrFromAddrPort_genType_fnRecorder {
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
					Result1 *net.TCPAddr
				}
				Sequence   uint32
				DoFn       MoqTCPAddrFromAddrPort_genType_doFn
				DoReturnFn MoqTCPAddrFromAddrPort_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTCPAddrFromAddrPort_genType) PrettyParams(params MoqTCPAddrFromAddrPort_genType_params) string {
	return fmt.Sprintf("TCPAddrFromAddrPort_genType(%#v)", params.Addr)
}

func (m *MoqTCPAddrFromAddrPort_genType) ParamsKey(params MoqTCPAddrFromAddrPort_genType_params, anyParams uint64) MoqTCPAddrFromAddrPort_genType_paramsKey {
	m.Scene.T.Helper()
	var addrUsed netip.AddrPort
	var addrUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	return MoqTCPAddrFromAddrPort_genType_paramsKey{
		Params: struct{ Addr netip.AddrPort }{
			Addr: addrUsed,
		},
		Hashes: struct{ Addr hash.Hash }{
			Addr: addrUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqTCPAddrFromAddrPort_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqTCPAddrFromAddrPort_genType) AssertExpectationsMet() {
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
