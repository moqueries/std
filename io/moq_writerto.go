// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package io

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that io.WriterTo is mocked completely
var _ io.WriterTo = (*MoqWriterTo_mock)(nil)

// MoqWriterTo holds the state of a moq of the WriterTo type
type MoqWriterTo struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWriterTo_mock

	ResultsByParams_WriteTo []MoqWriterTo_WriteTo_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			WriteTo struct {
				W moq.ParamIndexing
			}
		}
	}
	// MoqWriterTo_mock isolates the mock interface of the WriterTo type
}

type MoqWriterTo_mock struct {
	Moq *MoqWriterTo
}

// MoqWriterTo_recorder isolates the recorder interface of the WriterTo type
type MoqWriterTo_recorder struct {
	Moq *MoqWriterTo
}

// MoqWriterTo_WriteTo_params holds the params of the WriterTo type
type MoqWriterTo_WriteTo_params struct{ W io.Writer }

// MoqWriterTo_WriteTo_paramsKey holds the map key params of the WriterTo type
type MoqWriterTo_WriteTo_paramsKey struct {
	Params struct{ W io.Writer }
	Hashes struct{ W hash.Hash }
}

// MoqWriterTo_WriteTo_resultsByParams contains the results for a given set of
// parameters for the WriterTo type
type MoqWriterTo_WriteTo_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWriterTo_WriteTo_paramsKey]*MoqWriterTo_WriteTo_results
}

// MoqWriterTo_WriteTo_doFn defines the type of function needed when calling
// AndDo for the WriterTo type
type MoqWriterTo_WriteTo_doFn func(w io.Writer)

// MoqWriterTo_WriteTo_doReturnFn defines the type of function needed when
// calling DoReturnResults for the WriterTo type
type MoqWriterTo_WriteTo_doReturnFn func(w io.Writer) (n int64, err error)

// MoqWriterTo_WriteTo_results holds the results of the WriterTo type
type MoqWriterTo_WriteTo_results struct {
	Params  MoqWriterTo_WriteTo_params
	Results []struct {
		Values *struct {
			N   int64
			Err error
		}
		Sequence   uint32
		DoFn       MoqWriterTo_WriteTo_doFn
		DoReturnFn MoqWriterTo_WriteTo_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWriterTo_WriteTo_fnRecorder routes recorded function calls to the
// MoqWriterTo moq
type MoqWriterTo_WriteTo_fnRecorder struct {
	Params    MoqWriterTo_WriteTo_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWriterTo_WriteTo_results
	Moq       *MoqWriterTo
}

// MoqWriterTo_WriteTo_anyParams isolates the any params functions of the
// WriterTo type
type MoqWriterTo_WriteTo_anyParams struct {
	Recorder *MoqWriterTo_WriteTo_fnRecorder
}

// NewMoqWriterTo creates a new moq of the WriterTo type
func NewMoqWriterTo(scene *moq.Scene, config *moq.Config) *MoqWriterTo {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWriterTo{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWriterTo_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				WriteTo struct {
					W moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			WriteTo struct {
				W moq.ParamIndexing
			}
		}{
			WriteTo: struct {
				W moq.ParamIndexing
			}{
				W: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the WriterTo type
func (m *MoqWriterTo) Mock() *MoqWriterTo_mock { return m.Moq }

func (m *MoqWriterTo_mock) WriteTo(w io.Writer) (n int64, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqWriterTo_WriteTo_params{
		W: w,
	}
	var results *MoqWriterTo_WriteTo_results
	for _, resultsByParams := range m.Moq.ResultsByParams_WriteTo {
		paramsKey := m.Moq.ParamsKey_WriteTo(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_WriteTo(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_WriteTo(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_WriteTo(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(w)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(w)
	}
	return
}

// OnCall returns the recorder implementation of the WriterTo type
func (m *MoqWriterTo) OnCall() *MoqWriterTo_recorder {
	return &MoqWriterTo_recorder{
		Moq: m,
	}
}

func (m *MoqWriterTo_recorder) WriteTo(w io.Writer) *MoqWriterTo_WriteTo_fnRecorder {
	return &MoqWriterTo_WriteTo_fnRecorder{
		Params: MoqWriterTo_WriteTo_params{
			W: w,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqWriterTo_WriteTo_fnRecorder) Any() *MoqWriterTo_WriteTo_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteTo(r.Params))
		return nil
	}
	return &MoqWriterTo_WriteTo_anyParams{Recorder: r}
}

func (a *MoqWriterTo_WriteTo_anyParams) W() *MoqWriterTo_WriteTo_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqWriterTo_WriteTo_fnRecorder) Seq() *MoqWriterTo_WriteTo_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteTo(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWriterTo_WriteTo_fnRecorder) NoSeq() *MoqWriterTo_WriteTo_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteTo(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWriterTo_WriteTo_fnRecorder) ReturnResults(n int64, err error) *MoqWriterTo_WriteTo_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int64
			Err error
		}
		Sequence   uint32
		DoFn       MoqWriterTo_WriteTo_doFn
		DoReturnFn MoqWriterTo_WriteTo_doReturnFn
	}{
		Values: &struct {
			N   int64
			Err error
		}{
			N:   n,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqWriterTo_WriteTo_fnRecorder) AndDo(fn MoqWriterTo_WriteTo_doFn) *MoqWriterTo_WriteTo_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWriterTo_WriteTo_fnRecorder) DoReturnResults(fn MoqWriterTo_WriteTo_doReturnFn) *MoqWriterTo_WriteTo_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int64
			Err error
		}
		Sequence   uint32
		DoFn       MoqWriterTo_WriteTo_doFn
		DoReturnFn MoqWriterTo_WriteTo_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWriterTo_WriteTo_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWriterTo_WriteTo_resultsByParams
	for n, res := range r.Moq.ResultsByParams_WriteTo {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqWriterTo_WriteTo_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWriterTo_WriteTo_paramsKey]*MoqWriterTo_WriteTo_results{},
		}
		r.Moq.ResultsByParams_WriteTo = append(r.Moq.ResultsByParams_WriteTo, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_WriteTo) {
			copy(r.Moq.ResultsByParams_WriteTo[insertAt+1:], r.Moq.ResultsByParams_WriteTo[insertAt:0])
			r.Moq.ResultsByParams_WriteTo[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_WriteTo(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqWriterTo_WriteTo_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWriterTo_WriteTo_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWriterTo_WriteTo_fnRecorder {
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
					N   int64
					Err error
				}
				Sequence   uint32
				DoFn       MoqWriterTo_WriteTo_doFn
				DoReturnFn MoqWriterTo_WriteTo_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWriterTo) PrettyParams_WriteTo(params MoqWriterTo_WriteTo_params) string {
	return fmt.Sprintf("WriteTo(%#v)", params.W)
}

func (m *MoqWriterTo) ParamsKey_WriteTo(params MoqWriterTo_WriteTo_params, anyParams uint64) MoqWriterTo_WriteTo_paramsKey {
	m.Scene.T.Helper()
	var wUsed io.Writer
	var wUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.WriteTo.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	return MoqWriterTo_WriteTo_paramsKey{
		Params: struct{ W io.Writer }{
			W: wUsed,
		},
		Hashes: struct{ W hash.Hash }{
			W: wUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqWriterTo) Reset() { m.ResultsByParams_WriteTo = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWriterTo) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_WriteTo {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_WriteTo(results.Params))
			}
		}
	}
}