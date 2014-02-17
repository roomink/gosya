# Configure go app

Считываем файлы из заданной директории.
* основной файл `./settings.yml`
* С учётом заданного  значения env выбирается `./settings/{{env_name}}.yml`
* Локальные настройки `./settings.local.yml`

type Settings struct {
  Advert int `yaml:"advert"`
}
var settings Settings

gosya.Merge(&settings, "paht/to/conf/dir","development")

settings.Advert
