// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package mail

import (
	"fmt"
	"math/bits"
	"net/mail"
	"sync/atomic"
	"time"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that mail.Header_genType is mocked
// completely
var _ Header_genType = (*MoqHeader_genType_mock)(nil)

// Header_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Header_genType interface {
	Get(key string) string
	Date() (time.Time, error)
	AddressList(key string) ([]*mail.Address, error)
}

// MoqHeader_genType holds the state of a moq of the Header_genType type
type MoqHeader_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqHeader_genType_mock

	ResultsByParams_Get         []MoqHeader_genType_Get_resultsByParams
	ResultsByParams_Date        []MoqHeader_genType_Date_resultsByParams
	ResultsByParams_AddressList []MoqHeader_genType_AddressList_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Get struct {
				Key moq.ParamIndexing
			}
			Date        struct{}
			AddressList struct {
				Key moq.ParamIndexing
			}
		}
	}
	// MoqHeader_genType_mock isolates the mock interface of the Header_genType
}

// type
type MoqHeader_genType_mock struct {
	Moq *MoqHeader_genType
}

// MoqHeader_genType_recorder isolates the recorder interface of the
// Header_genType type
type MoqHeader_genType_recorder struct {
	Moq *MoqHeader_genType
}

// MoqHeader_genType_Get_params holds the params of the Header_genType type
type MoqHeader_genType_Get_params struct{ Key string }

// MoqHeader_genType_Get_paramsKey holds the map key params of the
// Header_genType type
type MoqHeader_genType_Get_paramsKey struct {
	Params struct{ Key string }
	Hashes struct{ Key hash.Hash }
}

// MoqHeader_genType_Get_resultsByParams contains the results for a given set
// of parameters for the Header_genType type
type MoqHeader_genType_Get_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqHeader_genType_Get_paramsKey]*MoqHeader_genType_Get_results
}

// MoqHeader_genType_Get_doFn defines the type of function needed when calling
// AndDo for the Header_genType type
type MoqHeader_genType_Get_doFn func(key string)

// MoqHeader_genType_Get_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Header_genType type
type MoqHeader_genType_Get_doReturnFn func(key string) string

// MoqHeader_genType_Get_results holds the results of the Header_genType type
type MoqHeader_genType_Get_results struct {
	Params  MoqHeader_genType_Get_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_Get_doFn
		DoReturnFn MoqHeader_genType_Get_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqHeader_genType_Get_fnRecorder routes recorded function calls to the
// MoqHeader_genType moq
type MoqHeader_genType_Get_fnRecorder struct {
	Params    MoqHeader_genType_Get_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqHeader_genType_Get_results
	Moq       *MoqHeader_genType
}

// MoqHeader_genType_Get_anyParams isolates the any params functions of the
// Header_genType type
type MoqHeader_genType_Get_anyParams struct {
	Recorder *MoqHeader_genType_Get_fnRecorder
}

// MoqHeader_genType_Date_params holds the params of the Header_genType type
type MoqHeader_genType_Date_params struct{}

// MoqHeader_genType_Date_paramsKey holds the map key params of the
// Header_genType type
type MoqHeader_genType_Date_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqHeader_genType_Date_resultsByParams contains the results for a given set
// of parameters for the Header_genType type
type MoqHeader_genType_Date_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqHeader_genType_Date_paramsKey]*MoqHeader_genType_Date_results
}

// MoqHeader_genType_Date_doFn defines the type of function needed when calling
// AndDo for the Header_genType type
type MoqHeader_genType_Date_doFn func()

// MoqHeader_genType_Date_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Header_genType type
type MoqHeader_genType_Date_doReturnFn func() (time.Time, error)

// MoqHeader_genType_Date_results holds the results of the Header_genType type
type MoqHeader_genType_Date_results struct {
	Params  MoqHeader_genType_Date_params
	Results []struct {
		Values *struct {
			Result1 time.Time
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_Date_doFn
		DoReturnFn MoqHeader_genType_Date_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqHeader_genType_Date_fnRecorder routes recorded function calls to the
// MoqHeader_genType moq
type MoqHeader_genType_Date_fnRecorder struct {
	Params    MoqHeader_genType_Date_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqHeader_genType_Date_results
	Moq       *MoqHeader_genType
}

// MoqHeader_genType_Date_anyParams isolates the any params functions of the
// Header_genType type
type MoqHeader_genType_Date_anyParams struct {
	Recorder *MoqHeader_genType_Date_fnRecorder
}

// MoqHeader_genType_AddressList_params holds the params of the Header_genType
// type
type MoqHeader_genType_AddressList_params struct{ Key string }

// MoqHeader_genType_AddressList_paramsKey holds the map key params of the
// Header_genType type
type MoqHeader_genType_AddressList_paramsKey struct {
	Params struct{ Key string }
	Hashes struct{ Key hash.Hash }
}

// MoqHeader_genType_AddressList_resultsByParams contains the results for a
// given set of parameters for the Header_genType type
type MoqHeader_genType_AddressList_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqHeader_genType_AddressList_paramsKey]*MoqHeader_genType_AddressList_results
}

// MoqHeader_genType_AddressList_doFn defines the type of function needed when
// calling AndDo for the Header_genType type
type MoqHeader_genType_AddressList_doFn func(key string)

// MoqHeader_genType_AddressList_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Header_genType type
type MoqHeader_genType_AddressList_doReturnFn func(key string) ([]*mail.Address, error)

// MoqHeader_genType_AddressList_results holds the results of the
// Header_genType type
type MoqHeader_genType_AddressList_results struct {
	Params  MoqHeader_genType_AddressList_params
	Results []struct {
		Values *struct {
			Result1 []*mail.Address
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_AddressList_doFn
		DoReturnFn MoqHeader_genType_AddressList_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqHeader_genType_AddressList_fnRecorder routes recorded function calls to
// the MoqHeader_genType moq
type MoqHeader_genType_AddressList_fnRecorder struct {
	Params    MoqHeader_genType_AddressList_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqHeader_genType_AddressList_results
	Moq       *MoqHeader_genType
}

// MoqHeader_genType_AddressList_anyParams isolates the any params functions of
// the Header_genType type
type MoqHeader_genType_AddressList_anyParams struct {
	Recorder *MoqHeader_genType_AddressList_fnRecorder
}

// NewMoqHeader_genType creates a new moq of the Header_genType type
func NewMoqHeader_genType(scene *moq.Scene, config *moq.Config) *MoqHeader_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqHeader_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqHeader_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Get struct {
					Key moq.ParamIndexing
				}
				Date        struct{}
				AddressList struct {
					Key moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Get struct {
				Key moq.ParamIndexing
			}
			Date        struct{}
			AddressList struct {
				Key moq.ParamIndexing
			}
		}{
			Get: struct {
				Key moq.ParamIndexing
			}{
				Key: moq.ParamIndexByValue,
			},
			Date: struct{}{},
			AddressList: struct {
				Key moq.ParamIndexing
			}{
				Key: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Header_genType type
func (m *MoqHeader_genType) Mock() *MoqHeader_genType_mock { return m.Moq }

func (m *MoqHeader_genType_mock) Get(key string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqHeader_genType_Get_params{
		Key: key,
	}
	var results *MoqHeader_genType_Get_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Get {
		paramsKey := m.Moq.ParamsKey_Get(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Get(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Get(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Get(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(key)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(key)
	}
	return
}

func (m *MoqHeader_genType_mock) Date() (result1 time.Time, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqHeader_genType_Date_params{}
	var results *MoqHeader_genType_Date_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Date {
		paramsKey := m.Moq.ParamsKey_Date(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Date(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Date(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Date(params))
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

func (m *MoqHeader_genType_mock) AddressList(key string) (result1 []*mail.Address, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqHeader_genType_AddressList_params{
		Key: key,
	}
	var results *MoqHeader_genType_AddressList_results
	for _, resultsByParams := range m.Moq.ResultsByParams_AddressList {
		paramsKey := m.Moq.ParamsKey_AddressList(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_AddressList(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_AddressList(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_AddressList(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(key)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(key)
	}
	return
}

// OnCall returns the recorder implementation of the Header_genType type
func (m *MoqHeader_genType) OnCall() *MoqHeader_genType_recorder {
	return &MoqHeader_genType_recorder{
		Moq: m,
	}
}

func (m *MoqHeader_genType_recorder) Get(key string) *MoqHeader_genType_Get_fnRecorder {
	return &MoqHeader_genType_Get_fnRecorder{
		Params: MoqHeader_genType_Get_params{
			Key: key,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqHeader_genType_Get_fnRecorder) Any() *MoqHeader_genType_Get_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Get(r.Params))
		return nil
	}
	return &MoqHeader_genType_Get_anyParams{Recorder: r}
}

func (a *MoqHeader_genType_Get_anyParams) Key() *MoqHeader_genType_Get_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqHeader_genType_Get_fnRecorder) Seq() *MoqHeader_genType_Get_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Get(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqHeader_genType_Get_fnRecorder) NoSeq() *MoqHeader_genType_Get_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Get(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqHeader_genType_Get_fnRecorder) ReturnResults(result1 string) *MoqHeader_genType_Get_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_Get_doFn
		DoReturnFn MoqHeader_genType_Get_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqHeader_genType_Get_fnRecorder) AndDo(fn MoqHeader_genType_Get_doFn) *MoqHeader_genType_Get_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqHeader_genType_Get_fnRecorder) DoReturnResults(fn MoqHeader_genType_Get_doReturnFn) *MoqHeader_genType_Get_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_Get_doFn
		DoReturnFn MoqHeader_genType_Get_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqHeader_genType_Get_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqHeader_genType_Get_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Get {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqHeader_genType_Get_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqHeader_genType_Get_paramsKey]*MoqHeader_genType_Get_results{},
		}
		r.Moq.ResultsByParams_Get = append(r.Moq.ResultsByParams_Get, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Get) {
			copy(r.Moq.ResultsByParams_Get[insertAt+1:], r.Moq.ResultsByParams_Get[insertAt:0])
			r.Moq.ResultsByParams_Get[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Get(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqHeader_genType_Get_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqHeader_genType_Get_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqHeader_genType_Get_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqHeader_genType_Get_doFn
				DoReturnFn MoqHeader_genType_Get_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqHeader_genType) PrettyParams_Get(params MoqHeader_genType_Get_params) string {
	return fmt.Sprintf("Get(%#v)", params.Key)
}

func (m *MoqHeader_genType) ParamsKey_Get(params MoqHeader_genType_Get_params, anyParams uint64) MoqHeader_genType_Get_paramsKey {
	m.Scene.T.Helper()
	var keyUsed string
	var keyUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Get.Key == moq.ParamIndexByValue {
			keyUsed = params.Key
		} else {
			keyUsedHash = hash.DeepHash(params.Key)
		}
	}
	return MoqHeader_genType_Get_paramsKey{
		Params: struct{ Key string }{
			Key: keyUsed,
		},
		Hashes: struct{ Key hash.Hash }{
			Key: keyUsedHash,
		},
	}
}

func (m *MoqHeader_genType_recorder) Date() *MoqHeader_genType_Date_fnRecorder {
	return &MoqHeader_genType_Date_fnRecorder{
		Params:   MoqHeader_genType_Date_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqHeader_genType_Date_fnRecorder) Any() *MoqHeader_genType_Date_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Date(r.Params))
		return nil
	}
	return &MoqHeader_genType_Date_anyParams{Recorder: r}
}

func (r *MoqHeader_genType_Date_fnRecorder) Seq() *MoqHeader_genType_Date_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Date(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqHeader_genType_Date_fnRecorder) NoSeq() *MoqHeader_genType_Date_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Date(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqHeader_genType_Date_fnRecorder) ReturnResults(result1 time.Time, result2 error) *MoqHeader_genType_Date_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 time.Time
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_Date_doFn
		DoReturnFn MoqHeader_genType_Date_doReturnFn
	}{
		Values: &struct {
			Result1 time.Time
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqHeader_genType_Date_fnRecorder) AndDo(fn MoqHeader_genType_Date_doFn) *MoqHeader_genType_Date_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqHeader_genType_Date_fnRecorder) DoReturnResults(fn MoqHeader_genType_Date_doReturnFn) *MoqHeader_genType_Date_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 time.Time
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_Date_doFn
		DoReturnFn MoqHeader_genType_Date_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqHeader_genType_Date_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqHeader_genType_Date_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Date {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqHeader_genType_Date_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqHeader_genType_Date_paramsKey]*MoqHeader_genType_Date_results{},
		}
		r.Moq.ResultsByParams_Date = append(r.Moq.ResultsByParams_Date, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Date) {
			copy(r.Moq.ResultsByParams_Date[insertAt+1:], r.Moq.ResultsByParams_Date[insertAt:0])
			r.Moq.ResultsByParams_Date[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Date(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqHeader_genType_Date_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqHeader_genType_Date_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqHeader_genType_Date_fnRecorder {
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
					Result1 time.Time
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqHeader_genType_Date_doFn
				DoReturnFn MoqHeader_genType_Date_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqHeader_genType) PrettyParams_Date(params MoqHeader_genType_Date_params) string {
	return fmt.Sprintf("Date()")
}

func (m *MoqHeader_genType) ParamsKey_Date(params MoqHeader_genType_Date_params, anyParams uint64) MoqHeader_genType_Date_paramsKey {
	m.Scene.T.Helper()
	return MoqHeader_genType_Date_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqHeader_genType_recorder) AddressList(key string) *MoqHeader_genType_AddressList_fnRecorder {
	return &MoqHeader_genType_AddressList_fnRecorder{
		Params: MoqHeader_genType_AddressList_params{
			Key: key,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqHeader_genType_AddressList_fnRecorder) Any() *MoqHeader_genType_AddressList_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_AddressList(r.Params))
		return nil
	}
	return &MoqHeader_genType_AddressList_anyParams{Recorder: r}
}

func (a *MoqHeader_genType_AddressList_anyParams) Key() *MoqHeader_genType_AddressList_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqHeader_genType_AddressList_fnRecorder) Seq() *MoqHeader_genType_AddressList_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_AddressList(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqHeader_genType_AddressList_fnRecorder) NoSeq() *MoqHeader_genType_AddressList_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_AddressList(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqHeader_genType_AddressList_fnRecorder) ReturnResults(result1 []*mail.Address, result2 error) *MoqHeader_genType_AddressList_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []*mail.Address
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_AddressList_doFn
		DoReturnFn MoqHeader_genType_AddressList_doReturnFn
	}{
		Values: &struct {
			Result1 []*mail.Address
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqHeader_genType_AddressList_fnRecorder) AndDo(fn MoqHeader_genType_AddressList_doFn) *MoqHeader_genType_AddressList_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqHeader_genType_AddressList_fnRecorder) DoReturnResults(fn MoqHeader_genType_AddressList_doReturnFn) *MoqHeader_genType_AddressList_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []*mail.Address
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqHeader_genType_AddressList_doFn
		DoReturnFn MoqHeader_genType_AddressList_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqHeader_genType_AddressList_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqHeader_genType_AddressList_resultsByParams
	for n, res := range r.Moq.ResultsByParams_AddressList {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqHeader_genType_AddressList_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqHeader_genType_AddressList_paramsKey]*MoqHeader_genType_AddressList_results{},
		}
		r.Moq.ResultsByParams_AddressList = append(r.Moq.ResultsByParams_AddressList, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_AddressList) {
			copy(r.Moq.ResultsByParams_AddressList[insertAt+1:], r.Moq.ResultsByParams_AddressList[insertAt:0])
			r.Moq.ResultsByParams_AddressList[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_AddressList(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqHeader_genType_AddressList_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqHeader_genType_AddressList_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqHeader_genType_AddressList_fnRecorder {
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
					Result1 []*mail.Address
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqHeader_genType_AddressList_doFn
				DoReturnFn MoqHeader_genType_AddressList_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqHeader_genType) PrettyParams_AddressList(params MoqHeader_genType_AddressList_params) string {
	return fmt.Sprintf("AddressList(%#v)", params.Key)
}

func (m *MoqHeader_genType) ParamsKey_AddressList(params MoqHeader_genType_AddressList_params, anyParams uint64) MoqHeader_genType_AddressList_paramsKey {
	m.Scene.T.Helper()
	var keyUsed string
	var keyUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.AddressList.Key == moq.ParamIndexByValue {
			keyUsed = params.Key
		} else {
			keyUsedHash = hash.DeepHash(params.Key)
		}
	}
	return MoqHeader_genType_AddressList_paramsKey{
		Params: struct{ Key string }{
			Key: keyUsed,
		},
		Hashes: struct{ Key hash.Hash }{
			Key: keyUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqHeader_genType) Reset() {
	m.ResultsByParams_Get = nil
	m.ResultsByParams_Date = nil
	m.ResultsByParams_AddressList = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqHeader_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Get {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Get(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Date {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Date(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_AddressList {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_AddressList(results.Params))
			}
		}
	}
}
