// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package netip

import (
	"fmt"
	"math/bits"
	"net/netip"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// AddrFromSlice_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type AddrFromSlice_genType func(slice []byte) (ip netip.Addr, ok bool)

// MoqAddrFromSlice_genType holds the state of a moq of the
// AddrFromSlice_genType type
type MoqAddrFromSlice_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAddrFromSlice_genType_mock

	ResultsByParams []MoqAddrFromSlice_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Slice moq.ParamIndexing
		}
	}
}

// MoqAddrFromSlice_genType_mock isolates the mock interface of the
// AddrFromSlice_genType type
type MoqAddrFromSlice_genType_mock struct {
	Moq *MoqAddrFromSlice_genType
}

// MoqAddrFromSlice_genType_params holds the params of the
// AddrFromSlice_genType type
type MoqAddrFromSlice_genType_params struct{ Slice []byte }

// MoqAddrFromSlice_genType_paramsKey holds the map key params of the
// AddrFromSlice_genType type
type MoqAddrFromSlice_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Slice hash.Hash }
}

// MoqAddrFromSlice_genType_resultsByParams contains the results for a given
// set of parameters for the AddrFromSlice_genType type
type MoqAddrFromSlice_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAddrFromSlice_genType_paramsKey]*MoqAddrFromSlice_genType_results
}

// MoqAddrFromSlice_genType_doFn defines the type of function needed when
// calling AndDo for the AddrFromSlice_genType type
type MoqAddrFromSlice_genType_doFn func(slice []byte)

// MoqAddrFromSlice_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the AddrFromSlice_genType type
type MoqAddrFromSlice_genType_doReturnFn func(slice []byte) (ip netip.Addr, ok bool)

// MoqAddrFromSlice_genType_results holds the results of the
// AddrFromSlice_genType type
type MoqAddrFromSlice_genType_results struct {
	Params  MoqAddrFromSlice_genType_params
	Results []struct {
		Values *struct {
			Ip netip.Addr
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqAddrFromSlice_genType_doFn
		DoReturnFn MoqAddrFromSlice_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAddrFromSlice_genType_fnRecorder routes recorded function calls to the
// MoqAddrFromSlice_genType moq
type MoqAddrFromSlice_genType_fnRecorder struct {
	Params    MoqAddrFromSlice_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAddrFromSlice_genType_results
	Moq       *MoqAddrFromSlice_genType
}

// MoqAddrFromSlice_genType_anyParams isolates the any params functions of the
// AddrFromSlice_genType type
type MoqAddrFromSlice_genType_anyParams struct {
	Recorder *MoqAddrFromSlice_genType_fnRecorder
}

// NewMoqAddrFromSlice_genType creates a new moq of the AddrFromSlice_genType
// type
func NewMoqAddrFromSlice_genType(scene *moq.Scene, config *moq.Config) *MoqAddrFromSlice_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAddrFromSlice_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAddrFromSlice_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Slice moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Slice moq.ParamIndexing
		}{
			Slice: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the AddrFromSlice_genType type
func (m *MoqAddrFromSlice_genType) Mock() AddrFromSlice_genType {
	return func(slice []byte) (_ netip.Addr, _ bool) {
		m.Scene.T.Helper()
		moq := &MoqAddrFromSlice_genType_mock{Moq: m}
		return moq.Fn(slice)
	}
}

func (m *MoqAddrFromSlice_genType_mock) Fn(slice []byte) (ip netip.Addr, ok bool) {
	m.Moq.Scene.T.Helper()
	params := MoqAddrFromSlice_genType_params{
		Slice: slice,
	}
	var results *MoqAddrFromSlice_genType_results
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
		result.DoFn(slice)
	}

	if result.Values != nil {
		ip = result.Values.Ip
		ok = result.Values.Ok
	}
	if result.DoReturnFn != nil {
		ip, ok = result.DoReturnFn(slice)
	}
	return
}

func (m *MoqAddrFromSlice_genType) OnCall(slice []byte) *MoqAddrFromSlice_genType_fnRecorder {
	return &MoqAddrFromSlice_genType_fnRecorder{
		Params: MoqAddrFromSlice_genType_params{
			Slice: slice,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAddrFromSlice_genType_fnRecorder) Any() *MoqAddrFromSlice_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAddrFromSlice_genType_anyParams{Recorder: r}
}

func (a *MoqAddrFromSlice_genType_anyParams) Slice() *MoqAddrFromSlice_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqAddrFromSlice_genType_fnRecorder) Seq() *MoqAddrFromSlice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAddrFromSlice_genType_fnRecorder) NoSeq() *MoqAddrFromSlice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAddrFromSlice_genType_fnRecorder) ReturnResults(ip netip.Addr, ok bool) *MoqAddrFromSlice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Ip netip.Addr
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqAddrFromSlice_genType_doFn
		DoReturnFn MoqAddrFromSlice_genType_doReturnFn
	}{
		Values: &struct {
			Ip netip.Addr
			Ok bool
		}{
			Ip: ip,
			Ok: ok,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAddrFromSlice_genType_fnRecorder) AndDo(fn MoqAddrFromSlice_genType_doFn) *MoqAddrFromSlice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAddrFromSlice_genType_fnRecorder) DoReturnResults(fn MoqAddrFromSlice_genType_doReturnFn) *MoqAddrFromSlice_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Ip netip.Addr
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqAddrFromSlice_genType_doFn
		DoReturnFn MoqAddrFromSlice_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAddrFromSlice_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAddrFromSlice_genType_resultsByParams
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
		results = &MoqAddrFromSlice_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAddrFromSlice_genType_paramsKey]*MoqAddrFromSlice_genType_results{},
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
		r.Results = &MoqAddrFromSlice_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAddrFromSlice_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAddrFromSlice_genType_fnRecorder {
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
					Ip netip.Addr
					Ok bool
				}
				Sequence   uint32
				DoFn       MoqAddrFromSlice_genType_doFn
				DoReturnFn MoqAddrFromSlice_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAddrFromSlice_genType) PrettyParams(params MoqAddrFromSlice_genType_params) string {
	return fmt.Sprintf("AddrFromSlice_genType(%#v)", params.Slice)
}

func (m *MoqAddrFromSlice_genType) ParamsKey(params MoqAddrFromSlice_genType_params, anyParams uint64) MoqAddrFromSlice_genType_paramsKey {
	m.Scene.T.Helper()
	var sliceUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Slice == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The slice parameter can't be indexed by value")
		}
		sliceUsedHash = hash.DeepHash(params.Slice)
	}
	return MoqAddrFromSlice_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Slice hash.Hash }{
			Slice: sliceUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAddrFromSlice_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAddrFromSlice_genType) AssertExpectationsMet() {
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