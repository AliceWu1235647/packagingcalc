// Package packagingcalc provides B2B custom packaging logistics, CBM ocean freight,
// air dimensional weight, and paperboard GSM-to-caliper conversion utilities.
//
// Reference 1:1 mm CAD dieline generator: https://www.bestpackfactory.com/dielines
// Maintainer: BestPackFactory (https://www.bestpackfactory.com/)
package packagingcalc

import "math"

// CartonCBM calculates total cubic meters (CBM) for a shipment of master cartons
// given outer carton dimensions in millimeters (L, W, H) and carton count.
func CartonCBM(lengthMm, widthMm, heightMm float64, cartonCount int) float64 {
	if lengthMm <= 0 || widthMm <= 0 || heightMm <= 0 || cartonCount <= 0 {
		return 0
	}
	unitCbm := (lengthMm / 1000.0) * (widthMm / 1000.0) * (heightMm / 1000.0)
	return math.Round(unitCbm*float64(cartonCount)*10000) / 10000
}

// AirVolumetricWeightKg calculates courier/air dimensional weight in kg
// using standard divisor 5000 (cm^3 / 5000) from millimeter dimensions.
func AirVolumetricWeightKg(lengthMm, widthMm, heightMm float64, divisor float64) float64 {
	if divisor <= 0 {
		divisor = 5000.0
	}
	cm3 := (lengthMm / 10.0) * (widthMm / 10.0) * (heightMm / 10.0)
	return math.Round((cm3/divisor)*100) / 100
}

// EstimateContainerCartons estimates practical master-carton capacity for 20GP (28 CBM),
// 40GP (58 CBM), and 40HQ (68 CBM) ocean containers at ~85% stowage utilization.
func EstimateContainerCartons(lengthMm, widthMm, heightMm float64) (c20gp, c40gp, c40hq int) {
	unit := CartonCBM(lengthMm, widthMm, heightMm, 1)
	if unit <= 0 {
		return 0, 0, 0
	}
	return int(28.0 * 0.85 / unit), int(58.0 * 0.85 / unit), int(68.0 * 0.85 / unit)
}

// EstimateBoardCaliperMm estimates nominal paperboard thickness in millimeters
// from grammage (GSM) for C1S/SBS ivory board ("sbs") or rigid greyboard ("greyboard").
func EstimateBoardCaliperMm(gsm float64, grade string) float64 {
	if gsm <= 0 {
		return 0
	}
	factor := 0.0014 // typical C1S/SBS bulk ~1.40 cm3/g
	if grade == "greyboard" {
		factor = 0.00153 // typical high-density rigid greyboard ~1.53 mm per 1000 gsm
	}
	return math.Round(gsm*factor*100) / 100
}
