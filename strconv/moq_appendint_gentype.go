// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package strconv

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// AppendInt_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type AppendInt_genType func(dst []byte, i int64, base int) []byte

// MoqAppendInt_genType holds the state of a moq of the AppendInt_genType type
type MoqAppendInt_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAppendInt_genType_mock

	ResultsByParams []MoqAppendInt_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dst    moq.ParamIndexing
			Param2 moq.ParamIndexing
			Base   moq.ParamIndexing
		}
	}
}

// MoqAppendInt_genType_mock isolates the mock interface of the
// AppendInt_genType type
type MoqAppendInt_genType_mock struct {
	Moq *MoqAppendInt_genType
}

// MoqAppendInt_genType_params holds the params of the AppendInt_genType type
type MoqAppendInt_genType_params struct {
	Dst    []byte
	Param2 int64
	Base   int
}

// MoqAppendInt_genType_paramsKey holds the map key params of the
// AppendInt_genType type
type MoqAppendInt_genType_paramsKey struct {
	Params struct {
		Param2 int64
		Base   int
	}
	Hashes struct {
		Dst    hash.Hash
		Param2 hash.Hash
		Base   hash.Hash
	}
}

// MoqAppendInt_genType_resultsByParams contains the results for a given set of
// parameters for the AppendInt_genType type
type MoqAppendInt_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAppendInt_genType_paramsKey]*MoqAppendInt_genType_results
}

// MoqAppendInt_genType_doFn defines the type of function needed when calling
// AndDo for the AppendInt_genType type
type MoqAppendInt_genType_doFn func(dst []byte, i int64, base int)

// MoqAppendInt_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the AppendInt_genType type
type MoqAppendInt_genType_doReturnFn func(dst []byte, i int64, base int) []byte

// MoqAppendInt_genType_results holds the results of the AppendInt_genType type
type MoqAppendInt_genType_results struct {
	Params  MoqAppendInt_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqAppendInt_genType_doFn
		DoReturnFn MoqAppendInt_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAppendInt_genType_fnRecorder routes recorded function calls to the
// MoqAppendInt_genType moq
type MoqAppendInt_genType_fnRecorder struct {
	Params    MoqAppendInt_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAppendInt_genType_results
	Moq       *MoqAppendInt_genType
}

// MoqAppendInt_genType_anyParams isolates the any params functions of the
// AppendInt_genType type
type MoqAppendInt_genType_anyParams struct {
	Recorder *MoqAppendInt_genType_fnRecorder
}

// NewMoqAppendInt_genType creates a new moq of the AppendInt_genType type
func NewMoqAppendInt_genType(scene *moq.Scene, config *moq.Config) *MoqAppendInt_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAppendInt_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAppendInt_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dst    moq.ParamIndexing
				Param2 moq.ParamIndexing
				Base   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dst    moq.ParamIndexing
			Param2 moq.ParamIndexing
			Base   moq.ParamIndexing
		}{
			Dst:    moq.ParamIndexByHash,
			Param2: moq.ParamIndexByValue,
			Base:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the AppendInt_genType type
func (m *MoqAppendInt_genType) Mock() AppendInt_genType {
	return func(dst []byte, param2 int64, base int) []byte {
		m.Scene.T.Helper()
		moq := &MoqAppendInt_genType_mock{Moq: m}
		return moq.Fn(dst, param2, base)
	}
}

func (m *MoqAppendInt_genType_mock) Fn(dst []byte, param2 int64, base int) (result1 []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqAppendInt_genType_params{
		Dst:    dst,
		Param2: param2,
		Base:   base,
	}
	var results *MoqAppendInt_genType_results
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
		result.DoFn(dst, param2, base)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(dst, param2, base)
	}
	return
}

func (m *MoqAppendInt_genType) OnCall(dst []byte, param2 int64, base int) *MoqAppendInt_genType_fnRecorder {
	return &MoqAppendInt_genType_fnRecorder{
		Params: MoqAppendInt_genType_params{
			Dst:    dst,
			Param2: param2,
			Base:   base,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAppendInt_genType_fnRecorder) Any() *MoqAppendInt_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAppendInt_genType_anyParams{Recorder: r}
}

func (a *MoqAppendInt_genType_anyParams) Dst() *MoqAppendInt_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqAppendInt_genType_anyParams) Param2() *MoqAppendInt_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqAppendInt_genType_anyParams) Base() *MoqAppendInt_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqAppendInt_genType_fnRecorder) Seq() *MoqAppendInt_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAppendInt_genType_fnRecorder) NoSeq() *MoqAppendInt_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAppendInt_genType_fnRecorder) ReturnResults(result1 []byte) *MoqAppendInt_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqAppendInt_genType_doFn
		DoReturnFn MoqAppendInt_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAppendInt_genType_fnRecorder) AndDo(fn MoqAppendInt_genType_doFn) *MoqAppendInt_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAppendInt_genType_fnRecorder) DoReturnResults(fn MoqAppendInt_genType_doReturnFn) *MoqAppendInt_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqAppendInt_genType_doFn
		DoReturnFn MoqAppendInt_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAppendInt_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAppendInt_genType_resultsByParams
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
		results = &MoqAppendInt_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAppendInt_genType_paramsKey]*MoqAppendInt_genType_results{},
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
		r.Results = &MoqAppendInt_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAppendInt_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAppendInt_genType_fnRecorder {
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
					Result1 []byte
				}
				Sequence   uint32
				DoFn       MoqAppendInt_genType_doFn
				DoReturnFn MoqAppendInt_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAppendInt_genType) PrettyParams(params MoqAppendInt_genType_params) string {
	return fmt.Sprintf("AppendInt_genType(%#v, %#v, %#v)", params.Dst, params.Param2, params.Base)
}

func (m *MoqAppendInt_genType) ParamsKey(params MoqAppendInt_genType_params, anyParams uint64) MoqAppendInt_genType_paramsKey {
	m.Scene.T.Helper()
	var dstUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Dst == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The dst parameter can't be indexed by value")
		}
		dstUsedHash = hash.DeepHash(params.Dst)
	}
	var param2Used int64
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	var baseUsed int
	var baseUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Base == moq.ParamIndexByValue {
			baseUsed = params.Base
		} else {
			baseUsedHash = hash.DeepHash(params.Base)
		}
	}
	return MoqAppendInt_genType_paramsKey{
		Params: struct {
			Param2 int64
			Base   int
		}{
			Param2: param2Used,
			Base:   baseUsed,
		},
		Hashes: struct {
			Dst    hash.Hash
			Param2 hash.Hash
			Base   hash.Hash
		}{
			Dst:    dstUsedHash,
			Param2: param2UsedHash,
			Base:   baseUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAppendInt_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAppendInt_genType) AssertExpectationsMet() {
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