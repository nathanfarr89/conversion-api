# conversion-api

### Endpoints

https://intelligent-mandarine-90254.herokuapp.com/temp/{number}

https://intelligent-mandarine-90254.herokuapp.com/weight/{number}

https://intelligent-mandarine-90254.herokuapp.com/measure/{number}

https://intelligent-mandarine-90254.herokuapp.com/feetmeters/{number}


Replace {number} with the number you want converted. 

/temp converts fahrenheit, celsius, and kelvin

/weight converts pounds and kilograms

/measure converts inches and centimeters

/feetmeters converts feet and meters

#### Example

https://intelligent-mandarine-90254.herokuapp.com/temp/100 gives the following:

{

  "number": 100,
  
  "ftcelsius": 37.77777777777778,
  
  "ctfahrenheit": 212,
  
  "ftkelvin": 310.9277777777778,
  
  "ctkelvin": 373.15,
  
  "ktcelsius": -173.14999999999998,
  
  "ktfahrenheit": -279.67
  
}
