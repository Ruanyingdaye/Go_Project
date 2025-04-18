package main


// 可以利用结构体进行结构封装
type loggingResponseWriter struct {
	ResponseWriter http.ResponseWriter
	StatusCode     int
}
func MidwareLogger(next http.Handler) http.Handler {{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{
			ResponseWriter: w,
		}
		// 定义了一个lrw
		// 本质就是http.ResponseWriter去写内容

		next.ServeHTTP(w, r)
		log.Printf(r.Method, r.URL.Path, time.Since(Start), lrw.statusCode)

		// output file
		file, _ := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		log.SetOutput(file)

		// output json
		logEntry := map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL.Path,
			"ip":       r.RemoteAddr,
			"status":   lrw.status,
			"duration": time.Since(start).Milliseconds(),
		}
		logJSON, _ := json.Marshal(logEntry)
		log.Println(string(logJSON))
	})
}

func (lrw *loggingResponseWriter)WriteHeader(code int){
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
// 手动写header添加status code的作用是？

func MidwareValidate(next http.Handler) http.Handler {

}

func MidwareLimitTraffic(next http.Handler) http.Handler {

}


func main(){



}