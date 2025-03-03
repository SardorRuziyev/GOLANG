<<<<<<< HEAD
/*
Hayvonlar va Ularning Ovozlari
 1. Animal nomli interface yarating. Unda quyidagi methodlar bo‘lsin:
 • MakeSound() string – Hayvonning chiqargan ovozini qaytarsin
 • Move() string – Hayvonning qanday harakatlanishini qaytarsin
 2. Uch xil hayvon yaratib, har biri Animal interfeysini implement qilsin:
 • Dog – Ovoz: "Woof!", Harakat: "Runs"
 • Cat – Ovoz: "Meow!", Harakat: "Jumps"
 • Bird – Ovoz: "Chirp!", Harakat: "Flies"
 3. Hayvonlar ro‘yxatini ([]Animal) yaratib, ularga turli hayvonlarni qo‘shing va ularning ovoz chiqarishi va harakatini ekranga chiqaring.

Qo‘shimcha qiyinlashtirish (optional)

✅ Hayvonlarga Name (string) maydoni qo‘shing va ularni nomi bilan chaqiring.
✅ Zoo nomli struct yarating, unda Animals []Animal bo‘lsin. ShowAnimals() methodini yozing, u barcha hayvonlarning ovozi va harakatini chiqarib bersin.

Natija qanday ko‘rinishi kerak?

Agar foydalanuvchi 3 ta hayvonni qo‘shsa va ularning ovozini tekshirsa, ekranga quyidagicha chiqadi:

Dog: Woof! - Runs
Cat: Meow! - Jumps
Bird: Chirp! - Flies
*/

package main

import "github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"

type Animal interface{
	MakeSound ()
	Move ()
}
type Dog struct{
	name string
}
func (d Dog) MakeSound() string {
	return "woof!"
}
func(d Dog) Move() string {
	return "runs"
}
func (d Dog) Name() string {
	return d.name
}
type Cat struct {
	name string
}
func (c Cat) MakeSound() string {
	return "meaow"

}
func (c Cat) Move() string {
	return "jumps"
}
func (c Cat) Name() string {
	return c.name
}
type Bird struct{
	name string
}
func (b Bird) MakeSound() string {
	return "Chirp!"
}

func (b Bird) Move() string {
	return "Flies"
}

func (b Bird) Name() string {
	return b.name
}
type Zoo struct {
	Animals []Animal
}
func main(){
	dog := Dog{name: "dog"}
	cat := Cat{name: "cat"}
	bird := Bird{name: "bird"}

}
=======
package main
>>>>>>> eec1064ec19710ca2f9fc2f6097371297182c3ff
