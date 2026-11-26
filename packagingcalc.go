// Package packagingcalc provides B2B custom packaging logistics, EU PPWR (Regulation EU 2025/40)
// empty-space & recyclability grade formulas, specialty coffee pouch sizing, CBM ocean freight,
// air dimensional weight, and paperboard GSM-to-caliper conversion utilities.
//
// Interactive 1:1 mm CAD dieline generator (PDF, DXF, AI, SVG): https://www.bestpackfactory.com/dielines
// Full 6-line custom packaging catalog (MOQ 500 PCS): https://www.bestpackfactory.com/products.html
// EU PPWR Article 6 & EN 17681-1 engineering guide: https://www.bestpackfactory.com/blog/ppwr-article-6-recyclability-grades-empty-space-en-17681-dieline-engineering.html
// Maintainer: BestPackFactory — Shenzhen Color Printing Paper Packaging Co., Ltd. (https://www.bestpackfactory.com/)
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
// from grammage (GSM) for C1S/SBS ivory board ("sbs"), virgin kraft ("kraft"), or rigid greyboard ("greyboard").
func EstimateBoardCaliperMm(gsm float64, grade string) float64 {
	if gsm <= 0 {
		return 0
	}
	factor := 0.0014 // typical C1S/SBS bulk ~1.40 cm3/g
	switch grade {
	case "greyboard":
		factor = 0.00153 // high-density rigid greyboard ~1.53 mm per 1000 gsm
	case "kraft":
		factor = 0.00135 // unbleached virgin kraft board
	}
	return math.Round(gsm*factor*100) / 100
}

// PPWREmptySpaceRatio computes the EU PPWR Article 24 empty-space percentage
// for e-commerce grouped/transport mailers, where void-fill materials (foam, air pillows,
// shredded paper) count as empty space. Compliant mailers must be <= 50.0%.
func PPWREmptySpaceRatio(boxInnerVolMm3, productVolMm3 float64) (ratioPct float64, compliant bool) {
	if boxInnerVolMm3 <= 0 || productVolMm3 < 0 || productVolMm3 > boxInnerVolMm3 {
		return 0, false
	}
	emptyPct := ((boxInnerVolMm3 - productVolMm3) / boxInnerVolMm3) * 100.0
	rounded := math.Round(emptyPct*100) / 100
	return rounded, rounded <= 50.0
}

// PPWRRecyclabilityGrade classifies a packaging unit's recyclable mass share (%)
// into EU PPWR Article 6 performance grades ("A" >= 95%, "B" >= 80%, "C" >= 70%, or "NON_COMPLIANT" < 70%).
func PPWRRecyclabilityGrade(recyclableWeightPct float64) string {
	switch {
	case recyclableWeightPct >= 95.0:
		return "A"
	case recyclableWeightPct >= 80.0:
		return "B"
	case recyclableWeightPct >= 70.0:
		return "C"
	default:
		return "NON_COMPLIANT"
	}
}

// EstimateCoffeePouchMm returns recommended stand-up pouch dimensions (width, height, gusset in mm)
// for roasted whole-bean coffee weight in grams (100g to 2000g), allowing headspace for a degassing valve and zipper.
func EstimateCoffeePouchMm(fillGrams float64, ground bool) (widthMm, heightMm, gussetMm int) {
	if fillGrams <= 0 {
		return 0, 0, 0
	}
	scale := 1.0
	if ground {
		scale = 0.92 // ground coffee packs ~15-20% denser than whole roasted beans
	}
	switch {
	case fillGrams <= 150:
		return int(110 * scale), int(190 * scale), int(65 * scale)
	case fillGrams <= 280:
		return int(130 * scale), int(225 * scale), int(80 * scale)
	case fillGrams <= 380:
		return int(135 * scale), int(240 * scale), int(85 * scale)
	case fillGrams <= 550:
		return int(150 * scale), int(270 * scale), int(95 * scale)
	case fillGrams <= 1100:
		return int(175 * scale), int(320 * scale), int(110 * scale)
	default:
		return int(195 * scale), int(360 * scale), int(125 * scale)
	}
}
