package neverhaveiever

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestController_TopicRecordsHandler(t *testing.T) {
	type fields struct {
		fileDataByName map[string][]byte
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected string
	}{
		{
			name: "Get file1",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3"),
					"file2": []byte("data2"),
					"file3": []byte(""),
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": ["file1"],"page": 1}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"records":["data1.1","data1.2","data1.3"],"count_page":1}`,
		},
		{
			name: "Get page out of range of file1",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3"),
					"file2": []byte("data2"),
					"file3": []byte(""),
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": ["file1"],"page": 100}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"records":[],"count_page":1}`,
		},
		{
			name: "Get second page of file1",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3"),
					"file2": []byte("data2"),
					"file3": []byte(""),
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": ["file1"],"page": 2,"page_size": 1}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"records":["data1.2"],"count_page":3}`,
		},
		{
			name: "Not found file",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3"),
					"file2": []byte("data2"),
					"file3": []byte(""),
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": ["notFoundFile"],"page": 1}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"error":"open games/never_have_i_ever/topics/notFoundFile: no such file or directory"}`,
		},
		{
			name: "Get empty file",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3"),
					"file2": []byte("data2"),
					"file3": []byte(""),
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": ["file3"],"page": 1}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"records":[""],"count_page":1}`,
		},
		{
			name: "Empty topic",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3"),
					"file2": []byte("data2"),
					"file3": []byte(""),
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": [],"page": 1}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"error":"topics is empty"}`,
		},
		{
			name: "Check default page_size in request",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3\ndata1.4\ndata1.5\ndata1.6\ndata1.7\ndata1.8\ndata1.9\ndata1.10\ndata1.11"),
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": ["file1"],"page": 1}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"records":["data1.1","data1.2","data1.3","data1.4","data1.5","data1.6","data1.7","data1.8","data1.9","data1.10"],"count_page":2}`,
		},
		{
			name: "Get some files",
			fields: fields{
				fileDataByName: map[string][]byte{
					"file1": []byte("data1.1\ndata1.2\ndata1.3"),
					"file2": []byte("data2"),
					"file3": nil,
				},
			},
			args: args{
				r: &http.Request{
					Method: "POST",
					Body:   io.NopCloser(bytes.NewReader([]byte(`{"topics": ["file1","file2","file3"],"page": 1}`))),
				},
				w: httptest.NewRecorder(),
			},
			expected: `{"records":["data1.1","data1.2","data1.3","data2"],"count_page":1}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &Controller{
				fileDataByName: tt.fields.fileDataByName,
			}
			controller.TopicRecordsHandler(tt.args.w, tt.args.r)

			res := tt.args.w.(*httptest.ResponseRecorder).Result()
			data, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.expected, string(data))
		})
	}
}
