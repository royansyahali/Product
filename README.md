Buat sebuah Rest API product dengan bahasa pemrograman go dengan RDMS mysql atau postgre

Buat 2 endpoint.

API Add product -> localhost:8080/product | data: {name : string, price : int, description : string, quantity : int}

API List product dengan sorting -> localhost:8080/product

Dengan spesifikasi product:

product id

product name

product price

product description

product quantity

Rest API bisa melakukan tambah product

Rest API bisa mengurutkan berdasarkan berikut:

product terbaru -> localhost:8080/product?order_by=created_at&order=desc

product harga termurah -> localhost:8080/product?order_by=price&order=asc

product harga terendah -> localhost:8080/product?order_by=price&order=asc

sort by product name (A-Z) dan (Z-A) -> localhost:8080/product?order_by=name&order=asc | localhost:8080/product?order_by=name&order=desc 

Database yang digunakan adalah Mysql

Redis sebagai Cache dengan Exp 5 Menit

Framework yang digunakan Gin
