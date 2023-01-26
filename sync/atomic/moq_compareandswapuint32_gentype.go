// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package atomic

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// CompareAndSwapUint32_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type CompareAndSwapUint32_genType func(addr *uint32, old, new uint32) (swapped bool)

// MoqCompareAndSwapUint32_genType holds the state of a moq of the
// CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCompareAndSwapUint32_genType_mock

	ResultsByParams []MoqCompareAndSwapUint32_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Addr moq.ParamIndexing
			Old  moq.ParamIndexing
			New  moq.ParamIndexing
		}
	}
}

// MoqCompareAndSwapUint32_genType_mock isolates the mock interface of the
// CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType_mock struct {
	Moq *MoqCompareAndSwapUint32_genType
}

// MoqCompareAndSwapUint32_genType_params holds the params of the
// CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType_params struct {
	Addr     *uint32
	Old, New uint32
}

// MoqCompareAndSwapUint32_genType_paramsKey holds the map key params of the
// CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType_paramsKey struct {
	Params struct {
		Addr     *uint32
		Old, New uint32
	}
	Hashes struct {
		Addr     hash.Hash
		Old, New hash.Hash
	}
}

// MoqCompareAndSwapUint32_genType_resultsByParams contains the results for a
// given set of parameters for the CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCompareAndSwapUint32_genType_paramsKey]*MoqCompareAndSwapUint32_genType_results
}

// MoqCompareAndSwapUint32_genType_doFn defines the type of function needed
// when calling AndDo for the CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType_doFn func(addr *uint32, old, new uint32)

// MoqCompareAndSwapUint32_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the CompareAndSwapUint32_genType
// type
type MoqCompareAndSwapUint32_genType_doReturnFn func(addr *uint32, old, new uint32) (swapped bool)

// MoqCompareAndSwapUint32_genType_results holds the results of the
// CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType_results struct {
	Params  MoqCompareAndSwapUint32_genType_params
	Results []struct {
		Values     *struct{ Swapped bool }
		Sequence   uint32
		DoFn       MoqCompareAndSwapUint32_genType_doFn
		DoReturnFn MoqCompareAndSwapUint32_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCompareAndSwapUint32_genType_fnRecorder routes recorded function calls to
// the MoqCompareAndSwapUint32_genType moq
type MoqCompareAndSwapUint32_genType_fnRecorder struct {
	Params    MoqCompareAndSwapUint32_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCompareAndSwapUint32_genType_results
	Moq       *MoqCompareAndSwapUint32_genType
}

// MoqCompareAndSwapUint32_genType_anyParams isolates the any params functions
// of the CompareAndSwapUint32_genType type
type MoqCompareAndSwapUint32_genType_anyParams struct {
	Recorder *MoqCompareAndSwapUint32_genType_fnRecorder
}

// NewMoqCompareAndSwapUint32_genType creates a new moq of the
// CompareAndSwapUint32_genType type
func NewMoqCompareAndSwapUint32_genType(scene *moq.Scene, config *moq.Config) *MoqCompareAndSwapUint32_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCompareAndSwapUint32_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCompareAndSwapUint32_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Addr moq.ParamIndexing
				Old  moq.ParamIndexing
				New  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Addr moq.ParamIndexing
			Old  moq.ParamIndexing
			New  moq.ParamIndexing
		}{
			Addr: moq.ParamIndexByHash,
			Old:  moq.ParamIndexByValue,
			New:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the CompareAndSwapUint32_genType type
func (m *MoqCompareAndSwapUint32_genType) Mock() CompareAndSwapUint32_genType {
	return func(addr *uint32, old, new uint32) (_ bool) {
		m.Scene.T.Helper()
		moq := &MoqCompareAndSwapUint32_genType_mock{Moq: m}
		return moq.Fn(addr, old, new)
	}
}

func (m *MoqCompareAndSwapUint32_genType_mock) Fn(addr *uint32, old, new uint32) (swapped bool) {
	m.Moq.Scene.T.Helper()
	params := MoqCompareAndSwapUint32_genType_params{
		Addr: addr,
		Old:  old,
		New:  new,
	}
	var results *MoqCompareAndSwapUint32_genType_results
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
		result.DoFn(addr, old, new)
	}

	if result.Values != nil {
		swapped = result.Values.Swapped
	}
	if result.DoReturnFn != nil {
		swapped = result.DoReturnFn(addr, old, new)
	}
	return
}

func (m *MoqCompareAndSwapUint32_genType) OnCall(addr *uint32, old, new uint32) *MoqCompareAndSwapUint32_genType_fnRecorder {
	return &MoqCompareAndSwapUint32_genType_fnRecorder{
		Params: MoqCompareAndSwapUint32_genType_params{
			Addr: addr,
			Old:  old,
			New:  new,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) Any() *MoqCompareAndSwapUint32_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCompareAndSwapUint32_genType_anyParams{Recorder: r}
}

func (a *MoqCompareAndSwapUint32_genType_anyParams) Addr() *MoqCompareAndSwapUint32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCompareAndSwapUint32_genType_anyParams) Old() *MoqCompareAndSwapUint32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqCompareAndSwapUint32_genType_anyParams) New() *MoqCompareAndSwapUint32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) Seq() *MoqCompareAndSwapUint32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) NoSeq() *MoqCompareAndSwapUint32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) ReturnResults(swapped bool) *MoqCompareAndSwapUint32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Swapped bool }
		Sequence   uint32
		DoFn       MoqCompareAndSwapUint32_genType_doFn
		DoReturnFn MoqCompareAndSwapUint32_genType_doReturnFn
	}{
		Values: &struct{ Swapped bool }{
			Swapped: swapped,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) AndDo(fn MoqCompareAndSwapUint32_genType_doFn) *MoqCompareAndSwapUint32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) DoReturnResults(fn MoqCompareAndSwapUint32_genType_doReturnFn) *MoqCompareAndSwapUint32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Swapped bool }
		Sequence   uint32
		DoFn       MoqCompareAndSwapUint32_genType_doFn
		DoReturnFn MoqCompareAndSwapUint32_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCompareAndSwapUint32_genType_resultsByParams
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
		results = &MoqCompareAndSwapUint32_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCompareAndSwapUint32_genType_paramsKey]*MoqCompareAndSwapUint32_genType_results{},
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
		r.Results = &MoqCompareAndSwapUint32_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCompareAndSwapUint32_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCompareAndSwapUint32_genType_fnRecorder {
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
				Values     *struct{ Swapped bool }
				Sequence   uint32
				DoFn       MoqCompareAndSwapUint32_genType_doFn
				DoReturnFn MoqCompareAndSwapUint32_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCompareAndSwapUint32_genType) PrettyParams(params MoqCompareAndSwapUint32_genType_params) string {
	return fmt.Sprintf("CompareAndSwapUint32_genType(%#v, %#v, %#v)", params.Addr, params.Old, params.New)
}

func (m *MoqCompareAndSwapUint32_genType) ParamsKey(params MoqCompareAndSwapUint32_genType_params, anyParams uint64) MoqCompareAndSwapUint32_genType_paramsKey {
	m.Scene.T.Helper()
	var addrUsed *uint32
	var addrUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	var oldUsed uint32
	var oldUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Old == moq.ParamIndexByValue {
			oldUsed = params.Old
		} else {
			oldUsedHash = hash.DeepHash(params.Old)
		}
	}
	var newUsed uint32
	var newUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.New == moq.ParamIndexByValue {
			newUsed = params.New
		} else {
			newUsedHash = hash.DeepHash(params.New)
		}
	}
	return MoqCompareAndSwapUint32_genType_paramsKey{
		Params: struct {
			Addr     *uint32
			Old, New uint32
		}{
			Addr: addrUsed,
			Old:  oldUsed,
			New:  newUsed,
		},
		Hashes: struct {
			Addr     hash.Hash
			Old, New hash.Hash
		}{
			Addr: addrUsedHash,
			Old:  oldUsedHash,
			New:  newUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCompareAndSwapUint32_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCompareAndSwapUint32_genType) AssertExpectationsMet() {
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