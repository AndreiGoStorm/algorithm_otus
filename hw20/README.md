## Домашнее задание №20
Программа сжатия файлов на основе алгоритма RLE. 

Использует два алгоритма simple RLE и Improved RLE

### 1) Создать бинарный файл 
Чтобы создать бинарный файл используем команду
```text
go build -o ./rle-app
```
или
```text
make build
```

### 2) Выбор алгоритма работы
Для Simple RLE (режим по умолчанию)
```text
./rle-app -from compress.txt
```
Для Improved RLE
```text
./rle-app -from compress.txt -mode
```

### 3) Вызов функции Compress/Decompress
Вызов Compress по умолчанию
```text
./rle-app -from compress.txt
```
Вызов Decompress
```text
./rle-app -from compress.rle -d
```
или для режима Improved RLE
```text
./rle-app -from compress.rle -mode -d
```

### 4) Список параметров
```text
  -d    Operation to decompress file. Compress operation by default.
  -from string
        File to read from path. Required.
  -mode
        Mode to choose algorithm: simple rle or improved rle. Simple rle by default.
```