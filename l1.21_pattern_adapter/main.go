// ----------------------------
// Adapter Pattern:
// WHEN TO USE: integrate incompatible classes/services via a uniform interface
// PROS: enables reuse, reduces coupling, unifies interface
// CONS: extra abstraction, slight overhead, can complicate design if overused
package main

import "fmt"

// Courier is a common interface for all delivery services.
type Courier interface {
	SendPackage(address string) string
}

// -------- WB -------
type WB struct{}

func (wb *WB) SendByWB(address string) string {
	return "order sent by WB delivery to " + address
}

// WBAdapter adapts WB to the Courier interface
type WBAdapter struct {
	wb *WB
}

func (wa *WBAdapter) SendPackage(address string) string {
	return wa.wb.SendByWB(address)
}

// -------- Ozon -------
type Ozon struct{}

func (o *Ozon) CreateDelivery(address string) string {
	return "order sent by Ozon delivery to " + address
}

// OzonAdapter adapts Ozon to the Courier interface
type OzonAdapter struct {
	o *Ozon
}

func (oa *OzonAdapter) SendPackage(address string) string {
	return oa.o.CreateDelivery(address)
}

// -------- Yandex -------
type Yandex struct{}

func (y *Yandex) MakeShipment(address string) string {
	return "order sent by Yandex delivery to " + address
}

// YandexAdapter adapts Yandex to the Courier interface
type YandexAdapter struct {
	y *Yandex
}

func (ya *YandexAdapter) SendPackage(address string) string {
	return ya.y.MakeShipment(address)
}

// -------- Business Layer --------
// ShippingSystem uses any Courier implementation
type ShippingSystem struct {
	courier Courier
}

func (s *ShippingSystem) SetCourier(courier Courier) {
	s.courier = courier
}

func (s *ShippingSystem) SendOrder(address string) {
	result := s.courier.SendPackage(address)
	fmt.Println(result)
}

func main() {
	sysShi := &ShippingSystem{}

	wb := &WB{}
	oz := &Ozon{}
	ya := &Yandex{}

	aWb := &WBAdapter{wb: wb}
	aOz := &OzonAdapter{o: oz}
	aYa := &YandexAdapter{y: ya}

	sysShi.SetCourier(aWb)
	sysShi.SendOrder("Moscow")

	sysShi.SetCourier(aOz)
	sysShi.SendOrder("Saint-Petersburg")

	sysShi.SetCourier(aYa)
	sysShi.SendOrder("Kazan")
}
