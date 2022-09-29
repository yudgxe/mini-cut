# mini-cut

$ cut опции путь_к_файлу

## Опции
-d string  
&emsp;&emsp;&emsp;использовать другой разделитель (default "\t")  
-f value  
&emsp;&emsp;&emsp;выбрать поле/поля    
  
  -s    только строки с разделителем
  
  ## Примеры
  go run .\cmd\main.go  -d "|" -f 1,2  data/1.txt  
  go run .\cmd\main.go -f 1,2,3 data/1.txt           
  go run .\cmd\main.go -s -f 1 data/1.txt            
