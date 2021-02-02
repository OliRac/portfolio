package main

/*Makes use of a local postgreSQL database to fill out resume data with the given language
Errors are handled in Database methods*/
func BuildResume(lang string) *ResumeData{
	var db Database

	db.Connect()
	defer db.Disconnect()

	headers := db.GetSectionHeaders(lang)

	edu := EduSection{
		Header: headers.Education,
		Entries: db.GetEducation(lang),
	} 

	know := KnowSection{
		Header: headers.Knowledge,
		Entries: db.GetKnowledge(lang),
	}

	exp := ExpSection{
		Header: headers.Experience,
		Entries: db.GetExperience(lang),
	}

	skills := SkillSection{
		Header: headers.Skills,
		Entries: db.GetSkill(),
	}

	about := db.GetAbout(lang)
	about.Header = headers.About	//Header is not set in GetAbout...

	hero := db.GetHero(lang)
	
	langBtn := "fr"

	//if user language is french, change lang button to english
	if lang == "fr"{
		langBtn = "en"
	}

	//Linking everything together
	data := ResumeData{
		Edus: &edu,
		Exps: &exp,
		Knows: &know,
		Skills: &skills,
		Person: hero,
		Info: about,
		NavbarItems: headers,
		LangButton: langBtn,
	}

	return &data
}