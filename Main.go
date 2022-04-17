package main
import "fmt"

func main(){

	var userNama string
	var dataArray = [3]string{"Lala","Lilil","Lulut"}

	// fmt.Printf("Semua Data =  %v\n",dataArray) //tampilan semua data yg ada di array
	// fmt.Printf("data pertama = %v\n",dataArray[1]) //tampilan data array sesuai urutan
	
	// fmt.Printf("Masukan Nama User = ")
	// fmt.Scan(&userNama)//inputan ke variabel biasa

	// fmt.Printf("Masukan Nama Array 2 = ")
	// fmt.Scan(&dataArray[2])//inputan ke variabel array spesifik

	// fmt.Printf("namanya %v\n",userNama)
	// fmt.Printf("data terakhir = %v\n",dataArray[2])//ngambil array berdasarkan urutan
	// fmt.Printf("data terakhir = %v\n",len(dataArray[2])) //ngambil panjang array
	// fmt.Printf("data terakhir = %v\n",len(dataArray)) //ngambil jumlah data array ada berapa
	// fmt.Printf("Semua Data =  %v\n",dataArray)

	//array appen
	var arrayAppend[]string

	fmt.Printf("Masukan Nama User = ")
	fmt.Scan(&userNama)//inputan ke variabel biasa
	arrayAppend = append(arrayAppend, userNama)

	fmt.Printf("Semua Data =  %v\n",arrayAppend) //tampilan semua data yg ada di array
}
