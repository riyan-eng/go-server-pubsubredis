package util

import "fmt"

type templateStruct struct{}

func NewTemplate() *templateStruct {
	return &templateStruct{}
}

func (t *templateStruct) EmailResetPassword(token string, expired string) (template string) {
	urlFe := fmt.Sprintf(`%v/setel-ulang-password/?token=%v`, "https://google.com", token)
	template = fmt.Sprintf(`
		<div style="flex: auto; text-align: center;">
			<h1>Sandbox Indonesia</h1>
			<p>Berikut adalah link untuk melakukan reset password: <a href="%v">%v</a></p>
			<p>hanya berlaku sampai %v, dan hanya bisa digunakan satu kali.</p>
		</div>
	`, urlFe, urlFe, expired)
	return
}
