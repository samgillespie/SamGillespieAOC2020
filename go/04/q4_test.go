package main

import "testing"

func TestPassport(t *testing.T) {

	t.Run("check_birth_year", func(t *testing.T) {
		pass := new(passport)
		// Invalid Inputs
		if pass.ValidBirthYear() == true {
			t.Fail()
		}
		pass.SetValue("byr", "1919")
		if pass.ValidBirthYear() == true {
			t.Fail()
		}
		pass.SetValue("byr", "2003")
		if pass.ValidBirthYear() == true {
			t.Fail()
		}

		pass.SetValue("byr", "1920")
		if pass.ValidBirthYear() == false {
			t.Fail()
		}
	})

	t.Run("check_issue_year", func(t *testing.T) {
		pass := new(passport)
		// Invalid Inputs
		if pass.ValidIssueYear() == true {
			t.Fail()
		}

		pass.SetValue("iyr", "1919")
		if pass.ValidIssueYear() == true {
			t.Fail()
		}
		pass.SetValue("iyr", "2021")
		if pass.ValidIssueYear() == true {
			t.Fail()
		}

		pass.SetValue("iyr", "2020")
		if pass.ValidIssueYear() == false {
			t.Fail()
		}

		pass.SetValue("iyr", "2010")
		if pass.ValidIssueYear() == false {
			t.Fail()
		}
	})

	t.Run("check_expiry_year", func(t *testing.T) {
		pass := new(passport)
		// Invalid Inputs
		if pass.ValidExpirationYear() == true {
			t.Fail()
		}

		pass.SetValue("eyr", "2019")
		if pass.ValidExpirationYear() == true {
			t.Fail()
		}
		pass.SetValue("eyr", "2031")
		if pass.ValidExpirationYear() == true {
			t.Fail()
		}

		pass.SetValue("eyr", "2020")
		if pass.ValidExpirationYear() == false {
			t.Fail()
		}

		pass.SetValue("eyr", "2030")
		if pass.ValidExpirationYear() == false {
			t.Fail()
		}
	})

	t.Run("check_ValidHeight", func(t *testing.T) {
		pass := new(passport)
		// Invalid Inputs
		if pass.ValidHeight() == true {
			t.Fail()
		}

		pass.SetValue("hgt", "149cm")
		if pass.ValidHeight() == true {
			t.Fail()
		}
		pass.SetValue("hgt", "194cm")
		if pass.ValidHeight() == true {
			t.Fail()
		}

		pass.SetValue("hgt", "58in")
		if pass.ValidHeight() == true {
			t.Fail()
		}
		pass.SetValue("hgt", "77in")
		if pass.ValidHeight() == true {
			t.Fail()
		}
		pass.SetValue("hgt", "76ow")
		if pass.ValidHeight() == true {
			t.Fail()
		}

		pass.SetValue("hgt", "59in")
		if pass.ValidHeight() == false {
			t.Fail()
		}
		pass.SetValue("hgt", "76in")
		if pass.ValidHeight() == false {
			t.Fail()
		}
		pass.SetValue("hgt", "150cm")
		if pass.ValidHeight() == false {
			t.Fail()
		}
		pass.SetValue("hgt", "193cm")
		if pass.ValidHeight() == false {
			t.Fail()
		}
	})

	t.Run("check_ValidHairColor", func(t *testing.T) {
		pass := new(passport)
		// Invalid Inputs
		if pass.ValidHairColor() == true {
			t.Fail()
		}

		pass.SetValue("hcl", "blue")
		if pass.ValidHairColor() == true {
			t.Fail()
		}
		pass.SetValue("hcl", "#1234zz")
		if pass.ValidHairColor() == true {
			t.Fail()
		}

		pass.SetValue("hcl", "#abcdef123")
		if pass.ValidHairColor() == true {
			t.Fail()
		}

		pass.SetValue("hcl", "#abcdef")
		if pass.ValidHairColor() == false {
			t.Fail()
		}

		pass.SetValue("hcl", "#09af56")
		if pass.ValidHairColor() == false {
			t.Fail()
		}
	})

	t.Run("check_ValidEyeColor", func(t *testing.T) {
		pass := new(passport)
		// Invalid Inputs
		if pass.ValidEyeColor() == true {
			t.Fail()
		}

		pass.SetValue("ecl", "blue")
		if pass.ValidEyeColor() == true {
			t.Fail()
		}

		pass.SetValue("ecl", "amber")
		if pass.ValidEyeColor() == true {
			t.Fail()
		}

		pass.SetValue("ecl", "amb")
		if pass.ValidEyeColor() == false {
			t.Fail()
		}

	})

	t.Run("check_ValidPassport", func(t *testing.T) {
		pass := new(passport)
		// Invalid Inputs
		if pass.ValidPassport() == true {
			t.Fail()
		}

		pass.SetValue("pid", "0123456789")
		if pass.ValidPassport() == true {
			t.Fail()
		}

		pass.SetValue("pid", "012345678")
		if pass.ValidPassport() == false {
			t.Fail()
		}
	})
}
