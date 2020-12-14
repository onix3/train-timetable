# Train Timetable
Смотри расписание движения поездов между двумя городами на удобной диаграмме

## Особенности
  - Диаграмма отражает рейсы между двумя городами в обоих направлениях
  - Добавлено 35 городов → [🗺](https://yandex.ru/maps/?um=constructor%3Ad7846cc6cf6516763b73d7d45ca4bff3188045e89037947249cd0afb2dad6f6d&source=constructorLink)
  - Настройка цветов темы
  
## Установка
Выполни команду:
```
go get -u github.com/onix3/train-timetable
```
Также в [Releases](https://github.com/onix3/train-timetable/releases) можно скачать и запустить **[exe](https://github.com/onix3/train-timetable/releases/download/v0.1.0/train-timetable.exe)**

## Примеры использования

![](https://user-images.githubusercontent.com/74375829/101986094-ed602d00-3c9c-11eb-91b3-4e51a04c6452.gif) | ![](https://user-images.githubusercontent.com/74375829/101986097-f18c4a80-3c9c-11eb-8bf9-77269de14cd8.png) | ![](https://user-images.githubusercontent.com/74375829/101986099-f3560e00-3c9c-11eb-8526-4176870282aa.png) | ![](https://user-images.githubusercontent.com/74375829/101986101-f3eea480-3c9c-11eb-8d98-a2dd44b74561.png)
------------ | ------------ | ------------ | ------------ 
![](https://user-images.githubusercontent.com/74375829/101986498-2a2d2380-3c9f-11eb-98c9-9a9007a365e0.gif) | ![](https://user-images.githubusercontent.com/74375829/101986110-fbae4900-3c9c-11eb-9353-73b27b36aed7.png) | ![](https://user-images.githubusercontent.com/74375829/101986113-fcdf7600-3c9c-11eb-9f5c-acc979e7c06a.png) | ![](https://user-images.githubusercontent.com/74375829/101986115-fd780c80-3c9c-11eb-9ad7-1471c1d98958.png)

## Личная предпосылка
Нечасто, но приходилось планировать поездку в другой город туда и обратно в тот же день. 
К тому же, нужно было выбрать подходящее "окно" времени пребывания в другом городе — и вместо сравнения рейсов в двух вкладках браузера мне было бы удобнее видеть расписание целиком на диаграмме по времени

## Детали
  - Критерий выбора городов — население больше 20 тысяч человек и наличие железнодорожной станции
  - Рейсы в прямом направлении обозначены зелёным узором, в обратном — синим
  - Нередко последние рейсы начинаются в текущие сутки и заканчиваются в следующие (например, 22:27 – 07:41). Диапазон следующих суток на диаграмме темнее фона
  - При нажатии F1 появится справка, при нажатии Enter последняя диаграмма сохранится на Рабочий стол. 
  - Во вкладке "Тема" добавлены виджеты, которых не хватает для демонстрации изменения цвета
  - Желая сделать размер окна приложения минимальным, диаграмма рейсов и карта городов представлены миниатюрами, по нажатию на которые изображение разворачивается на весь экран
  - Данные предоставлены сервисом [Яндекс.Расписания](https://rasp.yandex.by). Это упоминание добавлено и в приложение, дабы соблюсти условия пользования сервиса [•](https://yandex.ru/legal/rasp_api/index.html)
  - Кэширование:
    - Расписание кэшируется: запрос, повторённый в тот же день, не будет осуществлён, а будут использоваться данные, полученные по этому запросу ранее
    - Кэш расписания доступен и при отсутствии подключения к Интернету 
    - Неактуальный кэш (расписание за прошлые дни) удаляется
    - Кэшируются расписание, тема и последние выбранные города
    - При первом запуске приложения (при отсутствии файла конфигурации в `c:\Users\{user_name}\AppData\Roaming\fyne\train-timetable\`) используется тема по умолчанию

## Используемые библиотеки

Библиотека | Лицензия | Применение в проекте
------------ | ------------ | ------------
[fyne](https://fyne.io/) | BSD-3-Clause License | GUI
[lusingander/fyne-theme-generator](https://github.com/lusingander/fyne-theme-generator) | MIT License | Редактор темы (заимствован и переделан)
[fogleman/gg](https://github.com/fogleman/gg) | MIT License | Отрисовка диаграммы
[anthonynsimon/bild](https://github.com/anthonynsimon/bild) | MIT License | Изменение размера и обрезка изображения
