// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package multipart

import (
	"fmt"
	"math/bits"
	"mime/multipart"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that multipart.File is mocked
// completely
var _ multipart.File = (*MoqFile_mock)(nil)

// MoqFile holds the state of a moq of the File type
type MoqFile struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFile_mock

	ResultsByParams_Read   []MoqFile_Read_resultsByParams
	ResultsByParams_ReadAt []MoqFile_ReadAt_resultsByParams
	ResultsByParams_Seek   []MoqFile_Seek_resultsByParams
	ResultsByParams_Close  []MoqFile_Close_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Read struct {
				P moq.ParamIndexing
			}
			ReadAt struct {
				P   moq.ParamIndexing
				Off moq.ParamIndexing
			}
			Seek struct {
				Offset moq.ParamIndexing
				Whence moq.ParamIndexing
			}
			Close struct{}
		}
	}
}

// MoqFile_mock isolates the mock interface of the File type
type MoqFile_mock struct {
	Moq *MoqFile
}

// MoqFile_recorder isolates the recorder interface of the File type
type MoqFile_recorder struct {
	Moq *MoqFile
}

// MoqFile_Read_params holds the params of the File type
type MoqFile_Read_params struct{ P []byte }

// MoqFile_Read_paramsKey holds the map key params of the File type
type MoqFile_Read_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqFile_Read_resultsByParams contains the results for a given set of
// parameters for the File type
type MoqFile_Read_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_Read_paramsKey]*MoqFile_Read_results
}

// MoqFile_Read_doFn defines the type of function needed when calling AndDo for
// the File type
type MoqFile_Read_doFn func(p []byte)

// MoqFile_Read_doReturnFn defines the type of function needed when calling
// DoReturnResults for the File type
type MoqFile_Read_doReturnFn func(p []byte) (n int, err error)

// MoqFile_Read_results holds the results of the File type
type MoqFile_Read_results struct {
	Params  MoqFile_Read_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqFile_Read_doFn
		DoReturnFn MoqFile_Read_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_Read_fnRecorder routes recorded function calls to the MoqFile moq
type MoqFile_Read_fnRecorder struct {
	Params    MoqFile_Read_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_Read_results
	Moq       *MoqFile
}

// MoqFile_Read_anyParams isolates the any params functions of the File type
type MoqFile_Read_anyParams struct {
	Recorder *MoqFile_Read_fnRecorder
}

// MoqFile_ReadAt_params holds the params of the File type
type MoqFile_ReadAt_params struct {
	P   []byte
	Off int64
}

// MoqFile_ReadAt_paramsKey holds the map key params of the File type
type MoqFile_ReadAt_paramsKey struct {
	Params struct{ Off int64 }
	Hashes struct {
		P   hash.Hash
		Off hash.Hash
	}
}

// MoqFile_ReadAt_resultsByParams contains the results for a given set of
// parameters for the File type
type MoqFile_ReadAt_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_ReadAt_paramsKey]*MoqFile_ReadAt_results
}

// MoqFile_ReadAt_doFn defines the type of function needed when calling AndDo
// for the File type
type MoqFile_ReadAt_doFn func(p []byte, off int64)

// MoqFile_ReadAt_doReturnFn defines the type of function needed when calling
// DoReturnResults for the File type
type MoqFile_ReadAt_doReturnFn func(p []byte, off int64) (n int, err error)

// MoqFile_ReadAt_results holds the results of the File type
type MoqFile_ReadAt_results struct {
	Params  MoqFile_ReadAt_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqFile_ReadAt_doFn
		DoReturnFn MoqFile_ReadAt_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_ReadAt_fnRecorder routes recorded function calls to the MoqFile moq
type MoqFile_ReadAt_fnRecorder struct {
	Params    MoqFile_ReadAt_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_ReadAt_results
	Moq       *MoqFile
}

// MoqFile_ReadAt_anyParams isolates the any params functions of the File type
type MoqFile_ReadAt_anyParams struct {
	Recorder *MoqFile_ReadAt_fnRecorder
}

// MoqFile_Seek_params holds the params of the File type
type MoqFile_Seek_params struct {
	Offset int64
	Whence int
}

// MoqFile_Seek_paramsKey holds the map key params of the File type
type MoqFile_Seek_paramsKey struct {
	Params struct {
		Offset int64
		Whence int
	}
	Hashes struct {
		Offset hash.Hash
		Whence hash.Hash
	}
}

// MoqFile_Seek_resultsByParams contains the results for a given set of
// parameters for the File type
type MoqFile_Seek_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_Seek_paramsKey]*MoqFile_Seek_results
}

// MoqFile_Seek_doFn defines the type of function needed when calling AndDo for
// the File type
type MoqFile_Seek_doFn func(offset int64, whence int)

// MoqFile_Seek_doReturnFn defines the type of function needed when calling
// DoReturnResults for the File type
type MoqFile_Seek_doReturnFn func(offset int64, whence int) (int64, error)

// MoqFile_Seek_results holds the results of the File type
type MoqFile_Seek_results struct {
	Params  MoqFile_Seek_params
	Results []struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqFile_Seek_doFn
		DoReturnFn MoqFile_Seek_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_Seek_fnRecorder routes recorded function calls to the MoqFile moq
type MoqFile_Seek_fnRecorder struct {
	Params    MoqFile_Seek_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_Seek_results
	Moq       *MoqFile
}

// MoqFile_Seek_anyParams isolates the any params functions of the File type
type MoqFile_Seek_anyParams struct {
	Recorder *MoqFile_Seek_fnRecorder
}

// MoqFile_Close_params holds the params of the File type
type MoqFile_Close_params struct{}

// MoqFile_Close_paramsKey holds the map key params of the File type
type MoqFile_Close_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqFile_Close_resultsByParams contains the results for a given set of
// parameters for the File type
type MoqFile_Close_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFile_Close_paramsKey]*MoqFile_Close_results
}

// MoqFile_Close_doFn defines the type of function needed when calling AndDo
// for the File type
type MoqFile_Close_doFn func()

// MoqFile_Close_doReturnFn defines the type of function needed when calling
// DoReturnResults for the File type
type MoqFile_Close_doReturnFn func() error

// MoqFile_Close_results holds the results of the File type
type MoqFile_Close_results struct {
	Params  MoqFile_Close_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqFile_Close_doFn
		DoReturnFn MoqFile_Close_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFile_Close_fnRecorder routes recorded function calls to the MoqFile moq
type MoqFile_Close_fnRecorder struct {
	Params    MoqFile_Close_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFile_Close_results
	Moq       *MoqFile
}

// MoqFile_Close_anyParams isolates the any params functions of the File type
type MoqFile_Close_anyParams struct {
	Recorder *MoqFile_Close_fnRecorder
}

// NewMoqFile creates a new moq of the File type
func NewMoqFile(scene *moq.Scene, config *moq.Config) *MoqFile {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFile{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFile_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Read struct {
					P moq.ParamIndexing
				}
				ReadAt struct {
					P   moq.ParamIndexing
					Off moq.ParamIndexing
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
			ReadAt struct {
				P   moq.ParamIndexing
				Off moq.ParamIndexing
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
			ReadAt: struct {
				P   moq.ParamIndexing
				Off moq.ParamIndexing
			}{
				P:   moq.ParamIndexByHash,
				Off: moq.ParamIndexByValue,
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

// Mock returns the mock implementation of the File type
func (m *MoqFile) Mock() *MoqFile_mock { return m.Moq }

func (m *MoqFile_mock) Read(p []byte) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_Read_params{
		P: p,
	}
	var results *MoqFile_Read_results
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

func (m *MoqFile_mock) ReadAt(p []byte, off int64) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_ReadAt_params{
		P:   p,
		Off: off,
	}
	var results *MoqFile_ReadAt_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ReadAt {
		paramsKey := m.Moq.ParamsKey_ReadAt(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ReadAt(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ReadAt(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ReadAt(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(p, off)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(p, off)
	}
	return
}

func (m *MoqFile_mock) Seek(offset int64, whence int) (result1 int64, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_Seek_params{
		Offset: offset,
		Whence: whence,
	}
	var results *MoqFile_Seek_results
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

func (m *MoqFile_mock) Close() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqFile_Close_params{}
	var results *MoqFile_Close_results
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

// OnCall returns the recorder implementation of the File type
func (m *MoqFile) OnCall() *MoqFile_recorder {
	return &MoqFile_recorder{
		Moq: m,
	}
}

func (m *MoqFile_recorder) Read(p []byte) *MoqFile_Read_fnRecorder {
	return &MoqFile_Read_fnRecorder{
		Params: MoqFile_Read_params{
			P: p,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_Read_fnRecorder) Any() *MoqFile_Read_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	return &MoqFile_Read_anyParams{Recorder: r}
}

func (a *MoqFile_Read_anyParams) P() *MoqFile_Read_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFile_Read_fnRecorder) Seq() *MoqFile_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_Read_fnRecorder) NoSeq() *MoqFile_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Read(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_Read_fnRecorder) ReturnResults(n int, err error) *MoqFile_Read_fnRecorder {
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
		DoFn       MoqFile_Read_doFn
		DoReturnFn MoqFile_Read_doReturnFn
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

func (r *MoqFile_Read_fnRecorder) AndDo(fn MoqFile_Read_doFn) *MoqFile_Read_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_Read_fnRecorder) DoReturnResults(fn MoqFile_Read_doReturnFn) *MoqFile_Read_fnRecorder {
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
		DoFn       MoqFile_Read_doFn
		DoReturnFn MoqFile_Read_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_Read_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_Read_resultsByParams
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
		results = &MoqFile_Read_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_Read_paramsKey]*MoqFile_Read_results{},
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
		r.Results = &MoqFile_Read_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_Read_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_Read_fnRecorder {
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
				DoFn       MoqFile_Read_doFn
				DoReturnFn MoqFile_Read_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile) PrettyParams_Read(params MoqFile_Read_params) string {
	return fmt.Sprintf("Read(%#v)", params.P)
}

func (m *MoqFile) ParamsKey_Read(params MoqFile_Read_params, anyParams uint64) MoqFile_Read_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Read.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter of the Read function can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqFile_Read_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

func (m *MoqFile_recorder) ReadAt(p []byte, off int64) *MoqFile_ReadAt_fnRecorder {
	return &MoqFile_ReadAt_fnRecorder{
		Params: MoqFile_ReadAt_params{
			P:   p,
			Off: off,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_ReadAt_fnRecorder) Any() *MoqFile_ReadAt_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadAt(r.Params))
		return nil
	}
	return &MoqFile_ReadAt_anyParams{Recorder: r}
}

func (a *MoqFile_ReadAt_anyParams) P() *MoqFile_ReadAt_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqFile_ReadAt_anyParams) Off() *MoqFile_ReadAt_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqFile_ReadAt_fnRecorder) Seq() *MoqFile_ReadAt_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadAt(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_ReadAt_fnRecorder) NoSeq() *MoqFile_ReadAt_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadAt(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_ReadAt_fnRecorder) ReturnResults(n int, err error) *MoqFile_ReadAt_fnRecorder {
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
		DoFn       MoqFile_ReadAt_doFn
		DoReturnFn MoqFile_ReadAt_doReturnFn
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

func (r *MoqFile_ReadAt_fnRecorder) AndDo(fn MoqFile_ReadAt_doFn) *MoqFile_ReadAt_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_ReadAt_fnRecorder) DoReturnResults(fn MoqFile_ReadAt_doReturnFn) *MoqFile_ReadAt_fnRecorder {
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
		DoFn       MoqFile_ReadAt_doFn
		DoReturnFn MoqFile_ReadAt_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_ReadAt_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_ReadAt_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ReadAt {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqFile_ReadAt_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_ReadAt_paramsKey]*MoqFile_ReadAt_results{},
		}
		r.Moq.ResultsByParams_ReadAt = append(r.Moq.ResultsByParams_ReadAt, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ReadAt) {
			copy(r.Moq.ResultsByParams_ReadAt[insertAt+1:], r.Moq.ResultsByParams_ReadAt[insertAt:0])
			r.Moq.ResultsByParams_ReadAt[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ReadAt(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqFile_ReadAt_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_ReadAt_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_ReadAt_fnRecorder {
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
				DoFn       MoqFile_ReadAt_doFn
				DoReturnFn MoqFile_ReadAt_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile) PrettyParams_ReadAt(params MoqFile_ReadAt_params) string {
	return fmt.Sprintf("ReadAt(%#v, %#v)", params.P, params.Off)
}

func (m *MoqFile) ParamsKey_ReadAt(params MoqFile_ReadAt_params, anyParams uint64) MoqFile_ReadAt_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ReadAt.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter of the ReadAt function can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	var offUsed int64
	var offUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.ReadAt.Off == moq.ParamIndexByValue {
			offUsed = params.Off
		} else {
			offUsedHash = hash.DeepHash(params.Off)
		}
	}
	return MoqFile_ReadAt_paramsKey{
		Params: struct{ Off int64 }{
			Off: offUsed,
		},
		Hashes: struct {
			P   hash.Hash
			Off hash.Hash
		}{
			P:   pUsedHash,
			Off: offUsedHash,
		},
	}
}

func (m *MoqFile_recorder) Seek(offset int64, whence int) *MoqFile_Seek_fnRecorder {
	return &MoqFile_Seek_fnRecorder{
		Params: MoqFile_Seek_params{
			Offset: offset,
			Whence: whence,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_Seek_fnRecorder) Any() *MoqFile_Seek_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	return &MoqFile_Seek_anyParams{Recorder: r}
}

func (a *MoqFile_Seek_anyParams) Offset() *MoqFile_Seek_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqFile_Seek_anyParams) Whence() *MoqFile_Seek_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqFile_Seek_fnRecorder) Seq() *MoqFile_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_Seek_fnRecorder) NoSeq() *MoqFile_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Seek(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_Seek_fnRecorder) ReturnResults(result1 int64, result2 error) *MoqFile_Seek_fnRecorder {
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
		DoFn       MoqFile_Seek_doFn
		DoReturnFn MoqFile_Seek_doReturnFn
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

func (r *MoqFile_Seek_fnRecorder) AndDo(fn MoqFile_Seek_doFn) *MoqFile_Seek_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_Seek_fnRecorder) DoReturnResults(fn MoqFile_Seek_doReturnFn) *MoqFile_Seek_fnRecorder {
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
		DoFn       MoqFile_Seek_doFn
		DoReturnFn MoqFile_Seek_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_Seek_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_Seek_resultsByParams
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
		results = &MoqFile_Seek_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_Seek_paramsKey]*MoqFile_Seek_results{},
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
		r.Results = &MoqFile_Seek_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_Seek_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_Seek_fnRecorder {
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
				DoFn       MoqFile_Seek_doFn
				DoReturnFn MoqFile_Seek_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile) PrettyParams_Seek(params MoqFile_Seek_params) string {
	return fmt.Sprintf("Seek(%#v, %#v)", params.Offset, params.Whence)
}

func (m *MoqFile) ParamsKey_Seek(params MoqFile_Seek_params, anyParams uint64) MoqFile_Seek_paramsKey {
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
	return MoqFile_Seek_paramsKey{
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

func (m *MoqFile_recorder) Close() *MoqFile_Close_fnRecorder {
	return &MoqFile_Close_fnRecorder{
		Params:   MoqFile_Close_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqFile_Close_fnRecorder) Any() *MoqFile_Close_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	return &MoqFile_Close_anyParams{Recorder: r}
}

func (r *MoqFile_Close_fnRecorder) Seq() *MoqFile_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFile_Close_fnRecorder) NoSeq() *MoqFile_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFile_Close_fnRecorder) ReturnResults(result1 error) *MoqFile_Close_fnRecorder {
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
		DoFn       MoqFile_Close_doFn
		DoReturnFn MoqFile_Close_doReturnFn
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

func (r *MoqFile_Close_fnRecorder) AndDo(fn MoqFile_Close_doFn) *MoqFile_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFile_Close_fnRecorder) DoReturnResults(fn MoqFile_Close_doReturnFn) *MoqFile_Close_fnRecorder {
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
		DoFn       MoqFile_Close_doFn
		DoReturnFn MoqFile_Close_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFile_Close_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFile_Close_resultsByParams
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
		results = &MoqFile_Close_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFile_Close_paramsKey]*MoqFile_Close_results{},
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
		r.Results = &MoqFile_Close_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFile_Close_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFile_Close_fnRecorder {
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
				DoFn       MoqFile_Close_doFn
				DoReturnFn MoqFile_Close_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFile) PrettyParams_Close(params MoqFile_Close_params) string {
	return fmt.Sprintf("Close()")
}

func (m *MoqFile) ParamsKey_Close(params MoqFile_Close_params, anyParams uint64) MoqFile_Close_paramsKey {
	m.Scene.T.Helper()
	return MoqFile_Close_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqFile) Reset() {
	m.ResultsByParams_Read = nil
	m.ResultsByParams_ReadAt = nil
	m.ResultsByParams_Seek = nil
	m.ResultsByParams_Close = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFile) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Read {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Read(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_ReadAt {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ReadAt(results.Params))
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
