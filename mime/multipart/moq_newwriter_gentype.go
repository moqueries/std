// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package multipart

import (
	"fmt"
	"io"
	"math/bits"
	"mime/multipart"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewWriter_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewWriter_genType func(w io.Writer) *multipart.Writer

// MoqNewWriter_genType holds the state of a moq of the NewWriter_genType type
type MoqNewWriter_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewWriter_genType_mock

	ResultsByParams []MoqNewWriter_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			W moq.ParamIndexing
		}
	}
}

// MoqNewWriter_genType_mock isolates the mock interface of the
// NewWriter_genType type
type MoqNewWriter_genType_mock struct {
	Moq *MoqNewWriter_genType
}

// MoqNewWriter_genType_params holds the params of the NewWriter_genType type
type MoqNewWriter_genType_params struct{ W io.Writer }

// MoqNewWriter_genType_paramsKey holds the map key params of the
// NewWriter_genType type
type MoqNewWriter_genType_paramsKey struct {
	Params struct{ W io.Writer }
	Hashes struct{ W hash.Hash }
}

// MoqNewWriter_genType_resultsByParams contains the results for a given set of
// parameters for the NewWriter_genType type
type MoqNewWriter_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewWriter_genType_paramsKey]*MoqNewWriter_genType_results
}

// MoqNewWriter_genType_doFn defines the type of function needed when calling
// AndDo for the NewWriter_genType type
type MoqNewWriter_genType_doFn func(w io.Writer)

// MoqNewWriter_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewWriter_genType type
type MoqNewWriter_genType_doReturnFn func(w io.Writer) *multipart.Writer

// MoqNewWriter_genType_results holds the results of the NewWriter_genType type
type MoqNewWriter_genType_results struct {
	Params  MoqNewWriter_genType_params
	Results []struct {
		Values *struct {
			Result1 *multipart.Writer
		}
		Sequence   uint32
		DoFn       MoqNewWriter_genType_doFn
		DoReturnFn MoqNewWriter_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewWriter_genType_fnRecorder routes recorded function calls to the
// MoqNewWriter_genType moq
type MoqNewWriter_genType_fnRecorder struct {
	Params    MoqNewWriter_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewWriter_genType_results
	Moq       *MoqNewWriter_genType
}

// MoqNewWriter_genType_anyParams isolates the any params functions of the
// NewWriter_genType type
type MoqNewWriter_genType_anyParams struct {
	Recorder *MoqNewWriter_genType_fnRecorder
}

// NewMoqNewWriter_genType creates a new moq of the NewWriter_genType type
func NewMoqNewWriter_genType(scene *moq.Scene, config *moq.Config) *MoqNewWriter_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewWriter_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewWriter_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				W moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			W moq.ParamIndexing
		}{
			W: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewWriter_genType type
func (m *MoqNewWriter_genType) Mock() NewWriter_genType {
	return func(w io.Writer) *multipart.Writer {
		m.Scene.T.Helper()
		moq := &MoqNewWriter_genType_mock{Moq: m}
		return moq.Fn(w)
	}
}

func (m *MoqNewWriter_genType_mock) Fn(w io.Writer) (result1 *multipart.Writer) {
	m.Moq.Scene.T.Helper()
	params := MoqNewWriter_genType_params{
		W: w,
	}
	var results *MoqNewWriter_genType_results
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
		result.DoFn(w)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(w)
	}
	return
}

func (m *MoqNewWriter_genType) OnCall(w io.Writer) *MoqNewWriter_genType_fnRecorder {
	return &MoqNewWriter_genType_fnRecorder{
		Params: MoqNewWriter_genType_params{
			W: w,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewWriter_genType_fnRecorder) Any() *MoqNewWriter_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewWriter_genType_anyParams{Recorder: r}
}

func (a *MoqNewWriter_genType_anyParams) W() *MoqNewWriter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNewWriter_genType_fnRecorder) Seq() *MoqNewWriter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewWriter_genType_fnRecorder) NoSeq() *MoqNewWriter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewWriter_genType_fnRecorder) ReturnResults(result1 *multipart.Writer) *MoqNewWriter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Writer
		}
		Sequence   uint32
		DoFn       MoqNewWriter_genType_doFn
		DoReturnFn MoqNewWriter_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *multipart.Writer
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewWriter_genType_fnRecorder) AndDo(fn MoqNewWriter_genType_doFn) *MoqNewWriter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewWriter_genType_fnRecorder) DoReturnResults(fn MoqNewWriter_genType_doReturnFn) *MoqNewWriter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Writer
		}
		Sequence   uint32
		DoFn       MoqNewWriter_genType_doFn
		DoReturnFn MoqNewWriter_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewWriter_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewWriter_genType_resultsByParams
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
		results = &MoqNewWriter_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewWriter_genType_paramsKey]*MoqNewWriter_genType_results{},
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
		r.Results = &MoqNewWriter_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewWriter_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewWriter_genType_fnRecorder {
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
					Result1 *multipart.Writer
				}
				Sequence   uint32
				DoFn       MoqNewWriter_genType_doFn
				DoReturnFn MoqNewWriter_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewWriter_genType) PrettyParams(params MoqNewWriter_genType_params) string {
	return fmt.Sprintf("NewWriter_genType(%#v)", params.W)
}

func (m *MoqNewWriter_genType) ParamsKey(params MoqNewWriter_genType_params, anyParams uint64) MoqNewWriter_genType_paramsKey {
	m.Scene.T.Helper()
	var wUsed io.Writer
	var wUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	return MoqNewWriter_genType_paramsKey{
		Params: struct{ W io.Writer }{
			W: wUsed,
		},
		Hashes: struct{ W hash.Hash }{
			W: wUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewWriter_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewWriter_genType) AssertExpectationsMet() {
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
