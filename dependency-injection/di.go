package dependencyinjection

import (
	"fmt"
	"io"
	"log"
	"net/http"
)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)

}
// meaning http.ResponseWriter also implements the io.Writer interface
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

}