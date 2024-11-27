package utils

import (
	"math"
	"sort"
	"strconv"

	"github.com/expr-lang/expr"
)

type ExprEnv struct {
	Row           int     `expr:"ROW"`
	Col           int     `expr:"COL"`
	Depth         int     `expr:"DEPTH"`
	DepthMin      int     `expr:"DEPTHMIN"`
	DepthInterval int     `expr:"DEPTHINTERVAL"`
	BasePrice     float64 `expr:"BASEPRICE"`
}

type PriceComponent struct {
	Formula       string  `json:"formula" binding:"required"`
	BasePrice     float64 `json:"base_price" binding:"required"`
	Row           int     `json:"row,omitempty"`
	Col           int     `json:"col,omitempty"`
	Depth         int     `json:"depth,omitempty"`
	DepthMin      int     `json:"depth_min,omitempty"`
	DepthInterval int     `json:"depth_interval,omitempty"`
}

// CEILMULTI function to round to 'n' decimal places
func (ExprEnv) CEILMULTI(value, factor float64) float64 {
	if factor == 0 {
		return value // Avoid division by zero; return the original value
	}
	// Calculate the nearest multiple of factor
	multiple := math.Ceil(value/factor) * factor
	return multiple
}

// Calculates price
func Pricing(priceComponent PriceComponent) (*float64, error) {
	var price float64
	// Fill variables with values
	var env ExprEnv
	env.BasePrice = priceComponent.BasePrice
	env.Row = priceComponent.Row
	env.Col = priceComponent.Col
	env.Depth = priceComponent.Depth
	env.DepthMin = priceComponent.DepthMin
	env.DepthInterval = priceComponent.DepthInterval
	// Evaluate the expression
	res, err := expr.Eval(priceComponent.Formula, env)
	if err != nil {
		return nil, err
	}
	// Convert price to int or float64
	switch v := res.(type) {
	case int:
		priceInt := int(v)
		price = float64(priceInt)
	case float64:
		price = float64(v)
	default:
		return nil, err
	}
	return &price, nil
}

// FindIndexStep finds the largest number in the slice that is less than or equal to the input
func FindIndexStep(srcData []string, input int) int {
	var numbers []int
	for _, item := range srcData {
		if num, err := strconv.Atoi(item); err == nil {
			numbers = append(numbers, num)
		}
	}
	// Sort the numbers in ascending order
	sort.Ints(numbers)
	// Iterate through the sorted numbers
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] <= input {
			return numbers[i] // Return the first number that is less than or equal to input
		}
	}
	// Return -1 if no appropriate number is found
	return -1
}
