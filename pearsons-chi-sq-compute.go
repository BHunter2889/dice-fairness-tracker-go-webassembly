package main

import "math"

/**
 For clarity, the base computational `struct`s (per side aka 'option' and the total) will use the explicit 'PearsonsChiSq`
 naming convention.
 For simplicity, all `func`s will use the abbreviated `PCS` naming convention.
 */

const d20ChiTableLookup = 30.143

type DieConstantsBySides struct {
	NumberOfSides    int
	ChiSqTableValue  float64 // e.g. for a 5% significance error margin [X^2(degreesOfFreedom, 0.05)] OR X^2(#sides-1, 0.05)
	MinNumberOfRolls int     // chi-square requires a minimum to be applicable: `TotalNumberOfRolls >== 5 * NumberOfSides`
}

/**
	Die Constant struct instances for conventional unconventionally sided die (and a d6 of course).
 */
var d4, d6, d8, d10, d12, d20 = initDieConstants()

func initDieConstants() (d4, d6, d8, d10, d12, d20 DieConstantsBySides) {
	d4 = DieConstantsBySides{
		NumberOfSides:    4,
		ChiSqTableValue:  7.815, // TODO
		MinNumberOfRolls: 20,    // 5 rolls * 4 sides
	}

	d6 = DieConstantsBySides{
		NumberOfSides:    6,
		ChiSqTableValue:  11.070, // TODO
		MinNumberOfRolls: 30,     // 5 rolls * 6 sides
	}

	d8 = DieConstantsBySides{
		NumberOfSides:    8,
		ChiSqTableValue:  14.067, // TODO
		MinNumberOfRolls: 40,     // 5 rolls * 8 sides
	}

	d10 = DieConstantsBySides{
		NumberOfSides:    10,
		ChiSqTableValue:  16.919, // TODO
		MinNumberOfRolls: 50,     // 5 rolls * 10 sides
	}

	d12 = DieConstantsBySides{
		NumberOfSides:    12,
		ChiSqTableValue:  19.675, // TODO
		MinNumberOfRolls: 60,     // 5 rolls * 12 sides
	}

	d20 = DieConstantsBySides{
		NumberOfSides:    20,
		ChiSqTableValue:  30.143, // TODO
		MinNumberOfRolls: 100,    // 5 rolls * 20 sides
	}
	return
}

func GetDieConstantsBySides(numberOfSides int) (d DieConstantsBySides) {
	switch numberOfSides {
	case 4:
		d = d4
	case 6:
		d = d6
	case 8:
		d = d8
	case 10:
		d = d10
	case 12:
		d = d12
	case 20:
		d = d20
	default:
		d = d20
	}
	return
}

/**
	Pearson's Chi-Square Test requires the following values for each die side option (`PearsonsChiSqOption`)
	and for each die as a whole (`ComputedPearsonsChiSqValues`).
 */
type PearsonsChiSqOption struct {
	SideRollCount     int     `default:"0"`    // Should be passed in each time Compute is called, but stored for display.
	ExpectedRollCount float64 `default:"0.0"`  // Should be passed in each time Compute is called, but stored for display.
	Error             float64  `default:"0.0"` // `SideRollCount - ExpectedRollCount`
	SquaredError      float64  `default:"0.0"` // `math.Pow(Error, 2)` aka Error^2
}

type ComputedPearsonsChiSqValues struct {
	DieConstants         DieConstantsBySides
	OptionComputations   []PearsonsChiSqOption
	BalanceThreshold     float64 `default:"0.0"`   // Threshold for the SumSquaredError.
	ExpectedRollsPerSide float64 `default:"0.0"`   // `TotalRollCount/NumberOfSides`
	SumSquaredError      float64 `default:"0.0"`   // Total of all options SquaredError different from expected roll count. aka 'SSE'
	IsBalanced           bool    `default:"false"` // SSE <= BalanceThreshold
}

/**
	Methods for computing the associated Chi-Square values for each Side Option
 */

func (option *PearsonsChiSqOption) ComputeErrorAndSquaredError(
	currentSideRollCount int, expectedRollCount float64) (pcsqError float64, squaredError float64) {
	option.SideRollCount, option.ExpectedRollCount = currentSideRollCount, expectedRollCount

	option.Error = float64(option.SideRollCount) - option.ExpectedRollCount
	option.SquaredError = math.Pow(option.Error, 2)

	pcsqError, squaredError = option.Error, option.SquaredError
	return
}

/**
	Methods for Computing the Pearson's Chi-Square analysis
 */

func NewComputedPCSValues(numberOfSides int) (newCPCSValues *ComputedPearsonsChiSqValues) {
	newCPCSValues = &ComputedPearsonsChiSqValues{}
	newCPCSValues.DieConstants = GetDieConstantsBySides(numberOfSides)
	newCPCSValues.OptionComputations = make([]PearsonsChiSqOption, numberOfSides)
	for i := 0; i < numberOfSides; i++ {
		newCPCSValues.OptionComputations[i] = PearsonsChiSqOption{}
	}
	return
}

// Groups and calls all necessary compute functions.
func (cpcsv *ComputedPearsonsChiSqValues) ComputePChSqValues(
	currentRollCountTotal int, counts []int) (isBalanced bool, sse float64, balanceThreshold float64){
	cpcsv.ComputeExpectedRolls(currentRollCountTotal)
	cpcsv.ComputeBalanceThreshold(currentRollCountTotal)
	cpcsv.ComputeSumSquaredErrorIfMinRollCountMet(currentRollCountTotal, counts)
	isBalanced, sse, balanceThreshold = cpcsv.ComputeIsBalanced()
	return
}

func (cpcsv *ComputedPearsonsChiSqValues) ComputeExpectedRolls(currentRollCountTotal int) {
	cpcsv.ExpectedRollsPerSide = float64(currentRollCountTotal) / float64(cpcsv.DieConstants.NumberOfSides)
}

func (cpcsv *ComputedPearsonsChiSqValues) ComputeBalanceThreshold(currentRollCountTotal int) {
	cpcsv.BalanceThreshold = cpcsv.ExpectedRollsPerSide * cpcsv.DieConstants.ChiSqTableValue
}

func (cpcsv *ComputedPearsonsChiSqValues) ComputeSumSquaredErrorIfMinRollCountMet(currentRollCountTotal int, counts []int) (isMinRollCountMet bool) {
	if cpcsv.DieConstants.MinNumberOfRolls <= currentRollCountTotal {
		newSSE := 0.0
		for i, option := range cpcsv.OptionComputations {
			option.ComputeErrorAndSquaredError(counts[i], cpcsv.ExpectedRollsPerSide)
			newSSE += option.SquaredError
		}
		cpcsv.SumSquaredError = newSSE
		isMinRollCountMet = true
		return
	}
	isMinRollCountMet = false
	return
}

func (cpcsv *ComputedPearsonsChiSqValues) ComputeIsBalanced() (isBalanced bool, sse float64, balanceThreshold float64) {
	cpcsv.IsBalanced = cpcsv.SumSquaredError <= cpcsv.BalanceThreshold
	return cpcsv.IsBalanced, cpcsv.SumSquaredError, cpcsv.BalanceThreshold
}
