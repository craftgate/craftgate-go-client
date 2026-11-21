package adapter

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateReportFromTemplateRequest_JSON(t *testing.T) {
	start := time.Date(2025, 11, 19, 9, 15, 31, 622824000, time.UTC)
	end := time.Date(2025, 11, 27, 9, 15, 31, 622843000, time.UTC)

	tests := []struct {
		name    string
		request CreateReportFromTemplateRequest
		want    string
	}{
		{
			"report date filter and file format",
			CreateReportFromTemplateRequest{
				Filters:     []ReportFilter{ReportDateFilter(start, end)},
				FileFormats: []ReportFileType{ReportFileType_XLSX},
			},
			`{"filters":[{"name":"reportDate","values":["2025-11-19T09:15:31","2025-11-27T09:15:31"]}],"fileFormats":["XLSX"]}`,
		},
		{
			"file formats omitted when not set",
			CreateReportFromTemplateRequest{
				Filters: []ReportFilter{ReportDateFilter(start, end)},
			},
			`{"filters":[{"name":"reportDate","values":["2025-11-19T09:15:31","2025-11-27T09:15:31"]}]}`,
		},
		{
			"filters keep their order",
			CreateReportFromTemplateRequest{
				Filters: []ReportFilter{
					{Name: "status", Values: []interface{}{"SUCCESS"}},
					ReportDateFilter(start, end),
				},
				FileFormats: []ReportFileType{ReportFileType_CSV},
			},
			`{"filters":[{"name":"status","values":["SUCCESS"]},{"name":"reportDate","values":["2025-11-19T09:15:31","2025-11-27T09:15:31"]}],"fileFormats":["CSV"]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatalf("Marshal() error = %v", err)
			}
			if string(got) != tt.want {
				t.Errorf("Marshal() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestFileReporting_DownloadReport(t *testing.T) {
	want := []byte("id,amount\n1,10.50\n")
	var gotMethod, gotPath, gotFileType string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotFileType = r.Method, r.URL.Path, r.URL.Query().Get("fileType")
		w.Header().Set("Content-Type", "application/octet-stream")
		_, _ = w.Write(want)
	}))
	defer server.Close()

	client, err := New("api-key", "secret-key", server.URL)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	got, err := client.FileReporting.DownloadReport(context.Background(), RetrieveReportRequest{FileType: ReportFileType_CSV}, 42)
	if err != nil {
		t.Fatalf("DownloadReport() error = %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("method = %s, want GET", gotMethod)
	}
	if gotPath != "/file-reporting/v1/reports/42/download" {
		t.Errorf("path = %s", gotPath)
	}
	if gotFileType != "CSV" {
		t.Errorf("fileType = %q, want CSV", gotFileType)
	}
	if string(got) != string(want) {
		t.Errorf("body = %q, want %q", got, want)
	}
}
