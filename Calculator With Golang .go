package main

import (
   "fmt"
   "math"
)

var (
   menuutama, menu,cek int
   celcius,farenheit,reamur int
   sisi,alas,tinggi,kecepatan,percepatan,jarak,jari,kelvin,waktu,banyak_gelombang  float64
   cel,far,r1,r2,massa,volume,gravitasi,radian,kecepatan_awal,luas float64
   sin,cos,tan,gaya,usaha,tekanan,daya float64
   pi float64 = 3.14
)

func main() {
   fmt.Println("Aplikasi Tugas Golang ")
   fmt.Println("1. Luas Persegi")
   fmt.Println("2. Luas Segitiga")
   fmt.Println("3. Luas Lingkaran")
   fmt.Println("4. Sudut Cosinus, Sinus, dan Tangen")
   fmt.Println("5. Gerak Lurus Beraturan")
   fmt.Println("6. Gerak Lurus Berubah Beraturan")
   fmt.Println("7. Energi Kinetik dan Energi Potensial")
   fmt.Println("8. Frekuensi dan Getaran")
   fmt.Println("9. Massa Jenis")
   fmt.Println("10. Gaya, Tekanan, Usaha dan Daya")
   fmt.Println("11. Konversi Suhu")
   fmt.Println("12. Keluar")
   fmt.Println("=========================================")
   fmt.Print("Pilih salah satu menu diatas (masukkan nomor menu): ")
   fmt.Scanf("%d", &menuutama)

switch menuutama{
   case 1 :
      luas_persegi()
   case 2 :
      luas_segitiga()
   case 3 :
      luas_lingkaran()
   case 4 :
      sin_cos_tan()
   case 5 :
      GLB()
   case 6 :
      GLBB()
   case 7 :
      kinetik_potensial()
   case 8 :
      frekuensi_getaran()
   case 9 :
      massa_jenis()
   case 10 :
      gaya_tekanan_usaha_daya()
   case 11 :
      Suhu()
   default :
      exit()
   }
}

//=====================LUAS PERSEGI=======================//
func luas_persegi(){
   fmt.Println("Menghitung Luas Persegi")
   fmt.Print("Masukkan sisi persegi : ")
   fmt.Scanf("%f", &sisi)
   luas := sisi*sisi
   fmt.Println("Luas Persegi= ", luas)
   main()
}

//=======================LUAS SEGITIGA====================//
func luas_segitiga(){
   fmt.Println("Menghitung Luas Segitiga")
   fmt.Print("Masukkan alas segitiga : ")
   fmt.Scanf("%f", &alas)
   fmt.Print("Masukkan tinggi segitiga : ")
   fmt.Scanf("%f", &tinggi)
   luas := (alas*tinggi)/2
   fmt.Println("Luas Persegi= ", luas)
   main()
}

//===================LUAS LINGKARAN======================//
func luas_lingkaran(){
   fmt.Println("Menghitung Luas Lingkaran")
   fmt.Print("Masukkan jari-jari lingkaran : ")
   fmt.Scanf("%f", &jari)
   luas = pi*(jari*jari)
   fmt.Println("Luas Lingkaran= \n", luas)
   main()
}

//============SUDUT COSINUS, SINUS, DAN TANGEN=========//
func sin_cos_tan(){
   fmt.Print("Masukkan nilai dalam satuab radian = ")
   fmt.Scanf("%f", &radian)
   sin := math.Sin(radian)
   cos := math.Cos(radian)
   tan := math.Tan(radian)

   fmt.Println("Hasil hitung sudut Sin = ", sin)
   fmt.Println("Hasil hitung sudut Cos = ", cos)
   fmt.Println("Hasil hitung sudut Tan = ", tan)
   main()
}

//======================GLB======================//
func GLB(){
   fmt.Println(" Menghitung Kecepatan dalam Gerak Lurus Beraturan ")
   fmt.Print("Masukkan jarak = ")
   fmt.Scanf("%f", &jarak)
   fmt.Print("Masukkan waktu = ")
   fmt.Scanf("%f", &waktu)
   kecepatan := jarak/waktu
   fmt.Println("Kecepatan = ", kecepatan)
   main()
}

//=======================GLBB=====================//
func GLBB() {
   fmt.Println("Menghitung Gerak Lurus Berubah Beraturan")
   fmt.Println("Pilih Menu rumus ")
   fmt.Println("1. Jarak ")
   fmt.Println("2. Kecepatan Akhir")
   fmt.Println("3. Kecepatan Akhir kuadrat")
   fmt.Println("4. Menu Utama")
   fmt.Println("Masukkan pilihan menu (masukkan nomor menu) ")
   fmt.Scanf("%d", &menu)

   switch menu{
      case 1 :
         pilih1()
      case 2 :
          pilih2()
      case 3 :
         pilih3()
      default :
         main()
   }
}

func pilih1(){
   fmt.Print("Masukkan kecepatan awal = ")
   fmt.Scanf("%f", &kecepatan_awal)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &percepatan)
   fmt.Print("Masukkan nilai waktu= ")
   fmt.Scanf("%f", &waktu)
   GLBB := kecepatan_awal + 0.5 * percepatan*waktu*waktu
   fmt.Println("Nilai Gerak Lurus Berubah Beraturan = ", GLBB)
}

func pilih2(){
   fmt.Println("=====Nomor 2=====")
   fmt.Print("Masukkan kecepatan awal = ")
   fmt.Scanf("%f", &kecepatan_awal)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &percepatan)
   fmt.Print("Masukkan nilai waktu= ")
   fmt.Scanf("%f", &waktu)
   GLBB := kecepatan_awal-percepatan*waktu
   fmt.Println("Nilai Gerak Lurus Berubah Beraturan = ", GLBB)
}

func pilih3(){
   fmt.Println("=====Nomor 3=====")
   fmt.Print("Masukkan kecepatan awal = ")
   fmt.Scanf("%f", &kecepatan_awal)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &percepatan)
   fmt.Print("Masukkan nilai jarak= ")
   fmt.Scanf("%f", &jarak)
   GLBB := (kecepatan_awal*kecepatan_awal) + 2*percepatan*jarak
   fmt.Println("Nilai Gerak Lurus Berubah Beraturan = ", GLBB)
}

//========================EK dan EP==================//
func kinetik_potensial(){
   fmt.Println("1. Energi Kinetik")
   fmt.Println("2. Energi Potensial")
   fmt.Println("3. Menu Utama")
   fmt.Print("Pilih salah satu menu diatas ini (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

   if menu == 1 {
      kinetik()
   }else if menu == 2{
      potensial()
   }else{
      main()
   }
}

func kinetik(){
   fmt.Println("=====Mengukur Energi Kinetik=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &massa)
   fmt.Print("Masukkan nilai kecepatan = ")
   fmt.Scanf("%f", &kecepatan)
   kinetik := 0.5*(massa*(kecepatan*kecepatan))
   fmt.Println("Besar Energi Kinetik = \n", kinetik)
   kinetik_potensial()
}

func potensial(){
   fmt.Println("=====Mengukur Energi Potensial=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &massa)
   fmt.Print("Masukkan nilai gravitasi = ")
   fmt.Scanf("%f", &gravitasi)
   fmt.Print("Masukkan nilai tinggi = ")
   fmt.Scanf("%f", &tinggi)
   potensial := massa*gravitasi*tinggi
   fmt.Println("Besar Energi Potensial = \n", potensial)
   kinetik_potensial()
}

//=================FREKUENSI atau GETARAN=================//
func frekuensi_getaran() {
   fmt.Println("Mengitung frekuensi dan getaran ")
   fmt.Println("Pilih menu sesuai keinginan")
   fmt.Println("1. Frekuensi")
   fmt.Println("2. Getaran")
   fmt.Println("3. Menu Utama")
   fmt.Println("Masukkan pilihan menu (masukkan nomor menu) ")
   fmt.Scanf("%d", &menu)

   switch menu{
      case 1 :
         frekuensi()
      case 2 :
         getaran()
      default :
         main()
   }
}

func frekuensi(){
   fmt.Println("=====Mengukur Frekuensi=====")
   fmt.Print("Masukkan periode = ")
   fmt.Scanf("%f", &banyak_gelombang)
   frekuensi := 1 / banyak_gelombang
   fmt.Println("Nilai Frekuensi = ", frekuensi)
   main()
}

func getaran(){
   fmt.Println("=====Mengukur Getaran=====")
   fmt.Print("Masukkan jumlah getaran = ")
   fmt.Scanf("%f", &banyak_gelombang)
   fmt.Print("Masukkan nilai t = ")
   fmt.Scanf("%f", &waktu)
   frek := banyak_gelombang / waktu
   fmt.Println("Nilai Getaran = ", frek)
   main()
}

//======================MASSA JENIS==================//
func massa_jenis(){
   fmt.Println("=====Mengukur Massa Jenis=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &massa)
   fmt.Print("Masukkan nilai volume = ")
   fmt.Scanf("%f", &volume)
   massa := massa/volume
   fmt.Println("Nilai Massa Jenis = ", massa)
   main()
}

//===========GAYA, TEKANAN, USAHAN DAN DAYA==========//
func gaya_tekanan_usaha_daya(){
   fmt.Println("Menu")
   fmt.Println("1. Gaya")
   fmt.Println("2. Tekanan")
   fmt.Println("3. Usaha")
   fmt.Println("4. Daya")
   fmt.Println("5. Memu Utama")
   fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

switch menu{
   case 1 :
      Gaya();
   case 2 :
      Tekanan()
   case 3 :
      Usaha()
   case 4 :
      Daya()
   default :
      main()
   }
}

func Gaya(){
   fmt.Println("Menghitung besar gaya ")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &massa)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &percepatan)
   gaya := massa*percepatan
   fmt.Print("Besar Gaya = ", gaya)
   gaya_tekanan_usaha_daya()
}

func Tekanan(){
   fmt.Println("Menghitung besar tekanan ")
   fmt.Print("Masukkan nilai gaya = ")
   fmt.Scanf("%f", &gaya)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%d", &percepatan)
   tekanan := gaya/percepatan
   fmt.Print("Besar Tekanan = ", tekanan)
   gaya_tekanan_usaha_daya()
}

func Usaha(){
   fmt.Println("Menghitung besar usaha ")
   fmt.Print("Masukkan nilai gaya = ")
   fmt.Scanf("%f", &gaya)
   fmt.Print("Masukkan jarak = ")
   fmt.Scanf("%f", &jarak)
   usaha := gaya*jarak
   fmt.Print("Besar Usaha = ", usaha)
   gaya_tekanan_usaha_daya()
}

func Daya(){
   fmt.Println("Menghitung besar daya ")
   fmt.Print("Masukkan nilai usaha = ")
   fmt.Scanf("%f", &usaha)
   fmt.Print("Masukkan waktu = ")
   fmt.Scanf("%f", &waktu)
   daya := usaha/waktu
   fmt.Print("Besar Daya = ", daya)
   gaya_tekanan_usaha_daya()
}

//========================SUHU==================//
func Suhu(){
   fmt.Println("Konversi Suhu")
   fmt.Println("1. Celcius")
   fmt.Println("2. Farenheit")
   fmt.Println("3. Reamur")
   fmt.Println("4. Kelvin")
   fmt.Println("5. Menu Utama")
   fmt.Print("Pilih salah satu suhu yang ingin dikonversi pada menu diatas (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

switch menu{
   case 1 :
      fmt.Println("1. Celcius - Farenheit")
      fmt.Println("2. Celcius - Reamur")
      fmt.Println("3. Celcius - Kelvin")
      fmt.Println("4. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &celcius)

      switch celcius{
         case 1 :
            CF()
         case 2 :
            CR()
         case 3 :
            CK()
         default :
            Suhu()
      }
   case 2 :
      fmt.Println("1. Farenheit - Celcius")
      fmt.Println("2. Farenheit - Reamur")
      fmt.Println("3. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &farenheit)

      switch farenheit{
         case 1 :
            FC()
         case 2 :
            FR()
         default :
            Suhu()
      }
   case 3 :
      fmt.Println("1. Reamur - Celcius")
      fmt.Println("2. Reamur - Farenheit")
      fmt.Println("3. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &reamur)

      switch reamur{
         case 1 :
            RC()
         case 2 :
            RF()
         default :
            Suhu()
         }
   case 4 :
      fmt.Println("konversi Kelvin - Celcius")
      fmt.Print("Masukkan nilai suhu Kelvin = ")
      fmt.Scanf("%f", &kelvin)
      KC := kelvin-273
      fmt.Print("Nilai Suhu dalam Celcius = ", KC)
      Suhu()
   default :
      main()
}
}

//-----------------------------------------------------------------------------------------------//
//CELCIUS

func CF(){
   var cek int
   fmt.Println("konversi Celcius - Farenheit")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := ((9*cel)+32)/5
   fmt.Println("Nilai Suhu dalam Fareinheit =  ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

func CR(){
   var cek int
   fmt.Println("konversi Celcius - Reamur")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := cel*4/5
   fmt.Println("Nilai Suhu dalam Reamur = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
      exit()
   default :
      Suhu()
   }
}

func CK(){
   var cek int
   fmt.Println("konversi Celcius - Kelvin")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := cel+273
   fmt.Println("Nilai Suhu dalam Kelvin = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//--------------------------------------------------------------------------------------------//
// FARENHEIT

func FC(){
   var cek int
   fmt.Println("konversi Farenheit - Celcius")
   fmt.Print("Masukkan nilai suhu Farenheit = ")
   fmt.Scanf("%f", &far)
   hasil := 5/9*(far-32)
   fmt.Println("Nilai Suhu dalam Celcius = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
        exit()
      default :
         Suhu()
   }
}

func FR(){
   var cek int
   fmt.Println("konversi Farenheit - Reamur")
   fmt.Print("Masukkan nilai suhu Farenheit = ")
   fmt.Scanf("%f", &far)
   hasil := 4*(far-32)/9
   fmt.Println("Nilai Suhu dalam Reamur = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//-------------------------------------------------------------------------------------------//
// REAMUR

func RC(){
   var r1 float64
   var cek int

   fmt.Println("konversi Reamur - Celcius")
   fmt.Print("Masukkan nilai suhu Reamur = ")
   fmt.Scanf("%f", &r1)
   hasil := (r1*5)/4
   fmt.Println("Nilai Suhu dalam Celcius = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

func RF(){
   var r2 float64
   var cek int

   fmt.Println("konversi Reamur - Farenheit")
   fmt.Print("Masukkan nilai suhu Reamur = ")
   fmt.Scanf("%f", &r2)
   hasil := (r2*9)/4+32
   fmt.Println("Nilai Suhu dalam Farenheit = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//=========================EXIT=====================//
func exit(){
   fmt.Println("====Selesai===")
}