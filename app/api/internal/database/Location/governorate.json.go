package location

type governorate struct {
	ID      int    `json:"id"`
	Number  int    `json:"number"`
	Arabic  string `json:"ar"`
	English string `json:"en"`
}

var GovernorateData = []governorate{
	{ID: 1, Number: 1, Arabic: "القاهرة", English: "Cairo"},
	{ID: 2, Number: 2, Arabic: "الجيزة", English: "Giza"},
	{ID: 3, Number: 3, Arabic: "الاسكندرية", English: "Alexandria"},
	{ID: 4, Number: 4, Arabic: "الدقهلية", English: "Dakahlia"},
	{ID: 5, Number: 5, Arabic: "البحر الأحمر", English: "Red Sea"},
	{ID: 6, Number: 6, Arabic: "البحيرة", English: "Beheira"},
	{ID: 7, Number: 7, Arabic: "الفيوم", English: "Fayoum"},
	{ID: 8, Number: 8, Arabic: "الغربية", English: "Gharbiya"},
	{ID: 9, Number: 9, Arabic: "الإسماعلية", English: "Ismailia"},
	{ID: 10, Number: 10, Arabic: "المنوفية", English: "Menofia"},
	{ID: 11, Number: 11, Arabic: "المنيا", English: "Minya"},
	{ID: 12, Number: 12, Arabic: "القليوبية", English: "Qaliubiya"},
	{ID: 13, Number: 13, Arabic: "الوادي الجديد", English: "New Valley"},
	{ID: 14, Number: 14, Arabic: "السويس", English: "Suez"},
	{ID: 15, Number: 15, Arabic: "اسوان", English: "Aswan"},
	{ID: 16, Number: 16, Arabic: "اسيوط", English: "Assiut"},
	{ID: 17, Number: 17, Arabic: "بني سويف", English: "Beni Suef"},
	{ID: 18, Number: 18, Arabic: "بورسعيد", English: "Port Sanumber"},
	{ID: 19, Number: 19, Arabic: "دمياط", English: "Damietta"},
	{ID: 20, Number: 20, Arabic: "الشرقية", English: "Sharkia"},
	{ID: 21, Number: 21, Arabic: "جنوب سيناء", English: "South Sinai"},
	{ID: 22, Number: 22, Arabic: "كفر الشيخ", English: "Kafr Al sheikh"},
	{ID: 23, Number: 23, Arabic: "مطروح", English: "Matrouh"},
	{ID: 24, Number: 24, Arabic: "الاقصر", English: "Luxor"},
	{ID: 25, Number: 25, Arabic: "قنا", English: "Qena"},
	{ID: 26, Number: 26, Arabic: "شمال سيناء", English: "North Sinai"},
	{ID: 27, Number: 27, Arabic: "سوهاج", English: "Sohag"},
}
