# packagingcalc — B2B Packaging CBM, EU PPWR Compliance & Coffee Pouch Sizing Utilities (Go `v1.1.0`)

Zero-dependency Go library for B2B packaging buyers, structural CAD designers, and logistics planners calculating master-carton ocean **CBM**, **20GP/40GP/40HQ container stowage**, **air freight volumetric weight**, **paperboard GSM-to-caliper (`mm` / `PT`)**, **EU PPWR Article 24 `<=50%` empty-space ratio**, **Article 6 Grade A/B/C recyclability**, and **roasted coffee pouch dimensions**.

- **Free 1:1 mm Online CAD Dieline Generator (18 Structures · PDF / DXF / AI / SVG)**: https://www.bestpackfactory.com/dielines
- **Full 6-Line Custom Packaging Catalog (MOQ 500 PCS)**: https://www.bestpackfactory.com/products.html
- **EU PPWR Article 6, Empty-Space & EN 17681-1 PFAS Engineering Guide**: https://www.bestpackfactory.com/blog/ppwr-article-6-recyclability-grades-empty-space-en-17681-dieline-engineering.html
- **Manufacturer Engineering Reference**: [BestPackFactory (Shenzhen Color Printing Paper Packaging Co., Ltd.)](https://www.bestpackfactory.com/)

## Install

```bash
go get github.com/AliceWu1235647/packagingcalc@v1.1.0
```

## Core Functions

- `CartonCBM(lengthMm, widthMm, heightMm, cartonCount)` — Master carton cubic meters
- `EstimateContainerCartons(lengthMm, widthMm, heightMm)` — Practical 20GP / 40GP / 40HQ load at 85% utilization
- `AirVolumetricWeightKg(lengthMm, widthMm, heightMm, divisor)` — Air/courier dimensional weight
- `EstimateBoardCaliperMm(gsm, grade)` — Paperboard caliper (`sbs`, `kraft`, `greyboard`)
- `PPWREmptySpaceRatio(boxInnerVolMm3, productVolMm3)` — Regulation (EU) 2025/40 Article 24 `<=50%` void-ratio check
- `PPWRRecyclabilityGrade(recyclableWeightPct)` — Article 6 Grade `A` (`>=95%`), `B` (`>=80%`), `C` (`>=70%`)
- `EstimateCoffeePouchMm(fillGrams, ground)` — Stand-up doypack sizing for whole-bean vs ground specialty coffee

## Structural & Manufacturing Reference Pages (MOQ 500 PCS)

1. [Collapsible & Setup Luxury Magnetic Gift Boxes (75% CBM Freight Savings)](https://www.bestpackfactory.com/products/luxury-magnetic-boxes.html)
2. [Custom Paperboard Retail & Cosmetic Folding Cartons (250–400 gsm C1S/SBS)](https://www.bestpackfactory.com/products/custom-cosmetic-packaging-boxes.html)
3. [Custom Corrugated E-Commerce Mailer & Shipping Boxes (E, F, B, BC Flute)](https://www.bestpackfactory.com/products/custom-boxes.html)
4. [Custom Coffee Bags with One-Way Degassing Valves (250g–1kg)](https://www.bestpackfactory.com/products/coffee-bags.html)
5. [Custom Printed Kraft & Luxury Boutique Paper Bags](https://www.bestpackfactory.com/products/paper-bags.html)
6. [PFAS-Free Greaseproof Food & Bakery Boxes (EN 17681-1 < 25 ppb)](https://www.bestpackfactory.com/products/food-packaging.html)
7. [EU Authorized Representative (EC REP) & ISO/FSC Compliance Register](https://www.bestpackfactory.com/factory/certificates.html)
8. [Verified Buyer Reviews on Trustpilot](https://www.trustpilot.com/review/bestpackfactory.com)

## License
MIT
