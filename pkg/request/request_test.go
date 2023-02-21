package request;
//import "testing"
import (
    "testing"
    "fmt"
)
var body = Body{
    Model: "text-davinci-003",
    Temperature: 0,
    Top_P: 1,
    Presence_Pen: 0,
    Frequence_Pen: 0,
    Best_Of: 1,
}
func TestCheckFormat(t *testing.T) {
    t.Run("Invalid", TestInvalidFormat)
}
func TestInvalidFormat(t *testing.T) {
   t.Run("Model", TestInvalidModel) 
   t.Run("Temperature", TestInvalidTemperature)
   t.Run("topp", TestInvalidTopp)
   t.Run("Presence", TestInvalidPresence)
   t.Run("Frequence", TestInvalidFrequence)
   t.Run("BestOf", TestInvalidBestOf)
}
func Range[V float64 | int](minRange V, maxRange V) (min V, max V, invalid Body) {
    invalid = body
    var changeBy V
    changeBy ++ 
    min = minRange - changeBy
    max = maxRange + changeBy
    return
}
func GetErrorText[V float64 | int | string ](incorrect V, err error, expectedErr error) string {
    error := fmt.Sprintf("The incorrect value %v was checked in the CheckFormat function. It returned the error of %v when it was expected to return %v", incorrect, err, expectedErr)
    return error
}
func TestInvalidBestOf(t *testing.T) {
    min, max, invalid := Range(BestOfMin, BestOfMax)
    invalid.Best_Of = min
    err := CheckFormat(invalid)
    if err != ErrWrongBestOfRange{
       t.Errorf(GetErrorText(min, err, ErrWrongBestOfRange)) 
    }
    invalid.Best_Of = max
    err = CheckFormat(invalid)
    if err != ErrWrongBestOfRange  {
       t.Errorf(GetErrorText(max, err, ErrWrongBestOfRange)) 
    }
}
func TestInvalidFrequence(t *testing.T) {
    min, max, invalid := Range(FrequenceMin, PresenceMax)
    invalid.Frequence_Pen = min
    err := CheckFormat(invalid)
    if err != ErrWrongFrequenceRange{
       t.Errorf(GetErrorText(min, err, ErrWrongFrequenceRange)) 
    }
    invalid.Frequence_Pen = max
    err = CheckFormat(invalid)
    if err != ErrWrongFrequenceRange {
       t.Errorf(GetErrorText(max, err, ErrWrongFrequenceRange)) 
    }
}
func TestInvalidPresence(t *testing.T) {
    min, max, invalid := Range(PresenceMin, PresenceMax)
    invalid.Presence_Pen = min
    err := CheckFormat(invalid)
    if err != ErrWrongPresenceRange {
       t.Errorf(GetErrorText(min, err, ErrWrongPresenceRange)) 
    }
    invalid.Presence_Pen = max
    err = CheckFormat(invalid)
    if err != ErrWrongPresenceRange {
       t.Errorf(GetErrorText(max, err, ErrWrongPresenceRange)) 
    }
}
func TestInvalidTopp(t *testing.T) {
    min, max, invalid := Range(TopPMin, TopPMax)
    invalid.Top_P = min
    err := CheckFormat(invalid)
    if err != ErrWrongToppRange {
       t.Errorf(GetErrorText(min, err, ErrWrongToppRange)) 
    }
    invalid.Top_P = max
    err = CheckFormat(invalid)
    if err != ErrWrongToppRange {
       t.Errorf(GetErrorText(max, err, ErrWrongToppRange)) 
    }
}
func TestInvalidTemperature(t *testing.T) {
    min, max, invalid := Range(TempRangeMin, TempRangeMax)
    invalid.Temperature = min 
    err := CheckFormat(invalid)
    if err != ErrWrongTempRange {
        t.Errorf(GetErrorText(min, err, ErrWrongTempRange))
    }
    invalid.Temperature = max
    err = CheckFormat(invalid)
    if err != ErrWrongTempRange {
        t.Errorf(GetErrorText(max, err, ErrWrongTempRange))
    }
}

func TestInvalidModel(t *testing.T) {
    invalidModel := "text-wrong-model"
    invalid := body
    invalid.Model = invalidModel
    err := CheckFormat(invalid)
    if err != ErrWrongModel {
        t.Errorf(GetErrorText(invalidModel, err, ErrWrongModel))
    }
}
