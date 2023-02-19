// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package atomic

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"unsafe"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// StorePointer_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type StorePointer_genType func(addr *unsafe.Pointer, val unsafe.Pointer)

// MoqStorePointer_genType holds the state of a moq of the StorePointer_genType
// type
type MoqStorePointer_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStorePointer_genType_mock

	ResultsByParams []MoqStorePointer_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Addr moq.ParamIndexing
			Val  moq.ParamIndexing
		}
	}
}

// MoqStorePointer_genType_mock isolates the mock interface of the
// StorePointer_genType type
type MoqStorePointer_genType_mock struct {
	Moq *MoqStorePointer_genType
}

// MoqStorePointer_genType_params holds the params of the StorePointer_genType
// type
type MoqStorePointer_genType_params struct {
	Addr *unsafe.Pointer
	Val  unsafe.Pointer
}

// MoqStorePointer_genType_paramsKey holds the map key params of the
// StorePointer_genType type
type MoqStorePointer_genType_paramsKey struct {
	Params struct {
		Addr *unsafe.Pointer
		Val  unsafe.Pointer
	}
	Hashes struct {
		Addr hash.Hash
		Val  hash.Hash
	}
}

// MoqStorePointer_genType_resultsByParams contains the results for a given set
// of parameters for the StorePointer_genType type
type MoqStorePointer_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStorePointer_genType_paramsKey]*MoqStorePointer_genType_results
}

// MoqStorePointer_genType_doFn defines the type of function needed when
// calling AndDo for the StorePointer_genType type
type MoqStorePointer_genType_doFn func(addr *unsafe.Pointer, val unsafe.Pointer)

// MoqStorePointer_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the StorePointer_genType type
type MoqStorePointer_genType_doReturnFn func(addr *unsafe.Pointer, val unsafe.Pointer)

// MoqStorePointer_genType_results holds the results of the
// StorePointer_genType type
type MoqStorePointer_genType_results struct {
	Params  MoqStorePointer_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStorePointer_genType_doFn
		DoReturnFn MoqStorePointer_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStorePointer_genType_fnRecorder routes recorded function calls to the
// MoqStorePointer_genType moq
type MoqStorePointer_genType_fnRecorder struct {
	Params    MoqStorePointer_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStorePointer_genType_results
	Moq       *MoqStorePointer_genType
}

// MoqStorePointer_genType_anyParams isolates the any params functions of the
// StorePointer_genType type
type MoqStorePointer_genType_anyParams struct {
	Recorder *MoqStorePointer_genType_fnRecorder
}

// NewMoqStorePointer_genType creates a new moq of the StorePointer_genType
// type
func NewMoqStorePointer_genType(scene *moq.Scene, config *moq.Config) *MoqStorePointer_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStorePointer_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStorePointer_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Addr moq.ParamIndexing
				Val  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Addr moq.ParamIndexing
			Val  moq.ParamIndexing
		}{
			Addr: moq.ParamIndexByHash,
			Val:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the StorePointer_genType type
func (m *MoqStorePointer_genType) Mock() StorePointer_genType {
	return func(addr *unsafe.Pointer, val unsafe.Pointer) {
		m.Scene.T.Helper()
		moq := &MoqStorePointer_genType_mock{Moq: m}
		moq.Fn(addr, val)
	}
}

func (m *MoqStorePointer_genType_mock) Fn(addr *unsafe.Pointer, val unsafe.Pointer) {
	m.Moq.Scene.T.Helper()
	params := MoqStorePointer_genType_params{
		Addr: addr,
		Val:  val,
	}
	var results *MoqStorePointer_genType_results
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
		result.DoFn(addr, val)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(addr, val)
	}
	return
}

func (m *MoqStorePointer_genType) OnCall(addr *unsafe.Pointer, val unsafe.Pointer) *MoqStorePointer_genType_fnRecorder {
	return &MoqStorePointer_genType_fnRecorder{
		Params: MoqStorePointer_genType_params{
			Addr: addr,
			Val:  val,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqStorePointer_genType_fnRecorder) Any() *MoqStorePointer_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqStorePointer_genType_anyParams{Recorder: r}
}

func (a *MoqStorePointer_genType_anyParams) Addr() *MoqStorePointer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqStorePointer_genType_anyParams) Val() *MoqStorePointer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqStorePointer_genType_fnRecorder) Seq() *MoqStorePointer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStorePointer_genType_fnRecorder) NoSeq() *MoqStorePointer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStorePointer_genType_fnRecorder) ReturnResults() *MoqStorePointer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStorePointer_genType_doFn
		DoReturnFn MoqStorePointer_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStorePointer_genType_fnRecorder) AndDo(fn MoqStorePointer_genType_doFn) *MoqStorePointer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStorePointer_genType_fnRecorder) DoReturnResults(fn MoqStorePointer_genType_doReturnFn) *MoqStorePointer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStorePointer_genType_doFn
		DoReturnFn MoqStorePointer_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStorePointer_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStorePointer_genType_resultsByParams
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
		results = &MoqStorePointer_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStorePointer_genType_paramsKey]*MoqStorePointer_genType_results{},
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
		r.Results = &MoqStorePointer_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStorePointer_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStorePointer_genType_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqStorePointer_genType_doFn
				DoReturnFn MoqStorePointer_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStorePointer_genType) PrettyParams(params MoqStorePointer_genType_params) string {
	return fmt.Sprintf("StorePointer_genType(%#v, %#v)", params.Addr, params.Val)
}

func (m *MoqStorePointer_genType) ParamsKey(params MoqStorePointer_genType_params, anyParams uint64) MoqStorePointer_genType_paramsKey {
	m.Scene.T.Helper()
	var addrUsed *unsafe.Pointer
	var addrUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	var valUsed unsafe.Pointer
	var valUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Val == moq.ParamIndexByValue {
			valUsed = params.Val
		} else {
			valUsedHash = hash.DeepHash(params.Val)
		}
	}
	return MoqStorePointer_genType_paramsKey{
		Params: struct {
			Addr *unsafe.Pointer
			Val  unsafe.Pointer
		}{
			Addr: addrUsed,
			Val:  valUsed,
		},
		Hashes: struct {
			Addr hash.Hash
			Val  hash.Hash
		}{
			Addr: addrUsedHash,
			Val:  valUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqStorePointer_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStorePointer_genType) AssertExpectationsMet() {
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
