## Проектная работа на тему "Архиватор Хаффмана и LZ77"
Программа для сжатия файлов на основе алгоритмов Huffman, LZ77 и RLE. 

### 1) Создать бинарный файл 
Чтобы создать бинарный файл используем команду
```text
go build -o ./app
```
или
```text
make build
```

### 2) Выбор алгоритма работы
Для алгоритма `Huffman` (режим по умолчанию)
```text
./app -from compress.txt
```
Для `LZ77`
```text
./rle-app -from compress.txt -lz
```
Для `RLE`
```text
./rle-app -from compress.txt -rle
```

### 3) Вызов функции Compress/Decompress
Вызов `Compress` по умолчанию
```text
./rle-app -from compress.txt
```
Вызов `Decompress`
```text
./rle-app -from compress.rle -d
```
Пример вызова decompress для алгоритма LZ77
```text
./rle-app -from compress.txt -lz -d
```

### 4) Список параметров
```text
  -d    Action to decompress file. Compress action by default.
  -from string
        File to read from path. Required.
  -lz
        Mode to choose algorithm lz77. Huffman algorithm by default.
  -rle
        Mode to choose algorithm rle. Huffman algorithm by default.
```