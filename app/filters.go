package app

import (
  "regexp"
)




//PHONE
func isPhone(item string) bool{
  reg,_ := regexp.Compile(`^(?:\(?([0-9]{3})\)?[-. ]?)?([0-9]{3})[-. ]?([0-9]{4})$`)
  return reg.MatchString(item)
}



//EMAIL
func isEmail(item string) bool{
  reg,_ := regexp.Compile(`([a-z0-9][-a-z0-9_\+\.]*[a-z0-9])@([a-z0-9][-a-z0-9\.]*[a-z0-9]\.(biz|tv|com|edu|gov|info|jobs|mil|net|org|jp|in|ca|uk|es|br|kr|mx|de|fr)|([0-9]{1,3}\.{3}[0-9]{1,3}))`)
  return reg.MatchString(item)
}

//FORM IDs
func isFormTxt(item string) bool {
  reg,_ := regexp.Compile(`Txt$|txt$|\d$`)
  return reg.MatchString(item)
}


/** 
** REFERENCES 
** http://blog.stevenlevithan.com/archives/validate-phone-number
** http://www.webmonkey.com/2008/08/four_regular_expressions_to_check_email_addresses/
**/