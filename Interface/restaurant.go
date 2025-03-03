<<<<<<< HEAD
/*
Restoran buyurtmalarini boshqarish
Vazifa:
Restoran uchun buyurtma tizimini boshqarish uchun OrderManager interfeysini yarating, unda quyidagi metodlar bo‘lsin:

PlaceOrder(order string) string
CancelOrder(orderID int) string
OrderStatus(orderID int) string
Keyin ushbu interfeysni implement qiluvchi xizmat turlari:

DineIn: joyida ovqatlanish buyurtmalarini boshqaradi.
TakeAway: olib ketiladigan buyurtmalarni boshqaradi.
Delivery: yetkazib berish buyurtmalarini boshqaradi.
Har bir xizmat o‘ziga xos tarzda buyurtma statuslarini boshqarsin.
*/
package main 


type OrderManager interface {
	PlaceOrder(order string) string
	CancelOrder (orderID int) string
	OrderStatus(orderID int) string 
}

func main(){

}
=======
package main
>>>>>>> eec1064ec19710ca2f9fc2f6097371297182c3ff
