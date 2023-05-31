package pk2

import "fmt"
import "matheus/teste/pk1"

func Testprint() {
	pk1.Testprint()
	fmt.Println("chamei o Testprint do pk1 e printei no pk2");
}
