// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package io

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that io.PipeReader_starGenType is
// mocked completely
var _ PipeReader_starGenType = (*MoqPipeReader_starGenType_mock)(nil)

// PipeReader_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type PipeReader_starGenType interface {
	Read(data []byte) (n int, err error)
	Close() error
	CloseWithError(err error) error
}

// MoqPipeReader_starGenType holds the state of a moq of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPipeReader_starGenType_mock

	ResultsByParams_Read           []MoqPipeReader_starGenType_Read_resultsByParams
	ResultsByParams_Close          []MoqPipeReader_starGenType_Close_resultsByParams
	ResultsByParams_CloseWithError []MoqPipeReader_starGenType_CloseWithError_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Read struct {
				Data moq.ParamIndexing
			}
			Close          struct{}
			CloseWithError struct {
				Err moq.ParamIndexing
			}
		}
	}
	// MoqPipeReader_starGenType_mock isolates the mock interface of the
}

// PipeReader_starGenType type
type MoqPipeReader_starGenType_mock struct {
	Moq *MoqPipeReader_starGenType
}

// MoqPipeReader_starGenType_recorder isolates the recorder interface of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_recorder struct {
	Moq *MoqPipeReader_starGenType
}

// MoqPipeReader_starGenType_Read_params holds the params of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_Read_params struct{ Data []byte }

// MoqPipeReader_starGenType_Read_paramsKey holds the map key params of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_Read_paramsKey struct {
	Params struct{}
	Hashes struct{ Data hash.Hash }
}

// MoqPipeReader_starGenType_Read_resultsByParams contains the results for a
// given set of parameters for the PipeReader_starGenType type
type MoqPipeReader_starGenType_Read_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPipeReader_starGenType_Read_paramsKey]*MoqPipeReader_starGenType_Read_results
}

// MoqPipeReader_starGenType_Read_doFn defines the type of function needed when
// calling AndDo for the PipeReader_starGenType type
type MoqPipeReader_starGenType_Read_doFn func(data []byte)

// MoqPipeReader_starGenType_Read_doReturnFn defines the type of function
// needed when calling DoReturnResults for the PipeReader_starGenType type
type MoqPipeReader_starGenType_Read_doReturnFn func(data []byte) (n int, err error)

// MoqPipeReader_starGenType_Read_results holds the results of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_Read_results struct {
	Params  MoqPipeReader_starGenType_Read_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqPipeReader_starGenType_Read_doFn
		DoReturnFn MoqPipeReader_starGenType_Read_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPipeReader_starGenType_Read_fnRecorder routes recorded function calls to
// the MoqPipeReader_starGenType moq
type MoqPipeReader_starGenType_Read_fnRecorder struct {
	Params    MoqPipeReader_starGenType_Read_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPipeReader_starGenType_Read_results
	Moq       *MoqPipeReader_starGenType
}

// MoqPipeReader_starGenType_Read_anyParams isolates the any params functions
// of the PipeReader_starGenType type
type MoqPipeReader_starGenType_Read_anyParams struct {
	Recorder *MoqPipeReader_starGenType_Read_fnRecorder
}

// MoqPipeReader_starGenType_Close_params holds the params of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_Close_params struct{}

// MoqPipeReader_starGenType_Close_paramsKey holds the map key params of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_Close_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqPipeReader_starGenType_Close_resultsByParams contains the results for a
// given set of parameters for the PipeReader_starGenType type
type MoqPipeReader_starGenType_Close_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPipeReader_starGenType_Close_paramsKey]*MoqPipeReader_starGenType_Close_results
}

// MoqPipeReader_starGenType_Close_doFn defines the type of function needed
// when calling AndDo for the PipeReader_starGenType type
type MoqPipeReader_starGenType_Close_doFn func()

// MoqPipeReader_starGenType_Close_doReturnFn defines the type of function
// needed when calling DoReturnResults for the PipeReader_starGenType type
type MoqPipeReader_starGenType_Close_doReturnFn func() error

// MoqPipeReader_starGenType_Close_results holds the results of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_Close_results struct {
	Params  MoqPipeReader_starGenType_Close_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqPipeReader_starGenType_Close_doFn
		DoReturnFn MoqPipeReader_starGenType_Close_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPipeReader_starGenType_Close_fnRecorder routes recorded function calls to
// the MoqPipeReader_starGenType moq
type MoqPipeReader_starGenType_Close_fnRecorder struct {
	Params    MoqPipeReader_starGenType_Close_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPipeReader_starGenType_Close_results
	Moq       *MoqPipeReader_starGenType
}

// MoqPipeReader_starGenType_Close_anyParams isolates the any params functions
// of the PipeReader_starGenType type
type MoqPipeReader_starGenType_Close_anyParams struct {
	Recorder *MoqPipeReader_starGenType_Close_fnRecorder
}

// MoqPipeReader_starGenType_CloseWithError_params holds the params of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_CloseWithError_params struct{ Err error }

// MoqPipeReader_starGenType_CloseWithError_paramsKey holds the map key params
// of the PipeReader_starGenType type
type MoqPipeReader_starGenType_CloseWithError_paramsKey struct {
	Params struct{ Err error }
	Hashes struct{ Err hash.Hash }
}

// MoqPipeReader_starGenType_CloseWithError_resultsByParams contains the
// results for a given set of parameters for the PipeReader_starGenType type
type MoqPipeReader_starGenType_CloseWithError_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPipeReader_starGenType_CloseWithError_paramsKey]*MoqPipeReader_starGenType_CloseWithError_results
}

// MoqPipeReader_starGenType_CloseWithError_doFn defines the type of function
// needed when calling AndDo for the PipeReader_starGenType type
type MoqPipeReader_starGenType_CloseWithError_doFn func(err error)

// MoqPipeReader_starGenType_CloseWithError_doReturnFn defines the type of
// function needed when calling DoReturnResults for the PipeReader_starGenType
// type
type MoqPipeReader_starGenType_CloseWithError_doReturnFn func(err error) error

// MoqPipeReader_starGenType_CloseWithError_results holds the results of the
// PipeReader_starGenType type
type MoqPipeReader_starGenType_CloseWithError_results struct {
	Params  MoqPipeReader_starGenType_CloseWithError_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqPipeReader_starGenType_CloseWithError_doFn
		DoReturnFn MoqPipeReader_starGenType_CloseWithError_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPipeReader_starGenType_CloseWithError_fnRecorder routes recorded function
// calls to the MoqPipeReader_starGenType moq
type MoqPipeReader_starGenType_CloseWithError_fnRecorder struct {
	Params    MoqPipeReader_starGenType_CloseWithError_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPipeReader_starGenType_CloseWithError_results
	Moq       *MoqPipeReader_starGenType
}

// MoqPipeReader_starGenType_CloseWithError_anyParams isolates the any params
// functions of the PipeReader_starGenType type
type MoqPipeReader_starGenType_CloseWithError_anyParams struct {
	Recorder *MoqPipeReader_starGenType_CloseWithError_fnRecorder
}

// NewMoqPipeReader_starGenType creates a new moq of the PipeReader_starGenType
// type
func NewMoqPipeReader_starGenType(scene *moq.Scene, config *moq.Config) *MoqPipeReader_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPipeReader_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPipeReader_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Read struct {
					Data moq.ParamIndexing
				}
				Close          struct{}
				CloseWithError struct {
					Err moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Read struct {
				Data moq.ParamIndexing
			}
			Close          struct{}
			CloseWithError struct {
				Err moq.ParamIndexing
			}
		}{
			Read: struct {
				Data moq.ParamIndexing
			}{
				Data: moq.ParamIndexByHash,
			},
			Close: struct{}{},
			CloseWithError: struct {
				Err moq.ParamIndexing
			}{
				Err: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the PipeReader_starGenType type
func (m *MoqPipeReader_starGenType) Mock() *MoqPipeReader_starGenType_mock { return m.Moq }

func (m *MoqPipeReader_starGenType_mock) Read(data []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqPipeReader_starGenType_Read_params{
		Data: data,
	}
	var results *MoqPipeReader_starGenType_Read_results
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
		result.DoFn(data)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(data)
	}
	return
}

func (m *MoqPipeReader_starGenType_mock) Close() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqPipeReader_starGenType_Close_params{}
	var results *MoqPipeReader_starGenType_Close_results
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

func (m *MoqPipeReader_starGenType_mock) CloseWithError(err error) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqPipeReader_starGenType_CloseWithError_params{
		Err: err,
	}
	var results *MoqPipeReader_starGenType_CloseWithError_results
	for _, resultsByParams := range m.Moq.ResultsByParams_CloseWithError {
		paramsKey := m.Moq.ParamsKey_CloseWithError(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_CloseWithError(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_CloseWithError(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_CloseWithError(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(err)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(err)
	}
	return
}

// OnCall returns the recorder implementation of the PipeReader_starGenType
// type
func (m *MoqPipeReader_starGenType) OnCall() *MoqPipeReader_starGenType_recorder {
	return &MoqPipeReader_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqPipeReader_starGenType_recorder) Read(data []byte) *MoqPipeReader_starGenType_Read_fnRecorder {
	return &MoqPipeReader_starGenType_Read_fnRecorder{
		Params: MoqPipeReader_starGenType_Read_params{
			Data: data,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPipeReader_starGenType_Read_fnRecorder) Any() *MoqPipeReader_starGenType_Read_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	return &MoqPipeReader_starGenType_Read_anyParams{Recorder: r}
}

func (a *MoqPipeReader_starGenType_Read_anyParams) Data() *MoqPipeReader_starGenType_Read_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqPipeReader_starGenType_Read_fnRecorder) Seq() *MoqPipeReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPipeReader_starGenType_Read_fnRecorder) NoSeq() *MoqPipeReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPipeReader_starGenType_Read_fnRecorder) ReturnResults(n int, err error) *MoqPipeReader_starGenType_Read_fnRecorder {
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
		DoFn       MoqPipeReader_starGenType_Read_doFn
		DoReturnFn MoqPipeReader_starGenType_Read_doReturnFn
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

func (r *MoqPipeReader_starGenType_Read_fnRecorder) AndDo(fn MoqPipeReader_starGenType_Read_doFn) *MoqPipeReader_starGenType_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPipeReader_starGenType_Read_fnRecorder) DoReturnResults(fn MoqPipeReader_starGenType_Read_doReturnFn) *MoqPipeReader_starGenType_Read_fnRecorder {
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
		DoFn       MoqPipeReader_starGenType_Read_doFn
		DoReturnFn MoqPipeReader_starGenType_Read_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPipeReader_starGenType_Read_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPipeReader_starGenType_Read_resultsByParams
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
		results = &MoqPipeReader_starGenType_Read_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPipeReader_starGenType_Read_paramsKey]*MoqPipeReader_starGenType_Read_results{},
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
		r.Results = &MoqPipeReader_starGenType_Read_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPipeReader_starGenType_Read_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPipeReader_starGenType_Read_fnRecorder {
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
				DoFn       MoqPipeReader_starGenType_Read_doFn
				DoReturnFn MoqPipeReader_starGenType_Read_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPipeReader_starGenType) PrettyParams_Read(params MoqPipeReader_starGenType_Read_params) string {
	return fmt.Sprintf("Read(%#v)", params.Data)
}

func (m *MoqPipeReader_starGenType) ParamsKey_Read(params MoqPipeReader_starGenType_Read_params, anyParams uint64) MoqPipeReader_starGenType_Read_paramsKey {
	m.Scene.T.Helper()
	var dataUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Read.Data == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The data parameter of the Read function can't be indexed by value")
		}
		dataUsedHash = hash.DeepHash(params.Data)
	}
	return MoqPipeReader_starGenType_Read_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Data hash.Hash }{
			Data: dataUsedHash,
		},
	}
}

func (m *MoqPipeReader_starGenType_recorder) Close() *MoqPipeReader_starGenType_Close_fnRecorder {
	return &MoqPipeReader_starGenType_Close_fnRecorder{
		Params:   MoqPipeReader_starGenType_Close_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPipeReader_starGenType_Close_fnRecorder) Any() *MoqPipeReader_starGenType_Close_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	return &MoqPipeReader_starGenType_Close_anyParams{Recorder: r}
}

func (r *MoqPipeReader_starGenType_Close_fnRecorder) Seq() *MoqPipeReader_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPipeReader_starGenType_Close_fnRecorder) NoSeq() *MoqPipeReader_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPipeReader_starGenType_Close_fnRecorder) ReturnResults(result1 error) *MoqPipeReader_starGenType_Close_fnRecorder {
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
		DoFn       MoqPipeReader_starGenType_Close_doFn
		DoReturnFn MoqPipeReader_starGenType_Close_doReturnFn
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

func (r *MoqPipeReader_starGenType_Close_fnRecorder) AndDo(fn MoqPipeReader_starGenType_Close_doFn) *MoqPipeReader_starGenType_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPipeReader_starGenType_Close_fnRecorder) DoReturnResults(fn MoqPipeReader_starGenType_Close_doReturnFn) *MoqPipeReader_starGenType_Close_fnRecorder {
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
		DoFn       MoqPipeReader_starGenType_Close_doFn
		DoReturnFn MoqPipeReader_starGenType_Close_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPipeReader_starGenType_Close_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPipeReader_starGenType_Close_resultsByParams
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
		results = &MoqPipeReader_starGenType_Close_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPipeReader_starGenType_Close_paramsKey]*MoqPipeReader_starGenType_Close_results{},
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
		r.Results = &MoqPipeReader_starGenType_Close_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPipeReader_starGenType_Close_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPipeReader_starGenType_Close_fnRecorder {
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
				DoFn       MoqPipeReader_starGenType_Close_doFn
				DoReturnFn MoqPipeReader_starGenType_Close_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPipeReader_starGenType) PrettyParams_Close(params MoqPipeReader_starGenType_Close_params) string {
	return fmt.Sprintf("Close()")
}

func (m *MoqPipeReader_starGenType) ParamsKey_Close(params MoqPipeReader_starGenType_Close_params, anyParams uint64) MoqPipeReader_starGenType_Close_paramsKey {
	m.Scene.T.Helper()
	return MoqPipeReader_starGenType_Close_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqPipeReader_starGenType_recorder) CloseWithError(err error) *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
	return &MoqPipeReader_starGenType_CloseWithError_fnRecorder{
		Params: MoqPipeReader_starGenType_CloseWithError_params{
			Err: err,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) Any() *MoqPipeReader_starGenType_CloseWithError_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CloseWithError(r.Params))
		return nil
	}
	return &MoqPipeReader_starGenType_CloseWithError_anyParams{Recorder: r}
}

func (a *MoqPipeReader_starGenType_CloseWithError_anyParams) Err() *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) Seq() *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CloseWithError(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) NoSeq() *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CloseWithError(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) ReturnResults(result1 error) *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
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
		DoFn       MoqPipeReader_starGenType_CloseWithError_doFn
		DoReturnFn MoqPipeReader_starGenType_CloseWithError_doReturnFn
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

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) AndDo(fn MoqPipeReader_starGenType_CloseWithError_doFn) *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) DoReturnResults(fn MoqPipeReader_starGenType_CloseWithError_doReturnFn) *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
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
		DoFn       MoqPipeReader_starGenType_CloseWithError_doFn
		DoReturnFn MoqPipeReader_starGenType_CloseWithError_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPipeReader_starGenType_CloseWithError_resultsByParams
	for n, res := range r.Moq.ResultsByParams_CloseWithError {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqPipeReader_starGenType_CloseWithError_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPipeReader_starGenType_CloseWithError_paramsKey]*MoqPipeReader_starGenType_CloseWithError_results{},
		}
		r.Moq.ResultsByParams_CloseWithError = append(r.Moq.ResultsByParams_CloseWithError, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_CloseWithError) {
			copy(r.Moq.ResultsByParams_CloseWithError[insertAt+1:], r.Moq.ResultsByParams_CloseWithError[insertAt:0])
			r.Moq.ResultsByParams_CloseWithError[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_CloseWithError(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqPipeReader_starGenType_CloseWithError_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPipeReader_starGenType_CloseWithError_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPipeReader_starGenType_CloseWithError_fnRecorder {
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
				DoFn       MoqPipeReader_starGenType_CloseWithError_doFn
				DoReturnFn MoqPipeReader_starGenType_CloseWithError_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPipeReader_starGenType) PrettyParams_CloseWithError(params MoqPipeReader_starGenType_CloseWithError_params) string {
	return fmt.Sprintf("CloseWithError(%#v)", params.Err)
}

func (m *MoqPipeReader_starGenType) ParamsKey_CloseWithError(params MoqPipeReader_starGenType_CloseWithError_params, anyParams uint64) MoqPipeReader_starGenType_CloseWithError_paramsKey {
	m.Scene.T.Helper()
	var errUsed error
	var errUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.CloseWithError.Err == moq.ParamIndexByValue {
			errUsed = params.Err
		} else {
			errUsedHash = hash.DeepHash(params.Err)
		}
	}
	return MoqPipeReader_starGenType_CloseWithError_paramsKey{
		Params: struct{ Err error }{
			Err: errUsed,
		},
		Hashes: struct{ Err hash.Hash }{
			Err: errUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPipeReader_starGenType) Reset() {
	m.ResultsByParams_Read = nil
	m.ResultsByParams_Close = nil
	m.ResultsByParams_CloseWithError = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPipeReader_starGenType) AssertExpectationsMet() {
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
	for _, res := range m.ResultsByParams_CloseWithError {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_CloseWithError(results.Params))
			}
		}
	}
}
