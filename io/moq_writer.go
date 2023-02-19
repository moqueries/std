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

// The following type assertion assures that io.Writer is mocked completely
var _ io.Writer = (*MoqWriter_mock)(nil)

// MoqWriter holds the state of a moq of the Writer type
type MoqWriter struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWriter_mock

	ResultsByParams_Write []MoqWriter_Write_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Write struct {
				P moq.ParamIndexing
			}
		}
	}
	// MoqWriter_mock isolates the mock interface of the Writer type
}

type MoqWriter_mock struct {
	Moq *MoqWriter
}

// MoqWriter_recorder isolates the recorder interface of the Writer type
type MoqWriter_recorder struct {
	Moq *MoqWriter
}

// MoqWriter_Write_params holds the params of the Writer type
type MoqWriter_Write_params struct{ P []byte }

// MoqWriter_Write_paramsKey holds the map key params of the Writer type
type MoqWriter_Write_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqWriter_Write_resultsByParams contains the results for a given set of
// parameters for the Writer type
type MoqWriter_Write_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWriter_Write_paramsKey]*MoqWriter_Write_results
}

// MoqWriter_Write_doFn defines the type of function needed when calling AndDo
// for the Writer type
type MoqWriter_Write_doFn func(p []byte)

// MoqWriter_Write_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Writer type
type MoqWriter_Write_doReturnFn func(p []byte) (n int, err error)

// MoqWriter_Write_results holds the results of the Writer type
type MoqWriter_Write_results struct {
	Params  MoqWriter_Write_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqWriter_Write_doFn
		DoReturnFn MoqWriter_Write_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWriter_Write_fnRecorder routes recorded function calls to the MoqWriter
// moq
type MoqWriter_Write_fnRecorder struct {
	Params    MoqWriter_Write_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWriter_Write_results
	Moq       *MoqWriter
}

// MoqWriter_Write_anyParams isolates the any params functions of the Writer
// type
type MoqWriter_Write_anyParams struct {
	Recorder *MoqWriter_Write_fnRecorder
}

// NewMoqWriter creates a new moq of the Writer type
func NewMoqWriter(scene *moq.Scene, config *moq.Config) *MoqWriter {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWriter{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWriter_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Write struct {
					P moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Write struct {
				P moq.ParamIndexing
			}
		}{
			Write: struct {
				P moq.ParamIndexing
			}{
				P: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Writer type
func (m *MoqWriter) Mock() *MoqWriter_mock { return m.Moq }

func (m *MoqWriter_mock) Write(p []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqWriter_Write_params{
		P: p,
	}
	var results *MoqWriter_Write_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Write {
		paramsKey := m.Moq.ParamsKey_Write(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Write(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Write(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Write(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(p)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(p)
	}
	return
}

// OnCall returns the recorder implementation of the Writer type
func (m *MoqWriter) OnCall() *MoqWriter_recorder {
	return &MoqWriter_recorder{
		Moq: m,
	}
}

func (m *MoqWriter_recorder) Write(p []byte) *MoqWriter_Write_fnRecorder {
	return &MoqWriter_Write_fnRecorder{
		Params: MoqWriter_Write_params{
			P: p,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqWriter_Write_fnRecorder) Any() *MoqWriter_Write_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	return &MoqWriter_Write_anyParams{Recorder: r}
}

func (a *MoqWriter_Write_anyParams) P() *MoqWriter_Write_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqWriter_Write_fnRecorder) Seq() *MoqWriter_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWriter_Write_fnRecorder) NoSeq() *MoqWriter_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWriter_Write_fnRecorder) ReturnResults(n int, err error) *MoqWriter_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqWriter_Write_doFn
		DoReturnFn MoqWriter_Write_doReturnFn
	}{
		Values: &struct {
			N   int
			Err error
		}{
			N:   n,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqWriter_Write_fnRecorder) AndDo(fn MoqWriter_Write_doFn) *MoqWriter_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWriter_Write_fnRecorder) DoReturnResults(fn MoqWriter_Write_doReturnFn) *MoqWriter_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqWriter_Write_doFn
		DoReturnFn MoqWriter_Write_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWriter_Write_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWriter_Write_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Write {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqWriter_Write_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWriter_Write_paramsKey]*MoqWriter_Write_results{},
		}
		r.Moq.ResultsByParams_Write = append(r.Moq.ResultsByParams_Write, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Write) {
			copy(r.Moq.ResultsByParams_Write[insertAt+1:], r.Moq.ResultsByParams_Write[insertAt:0])
			r.Moq.ResultsByParams_Write[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Write(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqWriter_Write_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWriter_Write_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWriter_Write_fnRecorder {
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
					N   int
					Err error
				}
				Sequence   uint32
				DoFn       MoqWriter_Write_doFn
				DoReturnFn MoqWriter_Write_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWriter) PrettyParams_Write(params MoqWriter_Write_params) string {
	return fmt.Sprintf("Write(%#v)", params.P)
}

func (m *MoqWriter) ParamsKey_Write(params MoqWriter_Write_params, anyParams uint64) MoqWriter_Write_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Write.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter of the Write function can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqWriter_Write_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqWriter) Reset() { m.ResultsByParams_Write = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWriter) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Write {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Write(results.Params))
			}
		}
	}
}
