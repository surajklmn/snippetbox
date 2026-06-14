package main 
import(
	"net/http"
	"fmt"
)

func commonHeaders(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Security-Policty","default-src 'self;style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy","origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options","nosniff")
		w.Header().Set("X-Frame-Options","deny")
		w.Header().Set("X-XSS-Protection","0")

		//Our custom header
		w.Header().Set("Server","Go")
		//Any code before next.ServeHTTP will execute on the way down the chain
		next.ServeHTTP(w,r)
		//Any code before next.ServeHTTP will execute on the way up the chain
		//THE CHAIN: commonHeaders → servemux → application handler → servemux → commonHeaders
	})
}
