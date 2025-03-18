package cors

import "net/http"

func InterceptorHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем доступ с любых доменов
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Если это preflight-запрос (OPTIONS), сразу отдаем 200 OK
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Делаем вызов к следующему хэндлеру
		next.ServeHTTP(w, r)
	})
}
