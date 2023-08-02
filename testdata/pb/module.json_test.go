package pb_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"protoc-gen-go-json/testdata/pb"
	"testing"
)

func Assert(t *testing.T, args json.Marshaler, want string) {
	t.Helper()
	raw, err := args.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, want, string(raw))
}

func Asserts(t *testing.T, args json.Marshaler, want []string) {
	t.Helper()
	raw, err := args.MarshalJSON()
	require.NoError(t, err)
	got := string(raw)
	for i := 0; i < len(want); i++ {
		if want[i] == got {
			return
		}
	}
	require.Fail(t, got)
}

func TestOneof_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *pb.Oneof
		want string
	}{
		{name: "nil", args: nil, want: ""},
		{
			name: "empty",
			args: &pb.Oneof{},
			want: "{}",
		},
		{
			name: "not empty",
			args: &pb.Oneof{
				Oneof:   &pb.Oneof_String_{String_: &pb.String{Str: "123"}},
				NumberX: &pb.Number{U32: 100},
				StringX: &pb.String{Str: "123"},
			},
			want: `{"string":{"str":"123"},"numberX":{"u32":100},"stringX":{"str":"123"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Assert(t, tt.args, tt.want)
		})
	}
}

func TestMessage_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *pb.Message
		want string
	}{
		{
			name: "not empty",
			args: &pb.Message{
				Type:    pb.Type_BOOL,
				Number:  &pb.Number{},
				String_: &pb.String{Str: "msg1", Bytes: []byte{48, 59}},
				Bool:    &pb.Bool{B: true},
			},
			want: `{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Assert(t, tt.args, tt.want)
		})
	}
}

func TestMap_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args *pb.Map
		want []string
	}{
		{
			name: "not empty",
			args: &pb.Map{
				Numbers: map[uint32]*pb.Number{
					1: {U32: 0, U64: 0, S32: 0, S64: 0, Uf32: 0, Uf64: 0, Sf32: 0, Sf64: 0, I32: 0, I64: 0, F64: 0, F32: 0},
					2: {U32: 2, U64: 2, S32: 2, S64: 0, Uf32: 0, Uf64: 0, Sf32: 0, Sf64: 0, I32: 0, I64: 0, F64: 0, F32: 0},
					3: {U32: 3, U64: 3, S32: 2, S64: 0, Uf32: 0, Uf64: 0, Sf32: 0, Sf64: 0, I32: 0, I64: 0, F64: 0, F32: 0},
				},
				Strings: map[string]*pb.String{
					"sk1": {Str: "sk1", Bytes: []byte{48, 59}},
					"nil": {Str: "nil", Bytes: nil},
				},
				Bools: map[bool]*pb.Bool{true: {B: true}, false: {B: false}},
				Messages: map[string]*pb.Message{
					"msg1": {
						Type:    pb.Type_BOOL,
						Number:  &pb.Number{},
						String_: &pb.String{Str: "msg1", Bytes: []byte{48, 59}},
						Bool:    &pb.Bool{B: true},
					},
				},
				Arrays: map[string]*pb.Array{
					"arr1": {
						Numbers: []*pb.Number{{}, {}},
						Strings: []*pb.String{{Str: "arr1", Bytes: []byte{48, 59}}, {Str: "arr2", Bytes: []byte{48, 59}}},
						Bools:   []*pb.Bool{{B: true}, {B: false}},
						Messages: []*pb.Message{
							{
								Type:    pb.Type_BOOL,
								Number:  &pb.Number{},
								String_: &pb.String{Str: "arr1_msg1", Bytes: []byte{48, 59}},
								Bool:    &pb.Bool{B: true},
							},
							{
								Type:    pb.Type_BOOL,
								Number:  &pb.Number{},
								String_: &pb.String{Str: "arr1_msg2", Bytes: []byte{48, 59}},
								Bool:    &pb.Bool{B: true},
							},
						},
						Arrays: []*pb.Array{
							{
								Numbers: []*pb.Number{{}, {}},
								Strings: []*pb.String{{Str: "arr1", Bytes: []byte{48, 59}}, {Str: "arr2", Bytes: []byte{48, 59}}},
								Bools:   []*pb.Bool{{B: true}, {B: false}},
								Messages: []*pb.Message{
									{
										Type:    pb.Type_BOOL,
										Number:  &pb.Number{},
										String_: &pb.String{Str: "arr1_msg1", Bytes: []byte{48, 59}},
										Bool:    &pb.Bool{B: true},
									},
									{
										Type:    pb.Type_BOOL,
										Number:  &pb.Number{},
										String_: &pb.String{Str: "arr1_msg2", Bytes: []byte{48, 59}},
										Bool:    &pb.Bool{B: true},
									},
								},
								Arrays: nil,
								Types:  []pb.Type{pb.Type_BOOL, pb.Type_NUMBER, pb.Type_STRING},
								U32S:   []uint32{0, 1, 2, 3},
								Strs:   []string{"str1", "str2", "str3"},
							},
						},
						Types: []pb.Type{pb.Type_BOOL, pb.Type_NUMBER, pb.Type_STRING},
						U32S:  []uint32{0, 1, 2, 3},
						Strs:  []string{"str1", "str2", "str3"},
					},
				},
				Types:   map[int32]pb.Type{0: pb.Type_BOOL, 1: pb.Type_NUMBER, 2: pb.Type_STRING},
				U32S:    map[string]uint32{"u32_1": 1, "u32_2": 2, "u32_3": 3},
				Strs:    map[string]string{"str1": "str1", "str2": "str2"},
				Empties: map[string]*pb.Empty{"empty1": {}, "empty2": {}, "empty3": {}},
				Optionals: map[string]*pb.Optional{
					"optional1": {Number: &pb.Number{U32: 1}, String_: &pb.String{Str: "str1"}},
					"optional2": {Number: &pb.Number{U32: 2}, String_: &pb.String{Str: "str2"}},
				},
			},
			want: []string{
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"1":"NUMBER","2":"STRING","0":"BOOL"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_2":2,"u32_3":3,"u32_1":1},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty2":{},"empty3":{},"empty1":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_3":3,"u32_1":1,"u32_2":2},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty3":{},"empty1":{},"empty2":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_3":3,"u32_1":1,"u32_2":2},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"false":{"b":false},"true":{"b":true}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"2":"STRING","0":"BOOL","1":"NUMBER"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"nil":{"str":"nil"},"sk1":{"str":"sk1","bytes":"MDs="}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty3":{},"empty1":{},"empty2":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty2":{},"empty3":{},"empty1":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2},"1":{}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"3":{"u32":3,"u64":3,"s32":2},"1":{},"2":{"u32":2,"u64":2,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"nil":{"str":"nil"},"sk1":{"str":"sk1","bytes":"MDs="}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"3":{"u32":3,"u64":3,"s32":2},"1":{},"2":{"u32":2,"u64":2,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_3":3,"u32_1":1,"u32_2":2},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty1":{},"empty2":{},"empty3":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty2":{},"empty3":{},"empty1":{}},"optionals":{"optional2":{"number":{"u32":2},"string":{"str":"str2"}},"optional1":{"number":{"u32":1},"string":{"str":"str1"}}}}`,
				`{"numbers":{"1":{},"2":{"u32":2,"u64":2,"s32":2},"3":{"u32":3,"u64":3,"s32":2}},"strings":{"sk1":{"str":"sk1","bytes":"MDs="},"nil":{"str":"nil"}},"bools":{"true":{"b":true},"false":{"b":false}},"messages":{"msg1":{"type":"BOOL","number":{},"string":{"str":"msg1","bytes":"MDs="},"bool":{"b":true}}},"arrays":{"arr1":{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"arrays":[{"numbers":[{},{}],"strings":[{"str":"arr1","bytes":"MDs="},{"str":"arr2","bytes":"MDs="}],"bools":[{"b":true},{"b":false}],"messages":[{"type":"BOOL","number":{},"string":{"str":"arr1_msg1","bytes":"MDs="},"bool":{"b":true}},{"type":"BOOL","number":{},"string":{"str":"arr1_msg2","bytes":"MDs="},"bool":{"b":true}}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}],"types":["BOOL","NUMBER","STRING"],"u32s":[0,1,2,3],"strs":["str1","str2","str3"]}},"types":{"0":"BOOL","1":"NUMBER","2":"STRING"},"u32s":{"u32_1":1,"u32_2":2,"u32_3":3},"strs":{"str1":"str1","str2":"str2"},"empties":{"empty3":{},"empty1":{},"empty2":{}},"optionals":{"optional1":{"number":{"u32":1},"string":{"str":"str1"}},"optional2":{"number":{"u32":2},"string":{"str":"str2"}}}}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Asserts(t, tt.args, tt.want)
		})
	}
}
