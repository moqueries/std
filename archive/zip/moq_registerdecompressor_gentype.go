// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package zip

import (
	"archive/zip"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// RegisterDecompressor_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type RegisterDecompressor_genType func(method uint16, dcomp zip.Decompressor)

// MoqRegisterDecompressor_genType holds the state of a moq of the
// RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRegisterDecompressor_genType_mock

	ResultsByParams []MoqRegisterDecompressor_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Method moq.ParamIndexing
			Dcomp  moq.ParamIndexing
		}
	}
}

// MoqRegisterDecompressor_genType_mock isolates the mock interface of the
// RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType_mock struct {
	Moq *MoqRegisterDecompressor_genType
}

// MoqRegisterDecompressor_genType_params holds the params of the
// RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType_params struct {
	Method uint16
	Dcomp  zip.Decompressor
}

// MoqRegisterDecompressor_genType_paramsKey holds the map key params of the
// RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType_paramsKey struct {
	Params struct{ Method uint16 }
	Hashes struct {
		Method hash.Hash
		Dcomp  hash.Hash
	}
}

// MoqRegisterDecompressor_genType_resultsByParams contains the results for a
// given set of parameters for the RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRegisterDecompressor_genType_paramsKey]*MoqRegisterDecompressor_genType_results
}

// MoqRegisterDecompressor_genType_doFn defines the type of function needed
// when calling AndDo for the RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType_doFn func(method uint16, dcomp zip.Decompressor)

// MoqRegisterDecompressor_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the RegisterDecompressor_genType
// type
type MoqRegisterDecompressor_genType_doReturnFn func(method uint16, dcomp zip.Decompressor)

// MoqRegisterDecompressor_genType_results holds the results of the
// RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType_results struct {
	Params  MoqRegisterDecompressor_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqRegisterDecompressor_genType_doFn
		DoReturnFn MoqRegisterDecompressor_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRegisterDecompressor_genType_fnRecorder routes recorded function calls to
// the MoqRegisterDecompressor_genType moq
type MoqRegisterDecompressor_genType_fnRecorder struct {
	Params    MoqRegisterDecompressor_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRegisterDecompressor_genType_results
	Moq       *MoqRegisterDecompressor_genType
}

// MoqRegisterDecompressor_genType_anyParams isolates the any params functions
// of the RegisterDecompressor_genType type
type MoqRegisterDecompressor_genType_anyParams struct {
	Recorder *MoqRegisterDecompressor_genType_fnRecorder
}

// NewMoqRegisterDecompressor_genType creates a new moq of the
// RegisterDecompressor_genType type
func NewMoqRegisterDecompressor_genType(scene *moq.Scene, config *moq.Config) *MoqRegisterDecompressor_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRegisterDecompressor_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRegisterDecompressor_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Method moq.ParamIndexing
				Dcomp  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Method moq.ParamIndexing
			Dcomp  moq.ParamIndexing
		}{
			Method: moq.ParamIndexByValue,
			Dcomp:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the RegisterDecompressor_genType type
func (m *MoqRegisterDecompressor_genType) Mock() RegisterDecompressor_genType {
	return func(method uint16, dcomp zip.Decompressor) {
		m.Scene.T.Helper()
		moq := &MoqRegisterDecompressor_genType_mock{Moq: m}
		moq.Fn(method, dcomp)
	}
}

func (m *MoqRegisterDecompressor_genType_mock) Fn(method uint16, dcomp zip.Decompressor) {
	m.Moq.Scene.T.Helper()
	params := MoqRegisterDecompressor_genType_params{
		Method: method,
		Dcomp:  dcomp,
	}
	var results *MoqRegisterDecompressor_genType_results
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
		result.DoFn(method, dcomp)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(method, dcomp)
	}
	return
}

func (m *MoqRegisterDecompressor_genType) OnCall(method uint16, dcomp zip.Decompressor) *MoqRegisterDecompressor_genType_fnRecorder {
	return &MoqRegisterDecompressor_genType_fnRecorder{
		Params: MoqRegisterDecompressor_genType_params{
			Method: method,
			Dcomp:  dcomp,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) Any() *MoqRegisterDecompressor_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqRegisterDecompressor_genType_anyParams{Recorder: r}
}

func (a *MoqRegisterDecompressor_genType_anyParams) Method() *MoqRegisterDecompressor_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqRegisterDecompressor_genType_anyParams) Dcomp() *MoqRegisterDecompressor_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) Seq() *MoqRegisterDecompressor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) NoSeq() *MoqRegisterDecompressor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) ReturnResults() *MoqRegisterDecompressor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqRegisterDecompressor_genType_doFn
		DoReturnFn MoqRegisterDecompressor_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) AndDo(fn MoqRegisterDecompressor_genType_doFn) *MoqRegisterDecompressor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) DoReturnResults(fn MoqRegisterDecompressor_genType_doReturnFn) *MoqRegisterDecompressor_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqRegisterDecompressor_genType_doFn
		DoReturnFn MoqRegisterDecompressor_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRegisterDecompressor_genType_resultsByParams
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
		results = &MoqRegisterDecompressor_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRegisterDecompressor_genType_paramsKey]*MoqRegisterDecompressor_genType_results{},
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
		r.Results = &MoqRegisterDecompressor_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRegisterDecompressor_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRegisterDecompressor_genType_fnRecorder {
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
				DoFn       MoqRegisterDecompressor_genType_doFn
				DoReturnFn MoqRegisterDecompressor_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRegisterDecompressor_genType) PrettyParams(params MoqRegisterDecompressor_genType_params) string {
	return fmt.Sprintf("RegisterDecompressor_genType(%#v, %#v)", params.Method, params.Dcomp)
}

func (m *MoqRegisterDecompressor_genType) ParamsKey(params MoqRegisterDecompressor_genType_params, anyParams uint64) MoqRegisterDecompressor_genType_paramsKey {
	m.Scene.T.Helper()
	var methodUsed uint16
	var methodUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Method == moq.ParamIndexByValue {
			methodUsed = params.Method
		} else {
			methodUsedHash = hash.DeepHash(params.Method)
		}
	}
	var dcompUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Dcomp == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The dcomp parameter can't be indexed by value")
		}
		dcompUsedHash = hash.DeepHash(params.Dcomp)
	}
	return MoqRegisterDecompressor_genType_paramsKey{
		Params: struct{ Method uint16 }{
			Method: methodUsed,
		},
		Hashes: struct {
			Method hash.Hash
			Dcomp  hash.Hash
		}{
			Method: methodUsedHash,
			Dcomp:  dcompUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRegisterDecompressor_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRegisterDecompressor_genType) AssertExpectationsMet() {
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
