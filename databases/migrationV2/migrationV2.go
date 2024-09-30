package main

import (
	"thanapatjitmung/go-test-komgrip/config"
	"thanapatjitmung/go-test-komgrip/databases"
	"thanapatjitmung/go-test-komgrip/entities"

	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	mariaDb := databases.NewMariaDatabase(conf.MariaDB)
	tx := mariaDb.Begin()

	itemsAdding(tx)
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func itemsAdding(tx *gorm.DB) {
	items := []entities.Beer{
		{
			Name:     "SINGHA",
			Type:     "LAGER",
			Details:  "เบียร์สิงห์ ปริมาณแอลกอฮอล์ 5.0% เบียร์ลาเกอร์ระดับพรีเมี่ยมที่ ผลิตจาก วัตถุดิบชั้นเยี่ยม จึงเป็นเบียร์ที่ให้ รสชาติดี มีความเข้มข้นของบาร์เลย์ และฮอพแท้ 100% ",
			ImageURL: "https://i.pinimg.com/736x/73/cc/79/73cc79391b764ec40a5c77052bb846b9.jpg",
		},
		{
			Name:     "MAHANAKORN | HAZY IPA",
			Type:     "IPA",
			Details:  "เบียร์สไตล์ Hazy IPA ที่น้ำเบียร์มีความฉ่ำ ใช้ฮ็อปส์ Nelson Sauvin ร่วมกับ Azacca และ Idaho-7 ร่วมกับการดรายฮ็อป 2 ครั้งเพื่อให้กลิ่นโทน ลูกพีช เสาวรส ที่หนักแน่น ABV : 6.5% IBU : 35",
			ImageURL: "https://i.pinimg.com/736x/73/cc/79/73cc79391b764ec40a5c77052bb846b9.jpg",
		},
		{
			Name:     "MAHANAKORN | WHITE IPA",
			Type:     "IPA",
			Details:  "เบียร์ IPA ที่มีส่วนผสมของข้าวสาลี ให้กลิ่นโทนเขียว ยางสน แซมผลไม้เขตร้อน/ซิตรัส บอดี้บาง ABV : 6% IBU : 35",
			ImageURL: "https://i.pinimg.com/736x/73/cc/79/73cc79391b764ec40a5c77052bb846b9.jpg",
		},
		{
			Name:     "MAHANAKORN | SESSION IPA",
			Type:     "IPA",
			Details:  "เบียร์ IPA ที่ใช้เพียงฮ็อปส์สายพันธุ์ “แคชเมียร์” ที่ให้กลิ่นผลไม้เมืองร้อน, เมล่อน และซิตรัส ที่ชัดเจน ตามมาความขมที่นุ่มนวล และกลิ่นมะพร้าวอ่อนๆ ABV : 4.5% IBU : 40",
			ImageURL: "https://i.pinimg.com/736x/73/cc/79/73cc79391b764ec40a5c77052bb846b9.jpg",
		},
	}

	tx.CreateInBatches(items, len(items))
}
