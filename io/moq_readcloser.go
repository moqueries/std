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

// The following type assertion assures that io.ReadCloser is mocked completely
var _ io.ReadCloser = (*MoqReadCloser_mock)(nil)

// MoqReadCloser holds the state of a moq of the ReadCloser type
type MoqReadCloser struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReadCloser_mock

	ResultsByParams_Read  []MoqReadCloser_Read_resultsByParams
	ResultsByParams_Close []MoqReadCloser_Close_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Read struct {
				P moq.ParamIndexing
			}
			Close struct{}
		}
	}
}

// MoqReadCloser_mock isolates the mock interface of the ReadCloser type
type MoqReadCloser_mock struct {
	Moq *MoqReadCloser
}

// MoqReadCloser_recorder isolates the recorder interface of the ReadCloser
// type
type MoqReadCloser_recorder struct {
	Moq *MoqReadCloser
}

// MoqReadCloser_Read_params holds the params of the ReadCloser type
type MoqReadCloser_Read_params struct{ P []byte }

// MoqReadCloser_Read_paramsKey holds the map key params of the ReadCloser type
type MoqReadCloser_Read_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqReadCloser_Read_resultsByParams contains the results for a given set of
// parameters for the ReadCloser type
type MoqReadCloser_Read_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadCloser_Read_paramsKey]*MoqReadCloser_Read_results
}

// MoqReadCloser_Read_doFn defines the type of function needed when calling
// AndDo for the ReadCloser type
type MoqReadCloser_Read_doFn func(p []byte)

// MoqReadCloser_Read_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReadCloser type
type MoqReadCloser_Read_doReturnFn func(p []byte) (n int, err error)

// MoqReadCloser_Read_results holds the results of the ReadCloser type
type MoqReadCloser_Read_results struct {
	Params  MoqReadCloser_Read_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReadCloser_Read_doFn
		DoReturnFn MoqReadCloser_Read_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadCloser_Read_fnRecorder routes recorded function calls to the
// MoqReadCloser moq
type MoqReadCloser_Read_fnRecorder struct {
	Params    MoqReadCloser_Read_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadCloser_Read_results
	Moq       *MoqReadCloser
}

// MoqReadCloser_Read_anyParams isolates the any params functions of the
// ReadCloser type
type MoqReadCloser_Read_anyParams struct {
	Recorder *MoqReadCloser_Read_fnRecorder
}

// MoqReadCloser_Close_params holds the params of the ReadCloser type
type MoqReadCloser_Close_params struct{}

// MoqReadCloser_Close_paramsKey holds the map key params of the ReadCloser
// type
type MoqReadCloser_Close_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqReadCloser_Close_resultsByParams contains the results for a given set of
// parameters for the ReadCloser type
type MoqReadCloser_Close_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadCloser_Close_paramsKey]*MoqReadCloser_Close_results
}

// MoqReadCloser_Close_doFn defines the type of function needed when calling
// AndDo for the ReadCloser type
type MoqReadCloser_Close_doFn func()

// MoqReadCloser_Close_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReadCloser type
type MoqReadCloser_Close_doReturnFn func() error

// MoqReadCloser_Close_results holds the results of the ReadCloser type
type MoqReadCloser_Close_results struct {
	Params  MoqReadCloser_Close_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqReadCloser_Close_doFn
		DoReturnFn MoqReadCloser_Close_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadCloser_Close_fnRecorder routes recorded function calls to the
// MoqReadCloser moq
type MoqReadCloser_Close_fnRecorder struct {
	Params    MoqReadCloser_Close_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadCloser_Close_results
	Moq       *MoqReadCloser
}

// MoqReadCloser_Close_anyParams isolates the any params functions of the
// ReadCloser type
type MoqReadCloser_Close_anyParams struct {
	Recorder *MoqReadCloser_Close_fnRecorder
}

// NewMoqReadCloser creates a new moq of the ReadCloser type
func NewMoqReadCloser(scene *moq.Scene, config *moq.Config) *MoqReadCloser {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReadCloser{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReadCloser_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Read struct {
					P moq.ParamIndexing
				}
				Close struct{}
			}
		}{ParameterIndexing: struct {
			Read struct {
				P moq.ParamIndexing
			}
			Close struct{}
		}{
			Read: struct {
				P moq.ParamIndexing
			}{
				P: moq.ParamIndexByHash,
			},
			Close: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ReadCloser type
func (m *MoqReadCloser) Mock() *MoqReadCloser_mock { return m.Moq }

func (m *MoqReadCloser_mock) Read(p []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadCloser_Read_params{
		P: p,
	}
	var results *MoqReadCloser_Read_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Read {
		paramsKey := m.Moq.ParamsKey_Read(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Read(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Read(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Read(params))
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

func (m *MoqReadCloser_mock) Close() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadCloser_Close_params{}
	var results *MoqReadCloser_Close_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Close {
		paramsKey := m.Moq.ParamsKey_Close(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Close(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Close(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Close(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the ReadCloser type
func (m *MoqReadCloser) OnCall() *MoqReadCloser_recorder {
	return &MoqReadCloser_recorder{
		Moq: m,
	}
}

func (m *MoqReadCloser_recorder) Read(p []byte) *MoqReadCloser_Read_fnRecorder {
	return &MoqReadCloser_Read_fnRecorder{
		Params: MoqReadCloser_Read_params{
			P: p,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReadCloser_Read_fnRecorder) Any() *MoqReadCloser_Read_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	return &MoqReadCloser_Read_anyParams{Recorder: r}
}

func (a *MoqReadCloser_Read_anyParams) P() *MoqReadCloser_Read_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReadCloser_Read_fnRecorder) Seq() *MoqReadCloser_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadCloser_Read_fnRecorder) NoSeq() *MoqReadCloser_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadCloser_Read_fnRecorder) ReturnResults(n int, err error) *MoqReadCloser_Read_fnRecorder {
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
		DoFn       MoqReadCloser_Read_doFn
		DoReturnFn MoqReadCloser_Read_doReturnFn
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

func (r *MoqReadCloser_Read_fnRecorder) AndDo(fn MoqReadCloser_Read_doFn) *MoqReadCloser_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadCloser_Read_fnRecorder) DoReturnResults(fn MoqReadCloser_Read_doReturnFn) *MoqReadCloser_Read_fnRecorder {
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
		DoFn       MoqReadCloser_Read_doFn
		DoReturnFn MoqReadCloser_Read_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadCloser_Read_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadCloser_Read_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Read {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReadCloser_Read_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadCloser_Read_paramsKey]*MoqReadCloser_Read_results{},
		}
		r.Moq.ResultsByParams_Read = append(r.Moq.ResultsByParams_Read, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Read) {
			copy(r.Moq.ResultsByParams_Read[insertAt+1:], r.Moq.ResultsByParams_Read[insertAt:0])
			r.Moq.ResultsByParams_Read[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Read(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReadCloser_Read_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadCloser_Read_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadCloser_Read_fnRecorder {
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
				DoFn       MoqReadCloser_Read_doFn
				DoReturnFn MoqReadCloser_Read_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadCloser) PrettyParams_Read(params MoqReadCloser_Read_params) string {
	return fmt.Sprintf("Read(%#v)", params.P)
}

func (m *MoqReadCloser) ParamsKey_Read(params MoqReadCloser_Read_params, anyParams uint64) MoqReadCloser_Read_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Read.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter of the Read function can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqReadCloser_Read_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

func (m *MoqReadCloser_recorder) Close() *MoqReadCloser_Close_fnRecorder {
	return &MoqReadCloser_Close_fnRecorder{
		Params:   MoqReadCloser_Close_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReadCloser_Close_fnRecorder) Any() *MoqReadCloser_Close_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	return &MoqReadCloser_Close_anyParams{Recorder: r}
}

func (r *MoqReadCloser_Close_fnRecorder) Seq() *MoqReadCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadCloser_Close_fnRecorder) NoSeq() *MoqReadCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadCloser_Close_fnRecorder) ReturnResults(result1 error) *MoqReadCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqReadCloser_Close_doFn
		DoReturnFn MoqReadCloser_Close_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReadCloser_Close_fnRecorder) AndDo(fn MoqReadCloser_Close_doFn) *MoqReadCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadCloser_Close_fnRecorder) DoReturnResults(fn MoqReadCloser_Close_doReturnFn) *MoqReadCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqReadCloser_Close_doFn
		DoReturnFn MoqReadCloser_Close_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadCloser_Close_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadCloser_Close_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Close {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReadCloser_Close_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadCloser_Close_paramsKey]*MoqReadCloser_Close_results{},
		}
		r.Moq.ResultsByParams_Close = append(r.Moq.ResultsByParams_Close, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Close) {
			copy(r.Moq.ResultsByParams_Close[insertAt+1:], r.Moq.ResultsByParams_Close[insertAt:0])
			r.Moq.ResultsByParams_Close[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Close(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReadCloser_Close_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadCloser_Close_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadCloser_Close_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqReadCloser_Close_doFn
				DoReturnFn MoqReadCloser_Close_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadCloser) PrettyParams_Close(params MoqReadCloser_Close_params) string {
	return fmt.Sprintf("Close()")
}

func (m *MoqReadCloser) ParamsKey_Close(params MoqReadCloser_Close_params, anyParams uint64) MoqReadCloser_Close_paramsKey {
	m.Scene.T.Helper()
	return MoqReadCloser_Close_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqReadCloser) Reset() { m.ResultsByParams_Read = nil; m.ResultsByParams_Close = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReadCloser) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Read {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Read(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Close {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Close(results.Params))
			}
		}
	}
}