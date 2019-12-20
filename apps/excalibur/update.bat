cd E:\Excalibur\Document\CSV\带批注EXCEL\ExportCSV
svn update
cd E:\gopath\src\gitlab.com\megatech\serverex\apps\excalibur
copy E:\Excalibur\Document\CSV\带批注EXCEL\ExportCSV\*.csv E:\gopath\src\gitlab.com\megatech\serverex\apps\excalibur\data\csv\
make codegen
make
