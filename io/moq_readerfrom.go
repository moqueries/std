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

// The following type assertion assures that io.ReaderFrom is mocked completely
var _ io.ReaderFrom = (*MoqReaderFrom_mock)(nil)

// MoqReaderFrom holds the state of a moq of the ReaderFrom type
type MoqReaderFrom struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReaderFrom_mock

	ResultsByParams_ReadFrom []MoqReaderFrom_ReadFrom_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			ReadFrom struct {
				Param1 moq.ParamIndexing
			}
		}
	}
	// MoqReaderFrom_mock isolates the mock interface of the ReaderFrom type
}

type MoqReaderFrom_mock struct {
	Moq *MoqReaderFrom
}

// MoqReaderFrom_recorder isolates the recorder interface of the ReaderFrom
// type
type MoqReaderFrom_recorder struct {
	Moq *MoqReaderFrom
}

// MoqReaderFrom_ReadFrom_params holds the params of the ReaderFrom type
type MoqReaderFrom_ReadFrom_params struct{ Param1 io.Reader }

// MoqReaderFrom_ReadFrom_paramsKey holds the map key params of the ReaderFrom
// type
type MoqReaderFrom_ReadFrom_paramsKey struct {
	Params struct{ Param1 io.Reader }
	Hashes struct{ Param1 hash.Hash }
}

// MoqReaderFrom_ReadFrom_resultsByParams contains the results for a given set
// of parameters for the ReaderFrom type
type MoqReaderFrom_ReadFrom_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReaderFrom_ReadFrom_paramsKey]*MoqReaderFrom_ReadFrom_results
}

// MoqReaderFrom_ReadFrom_doFn defines the type of function needed when calling
// AndDo for the ReaderFrom type
type MoqReaderFrom_ReadFrom_doFn func(r io.Reader)

// MoqReaderFrom_ReadFrom_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReaderFrom type
type MoqReaderFrom_ReadFrom_doReturnFn func(r io.Reader) (n int64, err error)

// MoqReaderFrom_ReadFrom_results holds the results of the ReaderFrom type
type MoqReaderFrom_ReadFrom_results struct {
	Params  MoqReaderFrom_ReadFrom_params
	Results []struct {
		Values *struct {
			N   int64
			Err error
		}
		Sequence   uint32
		DoFn       MoqReaderFrom_ReadFrom_doFn
		DoReturnFn MoqReaderFrom_ReadFrom_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReaderFrom_ReadFrom_fnRecorder routes recorded function calls to the
// MoqReaderFrom moq
type MoqReaderFrom_ReadFrom_fnRecorder struct {
	Params    MoqReaderFrom_ReadFrom_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReaderFrom_ReadFrom_results
	Moq       *MoqReaderFrom
}

// MoqReaderFrom_ReadFrom_anyParams isolates the any params functions of the
// ReaderFrom type
type MoqReaderFrom_ReadFrom_anyParams struct {
	Recorder *MoqReaderFrom_ReadFrom_fnRecorder
}

// NewMoqReaderFrom creates a new moq of the ReaderFrom type
func NewMoqReaderFrom(scene *moq.Scene, config *moq.Config) *MoqReaderFrom {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReaderFrom{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReaderFrom_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				ReadFrom struct {
					Param1 moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			ReadFrom struct {
				Param1 moq.ParamIndexing
			}
		}{
			ReadFrom: struct {
				Param1 moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ReaderFrom type
func (m *MoqReaderFrom) Mock() *MoqReaderFrom_mock { return m.Moq }

func (m *MoqReaderFrom_mock) ReadFrom(param1 io.Reader) (n int64, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqReaderFrom_ReadFrom_params{
		Param1: param1,
	}
	var results *MoqReaderFrom_ReadFrom_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ReadFrom {
		paramsKey := m.Moq.ParamsKey_ReadFrom(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ReadFrom(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ReadFrom(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ReadFrom(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(param1)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(param1)
	}
	return
}

// OnCall returns the recorder implementation of the ReaderFrom type
func (m *MoqReaderFrom) OnCall() *MoqReaderFrom_recorder {
	return &MoqReaderFrom_recorder{
		Moq: m,
	}
}

func (m *MoqReaderFrom_recorder) ReadFrom(param1 io.Reader) *MoqReaderFrom_ReadFrom_fnRecorder {
	return &MoqReaderFrom_ReadFrom_fnRecorder{
		Params: MoqReaderFrom_ReadFrom_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReaderFrom_ReadFrom_fnRecorder) Any() *MoqReaderFrom_ReadFrom_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadFrom(r.Params))
		return nil
	}
	return &MoqReaderFrom_ReadFrom_anyParams{Recorder: r}
}

func (a *MoqReaderFrom_ReadFrom_anyParams) Param1() *MoqReaderFrom_ReadFrom_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReaderFrom_ReadFrom_fnRecorder) Seq() *MoqReaderFrom_ReadFrom_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadFrom(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReaderFrom_ReadFrom_fnRecorder) NoSeq() *MoqReaderFrom_ReadFrom_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadFrom(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReaderFrom_ReadFrom_fnRecorder) ReturnResults(n int64, err error) *MoqReaderFrom_ReadFrom_fnRecorder {
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
		DoFn       MoqReaderFrom_ReadFrom_doFn
		DoReturnFn MoqReaderFrom_ReadFrom_doReturnFn
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

func (r *MoqReaderFrom_ReadFrom_fnRecorder) AndDo(fn MoqReaderFrom_ReadFrom_doFn) *MoqReaderFrom_ReadFrom_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReaderFrom_ReadFrom_fnRecorder) DoReturnResults(fn MoqReaderFrom_ReadFrom_doReturnFn) *MoqReaderFrom_ReadFrom_fnRecorder {
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
		DoFn       MoqReaderFrom_ReadFrom_doFn
		DoReturnFn MoqReaderFrom_ReadFrom_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReaderFrom_ReadFrom_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReaderFrom_ReadFrom_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ReadFrom {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReaderFrom_ReadFrom_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReaderFrom_ReadFrom_paramsKey]*MoqReaderFrom_ReadFrom_results{},
		}
		r.Moq.ResultsByParams_ReadFrom = append(r.Moq.ResultsByParams_ReadFrom, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ReadFrom) {
			copy(r.Moq.ResultsByParams_ReadFrom[insertAt+1:], r.Moq.ResultsByParams_ReadFrom[insertAt:0])
			r.Moq.ResultsByParams_ReadFrom[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ReadFrom(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReaderFrom_ReadFrom_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReaderFrom_ReadFrom_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReaderFrom_ReadFrom_fnRecorder {
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
				DoFn       MoqReaderFrom_ReadFrom_doFn
				DoReturnFn MoqReaderFrom_ReadFrom_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReaderFrom) PrettyParams_ReadFrom(params MoqReaderFrom_ReadFrom_params) string {
	return fmt.Sprintf("ReadFrom(%#v)", params.Param1)
}

func (m *MoqReaderFrom) ParamsKey_ReadFrom(params MoqReaderFrom_ReadFrom_params, anyParams uint64) MoqReaderFrom_ReadFrom_paramsKey {
	m.Scene.T.Helper()
	var param1Used io.Reader
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ReadFrom.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqReaderFrom_ReadFrom_paramsKey{
		Params: struct{ Param1 io.Reader }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReaderFrom) Reset() { m.ResultsByParams_ReadFrom = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReaderFrom) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_ReadFrom {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ReadFrom(results.Params))
			}
		}
	}
}
