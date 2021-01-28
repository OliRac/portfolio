/*Sample queries*/
SELECT degree, institution, location, EXTRACT (YEAR FROM grad_year)
	FROM public.t_education 
	INNER JOIN public.t_language 
	ON t_education.lang_id = t_language.id
	WHERE t_language.code='en';