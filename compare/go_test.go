package compare

import (
	"reflect"
	"testing"
)

func TestCompareMaps(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		m1    map[string]interface{}
		m2    map[string]interface{}
		equal bool
	}{
		{
			name:  "Empty maps",
			m1:    map[string]interface{}{},
			m2:    map[string]interface{}{},
			equal: true,
		},
		{
			name: "Equal maps",
			m1: map[string]interface{}{
				"foo": "bar",
				"baz": 123,
			},
			m2: map[string]interface{}{
				"baz": 123,
				"foo": "bar",
			},
			equal: true,
		},
		{
			name: "Different values",
			m1: map[string]interface{}{
				"foo": "bar",
				"baz": 123,
			},
			m2: map[string]interface{}{
				"baz": 456,
				"foo": "qux",
			},
			equal: false,
		},
		{
			name: "Different keys",
			m1: map[string]interface{}{
				"foo": "bar",
			},
			m2: map[string]interface{}{
				"baz": "qux",
			},
			equal: false,
		},
		{
			name: "Nested maps",
			m1: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": 123,
				},
				"baz": "qux",
			},
			m2: map[string]interface{}{
				"baz": "qux",
				"foo": map[string]interface{}{
					"bar": 123,
				},
			},
			equal: true,
		},
		{
			name: "Different nested values",
			m1: map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": 123,
				},
				"baz": "qux",
			},
			m2: map[string]interface{}{
				"baz": "qux",
				"foo": map[string]interface{}{
					"bar": 456,
				},
			},
			equal: false,
		},
		{
			name: "Nested arrays",
			m1: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
			m2: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
			equal: true,
		},
		{
			name: "Different nested array values",
			m1: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
			m2: map[string]interface{}{
				"foo": []interface{}{"bar", "qux"},
			},
			equal: false,
		},
		{
			name: "Different nested array lengths",
			m1: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
			m2: map[string]interface{}{
				"foo": []interface{}{"bar", "baz", "qux"},
			},
			equal: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			equal := compareMaps(test.m1, test.m2)
			if equal != test.equal {
				t.Errorf("Expected equal=%v, but got equal=%v", test.equal, equal)
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	type args struct {
		m1 map[string]interface{}
		m2 map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "test mergeMaps",
			args: args{
				m1: map[string]interface{}{
					"key1": "value1",
					"key2": map[string]interface{}{
						"key21": "value21",
						"key22": "value22",
					},
					"key3": []map[string]interface{}{
						{
							"id":   "1",
							"name": "name1",
						},
						{
							"id":   "2",
							"name": "name2",
						},
						{
							"aaa": "bbb",
						},
					},
				},
				m2: map[string]interface{}{
					"key1": "value1",
					"key2": map[string]interface{}{
						"key21": "value21",
						"key23": "value23",
					},
					"key3": []map[string]interface{}{
						{
							"id":   "1",
							"name": "name1 ++",
						},
						{
							"id":   "3",
							"name": "name3",
						},
						{
							"name":  "name4",
							"extra": "xxx",
						},
					},
				},
			},
			want: map[string]interface{}{
				"key1": "value1",
				"key2": map[string]interface{}{
					"key21": "value21",
					"key22": "value22",
					"key23": "value23",
				},
				"key3": []map[string]interface{}{
					{
						"id":   "1",
						"name": "name1 ++",
					},
					{
						"id":   "2",
						"name": "name2",
					},
					{
						"aaa": "bbb",
					},
					{
						"id":   "3",
						"name": "name3",
					},
					{
						"name":  "name4",
						"extra": "xxx",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeMaps(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}

}
