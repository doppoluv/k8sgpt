package prompts

type Prompt struct {
	Template string
}

type LanguagePrompts map[string]Prompt

var languagePrompts = map[string]LanguagePrompts{
	"en": {
		"default_prompt": Prompt{Template: `Analyze the following Kubernetes error message.
				Provide a simple and accurate solution.
 				Do not suggest actions that could break the cluster. Be especially careful when suggesting the deletion or restart of any elements.
				If the issue is not critical, provide a solution in the following format: "error is insignificant".
				If you need to check the configuration files, try to specify the sections that need to be addressed.
 				Otherwise, provide the result in the following format, ensuring that the spaces are correct:
				Error: {Brief and clear explanation of the error}
				Solution: {Step-by-step solution of the error by points. Should not exceed 1000 characters. Use more specific kubectl, docker, curl, and other commands. Use only existing commands (for example, to pull an image, use docker pull instead of kubectl pull)}
				The points in the solution must be numbered in the correct order (for example, first point is checking the logs, and the last point is restarting the service)`},

		"prom_conf_prompt": Prompt{Template: `Simplify the following Prometheus error message.
				This error came when validating the Prometheus configuration file.
				Provide step by step instructions to fix, with suggestions, referencing Prometheus documentation if relevant.
				Write the output in the following format, observing the correctness of the spaces:
				Error: {Short and clear explanation of the error}
				Solution: {Step-by-step solution of the error point by point. It must not exceed 1000 characters. Try to give specific commands for each tip and action to identify and fix the problem}`},

		"prom_relabel_prompt": Prompt{Template: `Return your prompt, beginning with
				The following is a list of the form:
				job_name:
				{Prometheus job_name}
				relabel_configs:
				{Prometheus relabel_configs}
				kubernetes_sd_configs:
				{Prometheus service discovery config}

				For each job_name, describe the Kubernetes service and pod labels,
				namespaces, ports, and containers they match.
				Return the message:
				Discovered and parsed Prometheus scrape configurations.
				For targets to be scraped by Prometheus, ensure they are running with
				at least one of the following label sets:
				Then for each job, write this format:
				- Job: {job_name}
				  - Service Labels:
				    - {list of service labels}
				  - Pod Labels:
				    - {list of pod labels}
				  - Namespaces:
				    - {list of namespaces}
				  - Ports:
				    - {list of ports}
				  - Containers:
				    - {list of container names}
				Make sure to use the correct spacing.`},

		"kyverno_prompt": Prompt{Template: `Simplify the following Kyverno warnings message.
				Provide the most probable solution as a kubectl command.
				Write the output in the following format, for the solution, only show the kubectl command, observing the correctness of the spaces:
				Error: {Detailed explanation error here}
				Solution: {kubectl command}`},

		"raw_prompt": Prompt{Template: `{
				"model": "%s"
				"language": "%s",
				"prompt": "%s",
				"message": "%s"
				}`},
	},

	"ru": {
		"default_prompt": Prompt{Template: `Проанализируй следующее сообщение об ошибке Kubernetes.
				Сформулируй простое и точное решение.
				Не предлагай действий, которые могут сломать кластер. Будь особенно осторожен, предлагая удаление или перезапуск каких-либо элементов.
				Если проблема некритична, выведи решение в следующем формате: ошибка несущественна.
				При необходимости проверки конфигурационных файлов, старайся указывать секции, на которые нужно обратить внимание
				Иначе результат представь в следующем формате, соблюдая правильность пробелов, название элементов оставляй на английском:
				Ошибка: {Краткое и понятное объяснение ошибки}
				Решение: {Пошаговое решение ошибки по пунктам. Не должно превышать 1000 символов. Используй больше конкретных команд kubectl, docker, curl и др. Используй только существующие команды (например для pull образа нужно использовать docker pull, а не kubectl pull)}
				Пункты в решении должны быть пронумерованы в правильном порядке (например 1 пункт - проверка логов, последний пункт - перезагрузка сервиса)`},

		"prom_conf_prompt": Prompt{Template: `Упрости следующее сообщение об ошибке Prometheus.
				Эта ошибка возникла при проверке конфигурационного файла Prometheus.
				Предоставь пошаговые инструкции по исправлению с предложениями, со ссылками на документацию Prometheus, если это уместно.
				Результат представь в следующем формате, соблюдая правильность пробелов, название элементов оставляй на английском:
				Ошибка: {Краткое и понятное объяснение ошибки}
				Решение: {Пошаговое решение ошибки по пунктам. Не должно превышать 1000 символов. Старайся давать конкретные команды к каждому совету и действия для выявления и исправления проблемы}`},

		"prom_relabel_prompt": Prompt{Template: `Верни ответ, начиная с:
				Следующее представляет собой список в форме:
				job_name:
				{Имя задачи Prometheus}
				relabel_configs:
				{Конфигурации перемаркировки Prometheus}
				kubernetes_sd_configs:
				{Конфигурация обнаружения сервисов Prometheus}
				Для каждого job_name опиши соответствующие метки сервисов и подов Kubernetes,
				пространства имен, порты и контейнеры, которые они охватывают.
				Верни сообщение:
				Обнаружены и разобраны конфигурации сканирования Prometheus.
				Чтобы цели могли быть сканированы Prometheus, убедитесь, что они работают с
				хотя бы одним из следующих наборов меток:
				Затем для каждой задачи выведи в таком формате:
				- Задача: {job_name}
				  - Метки сервиса:
					- {список меток сервиса}
				  - Метки пода:
					- {список меток пода}
				  - Пространства имен:
					- {список пространств имен}
				  - Порты:
					- {список портов}
				  - Контейнеры:
					- {список имен контейнеров}
				Соблюдай правильность пробелов, название элементов оставляй на английском`},

		"kyverno_prompt": Prompt{Template: `Упрости следующее предупреждение Kyverno.
				Предложи наиболее вероятное решение в виде команды kubectl.
				Выведи результат в следующем формате, для решения покажи только команду kubectl, соблюдай правильность пробелов, название элементов оставляй на английском:
				Ошибка: {Подробное объяснение ошибки}
				Решение: {команда kubectl}`},

		// raw промт взят с https://developers.sber.ru/docs/ru/gigachat/prompts-hub/overview
		"raw_prompt": Prompt{Template: `{
				"model": "%s"
				"language": "%s",
				"messages": [
						{
							"role": "system",
							"content": "%s"
						},
						{
							"role": "user",
							"content": "%s"
						}
				]
				}`},
	},
}
