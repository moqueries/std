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

// The following type assertion assures that io.WriteSeeker is mocked
// completely
var _ io.WriteSeeker = (*MoqWriteSeeker_mock)(nil)

// MoqWriteSeeker holds the state of a moq of the WriteSeeker type
type MoqWriteSeeker struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWriteSeeker_mock

	ResultsByParams_Write []MoqWriteSeeker_Write_resultsByParams
	ResultsByParams_Seek  []MoqWriteSeeker_Seek_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Write struct {
				P moq.ParamIndexing
			}
			Seek struct {
				Offset moq.ParamIndexing
				Whence moq.ParamIndexing
			}
		}
	}
	// MoqWriteSeeker_mock isolates the mock interface of the WriteSeeker type
}

type MoqWriteSeeker_mock struct {
	Moq *MoqWriteSeeker
}

// MoqWriteSeeker_recorder isolates the recorder interface of the WriteSeeker
// type
type MoqWriteSeeker_recorder struct {
	Moq *MoqWriteSeeker
}

// MoqWriteSeeker_Write_params holds the params of the WriteSeeker type
type MoqWriteSeeker_Write_params struct{ P []byte }

// MoqWriteSeeker_Write_paramsKey holds the map key params of the WriteSeeker
// type
type MoqWriteSeeker_Write_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqWriteSeeker_Write_resultsByParams contains the results for a given set of
// parameters for the WriteSeeker type
type MoqWriteSeeker_Write_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWriteSeeker_Write_paramsKey]*MoqWriteSeeker_Write_results
}

// MoqWriteSeeker_Write_doFn defines the type of function needed when calling
// AndDo for the WriteSeeker type
type MoqWriteSeeker_Write_doFn func(p []byte)

// MoqWriteSeeker_Write_doReturnFn defines the type of function needed when
// calling DoReturnResults for the WriteSeeker type
type MoqWriteSeeker_Write_doReturnFn func(p []byte) (n int, err error)

// MoqWriteSeeker_Write_results holds the results of the WriteSeeker type
type MoqWriteSeeker_Write_results struct {
	Params  MoqWriteSeeker_Write_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqWriteSeeker_Write_doFn
		DoReturnFn MoqWriteSeeker_Write_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWriteSeeker_Write_fnRecorder routes recorded function calls to the
// MoqWriteSeeker moq
type MoqWriteSeeker_Write_fnRecorder struct {
	Params    MoqWriteSeeker_Write_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWriteSeeker_Write_results
	Moq       *MoqWriteSeeker
}

// MoqWriteSeeker_Write_anyParams isolates the any params functions of the
// WriteSeeker type
type MoqWriteSeeker_Write_anyParams struct {
	Recorder *MoqWriteSeeker_Write_fnRecorder
}

// MoqWriteSeeker_Seek_params holds the params of the WriteSeeker type
type MoqWriteSeeker_Seek_params struct {
	Offset int64
	Whence int
}

// MoqWriteSeeker_Seek_paramsKey holds the map key params of the WriteSeeker
// type
type MoqWriteSeeker_Seek_paramsKey struct {
	Params struct {
		Offset int64
		Whence int
	}
	Hashes struct {
		Offset hash.Hash
		Whence hash.Hash
	}
}

// MoqWriteSeeker_Seek_resultsByParams contains the results for a given set of
// parameters for the WriteSeeker type
type MoqWriteSeeker_Seek_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWriteSeeker_Seek_paramsKey]*MoqWriteSeeker_Seek_results
}

// MoqWriteSeeker_Seek_doFn defines the type of function needed when calling
// AndDo for the WriteSeeker type
type MoqWriteSeeker_Seek_doFn func(offset int64, whence int)

// MoqWriteSeeker_Seek_doReturnFn defines the type of function needed when
// calling DoReturnResults for the WriteSeeker type
type MoqWriteSeeker_Seek_doReturnFn func(offset int64, whence int) (int64, error)

// MoqWriteSeeker_Seek_results holds the results of the WriteSeeker type
type MoqWriteSeeker_Seek_results struct {
	Params  MoqWriteSeeker_Seek_params
	Results []struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqWriteSeeker_Seek_doFn
		DoReturnFn MoqWriteSeeker_Seek_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWriteSeeker_Seek_fnRecorder routes recorded function calls to the
// MoqWriteSeeker moq
type MoqWriteSeeker_Seek_fnRecorder struct {
	Params    MoqWriteSeeker_Seek_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWriteSeeker_Seek_results
	Moq       *MoqWriteSeeker
}

// MoqWriteSeeker_Seek_anyParams isolates the any params functions of the
// WriteSeeker type
type MoqWriteSeeker_Seek_anyParams struct {
	Recorder *MoqWriteSeeker_Seek_fnRecorder
}

// NewMoqWriteSeeker creates a new moq of the WriteSeeker type
func NewMoqWriteSeeker(scene *moq.Scene, config *moq.Config) *MoqWriteSeeker {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWriteSeeker{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWriteSeeker_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Write struct {
					P moq.ParamIndexing
				}
				Seek struct {
					Offset moq.ParamIndexing
					Whence moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Write struct {
				P moq.ParamIndexing
			}
			Seek struct {
				Offset moq.ParamIndexing
				Whence moq.ParamIndexing
			}
		}{
			Write: struct {
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
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the WriteSeeker type
func (m *MoqWriteSeeker) Mock() *MoqWriteSeeker_mock { return m.Moq }

func (m *MoqWriteSeeker_mock) Write(p []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqWriteSeeker_Write_params{
		P: p,
	}
	var results *MoqWriteSeeker_Write_results
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

func (m *MoqWriteSeeker_mock) Seek(offset int64, whence int) (result1 int64, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqWriteSeeker_Seek_params{
		Offset: offset,
		Whence: whence,
	}
	var results *MoqWriteSeeker_Seek_results
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

// OnCall returns the recorder implementation of the WriteSeeker type
func (m *MoqWriteSeeker) OnCall() *MoqWriteSeeker_recorder {
	return &MoqWriteSeeker_recorder{
		Moq: m,
	}
}

func (m *MoqWriteSeeker_recorder) Write(p []byte) *MoqWriteSeeker_Write_fnRecorder {
	return &MoqWriteSeeker_Write_fnRecorder{
		Params: MoqWriteSeeker_Write_params{
			P: p,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqWriteSeeker_Write_fnRecorder) Any() *MoqWriteSeeker_Write_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	return &MoqWriteSeeker_Write_anyParams{Recorder: r}
}

func (a *MoqWriteSeeker_Write_anyParams) P() *MoqWriteSeeker_Write_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqWriteSeeker_Write_fnRecorder) Seq() *MoqWriteSeeker_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWriteSeeker_Write_fnRecorder) NoSeq() *MoqWriteSeeker_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Write(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWriteSeeker_Write_fnRecorder) ReturnResults(n int, err error) *MoqWriteSeeker_Write_fnRecorder {
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
		DoFn       MoqWriteSeeker_Write_doFn
		DoReturnFn MoqWriteSeeker_Write_doReturnFn
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

func (r *MoqWriteSeeker_Write_fnRecorder) AndDo(fn MoqWriteSeeker_Write_doFn) *MoqWriteSeeker_Write_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWriteSeeker_Write_fnRecorder) DoReturnResults(fn MoqWriteSeeker_Write_doReturnFn) *MoqWriteSeeker_Write_fnRecorder {
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
		DoFn       MoqWriteSeeker_Write_doFn
		DoReturnFn MoqWriteSeeker_Write_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWriteSeeker_Write_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWriteSeeker_Write_resultsByParams
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
		results = &MoqWriteSeeker_Write_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWriteSeeker_Write_paramsKey]*MoqWriteSeeker_Write_results{},
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
		r.Results = &MoqWriteSeeker_Write_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWriteSeeker_Write_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWriteSeeker_Write_fnRecorder {
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
				DoFn       MoqWriteSeeker_Write_doFn
				DoReturnFn MoqWriteSeeker_Write_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWriteSeeker) PrettyParams_Write(params MoqWriteSeeker_Write_params) string {
	return fmt.Sprintf("Write(%#v)", params.P)
}

func (m *MoqWriteSeeker) ParamsKey_Write(params MoqWriteSeeker_Write_params, anyParams uint64) MoqWriteSeeker_Write_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Write.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter of the Write function can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqWriteSeeker_Write_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

func (m *MoqWriteSeeker_recorder) Seek(offset int64, whence int) *MoqWriteSeeker_Seek_fnRecorder {
	return &MoqWriteSeeker_Seek_fnRecorder{
		Params: MoqWriteSeeker_Seek_params{
			Offset: offset,
			Whence: whence,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqWriteSeeker_Seek_fnRecorder) Any() *MoqWriteSeeker_Seek_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	return &MoqWriteSeeker_Seek_anyParams{Recorder: r}
}

func (a *MoqWriteSeeker_Seek_anyParams) Offset() *MoqWriteSeeker_Seek_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqWriteSeeker_Seek_anyParams) Whence() *MoqWriteSeeker_Seek_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqWriteSeeker_Seek_fnRecorder) Seq() *MoqWriteSeeker_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWriteSeeker_Seek_fnRecorder) NoSeq() *MoqWriteSeeker_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWriteSeeker_Seek_fnRecorder) ReturnResults(result1 int64, result2 error) *MoqWriteSeeker_Seek_fnRecorder {
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
		DoFn       MoqWriteSeeker_Seek_doFn
		DoReturnFn MoqWriteSeeker_Seek_doReturnFn
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

func (r *MoqWriteSeeker_Seek_fnRecorder) AndDo(fn MoqWriteSeeker_Seek_doFn) *MoqWriteSeeker_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWriteSeeker_Seek_fnRecorder) DoReturnResults(fn MoqWriteSeeker_Seek_doReturnFn) *MoqWriteSeeker_Seek_fnRecorder {
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
		DoFn       MoqWriteSeeker_Seek_doFn
		DoReturnFn MoqWriteSeeker_Seek_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWriteSeeker_Seek_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWriteSeeker_Seek_resultsByParams
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
		results = &MoqWriteSeeker_Seek_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWriteSeeker_Seek_paramsKey]*MoqWriteSeeker_Seek_results{},
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
		r.Results = &MoqWriteSeeker_Seek_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWriteSeeker_Seek_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWriteSeeker_Seek_fnRecorder {
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
				DoFn       MoqWriteSeeker_Seek_doFn
				DoReturnFn MoqWriteSeeker_Seek_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWriteSeeker) PrettyParams_Seek(params MoqWriteSeeker_Seek_params) string {
	return fmt.Sprintf("Seek(%#v, %#v)", params.Offset, params.Whence)
}

func (m *MoqWriteSeeker) ParamsKey_Seek(params MoqWriteSeeker_Seek_params, anyParams uint64) MoqWriteSeeker_Seek_paramsKey {
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
	return MoqWriteSeeker_Seek_paramsKey{
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

// Reset resets the state of the moq
func (m *MoqWriteSeeker) Reset() { m.ResultsByParams_Write = nil; m.ResultsByParams_Seek = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWriteSeeker) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Write {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Write(results.Params))
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
}
