package coincap

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// loggingRoundTripper - структура для логирования HTTP-запросов и ответов.
type loggingRoundTripper struct {
	logger io.Writer //  записи логов
	next   http.RoundTripper
}

// HTTP-запрос с логированием.
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	// Передаем запрос дальше
	return l.next.RoundTrip(r)
}
