package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"edugo/services"
	"edugo/storage"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nWelcome to EduGo!")
		fmt.Println("1. Kurslarni ko‘rish")
		fmt.Println("2. Ro‘yxatdan o‘tish")
		fmt.Println("3. Kirish")
		fmt.Println("4. Kurs sotib olish")
		fmt.Println("5. Sharh qoldirish")
		fmt.Println("0. Chiqish")

		fmt.Print("> Tanlang: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		data, _ := storage.LoadData()

		switch choice {
		case "1":
			fmt.Println("📚 Mavjud kurslar:")
			services.GetAllCourses(data)

		case "2":
			fmt.Print("✍ Ismingiz: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("🔑 Parol: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			user, err := services.RegisterUser(name, password, data)
			if err != nil {
				fmt.Println("❌ Xatolik:", err)
			} else {
				storage.SaveData(data)
				fmt.Printf("✅ Ro‘yxatdan o‘tish muvaffaqiyatli! ID: %d\n", user.ID)
			}

		case "3":
			fmt.Print("👤 Ism: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("🔑 Parol: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			user, err := services.AuthenticateUser(name, password, data)
			if err != nil {
				fmt.Println("❌ Login xato:", err)
			} else {
				fmt.Printf("✅ Xush kelibsiz, %s!\n", user.Name)
			}

		case "4":
			fmt.Print("🆔 Foydalanuvchi ID: ")
			uidStr, _ := reader.ReadString('\n')
			uid, _ := strconv.Atoi(strings.TrimSpace(uidStr))

			fmt.Print("🎯 Kurs ID: ")
			cidStr, _ := reader.ReadString('\n')
			cid, _ := strconv.Atoi(strings.TrimSpace(cidStr))

			err := services.EnrollCourse(data, uid, cid)
			if err != nil {
				fmt.Println("❌ Kurs topilmadi:", err)
			} else {
				for i := range data.Users {
					if data.Users[i].ID == uid {
						services.BuyCourse(&data.Users[i], cid)
						break
					}
				}
				storage.SaveData(data)
				fmt.Println("💳 Kurs sotib olindi va ro‘yxatdan o‘tildi!")
			}

		case "5":
			fmt.Print("🆔 Foydalanuvchi ID: ")
			uidStr, _ := reader.ReadString('\n')
			uid, _ := strconv.Atoi(strings.TrimSpace(uidStr))

			fmt.Print("🎯 Kurs ID: ")
			cidStr, _ := reader.ReadString('\n')
			cid, _ := strconv.Atoi(strings.TrimSpace(cidStr))

			fmt.Print("✍ Sharh: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			fmt.Print("⭐ Reyting (1-5): ")
			ratingStr, _ := reader.ReadString('\n')
			rating, _ := strconv.Atoi(strings.TrimSpace(ratingStr))

			services.AddReview(data, cid, uid, rating, text)
			storage.SaveData(data)
			fmt.Println("✅ Sharhingiz saqlandi!")

		case "0":
			fmt.Println("👋 Dasturdan chiqilmoqda...")
			return

		default:
			fmt.Println("⚠️ Noto‘g‘ri tanlov!")
		}
	}
}
