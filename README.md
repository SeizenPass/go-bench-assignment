# Amiran Kurman SE-1901   
Использовался easyjson для кодогенерации структуры User.  
Результаты запуска на моем ноутбуке за 1 секунду (ну дефолтно, да).
```shell script
BenchmarkSlow-8                2         644900900 ns/op        19767360 B/op     189814 allocs/op
BenchmarkFast-8               68          19047566 ns/op         2247012 B/op       7328 allocs/op
```  
В Makefile можете найти все нужные команды:
```shell script
gen - для бенчмарка и получения результатов
mem - для профилирования по памяти
cpu - для профилирования по использованию процессора
json - для генерации easyjson
``` 