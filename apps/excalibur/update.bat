cd E:\Excalibur\Document\CSV\����עEXCEL\ExportCSV
svn update
cd E:\gopath\src\gitlab.com\megatech\serverex\apps\excalibur
copy E:\Excalibur\Document\CSV\����עEXCEL\ExportCSV\*.csv E:\gopath\src\gitlab.com\megatech\serverex\apps\excalibur\data\csv\
make codegen
make
