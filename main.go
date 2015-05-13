package logstash

// https://github.com/bobrik/logstasher/blob/master/writer.go
import (
	"encoding/json"
	"io"
	"time"
)

type Writer struct {
	w      io.Writer
	fields map[string]interface{}
}

func New(w io.Writer, f map[string]interface{}) Writer {
	return Writer{
		w:      w,
		fields: f,
	}
}

func (w Writer) Write(p []byte) (n int, err error) {
	event := map[string]interface{}{
		"@timestamp": time.Now(),
		"message":    string(p),
	}

	for k, v := range w.fields {
		event[k] = v
	}

	return len(p), json.NewEncoder(w.w).Encode(event)
}
