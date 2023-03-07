package request

import (
	"errors"
	"fmt"
)

// enum types
type Model = string

// enums for the models
const (
	Turbo   Model = "gpt-3.5-turbo"
	Davinci Model = "text-davinci-003"
	Curie   Model = "text-curie-001"
	Babbage Model = "text-babbage-001"
	Ada     Model = "text-ada-001"
)

// variables for the minimum and maximum ranges
var (
	TempRangeMin float64 = 0
	TempRangeMax float64 = 1
	TopPMin      float64 = 0
	TopPMax      float64 = 1
	PresenceMin  float64 = -2
	PresenceMax  float64 = 2
	FrequenceMin float64 = -2
	FrequenceMax float64 = 2
	BestOfMin    int     = 1
	BestOfMax    int     = 20
)

// the errors to use for wrong formats
var (
	ErrWrongModel          error = errors.New(fmt.Sprintf("The model you entered was not correct"))
	ErrWrongTempRange      error = errors.New(fmt.Sprintf("The temperature is not in the correct range (%v <= temp <= %v)", TempRangeMin, TempRangeMax))
	ErrWrongToppRange      error = errors.New(fmt.Sprintf("The top_p is not in the correct range (%v <= top_p <= %v)", TopPMin, TopPMax))
	ErrWrongPresenceRange  error = errors.New(fmt.Sprintf("The presence penalty is not in the correct range (%v <= presence <= %v)", PresenceMin, PresenceMax))
	ErrWrongFrequenceRange error = errors.New(fmt.Sprintf("The presence penalty is not in the correct range (%v <= frequence <= %v)", FrequenceMin, FrequenceMax))
	ErrWrongBestOfRange    error = errors.New(fmt.Sprintf("The best of variable is not in the correct range (%v <= best of  <= %v)", BestOfMin, BestOfMax))
	ErrWrongKey            error = errors.New("The API Kry you entered is either wrong or hasn't been set up with a paid account of GPT. You must sign up for a paid account at Openai GPT.")
)
