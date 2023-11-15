package main

import (
   "net/http/httptest"
   "testing"
)

func TestHelloWorld(t *testing.T) {
   req := httptest.NewRequest("GET", "/", nil)
   w := httptest.NewRecorder()
   s := Server{}

   s.ServeHTTP(w, req)

   if w.Result().StatusCode != 200 {
      t.Fatalf("unexpected status code %d", w.Result().StatusCode)
}
body := w.Body.String()
if body != `{"message": "hello world"}` {
    t.Fatalf("unexpected body received: %s", body)
}<span></span>
}