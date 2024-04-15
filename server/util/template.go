package util

import (
	"fmt"

	"server/env"

	"github.com/golang-jwt/jwt/v5"
)

type templateStruct struct{}

func NewTemplate() *templateStruct {
	return &templateStruct{}
}

func (t *templateStruct) EmailResetPassword(token string, expired *jwt.NumericDate) (template string) {
	urlFe := fmt.Sprintf(`%v/setel-ulang-password/?token=%v`, env.NewEnvironment().SERVER_HOST_FE, token)
	template = fmt.Sprintf(`
		<div style="flex: auto; text-align: center;">
			<h1>Sandbox Indonesia</h1>
			<p>Berikut adalah link untuk melakukan reset password: <a href="%v">%v</a></p>
			<p>hanya berlaku sampai %v, dan hanya bisa digunakan satu kali.</p>
		</div>
	`, urlFe, urlFe, expired)
	return
}

func (t *templateStruct) PDFExample() (template string) {
	template = `
	<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body style="font-family: 'Times New Roman', Times, serif; margin: 0; font-size: 12pt; background-color: white;">
    <table style="vertical-align: top; width: 100%;">
        <tbody>
            <tr>
                <td style="text-align: right; width: 15%;">
                    
                </td>
                <td style="text-align: center; color: #323296;">
                    <div style="font-size: 14pt;"><b>KEMENTERIAN PERTANIAN</b></div>
                    <div style="font-size: 14pt;"><b>DIREKTORAT JENDERAL TANAMAN PANGAN</b></div>
                    <div style="font-size: 12pt;"><b>BALAI BESAR PERAMALAN ORGANISME PENGGANGGU TUMBUHAN</b></div>
                    <div style="font-size: 10pt;">JALAN RAYA KALIASIN TROMOL POS 1 JATISARI KARAWANG 413374</div>
                    <div style="font-size: 10pt;">TELEPON / FAKSIMILI: (0264) 360581, 360368 e-mail:
                        bbpopt@pertanian.go.id</div>
                    <div style="font-size: 10pt;">WEBSITE: http://bbpopt.tanamanpangan.pertanian.go.id</div>
                </td>
            </tr>
        </tbody>
    </table>

    <hr style="height: 2px; border-top: solid 3px #323296; border-bottom: solid 5px #323296; background: white;">
    <br>
    <div style="text-align: center;">
        <div><b>NOTA DINAS</b></div>
    </div>
    <br>
    <table width="100%">
        <tbody>
            <tr>
                <td style="width: 60px;">Yth.</td>
                <td>:</td>
                <td>{{ke}}</td>
            </tr>
            <tr>
                <td style="width: 60px;">Dari</td>
                <td>:</td>
                <td>{{dari}}</td>
            </tr>
            <tr>
                <td style="width: 60px;">Hal</td>
                <td>:</td>
                <td>{{hal}}</td>
            </tr>
            <tr>
                <td style="width: 60px;">Tanggal</td>
                <td>:</td>
                <td>{{tanggal_nota_dinas}}</td>
            </tr>
        </tbody>
    </table>
    <hr style="height: 1px; background-color: black; border: none;">
    <div style="text-align: justify;">
        Dalam rangka pelaksanaan kegiatan Balai Besar Peramalan Organisme Pengganggu Tumbuhan Tahun {{tahun}}, bersama ini kami mengajukan penggunaan anggaran kegiatan sebagai berikut:
    </div>
    <br>
    <table style="padding-left: 15px;">
        <tbody style="vertical-align: top;">
            <tr>
                <td style="width: 10px;">1.</td>
                <td style="width: 250px;">Program ({{program_kode}})</td>
                <td style="width: 5px;">:</td>
                <td>{{program_nama}}</td>
            </tr>
            <tr>
                <td>2.</td>
                <td>Kegiatan ({{kegiatan_kode}})</td>
                <td>:</td>
                <td>{{kegiatan_nama}}</td>
            </tr>
            <tr>
                <td>3.</td>
                <td>Output ({{kegiatan_kode}}.{{output_kode}}.{{sub_output_kode}})</td>
                <td>:</td>
                <td>{{output_nama}}</td>
            </tr>
            <tr>
                <td>4.</td>
                <td>Komponen ({{komponen_kode}})</td>
                <td>:</td>
                <td>{{komponen_nama}}</td>
            </tr>
            <tr>
                <td>5.</td>
                <td>Sub Komponen ({{komponen_kode}}.{{sub_komponen_kode}})</td>
                <td>:</td>
                <td>{{sub_komponen_nama}}</td>
            </tr>
            <tr>
                <td>6.</td>
                <td>M.A.K / Akun</td>
                <td>:</td>
                <td>{{sub_komponen_kode}}.{{mak_kode}}</td>
            </tr>
            <tr>
                <td>7.</td>
                <td>Sub Kegiatan</td>
                <td>:</td>
                <td>{{sub_kegiatan}}</td>
            </tr>
            <tr>
                <td>8.</td>
                <td>Pagu Anggaran</td>
                <td>:</td>
                <td>Rp {{pagu_anggaran}},-</td>
            </tr>
            <tr>
                <td>9.</td>
                <td>Nama dan Waktu Pelaksanaan</td>
                <td>:</td>
                <td>Terlampir</td>
            </tr>
            <tr>
                <td>10.</td>
                <td>Kartu</td>
                <td>:</td>
                <td></td>
            </tr>
        </tbody>
    </table>
    <br>
    <div>
        Sehubungan dengan kegiatan tersebut di atas, maka kami mohon persetujuannya untuk pencairan dana.
    </div>
    <div>
        Demikian kami sampaikan, atas perhatiannya kami ucapkan terima kasih.
    </div>

    <br>

    <table width="100%" style="text-align: center;">
        <tbody>
            <tr>
                <td width="50%">
                    <div>Menyetujui,</div>
                </td>
                <td width="50%"></td>
            </tr>
            <tr>
                <td>Kepala Balai</td>
                <td>{{ketua_kelompok_jabatan}}</td>
            </tr>
            <tr>
            </tr>
            <tr>
                <td>{{kepala_balai_nama}}</td>
                <td>{{ketua_kelompok_nama}}</td>
            </tr>
            <tr>
                <td>{{kepala_balai_nip}}</td>
                <td>{{ketua_kelompok_nip}}</td>
            </tr>
        </tbody>
    </table>

</body>

</html>
	`
	return
}
