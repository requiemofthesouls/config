# Wrapper Config

## Описание
Данный модуль нужен для конфигурации сервиса.

Внутри модуля реализована обертка для пакетов [viper](https://github.com/spf13/viper)

В нем описан interface Wrapper с необходимыми для работы методами, он может обновляться при необходимости

Конфигурационный файл сервиса должен быть формата yaml.

Долнительно в модуле реализована работа с ENV. 
Например, в модуле для переменных окружений используется префикс "ENV".
Реестр environments будет искать env переменные, начинающиеся с "ENV_".

## Пример использования

### Viper

Использование через конструктор.
``` go
package main

import (
	"log"

	"github.com/requiemofthesouls/container/config"
)

func main() {
	var (
		cfg config.Wrapper
		err error
	)
	if cfg, err = config.New("./cfg/config.yaml"); err != nil {
		log.Fatal(err)
	}
	
	log.Println(cfg.Get("service-name"))
}
```
Использование через definition.
``` go
package main

import (
	"log"

	"github.com/requiemofthesouls/container"

	"github.com/requiemofthesouls/container/config"
	cfgCont "github.com/requiemofthesouls/container/config/container"
)

func main() {
	var cfg config.Wrapper
	if err := container.Container.Fill(cfgCont.DIWrapper, &conf); err != nil {
		log.Fatal(err)
	}

	log.Println(conf.Get("service-name"))
}
```
## Зависимости от модулей
- [container](https://github.com/requiemofthesouls/container/-/blob/main/README.md)
