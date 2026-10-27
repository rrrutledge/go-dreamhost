package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetProcessedRespBodyKeepsEscapedQuotes(t *testing.T) {
	body := `{"result":"success","data":[{"zone":"example.com","value":"0 issue \"letsencrypt.org\"","account_id":"1","comment":"","type":"CAA","record":"example.com","editable":"1"}]}`
	resp := http.Response{Body: io.NopCloser(strings.NewReader(body))}

	processed, err := getProcessedRespBody(resp)
	if err != nil {
		t.Fatalf("getProcessedRespBody: %v", err)
	}

	var list listDNSRecordsResponse
	if err := json.Unmarshal(processed, &list); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(list.Records) != 1 {
		t.Fatalf("got %d records, want 1", len(list.Records))
	}
	if want := `0 issue "letsencrypt.org"`; list.Records[0].Value != want {
		t.Errorf("value = %q, want %q", list.Records[0].Value, want)
	}
}
