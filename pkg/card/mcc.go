package card

func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	mcc := map[string]string{
		"5011": "Супермаркеты",
		"5014": "Рестораны",
		"5015": "Заправки",
	}

	result := "Категория не указана"

	for i := range mcc {
		if i == code {
			result = mcc[i]
		}
	}
	return result
}

