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

// DialUnix_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type DialUnix_genType func(network string, laddr, raddr *net.UnixAddr) (*net.UnixConn, error)

// MoqDialUnix_genType holds the state of a moq of the DialUnix_genType type
type MoqDialUnix_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDialUnix_genType_mock

	ResultsByParams []MoqDialUnix_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Network moq.ParamIndexing
			Laddr   moq.ParamIndexing
			Raddr   moq.ParamIndexing
		}
	}
}

// MoqDialUnix_genType_mock isolates the mock interface of the DialUnix_genType
// type
type MoqDialUnix_genType_mock struct {
	Moq *MoqDialUnix_genType
}

// MoqDialUnix_genType_params holds the params of the DialUnix_genType type
type MoqDialUnix_genType_params struct {
	Network      string
	Laddr, Raddr *net.UnixAddr
}

// MoqDialUnix_genType_paramsKey holds the map key params of the
// DialUnix_genType type
type MoqDialUnix_genType_paramsKey struct {
	Params struct {
		Network      string
		Laddr, Raddr *net.UnixAddr
	}
	Hashes struct {
		Network      hash.Hash
		Laddr, Raddr hash.Hash
	}
}

// MoqDialUnix_genType_resultsByParams contains the results for a given set of
// parameters for the DialUnix_genType type
type MoqDialUnix_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDialUnix_genType_paramsKey]*MoqDialUnix_genType_results
}

// MoqDialUnix_genType_doFn defines the type of function needed when calling
// AndDo for the DialUnix_genType type
type MoqDialUnix_genType_doFn func(network string, laddr, raddr *net.UnixAddr)

// MoqDialUnix_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the DialUnix_genType type
type MoqDialUnix_genType_doReturnFn func(network string, laddr, raddr *net.UnixAddr) (*net.UnixConn, error)

// MoqDialUnix_genType_results holds the results of the DialUnix_genType type
type MoqDialUnix_genType_results struct {
	Params  MoqDialUnix_genType_params
	Results []struct {
		Values *struct {
			Result1 *net.UnixConn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialUnix_genType_doFn
		DoReturnFn MoqDialUnix_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDialUnix_genType_fnRecorder routes recorded function calls to the
// MoqDialUnix_genType moq
type MoqDialUnix_genType_fnRecorder struct {
	Params    MoqDialUnix_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDialUnix_genType_results
	Moq       *MoqDialUnix_genType
}

// MoqDialUnix_genType_anyParams isolates the any params functions of the
// DialUnix_genType type
type MoqDialUnix_genType_anyParams struct {
	Recorder *MoqDialUnix_genType_fnRecorder
}

// NewMoqDialUnix_genType creates a new moq of the DialUnix_genType type
func NewMoqDialUnix_genType(scene *moq.Scene, config *moq.Config) *MoqDialUnix_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDialUnix_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDialUnix_genType_mock{},

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

// Mock returns the moq implementation of the DialUnix_genType type
func (m *MoqDialUnix_genType) Mock() DialUnix_genType {
	return func(network string, laddr, raddr *net.UnixAddr) (*net.UnixConn, error) {
		m.Scene.T.Helper()
		moq := &MoqDialUnix_genType_mock{Moq: m}
		return moq.Fn(network, laddr, raddr)
	}
}

func (m *MoqDialUnix_genType_mock) Fn(network string, laddr, raddr *net.UnixAddr) (result1 *net.UnixConn, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqDialUnix_genType_params{
		Network: network,
		Laddr:   laddr,
		Raddr:   raddr,
	}
	var results *MoqDialUnix_genType_results
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

func (m *MoqDialUnix_genType) OnCall(network string, laddr, raddr *net.UnixAddr) *MoqDialUnix_genType_fnRecorder {
	return &MoqDialUnix_genType_fnRecorder{
		Params: MoqDialUnix_genType_params{
			Network: network,
			Laddr:   laddr,
			Raddr:   raddr,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDialUnix_genType_fnRecorder) Any() *MoqDialUnix_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDialUnix_genType_anyParams{Recorder: r}
}

func (a *MoqDialUnix_genType_anyParams) Network() *MoqDialUnix_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqDialUnix_genType_anyParams) Laddr() *MoqDialUnix_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqDialUnix_genType_anyParams) Raddr() *MoqDialUnix_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqDialUnix_genType_fnRecorder) Seq() *MoqDialUnix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDialUnix_genType_fnRecorder) NoSeq() *MoqDialUnix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDialUnix_genType_fnRecorder) ReturnResults(result1 *net.UnixConn, result2 error) *MoqDialUnix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.UnixConn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialUnix_genType_doFn
		DoReturnFn MoqDialUnix_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *net.UnixConn
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDialUnix_genType_fnRecorder) AndDo(fn MoqDialUnix_genType_doFn) *MoqDialUnix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDialUnix_genType_fnRecorder) DoReturnResults(fn MoqDialUnix_genType_doReturnFn) *MoqDialUnix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *net.UnixConn
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDialUnix_genType_doFn
		DoReturnFn MoqDialUnix_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDialUnix_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDialUnix_genType_resultsByParams
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
		results = &MoqDialUnix_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDialUnix_genType_paramsKey]*MoqDialUnix_genType_results{},
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
		r.Results = &MoqDialUnix_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDialUnix_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDialUnix_genType_fnRecorder {
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
					Result1 *net.UnixConn
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqDialUnix_genType_doFn
				DoReturnFn MoqDialUnix_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDialUnix_genType) PrettyParams(params MoqDialUnix_genType_params) string {
	return fmt.Sprintf("DialUnix_genType(%#v, %#v, %#v)", params.Network, params.Laddr, params.Raddr)
}

func (m *MoqDialUnix_genType) ParamsKey(params MoqDialUnix_genType_params, anyParams uint64) MoqDialUnix_genType_paramsKey {
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
	var laddrUsed *net.UnixAddr
	var laddrUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Laddr == moq.ParamIndexByValue {
			laddrUsed = params.Laddr
		} else {
			laddrUsedHash = hash.DeepHash(params.Laddr)
		}
	}
	var raddrUsed *net.UnixAddr
	var raddrUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Raddr == moq.ParamIndexByValue {
			raddrUsed = params.Raddr
		} else {
			raddrUsedHash = hash.DeepHash(params.Raddr)
		}
	}
	return MoqDialUnix_genType_paramsKey{
		Params: struct {
			Network      string
			Laddr, Raddr *net.UnixAddr
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
func (m *MoqDialUnix_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDialUnix_genType) AssertExpectationsMet() {
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
