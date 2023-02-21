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

// The following type assertion assures that io.ReadSeekCloser is mocked
// completely
var _ io.ReadSeekCloser = (*MoqReadSeekCloser_mock)(nil)

// MoqReadSeekCloser holds the state of a moq of the ReadSeekCloser type
type MoqReadSeekCloser struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReadSeekCloser_mock

	ResultsByParams_Read  []MoqReadSeekCloser_Read_resultsByParams
	ResultsByParams_Seek  []MoqReadSeekCloser_Seek_resultsByParams
	ResultsByParams_Close []MoqReadSeekCloser_Close_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Read struct {
				P moq.ParamIndexing
			}
			Seek struct {
				Offset moq.ParamIndexing
				Whence moq.ParamIndexing
			}
			Close struct{}
		}
	}
}

// MoqReadSeekCloser_mock isolates the mock interface of the ReadSeekCloser
// type
type MoqReadSeekCloser_mock struct {
	Moq *MoqReadSeekCloser
}

// MoqReadSeekCloser_recorder isolates the recorder interface of the
// ReadSeekCloser type
type MoqReadSeekCloser_recorder struct {
	Moq *MoqReadSeekCloser
}

// MoqReadSeekCloser_Read_params holds the params of the ReadSeekCloser type
type MoqReadSeekCloser_Read_params struct{ P []byte }

// MoqReadSeekCloser_Read_paramsKey holds the map key params of the
// ReadSeekCloser type
type MoqReadSeekCloser_Read_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqReadSeekCloser_Read_resultsByParams contains the results for a given set
// of parameters for the ReadSeekCloser type
type MoqReadSeekCloser_Read_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadSeekCloser_Read_paramsKey]*MoqReadSeekCloser_Read_results
}

// MoqReadSeekCloser_Read_doFn defines the type of function needed when calling
// AndDo for the ReadSeekCloser type
type MoqReadSeekCloser_Read_doFn func(p []byte)

// MoqReadSeekCloser_Read_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReadSeekCloser type
type MoqReadSeekCloser_Read_doReturnFn func(p []byte) (n int, err error)

// MoqReadSeekCloser_Read_results holds the results of the ReadSeekCloser type
type MoqReadSeekCloser_Read_results struct {
	Params  MoqReadSeekCloser_Read_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReadSeekCloser_Read_doFn
		DoReturnFn MoqReadSeekCloser_Read_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadSeekCloser_Read_fnRecorder routes recorded function calls to the
// MoqReadSeekCloser moq
type MoqReadSeekCloser_Read_fnRecorder struct {
	Params    MoqReadSeekCloser_Read_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadSeekCloser_Read_results
	Moq       *MoqReadSeekCloser
}

// MoqReadSeekCloser_Read_anyParams isolates the any params functions of the
// ReadSeekCloser type
type MoqReadSeekCloser_Read_anyParams struct {
	Recorder *MoqReadSeekCloser_Read_fnRecorder
}

// MoqReadSeekCloser_Seek_params holds the params of the ReadSeekCloser type
type MoqReadSeekCloser_Seek_params struct {
	Offset int64
	Whence int
}

// MoqReadSeekCloser_Seek_paramsKey holds the map key params of the
// ReadSeekCloser type
type MoqReadSeekCloser_Seek_paramsKey struct {
	Params struct {
		Offset int64
		Whence int
	}
	Hashes struct {
		Offset hash.Hash
		Whence hash.Hash
	}
}

// MoqReadSeekCloser_Seek_resultsByParams contains the results for a given set
// of parameters for the ReadSeekCloser type
type MoqReadSeekCloser_Seek_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadSeekCloser_Seek_paramsKey]*MoqReadSeekCloser_Seek_results
}

// MoqReadSeekCloser_Seek_doFn defines the type of function needed when calling
// AndDo for the ReadSeekCloser type
type MoqReadSeekCloser_Seek_doFn func(offset int64, whence int)

// MoqReadSeekCloser_Seek_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReadSeekCloser type
type MoqReadSeekCloser_Seek_doReturnFn func(offset int64, whence int) (int64, error)

// MoqReadSeekCloser_Seek_results holds the results of the ReadSeekCloser type
type MoqReadSeekCloser_Seek_results struct {
	Params  MoqReadSeekCloser_Seek_params
	Results []struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadSeekCloser_Seek_doFn
		DoReturnFn MoqReadSeekCloser_Seek_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadSeekCloser_Seek_fnRecorder routes recorded function calls to the
// MoqReadSeekCloser moq
type MoqReadSeekCloser_Seek_fnRecorder struct {
	Params    MoqReadSeekCloser_Seek_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadSeekCloser_Seek_results
	Moq       *MoqReadSeekCloser
}

// MoqReadSeekCloser_Seek_anyParams isolates the any params functions of the
// ReadSeekCloser type
type MoqReadSeekCloser_Seek_anyParams struct {
	Recorder *MoqReadSeekCloser_Seek_fnRecorder
}

// MoqReadSeekCloser_Close_params holds the params of the ReadSeekCloser type
type MoqReadSeekCloser_Close_params struct{}

// MoqReadSeekCloser_Close_paramsKey holds the map key params of the
// ReadSeekCloser type
type MoqReadSeekCloser_Close_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqReadSeekCloser_Close_resultsByParams contains the results for a given set
// of parameters for the ReadSeekCloser type
type MoqReadSeekCloser_Close_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadSeekCloser_Close_paramsKey]*MoqReadSeekCloser_Close_results
}

// MoqReadSeekCloser_Close_doFn defines the type of function needed when
// calling AndDo for the ReadSeekCloser type
type MoqReadSeekCloser_Close_doFn func()

// MoqReadSeekCloser_Close_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReadSeekCloser type
type MoqReadSeekCloser_Close_doReturnFn func() error

// MoqReadSeekCloser_Close_results holds the results of the ReadSeekCloser type
type MoqReadSeekCloser_Close_results struct {
	Params  MoqReadSeekCloser_Close_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqReadSeekCloser_Close_doFn
		DoReturnFn MoqReadSeekCloser_Close_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadSeekCloser_Close_fnRecorder routes recorded function calls to the
// MoqReadSeekCloser moq
type MoqReadSeekCloser_Close_fnRecorder struct {
	Params    MoqReadSeekCloser_Close_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadSeekCloser_Close_results
	Moq       *MoqReadSeekCloser
}

// MoqReadSeekCloser_Close_anyParams isolates the any params functions of the
// ReadSeekCloser type
type MoqReadSeekCloser_Close_anyParams struct {
	Recorder *MoqReadSeekCloser_Close_fnRecorder
}

// NewMoqReadSeekCloser creates a new moq of the ReadSeekCloser type
func NewMoqReadSeekCloser(scene *moq.Scene, config *moq.Config) *MoqReadSeekCloser {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReadSeekCloser{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReadSeekCloser_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Read struct {
					P moq.ParamIndexing
				}
				Seek struct {
					Offset moq.ParamIndexing
					Whence moq.ParamIndexing
				}
				Close struct{}
			}
		}{ParameterIndexing: struct {
			Read struct {
				P moq.ParamIndexing
			}
			Seek struct {
				Offset moq.ParamIndexing
				Whence moq.ParamIndexing
			}
			Close struct{}
		}{
			Read: struct {
				P moq.ParamIndexing
			}{
				P: moq.ParamIndexByHash,
			},
			Seek: struct {
				Offset moq.ParamIndexing
				Whence moq.ParamIndexing
			}{
				Offset: moq.ParamIndexByValue,
				Whence: moq.ParamIndexByValue,
			},
			Close: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ReadSeekCloser type
func (m *MoqReadSeekCloser) Mock() *MoqReadSeekCloser_mock { return m.Moq }

func (m *MoqReadSeekCloser_mock) Read(p []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadSeekCloser_Read_params{
		P: p,
	}
	var results *MoqReadSeekCloser_Read_results
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

func (m *MoqReadSeekCloser_mock) Seek(offset int64, whence int) (result1 int64, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadSeekCloser_Seek_params{
		Offset: offset,
		Whence: whence,
	}
	var results *MoqReadSeekCloser_Seek_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Seek {
		paramsKey := m.Moq.ParamsKey_Seek(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Seek(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Seek(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Seek(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(offset, whence)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(offset, whence)
	}
	return
}

func (m *MoqReadSeekCloser_mock) Close() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadSeekCloser_Close_params{}
	var results *MoqReadSeekCloser_Close_results
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

// OnCall returns the recorder implementation of the ReadSeekCloser type
func (m *MoqReadSeekCloser) OnCall() *MoqReadSeekCloser_recorder {
	return &MoqReadSeekCloser_recorder{
		Moq: m,
	}
}

func (m *MoqReadSeekCloser_recorder) Read(p []byte) *MoqReadSeekCloser_Read_fnRecorder {
	return &MoqReadSeekCloser_Read_fnRecorder{
		Params: MoqReadSeekCloser_Read_params{
			P: p,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReadSeekCloser_Read_fnRecorder) Any() *MoqReadSeekCloser_Read_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	return &MoqReadSeekCloser_Read_anyParams{Recorder: r}
}

func (a *MoqReadSeekCloser_Read_anyParams) P() *MoqReadSeekCloser_Read_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReadSeekCloser_Read_fnRecorder) Seq() *MoqReadSeekCloser_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadSeekCloser_Read_fnRecorder) NoSeq() *MoqReadSeekCloser_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadSeekCloser_Read_fnRecorder) ReturnResults(n int, err error) *MoqReadSeekCloser_Read_fnRecorder {
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
		DoFn       MoqReadSeekCloser_Read_doFn
		DoReturnFn MoqReadSeekCloser_Read_doReturnFn
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

func (r *MoqReadSeekCloser_Read_fnRecorder) AndDo(fn MoqReadSeekCloser_Read_doFn) *MoqReadSeekCloser_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadSeekCloser_Read_fnRecorder) DoReturnResults(fn MoqReadSeekCloser_Read_doReturnFn) *MoqReadSeekCloser_Read_fnRecorder {
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
		DoFn       MoqReadSeekCloser_Read_doFn
		DoReturnFn MoqReadSeekCloser_Read_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadSeekCloser_Read_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadSeekCloser_Read_resultsByParams
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
		results = &MoqReadSeekCloser_Read_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadSeekCloser_Read_paramsKey]*MoqReadSeekCloser_Read_results{},
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
		r.Results = &MoqReadSeekCloser_Read_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadSeekCloser_Read_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadSeekCloser_Read_fnRecorder {
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
				DoFn       MoqReadSeekCloser_Read_doFn
				DoReturnFn MoqReadSeekCloser_Read_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadSeekCloser) PrettyParams_Read(params MoqReadSeekCloser_Read_params) string {
	return fmt.Sprintf("Read(%#v)", params.P)
}

func (m *MoqReadSeekCloser) ParamsKey_Read(params MoqReadSeekCloser_Read_params, anyParams uint64) MoqReadSeekCloser_Read_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Read.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter of the Read function can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqReadSeekCloser_Read_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

func (m *MoqReadSeekCloser_recorder) Seek(offset int64, whence int) *MoqReadSeekCloser_Seek_fnRecorder {
	return &MoqReadSeekCloser_Seek_fnRecorder{
		Params: MoqReadSeekCloser_Seek_params{
			Offset: offset,
			Whence: whence,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) Any() *MoqReadSeekCloser_Seek_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	return &MoqReadSeekCloser_Seek_anyParams{Recorder: r}
}

func (a *MoqReadSeekCloser_Seek_anyParams) Offset() *MoqReadSeekCloser_Seek_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqReadSeekCloser_Seek_anyParams) Whence() *MoqReadSeekCloser_Seek_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) Seq() *MoqReadSeekCloser_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) NoSeq() *MoqReadSeekCloser_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) ReturnResults(result1 int64, result2 error) *MoqReadSeekCloser_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadSeekCloser_Seek_doFn
		DoReturnFn MoqReadSeekCloser_Seek_doReturnFn
	}{
		Values: &struct {
			Result1 int64
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) AndDo(fn MoqReadSeekCloser_Seek_doFn) *MoqReadSeekCloser_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) DoReturnResults(fn MoqReadSeekCloser_Seek_doReturnFn) *MoqReadSeekCloser_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadSeekCloser_Seek_doFn
		DoReturnFn MoqReadSeekCloser_Seek_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadSeekCloser_Seek_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Seek {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReadSeekCloser_Seek_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadSeekCloser_Seek_paramsKey]*MoqReadSeekCloser_Seek_results{},
		}
		r.Moq.ResultsByParams_Seek = append(r.Moq.ResultsByParams_Seek, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Seek) {
			copy(r.Moq.ResultsByParams_Seek[insertAt+1:], r.Moq.ResultsByParams_Seek[insertAt:0])
			r.Moq.ResultsByParams_Seek[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Seek(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReadSeekCloser_Seek_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadSeekCloser_Seek_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadSeekCloser_Seek_fnRecorder {
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
					Result1 int64
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqReadSeekCloser_Seek_doFn
				DoReturnFn MoqReadSeekCloser_Seek_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadSeekCloser) PrettyParams_Seek(params MoqReadSeekCloser_Seek_params) string {
	return fmt.Sprintf("Seek(%#v, %#v)", params.Offset, params.Whence)
}

func (m *MoqReadSeekCloser) ParamsKey_Seek(params MoqReadSeekCloser_Seek_params, anyParams uint64) MoqReadSeekCloser_Seek_paramsKey {
	m.Scene.T.Helper()
	var offsetUsed int64
	var offsetUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Seek.Offset == moq.ParamIndexByValue {
			offsetUsed = params.Offset
		} else {
			offsetUsedHash = hash.DeepHash(params.Offset)
		}
	}
	var whenceUsed int
	var whenceUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Seek.Whence == moq.ParamIndexByValue {
			whenceUsed = params.Whence
		} else {
			whenceUsedHash = hash.DeepHash(params.Whence)
		}
	}
	return MoqReadSeekCloser_Seek_paramsKey{
		Params: struct {
			Offset int64
			Whence int
		}{
			Offset: offsetUsed,
			Whence: whenceUsed,
		},
		Hashes: struct {
			Offset hash.Hash
			Whence hash.Hash
		}{
			Offset: offsetUsedHash,
			Whence: whenceUsedHash,
		},
	}
}

func (m *MoqReadSeekCloser_recorder) Close() *MoqReadSeekCloser_Close_fnRecorder {
	return &MoqReadSeekCloser_Close_fnRecorder{
		Params:   MoqReadSeekCloser_Close_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReadSeekCloser_Close_fnRecorder) Any() *MoqReadSeekCloser_Close_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	return &MoqReadSeekCloser_Close_anyParams{Recorder: r}
}

func (r *MoqReadSeekCloser_Close_fnRecorder) Seq() *MoqReadSeekCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadSeekCloser_Close_fnRecorder) NoSeq() *MoqReadSeekCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadSeekCloser_Close_fnRecorder) ReturnResults(result1 error) *MoqReadSeekCloser_Close_fnRecorder {
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
		DoFn       MoqReadSeekCloser_Close_doFn
		DoReturnFn MoqReadSeekCloser_Close_doReturnFn
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

func (r *MoqReadSeekCloser_Close_fnRecorder) AndDo(fn MoqReadSeekCloser_Close_doFn) *MoqReadSeekCloser_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadSeekCloser_Close_fnRecorder) DoReturnResults(fn MoqReadSeekCloser_Close_doReturnFn) *MoqReadSeekCloser_Close_fnRecorder {
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
		DoFn       MoqReadSeekCloser_Close_doFn
		DoReturnFn MoqReadSeekCloser_Close_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadSeekCloser_Close_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadSeekCloser_Close_resultsByParams
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
		results = &MoqReadSeekCloser_Close_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadSeekCloser_Close_paramsKey]*MoqReadSeekCloser_Close_results{},
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
		r.Results = &MoqReadSeekCloser_Close_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadSeekCloser_Close_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadSeekCloser_Close_fnRecorder {
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
				DoFn       MoqReadSeekCloser_Close_doFn
				DoReturnFn MoqReadSeekCloser_Close_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadSeekCloser) PrettyParams_Close(params MoqReadSeekCloser_Close_params) string {
	return fmt.Sprintf("Close()")
}

func (m *MoqReadSeekCloser) ParamsKey_Close(params MoqReadSeekCloser_Close_params, anyParams uint64) MoqReadSeekCloser_Close_paramsKey {
	m.Scene.T.Helper()
	return MoqReadSeekCloser_Close_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqReadSeekCloser) Reset() {
	m.ResultsByParams_Read = nil
	m.ResultsByParams_Seek = nil
	m.ResultsByParams_Close = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReadSeekCloser) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Read {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Read(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Seek {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Seek(results.Params))
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
