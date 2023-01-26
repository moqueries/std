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

// The following type assertion assures that multipart.Reader_starGenType is
// mocked completely
var _ Reader_starGenType = (*MoqReader_starGenType_mock)(nil)

// Reader_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Reader_starGenType interface {
	ReadForm(maxMemory int64) (*multipart.Form, error)
	NextPart() (*multipart.Part, error)
	NextRawPart() (*multipart.Part, error)
}

// MoqReader_starGenType holds the state of a moq of the Reader_starGenType
// type
type MoqReader_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReader_starGenType_mock

	ResultsByParams_ReadForm    []MoqReader_starGenType_ReadForm_resultsByParams
	ResultsByParams_NextPart    []MoqReader_starGenType_NextPart_resultsByParams
	ResultsByParams_NextRawPart []MoqReader_starGenType_NextRawPart_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			ReadForm struct {
				MaxMemory moq.ParamIndexing
			}
			NextPart    struct{}
			NextRawPart struct{}
		}
	}
}

// MoqReader_starGenType_mock isolates the mock interface of the
// Reader_starGenType type
type MoqReader_starGenType_mock struct {
	Moq *MoqReader_starGenType
}

// MoqReader_starGenType_recorder isolates the recorder interface of the
// Reader_starGenType type
type MoqReader_starGenType_recorder struct {
	Moq *MoqReader_starGenType
}

// MoqReader_starGenType_ReadForm_params holds the params of the
// Reader_starGenType type
type MoqReader_starGenType_ReadForm_params struct{ MaxMemory int64 }

// MoqReader_starGenType_ReadForm_paramsKey holds the map key params of the
// Reader_starGenType type
type MoqReader_starGenType_ReadForm_paramsKey struct {
	Params struct{ MaxMemory int64 }
	Hashes struct{ MaxMemory hash.Hash }
}

// MoqReader_starGenType_ReadForm_resultsByParams contains the results for a
// given set of parameters for the Reader_starGenType type
type MoqReader_starGenType_ReadForm_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReader_starGenType_ReadForm_paramsKey]*MoqReader_starGenType_ReadForm_results
}

// MoqReader_starGenType_ReadForm_doFn defines the type of function needed when
// calling AndDo for the Reader_starGenType type
type MoqReader_starGenType_ReadForm_doFn func(maxMemory int64)

// MoqReader_starGenType_ReadForm_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Reader_starGenType type
type MoqReader_starGenType_ReadForm_doReturnFn func(maxMemory int64) (*multipart.Form, error)

// MoqReader_starGenType_ReadForm_results holds the results of the
// Reader_starGenType type
type MoqReader_starGenType_ReadForm_results struct {
	Params  MoqReader_starGenType_ReadForm_params
	Results []struct {
		Values *struct {
			Result1 *multipart.Form
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_ReadForm_doFn
		DoReturnFn MoqReader_starGenType_ReadForm_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReader_starGenType_ReadForm_fnRecorder routes recorded function calls to
// the MoqReader_starGenType moq
type MoqReader_starGenType_ReadForm_fnRecorder struct {
	Params    MoqReader_starGenType_ReadForm_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReader_starGenType_ReadForm_results
	Moq       *MoqReader_starGenType
}

// MoqReader_starGenType_ReadForm_anyParams isolates the any params functions
// of the Reader_starGenType type
type MoqReader_starGenType_ReadForm_anyParams struct {
	Recorder *MoqReader_starGenType_ReadForm_fnRecorder
}

// MoqReader_starGenType_NextPart_params holds the params of the
// Reader_starGenType type
type MoqReader_starGenType_NextPart_params struct{}

// MoqReader_starGenType_NextPart_paramsKey holds the map key params of the
// Reader_starGenType type
type MoqReader_starGenType_NextPart_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqReader_starGenType_NextPart_resultsByParams contains the results for a
// given set of parameters for the Reader_starGenType type
type MoqReader_starGenType_NextPart_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReader_starGenType_NextPart_paramsKey]*MoqReader_starGenType_NextPart_results
}

// MoqReader_starGenType_NextPart_doFn defines the type of function needed when
// calling AndDo for the Reader_starGenType type
type MoqReader_starGenType_NextPart_doFn func()

// MoqReader_starGenType_NextPart_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Reader_starGenType type
type MoqReader_starGenType_NextPart_doReturnFn func() (*multipart.Part, error)

// MoqReader_starGenType_NextPart_results holds the results of the
// Reader_starGenType type
type MoqReader_starGenType_NextPart_results struct {
	Params  MoqReader_starGenType_NextPart_params
	Results []struct {
		Values *struct {
			Result1 *multipart.Part
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_NextPart_doFn
		DoReturnFn MoqReader_starGenType_NextPart_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReader_starGenType_NextPart_fnRecorder routes recorded function calls to
// the MoqReader_starGenType moq
type MoqReader_starGenType_NextPart_fnRecorder struct {
	Params    MoqReader_starGenType_NextPart_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReader_starGenType_NextPart_results
	Moq       *MoqReader_starGenType
}

// MoqReader_starGenType_NextPart_anyParams isolates the any params functions
// of the Reader_starGenType type
type MoqReader_starGenType_NextPart_anyParams struct {
	Recorder *MoqReader_starGenType_NextPart_fnRecorder
}

// MoqReader_starGenType_NextRawPart_params holds the params of the
// Reader_starGenType type
type MoqReader_starGenType_NextRawPart_params struct{}

// MoqReader_starGenType_NextRawPart_paramsKey holds the map key params of the
// Reader_starGenType type
type MoqReader_starGenType_NextRawPart_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqReader_starGenType_NextRawPart_resultsByParams contains the results for a
// given set of parameters for the Reader_starGenType type
type MoqReader_starGenType_NextRawPart_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReader_starGenType_NextRawPart_paramsKey]*MoqReader_starGenType_NextRawPart_results
}

// MoqReader_starGenType_NextRawPart_doFn defines the type of function needed
// when calling AndDo for the Reader_starGenType type
type MoqReader_starGenType_NextRawPart_doFn func()

// MoqReader_starGenType_NextRawPart_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Reader_starGenType type
type MoqReader_starGenType_NextRawPart_doReturnFn func() (*multipart.Part, error)

// MoqReader_starGenType_NextRawPart_results holds the results of the
// Reader_starGenType type
type MoqReader_starGenType_NextRawPart_results struct {
	Params  MoqReader_starGenType_NextRawPart_params
	Results []struct {
		Values *struct {
			Result1 *multipart.Part
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_NextRawPart_doFn
		DoReturnFn MoqReader_starGenType_NextRawPart_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReader_starGenType_NextRawPart_fnRecorder routes recorded function calls
// to the MoqReader_starGenType moq
type MoqReader_starGenType_NextRawPart_fnRecorder struct {
	Params    MoqReader_starGenType_NextRawPart_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReader_starGenType_NextRawPart_results
	Moq       *MoqReader_starGenType
}

// MoqReader_starGenType_NextRawPart_anyParams isolates the any params
// functions of the Reader_starGenType type
type MoqReader_starGenType_NextRawPart_anyParams struct {
	Recorder *MoqReader_starGenType_NextRawPart_fnRecorder
}

// NewMoqReader_starGenType creates a new moq of the Reader_starGenType type
func NewMoqReader_starGenType(scene *moq.Scene, config *moq.Config) *MoqReader_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReader_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReader_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				ReadForm struct {
					MaxMemory moq.ParamIndexing
				}
				NextPart    struct{}
				NextRawPart struct{}
			}
		}{ParameterIndexing: struct {
			ReadForm struct {
				MaxMemory moq.ParamIndexing
			}
			NextPart    struct{}
			NextRawPart struct{}
		}{
			ReadForm: struct {
				MaxMemory moq.ParamIndexing
			}{
				MaxMemory: moq.ParamIndexByValue,
			},
			NextPart:    struct{}{},
			NextRawPart: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Reader_starGenType type
func (m *MoqReader_starGenType) Mock() *MoqReader_starGenType_mock { return m.Moq }

func (m *MoqReader_starGenType_mock) ReadForm(maxMemory int64) (result1 *multipart.Form, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReader_starGenType_ReadForm_params{
		MaxMemory: maxMemory,
	}
	var results *MoqReader_starGenType_ReadForm_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ReadForm {
		paramsKey := m.Moq.ParamsKey_ReadForm(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ReadForm(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ReadForm(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ReadForm(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(maxMemory)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(maxMemory)
	}
	return
}

func (m *MoqReader_starGenType_mock) NextPart() (result1 *multipart.Part, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReader_starGenType_NextPart_params{}
	var results *MoqReader_starGenType_NextPart_results
	for _, resultsByParams := range m.Moq.ResultsByParams_NextPart {
		paramsKey := m.Moq.ParamsKey_NextPart(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_NextPart(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_NextPart(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_NextPart(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

func (m *MoqReader_starGenType_mock) NextRawPart() (result1 *multipart.Part, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReader_starGenType_NextRawPart_params{}
	var results *MoqReader_starGenType_NextRawPart_results
	for _, resultsByParams := range m.Moq.ResultsByParams_NextRawPart {
		paramsKey := m.Moq.ParamsKey_NextRawPart(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_NextRawPart(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_NextRawPart(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_NextRawPart(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Reader_starGenType type
func (m *MoqReader_starGenType) OnCall() *MoqReader_starGenType_recorder {
	return &MoqReader_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqReader_starGenType_recorder) ReadForm(maxMemory int64) *MoqReader_starGenType_ReadForm_fnRecorder {
	return &MoqReader_starGenType_ReadForm_fnRecorder{
		Params: MoqReader_starGenType_ReadForm_params{
			MaxMemory: maxMemory,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) Any() *MoqReader_starGenType_ReadForm_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadForm(r.Params))
		return nil
	}
	return &MoqReader_starGenType_ReadForm_anyParams{Recorder: r}
}

func (a *MoqReader_starGenType_ReadForm_anyParams) MaxMemory() *MoqReader_starGenType_ReadForm_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) Seq() *MoqReader_starGenType_ReadForm_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadForm(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) NoSeq() *MoqReader_starGenType_ReadForm_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ReadForm(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) ReturnResults(result1 *multipart.Form, result2 error) *MoqReader_starGenType_ReadForm_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Form
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_ReadForm_doFn
		DoReturnFn MoqReader_starGenType_ReadForm_doReturnFn
	}{
		Values: &struct {
			Result1 *multipart.Form
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) AndDo(fn MoqReader_starGenType_ReadForm_doFn) *MoqReader_starGenType_ReadForm_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) DoReturnResults(fn MoqReader_starGenType_ReadForm_doReturnFn) *MoqReader_starGenType_ReadForm_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Form
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_ReadForm_doFn
		DoReturnFn MoqReader_starGenType_ReadForm_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReader_starGenType_ReadForm_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ReadForm {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReader_starGenType_ReadForm_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReader_starGenType_ReadForm_paramsKey]*MoqReader_starGenType_ReadForm_results{},
		}
		r.Moq.ResultsByParams_ReadForm = append(r.Moq.ResultsByParams_ReadForm, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ReadForm) {
			copy(r.Moq.ResultsByParams_ReadForm[insertAt+1:], r.Moq.ResultsByParams_ReadForm[insertAt:0])
			r.Moq.ResultsByParams_ReadForm[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ReadForm(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReader_starGenType_ReadForm_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReader_starGenType_ReadForm_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReader_starGenType_ReadForm_fnRecorder {
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
					Result1 *multipart.Form
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqReader_starGenType_ReadForm_doFn
				DoReturnFn MoqReader_starGenType_ReadForm_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReader_starGenType) PrettyParams_ReadForm(params MoqReader_starGenType_ReadForm_params) string {
	return fmt.Sprintf("ReadForm(%#v)", params.MaxMemory)
}

func (m *MoqReader_starGenType) ParamsKey_ReadForm(params MoqReader_starGenType_ReadForm_params, anyParams uint64) MoqReader_starGenType_ReadForm_paramsKey {
	m.Scene.T.Helper()
	var maxMemoryUsed int64
	var maxMemoryUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ReadForm.MaxMemory == moq.ParamIndexByValue {
			maxMemoryUsed = params.MaxMemory
		} else {
			maxMemoryUsedHash = hash.DeepHash(params.MaxMemory)
		}
	}
	return MoqReader_starGenType_ReadForm_paramsKey{
		Params: struct{ MaxMemory int64 }{
			MaxMemory: maxMemoryUsed,
		},
		Hashes: struct{ MaxMemory hash.Hash }{
			MaxMemory: maxMemoryUsedHash,
		},
	}
}

func (m *MoqReader_starGenType_recorder) NextPart() *MoqReader_starGenType_NextPart_fnRecorder {
	return &MoqReader_starGenType_NextPart_fnRecorder{
		Params:   MoqReader_starGenType_NextPart_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) Any() *MoqReader_starGenType_NextPart_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NextPart(r.Params))
		return nil
	}
	return &MoqReader_starGenType_NextPart_anyParams{Recorder: r}
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) Seq() *MoqReader_starGenType_NextPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NextPart(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) NoSeq() *MoqReader_starGenType_NextPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NextPart(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) ReturnResults(result1 *multipart.Part, result2 error) *MoqReader_starGenType_NextPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Part
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_NextPart_doFn
		DoReturnFn MoqReader_starGenType_NextPart_doReturnFn
	}{
		Values: &struct {
			Result1 *multipart.Part
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) AndDo(fn MoqReader_starGenType_NextPart_doFn) *MoqReader_starGenType_NextPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) DoReturnResults(fn MoqReader_starGenType_NextPart_doReturnFn) *MoqReader_starGenType_NextPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Part
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_NextPart_doFn
		DoReturnFn MoqReader_starGenType_NextPart_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReader_starGenType_NextPart_resultsByParams
	for n, res := range r.Moq.ResultsByParams_NextPart {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReader_starGenType_NextPart_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReader_starGenType_NextPart_paramsKey]*MoqReader_starGenType_NextPart_results{},
		}
		r.Moq.ResultsByParams_NextPart = append(r.Moq.ResultsByParams_NextPart, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_NextPart) {
			copy(r.Moq.ResultsByParams_NextPart[insertAt+1:], r.Moq.ResultsByParams_NextPart[insertAt:0])
			r.Moq.ResultsByParams_NextPart[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_NextPart(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReader_starGenType_NextPart_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReader_starGenType_NextPart_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReader_starGenType_NextPart_fnRecorder {
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
					Result1 *multipart.Part
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqReader_starGenType_NextPart_doFn
				DoReturnFn MoqReader_starGenType_NextPart_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReader_starGenType) PrettyParams_NextPart(params MoqReader_starGenType_NextPart_params) string {
	return fmt.Sprintf("NextPart()")
}

func (m *MoqReader_starGenType) ParamsKey_NextPart(params MoqReader_starGenType_NextPart_params, anyParams uint64) MoqReader_starGenType_NextPart_paramsKey {
	m.Scene.T.Helper()
	return MoqReader_starGenType_NextPart_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqReader_starGenType_recorder) NextRawPart() *MoqReader_starGenType_NextRawPart_fnRecorder {
	return &MoqReader_starGenType_NextRawPart_fnRecorder{
		Params:   MoqReader_starGenType_NextRawPart_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) Any() *MoqReader_starGenType_NextRawPart_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NextRawPart(r.Params))
		return nil
	}
	return &MoqReader_starGenType_NextRawPart_anyParams{Recorder: r}
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) Seq() *MoqReader_starGenType_NextRawPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NextRawPart(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) NoSeq() *MoqReader_starGenType_NextRawPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NextRawPart(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) ReturnResults(result1 *multipart.Part, result2 error) *MoqReader_starGenType_NextRawPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Part
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_NextRawPart_doFn
		DoReturnFn MoqReader_starGenType_NextRawPart_doReturnFn
	}{
		Values: &struct {
			Result1 *multipart.Part
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) AndDo(fn MoqReader_starGenType_NextRawPart_doFn) *MoqReader_starGenType_NextRawPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) DoReturnResults(fn MoqReader_starGenType_NextRawPart_doReturnFn) *MoqReader_starGenType_NextRawPart_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *multipart.Part
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReader_starGenType_NextRawPart_doFn
		DoReturnFn MoqReader_starGenType_NextRawPart_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReader_starGenType_NextRawPart_resultsByParams
	for n, res := range r.Moq.ResultsByParams_NextRawPart {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReader_starGenType_NextRawPart_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReader_starGenType_NextRawPart_paramsKey]*MoqReader_starGenType_NextRawPart_results{},
		}
		r.Moq.ResultsByParams_NextRawPart = append(r.Moq.ResultsByParams_NextRawPart, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_NextRawPart) {
			copy(r.Moq.ResultsByParams_NextRawPart[insertAt+1:], r.Moq.ResultsByParams_NextRawPart[insertAt:0])
			r.Moq.ResultsByParams_NextRawPart[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_NextRawPart(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReader_starGenType_NextRawPart_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReader_starGenType_NextRawPart_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReader_starGenType_NextRawPart_fnRecorder {
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
					Result1 *multipart.Part
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqReader_starGenType_NextRawPart_doFn
				DoReturnFn MoqReader_starGenType_NextRawPart_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReader_starGenType) PrettyParams_NextRawPart(params MoqReader_starGenType_NextRawPart_params) string {
	return fmt.Sprintf("NextRawPart()")
}

func (m *MoqReader_starGenType) ParamsKey_NextRawPart(params MoqReader_starGenType_NextRawPart_params, anyParams uint64) MoqReader_starGenType_NextRawPart_paramsKey {
	m.Scene.T.Helper()
	return MoqReader_starGenType_NextRawPart_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqReader_starGenType) Reset() {
	m.ResultsByParams_ReadForm = nil
	m.ResultsByParams_NextPart = nil
	m.ResultsByParams_NextRawPart = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReader_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_ReadForm {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ReadForm(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_NextPart {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_NextPart(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_NextRawPart {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_NextRawPart(results.Params))
			}
		}
	}
}
