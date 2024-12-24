package httputil

import "net/http"

type WrappedResponseWriter struct {
	StatusCode  int
	ResponseLen int

	w           http.ResponseWriter
	wroteHeader bool
}

// 创建一个新的WrappedResponseWriter实例
func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	// 返回一个WrappedResponseWriter实例，初始状态码为200
	return &WrappedResponseWriter{
		StatusCode: 200,
		w:          w,
	}
}

func (w *WrappedResponseWriter) Header() http.Header {
	return w.w.Header()
}

// Write 方法用于将数据写入到底层的 http.ResponseWriter 中，并追踪已写入数据的长度。
// 该方法实现了 http.ResponseWriter 接口的 Write 方法。
// 参数:
//   - bytes []byte: 要写入的数据，以字节切片形式表示。
//
// 返回值:
//   - int: 成功写入的字节数。
//   - error: 如果写入过程中发生错误，返回该错误；如果没有错误，则返回 nil。
func (w *WrappedResponseWriter) Write(bytes []byte) (int, error) {
	// 调用底层的 http.ResponseWriter 的 Write 方法写入数据。
	n, err := w.w.Write(bytes)
	// 更新记录的已写入数据长度。
	w.ResponseLen += n
	// 返回写入的字节数和可能的错误。
	return n, err
}

// WriteHeader sets the HTTP status code to StatusCode and
// writes the headers. Only the first call to WriteHeader has
// any effect, and subsequent calls are ignored.
func (w *WrappedResponseWriter) WriteHeader(StatusCode int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true
	w.StatusCode = StatusCode
	w.w.WriteHeader(StatusCode)
}
