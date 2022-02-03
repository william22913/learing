package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/william22913/learning/pdf/util"
)

func main() {
	pdf := initPDF()

	// ------------------------------
	// Creating Transaction Header,
	// Passing the struct to this function to fill the transaction header
	header, data := getTransactionHeader()
	a := util.AddTable(pdf, header, data, 5, false, false, 1, 3, 1)
	// ------------------------------

	// ------------------------------
	// Creating Store name ,
	// Passing the struct/Storename to this function to fill the transaction header
	header, data = getStoreName()

	// +5 for give the space from store name to transaction header
	a = a + 5
	a = util.AddTable(pdf, header, data, a, false, false, 1, 3, 1)
	// ------------------------------

	header, data = getTransactionTable()
	a = util.AddTable(pdf, header, data, a, true, true, 1, 3, 1)

	a = createDiscountTable(pdf, a)
	err := pdf.OutputFileAndClose("./test.pdf")

	fmt.Println(err, a)
}

func createDiscountTable(pdf *gofpdf.Fpdf, currentPosition float64) float64 {
	header, data := getTransactionDetail()
	output := util.AddTable(pdf, header, data, currentPosition, true, false, 1, 3, 1)

	header, data = getDiscount()
	currentPosition = util.AddTable(pdf, header, data, currentPosition, true, false, 1, 3.65, 8+25+50+10+10+10+18+15+1)

	header, data = getTotalAmount()
	_ = util.AddTable(pdf, header, data, currentPosition, true, false, 1, 3.7, 8+25+50+10+10+10+18+15+1)

	return output
}

func getTotalAmount() (header []util.PDFTableHeader, data [][]string) {
	data = append(data,
		[]string{"Total Amount :", "1,103,851.782"},
	)

	header = []util.PDFTableHeader{}
	header = append(header,
		util.PDFTableHeader{
			ColumnWidth: 25,
			IsBold:      true,
		}, util.PDFTableHeader{
			ColumnWidth: 25,
			IsBold:      true,
		},
	)
	return
}

func getDiscount() (header []util.PDFTableHeader, data [][]string) {
	data = append(data,
		[]string{"Discount :", "10,136.38"},
		[]string{"Total Tax :", "100,000.162"},
	)

	header = []util.PDFTableHeader{}
	header = append(header,
		util.PDFTableHeader{
			ColumnWidth: 25,
		}, util.PDFTableHeader{
			ColumnWidth: 25,
		},
	)
	return
}

func getTransactionTable() (header []util.PDFTableHeader, data [][]string) {
	data = append(data,
		[]string{"1", "828381232113545", "Test Barang 1", "1", "12", "CT", "245,455", "1%", "243,000.45", "267,300.495"},
		[]string{"2", "8283812321", "Test Barang 2", "1", "12", "CT", "245,455", "1%", "243,000.45", "267,300.495"},
		[]string{"3", "8283812321", "Test Barang 3 Test Barang 3 Test Barang 3", "1", "12", "CT", "245,455", "1%", "243,000.45", "267,300.495"},
		[]string{"4", "8283812321", "Test Barang 4", "1", "12", "CT", "245,455", "1%", "243,000.45", "267,300.495"},
	)

	header = []util.PDFTableHeader{}
	header = append(header,
		util.PDFTableHeader{
			Name:        "No.",
			ColumnWidth: 8,
		}, util.PDFTableHeader{
			Name:        "Barcode",
			ColumnWidth: 25,
		}, util.PDFTableHeader{
			Name:        "Item Name",
			ColumnWidth: 50,
		}, util.PDFTableHeader{
			Name:        "Qty",
			ColumnWidth: 10,
		}, util.PDFTableHeader{
			Name:        "Qty Base",
			ColumnWidth: 10,
		}, util.PDFTableHeader{
			Name:        "Unit",
			ColumnWidth: 10,
		}, util.PDFTableHeader{
			Name:        "Price",
			ColumnWidth: 18,
		}, util.PDFTableHeader{
			Name:        "Discount",
			ColumnWidth: 15,
		}, util.PDFTableHeader{
			Name:        "Sub Total",
			ColumnWidth: 25,
		}, util.PDFTableHeader{
			Name:        "Sub Total + PPN",
			ColumnWidth: 25,
		},
	)
	return
}

func getTransactionDetail() (header []util.PDFTableHeader, data [][]string) {
	text :=
		"Ship To: \n" +
			"JL. Pemuda Siliwangi No. 126, RT. 004 RW 001, Kelurahan Bojong Rawalumbu, Kecamatan Rawa Lumbu, Kota Bekasi, Provinsi Jawa Barat\n" +
			"Notes:\n" +
			"Apabila Terdapat Perbedaan Harga, Barang Jangan Dikirim dan Segera Hubungi MD yang Bersangkutan"
	data = append(data, []string{text})

	header = []util.PDFTableHeader{}
	header = append(header,
		util.PDFTableHeader{
			//you can adjust this by looking up transaction header's column width
			ColumnWidth: 8 + 25 + 50 + 10 + 10 + 10 + 18 + 15,
		},
	)

	return
}

func getStoreName() (header []util.PDFTableHeader, data [][]string) {

	data = append(data, []string{"PO GS KEMANG"})

	header = []util.PDFTableHeader{}
	header = append(header,
		util.PDFTableHeader{
			//you can adjust this by looking up transaction header's column width
			ColumnWidth: 40 * 4,
		},
	)

	return

}

func getTransactionHeader() (header []util.PDFTableHeader, data [][]string) {

	data = append(data, []string{"Purchase Order No.", ": PO-GS004/2201/85586", "Print Date", ": 31/01/2022 09:15"})
	data = append(data, []string{"Trans. Date.", ": 31/01/2022", "Due Date", ": 14/02/2022"})
	data = append(data, []string{"Vendor", ": Nestle", "Created By", ": BAGAS"})
	data = append(data, []string{"Payment Type", ": TUNAI", "Confirmed By", ": BAGAS"})
	data = append(data, []string{"Expired Date", ": 05/02/2022", "Delivery Date", ": 31/01/2022"})

	header = append(header,
		util.PDFTableHeader{
			ColumnWidth: 40,
		}, util.PDFTableHeader{
			ColumnWidth: 40,
		}, util.PDFTableHeader{
			ColumnWidth: 40,
		}, util.PDFTableHeader{
			ColumnWidth: 40,
		},
	)

	return
}

func initPDF() *gofpdf.Fpdf {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 7)

	return pdf
}
